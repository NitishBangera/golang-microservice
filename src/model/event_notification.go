package model

// Eventnotification structure
type Eventnotification struct {
	ID             string                 `json:"id"`
	Ttid           string                 `json:"ttid"`
	Type           string                 `json:"type"`
	Origin         string                 `json:"origin"`
	EventData      map[string]interface{} `json:"event_data"`
	ProcessingTime int64                  `json:"processing_time"`
	ReferenceID    string                 `json:"reference_id"`
}
