package slack_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/danryan/go-slack/slack"
	"github.com/stretchr/testify/assert"
)

func TestClientErrorHandling(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/fake.method", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"ok":false,"error":"unknown_method"}`)
	})

	var out map[string]interface{}
	_, err := client.Get("fake.method", out)
	e := slack.Error(err)

	if assert.NotNil(t, e) {
		assert.IsType(t, &slack.ErrorResponse{}, e)
		assert.False(t, e.Ok)
		assert.Equal(t, "unknown_method", e.Message)
		assert.Equal(t, 200, e.Response.StatusCode) // :'(
	}
}
