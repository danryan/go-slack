package slack

// HistoryOptions is an options type that modifies the history query for channels, groups, and IMs.
type HistoryOptions struct {
	Channel string `url:"channel"`
	Latest  string `url:"latest,omitempty"`
	Oldest  string `url:"oldest,omitempty"`
	Count   int    `url:"count,omitempty"`
}

// History is the result type for any history request.
type History struct {
	Messages []Message `json:"messages,omitempty"`
	HasMore  bool      `json:"has_more,omitempty"`
}
