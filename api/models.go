package api

// Common json response data
type Common struct {
	ErrorText  string `json:"errorText,omitempty"`
	Error      string `json:"error,omitempty"`
	ServerDate string `json:"serverdate,omitempty"`
	ServerTime string `json:"servertime,omitempty"`
}

type StopLocation struct {
	ID     string `json:"id"`
	Lon    string `json:"lon"`
	Idx    string `json:"idx"`
	Weight string `json:"weight"`
	Name   string `json:"name"`
	Track  string `json:"track"`
	Lat    string `json:"lat"`
}

type CoordLocation struct {
	Lon  string `json:"lon"`
	Idx  string `json:"idx"`
	Name string `json:"name"`
	Type string `json:"type"`
	Lat  string `json:"Lat"`
}

// GET locations.name request
type getLocationRequest struct {
	Input         string `json:"input"`
	Format        string `json:"format"`
	JsonPCallback string `json:"jsonpCallback"`
}

// GET locations.name response
type getLocationResponse struct {
	Common                    *Common
	StopLocations             []*StopLocation  `json:"StopLocation"`
	CoordLocations            []*CoordLocation `json:"CoordLocation"`
	NoNamespaceSchemaLocation string           `json:"noNamespaceSchemaLocation"`
}

type getTripRequest struct {
	OriginID      string `json:"originId"`
	DestID        string `json:"destId"`
	Format        string `json:"format"`
	JsonPCallback string `json:"jsonpCallback"`
}

type Leg struct {
	Name             string   `json:"name,omitempty"`
	Sname            string   `json:"sname,omitempty"`
	JourneyNumber    string   `json:"journeyNumber,omitempty"`
	Type             string   `json:"type,omitempty"`
	ID               string   `json:"id,omitempty"`
	Direction        string   `json:"direction,omitempty"`
	FgColor          string   `json:"fgColor,omitempty"`
	BgColor          string   `json:"bgColor,omitempty"`
	Stroke           string   `json:"stroke,omitempty"`
	Accessibility    string   `json:"accessibility,omitempty"`
	Origin           *Details `json:"Origin,omitempty"`
	Destination      *Details `json:"Destination,omitempty"`
	JourneyDetailRef *Ref     `json:"JourneyDetailRef,omitempty"`
	GeometryRef      *Ref     `json:"GeometryRef,omitempty"`
	Booking          bool     `json:"booking,omitempty"`
	Notes            *Notes   `json:"Notes,omitempty"`
	Night            bool     `json:"night,omitempty"`
	Reachable        bool     `json:"reachable,omitempty"`
	Cancelled        bool     `json:"cancelled,omitempty"`
}

type Ref struct {
	Ref string `json:"ref,omitempty"`
}

type Details struct {
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	RouteIdx   string `json:"routeIdx,omitempty"`
	Time       string `json:"time,omitempty"`
	Date       string `json:"date,omitempty"`
	Track      string `json:"track,omitempty"`
	RtTime     string `json:"rtTime,omitempty"`
	RtDate     string `json:"rtDate,omitempty"`
	Sym        string `json:"$,omitempty"`
	Cancelled  bool   `json:"cancelled,omitempty"`
	RtTrack    string `json:"rtTrack,omitempty"`
	Notes      *Notes `json:"Notes,omitempty"`
	DirectDate string `json:"directdate,omitempty"`
	DirectTime string `json:"directtime,omitempty"`
}

type Notes struct {
	Note []*Note `json:"Note,omitempty"`
}

type Note struct {
	Priority string `json:"priority,omitempty"`
	Severity string `json:"severity,omitemptyy"`
	Key      string `json:"key,omitempty"`
}

type Trip struct {
	Leg            *Leg   `json:"Leg,omitempty"`
	TravelWarrenty bool   `json:"travelWarrenty,omitempty"`
	Valid          bool   `json:"valid,omitempty"`
	Alternative    bool   `json:"alternative,omitempty"`
	Type           string `json:"type,omitempty"`
}

type Trips struct {
	Common                    *Common
	Trip                      []*Trip `json:"Trip,omitempty"`
	NoNamespaceSchemaLocation string  `json:"noNamespaceSchemaLocation,omitempty"`
}
type getTripResponse struct {
	TripList *Trips `json:"TripList,omitempty"`
}

// DialogFlowRequest struct
type DialogFlowRequest struct {
	Result struct {
		Action string `json:"action"`
	} `json:"result"`
	OriginalRequest DialogFlowOriginalRequest `json:"originalRequest"`
}

// DialogFlowOriginalRequest struct
type DialogFlowOriginalRequest struct {
	Data DialogFlowOriginalRequestData `json:"data"`
}

// DialogFlowOriginalRequestData struct
type DialogFlowOriginalRequestData struct {
	Device DialogFlowOriginalRequestDevice `json:"device"`
}

// DialogFlowOriginalRequestDevice struct
type DialogFlowOriginalRequestDevice struct {
	Location DialogFlowOriginalRequestLocation `json:"location"`
}

// DialogFlowOriginalRequestLocation struct
type DialogFlowOriginalRequestLocation struct {
	Coordinates DialogFlowOriginalRequestCoordinates `json:"coordinates"`
}

// DialogFlowOriginalRequestCoordinates struct
type DialogFlowOriginalRequestCoordinates struct {
	Lat  float32 `json:"latitude"`
	Long float32 `json:"longitude"`
}

// DialogFlowResponse struct
type DialogFlowResponse struct {
	Speech string `json:"speech"`
}

// DialogFlowLocationResponse struct
type DialogFlowLocationResponse struct {
	Speech string                 `json:"speech"`
	Data   DialogFlowResponseData `json:"data"`
}

// DialogFlowResponseData struct
type DialogFlowResponseData struct {
	Google DialogFlowResponseGoogle `json:"google"`
}

// DialogFlowResponseGoogle struct
type DialogFlowResponseGoogle struct {
	ExpectUserResponse bool                           `json:"expectUserResponse"`
	IsSsml             bool                           `json:"isSsml"`
	SystemIntent       DialogFlowResponseSystemIntent `json:"systemIntent"`
}

// DialogFlowResponseSystemIntent struct
type DialogFlowResponseSystemIntent struct {
	Intent string                             `json:"intent"`
	Data   DialogFlowResponseSystemIntentData `json:"data"`
}

// DialogFlowResponseSystemIntentData struct
type DialogFlowResponseSystemIntentData struct {
	Type        string   `json:"@type"`
	OptContext  string   `json:"optContext"`
	Permissions []string `json:"permissions"`
}

/*
{
  "payload": {
    "google": {
      "expectUserResponse": true,
      "richResponse": {
        "items": [
          {
            "simpleResponse": {
              "textToSpeech": "this is a Google Assistant response"
            }
          }
        ]
      }
    }
  }
}*/

type WebHookResponse struct {
	Payload *Payload `json:"payload"`
}

type Payload struct {
	Google *Google `json:"google"`
}

type Google struct {
	ExpectUserResponse bool          `json:"expectUserResponse"`
	RichResponse       *RichResponse `json:"richResponse"`
}

type RichResponse struct {
	Items []*RichResponseItem `json:"items"`
}

type RichResponseItem struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}
