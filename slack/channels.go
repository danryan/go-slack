package slack

import "net/http"

// ChannelsService handles channel-related API requests.
type ChannelsService struct {
	client *Client
}

// Channel represents a Slack channel.
type Channel struct {
	ID          string           `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	IsChannel   bool             `json:"is_channel,omitempty"`
	Created     Timestamp        `json:"created,omitempty"`
	Creator     string           `json:"creator,omitempty"`
	IsArchived  bool             `json:"is_archived,omitempty"`
	IsGeneral   bool             `json:"is_general,omitempty"`
	IsMember    bool             `json:"is_member,omitempty"`
	Topic       Topic            `json:"topic,omitempty"`
	Members     []string         `json:"members,omitempty"`
	Purpose     Purpose          `json:"purpose,omitempty"`
	NumMembers  int              `json:"num_members,omitempty"`
	Latest      Message          `json:"latest,omitempty"`
	LastRead    MessageTimestamp `json:"last_read,omitempty"`
	UnreadCount int              `json:"unread_count,omitempty"`
}

type channel struct {
	*Channel `json:"channel,omitempty"`
}

// Info provides a generic type for additional data
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

type ChannelOptions struct {
	Channel string `url:"channel,omitempty"`
	Name    string `url:"name,omitempty"`
	Pretty  bool   `url:"pretty,int,omitempty"`
}

// List returns a slice of Channels, the raw http.Response, and an optional error.
func (s *ChannelsService) List() ([]Channel, *http.Response, error) {
	var ch *channels

	res, err := s.client.Get("channels.list", &ch)
	if err != nil {
		return nil, res, err
	}

	return ch.Channels, res, nil
}

// History retrieves the message history of a channel.
func (s *ChannelsService) History(ch string, opts *HistoryOptions) (*History, *http.Response, error) {
	var h *History

	opts.Channel = ch
	path, err := addOptions("channels.history", opts)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Get(path, &h)
	if err != nil {
		return nil, res, err
	}

	return h, res, nil
}

// Info retrieves a channel.
func (s *ChannelsService) Info(id string, opts *ChannelOptions) (*Channel, *http.Response, error) {
	var ch *channel

	if opts == nil {
		opts = &ChannelOptions{}
	}
	opts.Channel = id

	path, err := addOptions("channels.info", opts)

	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Get(path, &ch)
	if err != nil {
		return nil, res, err
	}

	return ch.Channel, res, nil
}

// Rename sets a channel's name, optionally returning an error.
func (s *ChannelsService) Rename(id, name string, opts *ChannelOptions) (*Channel, *http.Response, error) {
	var ch *channel

	if opts == nil {
		opts = &ChannelOptions{}
	}
	opts.Channel = id
	opts.Name = name

	path, err := addOptions("channels.rename", opts)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Get(path, &ch)
	if err != nil {
		return nil, res, err
	}

	return ch.Channel, res, nil
}
