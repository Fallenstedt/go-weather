package weather

type alertsResponse struct {
	Alerts  []Alerts `json:"features"`
	Title   string   `json:"title"`
	Updated string   `json:"updated"`
}

type Alerts struct {
	Id         string `json:"id"`
	Properties AlertProperties `json:"properties"`
}

type AlertProperties struct {
	Sent        string `json:"sent"`
	Effective   string `json:"effective"`
	Onset       string `json:"onset"`
	Expires     string `json:"expires"`
	Ends        string `json:"ends"`
	Status      string `json:"status"`
	MessageType string `json:"messageType"`
	Category    string `json:"category"`
	Severity    string `json:"severity"`
	Certainty   string `json:"certainty"`
	Urgency     string `json:"urgency"`
	Event       string `json:"event"`
	SenderName  string `json:"senderName"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	Instruction string `json:"instruction,omitempty"`
	Response    string `json:"response"`
}