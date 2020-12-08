package model

type Eventnotification struct {
	Id              string
	Ttid            string
	Type            string
	Origin          string
	Event_data      map[string]interface{}
	Processing_time int64
	Reference_id    string
}
