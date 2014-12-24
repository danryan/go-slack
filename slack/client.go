package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type Client struct {
	Team    string
	APIKey  string
	BaseURL *url.URL

	Channels *ChannelsService

	client *http.Client
}

func NewClient(team, key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse("https://slack.com/api/")
	c := &Client{
		client:  httpClient,
		APIKey:  key,
		Team:    team,
		BaseURL: u,
	}

	c.Channels = &ChannelsService{client: c}

	return c
}

type clientOptions struct {
	Token string `url:"token,omitempty"`
}

// NewRequest builds an http.Request, resolves relative URLs, and sets HTTP headers
func (c *Client) NewRequest(meth string, path string, input interface{}) (*http.Request, error) {
	// opt := &ClientOptions{c.APIKey}

	uri, err := addOptions(path, clientOptions{c.APIKey})
	if err != nil {
		return nil, err
	}

	ref, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(ref)

	buf := new(bytes.Buffer)
	if input != nil {
		if err := json.NewEncoder(buf).Encode(input); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(meth, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) Get(path string, output interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, output)
}

type responseClone struct {
	body *bytes.Buffer
	dup  *bytes.Buffer
}

// Do performs the request
func (c *Client) Do(req *http.Request, output interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, bodydup, err := cloneResponseBody(res)
	if err != nil {
		return res, err
	}

	if err := checkResponse(res); err != nil {
		return res, err
	}

	if err := checkError(bodydup, res); err != nil {
		return res, err
	}

	if output != nil {
		if w, ok := output.(io.Writer); ok {
			io.Copy(w, body)
		} else {
			err = json.NewDecoder(body).Decode(output)
		}
	}

	return res, err
}

// cloneResponseBody copies response body of http.Response r, and writes it to two bytes.Buffers,
// returning said buffers. This is necessary because slack returns `200 OK` even if the request
// fails, and makes us check the response body for errors.
func cloneResponseBody(r *http.Response) (io.Reader, io.Reader, error) {
	buf1, buf2 := new(bytes.Buffer), new(bytes.Buffer)

	defer r.Body.Close()

	if _, err := io.Copy(io.MultiWriter(buf1, buf2), r.Body); err != nil {
		return nil, nil, err
	}

	return buf1, buf2, nil
}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}
	u.RawQuery = u.RawQuery + `&` + qs.Encode()

	return u.String(), nil
}

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
