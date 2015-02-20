package slack_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/danryan/go-slack/slack"
	"github.com/stretchr/testify/assert"
)

func TestChannels_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels.list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fixtures["channels.list"])
	})

	ch, _, err := client.Channels.List()
	if assert.Nil(t, err) {
		assert.Len(t, ch, 3)
		assert.Equal(t, "C0000AAAA", ch[0].ID)
	}
}

func TestChannels_Info(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels.info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fixtures["channels.info"])
	})

	ch, _, err := client.Channels.Info("C0000AAAA", nil)
	if assert.Nil(t, err) {
		assert.Equal(t, "C0000AAAA", ch.ID)
	}
}

func TestChannels_Info_invalid(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels.info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"ok":false,"error":"channel_not_found"}`)
	})

	ch, _, err := client.Channels.Info("C0000AAAA", nil)
	e := slack.Error(err)

	if assert.NotNil(t, e) {
		assert.Nil(t, ch)
		assert.IsType(t, &slack.ErrorResponse{}, e)
		assert.False(t, e.Ok)
		assert.Equal(t, "channel_not_found", e.Message)
		assert.Equal(t, 200, e.Response.StatusCode) // :'(
	}
}

func TestChannels_Rename(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels.rename", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fixtures["channels.rename"])
	})

	ch, _, err := client.Channels.Rename("C0000AAAA", "asdf", nil)
	if assert.Nil(t, err) {
		assert.Equal(t, "asdf", ch.Name)
	}
}
