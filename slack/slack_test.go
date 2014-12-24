package slack_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/danryan/go-slack/slack"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the client being tested.
	client *slack.Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server

	fixtures = make(map[string]string)

	fixtureNames = []string{
		"channels.list",
		"channels.info",
	}
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = slack.New("team", "api_token")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func readTestFixture(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	s := string(b)

	return s, nil
}

func init() {
	for _, n := range fixtureNames {
		f, _ := readTestFixture(fmt.Sprintf("../test/%v.json", n))
		fixtures[n] = f
	}
}
