package slack

import "net/http"

type ChannelsService struct {
	client *Client
}

type Channel struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	IsChannel  bool      `json:"is_channel,omitempty"`
	Created    Timestamp `json:"created,omitempty"`
	Creator    string    `json:"creator,omitempty"`
	IsArchived bool      `json:"is_archived,omitempty"`
	IsGeneral  bool      `json:"is_general,omitempty"`
	IsMember   bool      `json:"is_member,omitempty"`
	Topic      Topic     `json:"topic,omitempty"`
	Members    []string  `json:"members,omitempty"`
	Purpose    Purpose   `json:"purpose,omitempty"`
	NumMembers int       `json:"num_members,omitempty"`
	Latest     Message   `json:"latest,omitempty"`
}

type Info struct {
	Value   string `json:"value,omitempty"`
	Creator string `json:"creator,omitempty"`
	LastSet Timestamp
}

// Topic type for Channel.Topic
type Topic Info

// Purpose type for Channel.Purpose
type Purpose Info

type channels struct {
	Channels []Channel `json:"channels,omitempty"`
}

// List returns a slice of Channels, the raw http.Response, and an optional error.
func (s *ChannelsService) List() ([]Channel, *http.Response, error) {
	ch := new(channels)

	res, err := s.client.Get("channels.list", &ch)
	if err != nil {
		return nil, res, err
	}

	return ch.Channels, res, nil
}

// History retrieves the message history of a channel.
func (s *ChannelsService) History(ch string, opts *HistoryOptions) (h History, res *http.Response, err error) {
	var path string

	opts.Channel = ch
	if path, err = addOptions("channels.history", opts); err != nil {
		return nil, nil, err
	}

	if res, err = s.client.Get(path, &h); err != nil {
		return nil, res, err
	}

	return h, res, nil
}
