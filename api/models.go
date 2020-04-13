package api

// Common json response data
type common struct {
	ErrorText  string `json:"errorText"`
	Error      string `json:"error"`
	ServerDate string `json:"serverdate"`
	ServerTime string `json:"servertime"`
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
	Common                    common
	StopLocations             []StopLocation  `json:"StopLocation"`
	CoordLocations            []CoordLocation `json:"CoordLocation"`
	NoNamespaceSchemaLocation string          `json:"noNamespaceSchemaLocation"`
}
