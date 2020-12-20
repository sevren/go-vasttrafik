package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sevren/go-vasttrafik/api"
	"github.com/sevren/go-vasttrafik/auth"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Debug bool `env:"DEBUG"`
	VT    auth.Config
	API   api.Config
}

type intent struct {
	DisplayName string `json:"displayName"`
}

type queryResult struct {
	Intent       intent `json:"intent"`
	LanguageCode string `json:"languageCode"`
}

type text struct {
	Text []string `json:"text"`
}

type message struct {
	Text text `json:"text"`
}

// webhookRequest is used to unmarshal a WebhookRequest JSON object. Note that
// not all members need to be defined--just those that you need to process.
// As an alternative, you could use the types provided by
// the Dialogflow protocol buffers.
type webhookRequest struct {
	Session     string      `json:"session"`
	ResponseID  string      `json:"responseId"`
	QueryResult queryResult `json:"queryResult"`
}

// webhookResponse is used to marshal a WebhookResponse JSON object. Note that
// not all members need to be defined--just those that you need to process.
// As an alternative, you could use the types provided by
// the Dialogflow protocol buffers.
type webhookResponse struct {
	FulfillmentMessages []message `json:"fulfillmentMessages"`
}

func checkConfig(cfg *config) error {
	if cfg.VT.Key == "" {
		return fmt.Errorf("VT_KEY can't be empty")
	}

	if cfg.VT.Secret == "" {
		return fmt.Errorf("VT_SECRET can't be empty")
	}

	return nil
}

func main() {
	log.Info("Initalizing go-vasttrafik...")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Could not parse config %s", err)
	}

	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	if err := checkConfig(&cfg); err != nil {
		log.Fatal(err)
	}

	a, err := auth.GetAccessToken(cfg.VT.Key, cfg.VT.Secret)
	if err != nil {
		log.Fatal(err)
	}
	cl := api.New(cfg.API.BaseURL, a)

	log.Info(cl.Auth.Token)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", HandleWebhookRequest(cl))
	log.Fatal(http.ListenAndServeTLS(":3000", "./cert.pem", "./private.key", r))
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "ERROR: %v", err)
}

func returnAPIErrorMessage(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(api.DialogFlowResponse{
		Speech: err.Error(),
	})
}

// welcome creates a response for the welcome intent.
func welcome(request webhookRequest) (webhookResponse, error) {
	response := webhookResponse{
		FulfillmentMessages: []message{
			{
				Text: text{
					Text: []string{"Welcome from Dialogflow Go Webhook"},
				},
			},
		},
	}
	return response, nil
}

// HandleWebhookRequest handles WebhookRequest and sends the WebhookResponse.
func HandleWebhookRequest(cl *api.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request webhookRequest
		var response webhookResponse
		var err error

		// Read input JSON
		if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
			handleError(w, err)
			return
		}
		log.Printf("Request: %+v", request)

		log.Infof("The language is %s", request.QueryResult.LanguageCode)

		// Call intent handler
		switch intent := request.QueryResult.Intent.DisplayName; intent {
		case "Check-bus":
			response, err = welcome(request)
		default:
			err = fmt.Errorf("Unknown intent: %s", intent)
		}
		if err != nil {
			handleError(w, err)
			return
		}
		log.Printf("Response: %+v", response)

		resp, err := cl.GetTrip()
		if err != nil {
			returnAPIErrorMessage(w, err)
		}

		now, err := time.Parse("15:04", time.Now().Format("15:04"))

		originTime, err := time.Parse("15:04", resp.TripList.Trip[0].Leg.Origin.Time)
		destTime, err := time.Parse("15:04", resp.TripList.Trip[0].Leg.Destination.Time)

		log.Infof("You have %s to get to the stop", originTime.Sub(now))
		if err == nil {
			log.Infof("The trip will take %s minutes", destTime.Sub(originTime))
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Println(originTime.Sub(now).Minutes())

		items := []*api.RichResponseItem{}
		speech := ""
		if request.QueryResult.LanguageCode == "en" {
			speech = fmt.Sprintf("The next trip is at %s, You have %s minutes to get to the stop", resp.TripList.Trip[0].Leg.Origin.Time, fmt.Sprintf("%.0f", originTime.Sub(now).Minutes()))
		} else if request.QueryResult.LanguageCode == "ru" {
			speech = fmt.Sprintf("Следющий автовус в %s, У вас есть %s минуты чтобы добраться до остановки", resp.TripList.Trip[0].Leg.Origin.Time, fmt.Sprintf("%.0f", originTime.Sub(now).Minutes()))
		}

		items = append(items, &api.RichResponseItem{SimpleResponse: &api.SimpleResponse{
			TextToSpeech: speech,
		}})

		webHookResp := api.WebHookResponse{
			Payload: &api.Payload{
				Google: &api.Google{
					ExpectUserResponse: false,
					RichResponse: &api.RichResponse{
						Items: items,
					},
				},
			},
		}

		if err = json.NewEncoder(w).Encode(&webHookResp); err != nil {
			handleError(w, err)
			return
		}
	}
}
