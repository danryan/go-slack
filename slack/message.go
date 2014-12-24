package slack

// Message type
type Message struct {
	Text      string           `json:"text,omitempty"`
	Username  string           `json:"username,omitempty"`
	Type      string           `json:"type,omitempty"`
	Subtype   string           `json:"subtype,omitempty"`
	Timestamp MessageTimestamp `json:"ts,omitempty"`
}
