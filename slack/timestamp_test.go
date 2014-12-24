package slack_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/danryan/go-slack/slack"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp_MarshalJSON(t *testing.T) {
	now := int64(1419192595)

	ts := slack.Timestamp(time.Unix(now, 0))
	b, err := ts.MarshalJSON()
	assert.Nil(t, err)

	actual, _ := strconv.Atoi(string(b))
	assert.Equal(t, 1419192595, actual)
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {

}
