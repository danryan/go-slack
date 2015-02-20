package slack

// New creates a new slack.Client and returns a pointer to it.
func New(team, key string) *Client {
	return NewClient(team, key, nil)
}

// Error takes an error interface and attempts to assert it is of type *slack.ErrorResponse,
// returning the ErrorResponse if possible, or nil if not.
func Error(err error) *ErrorResponse {
	if e, ok := err.(*ErrorResponse); ok {
		return e
	}

	return nil
}
