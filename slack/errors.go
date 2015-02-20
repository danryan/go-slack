package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response
	Ok       bool   `json:"ok"`
	Message  string `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v", r.Response.Request.Method, r.Response.Request.URL.Path, r.Message)
}

func checkError(r io.Reader, res *http.Response) error {
	v := &ErrorResponse{Response: res}

	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	if !v.Ok {
		return v
	}

	return nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return &ErrorResponse{Response: r}
}
