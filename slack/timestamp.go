package slack

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Timestamp is a custom time.Time type to convert Slack's use of seconds-since-epoch timestamps
type Timestamp time.Time

// MarshalJSON implements the json encoder interface
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalJSON implements the json decoder interface
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}

func (t *Timestamp) String() string {
	i := time.Time(*t).Unix()

	return strconv.FormatInt(i, 10)
}

type MessageTimestamp time.Time

func (t MessageTimestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(t)
	sec := strconv.FormatInt(ts.Unix(), 10)
	msec := strconv.FormatInt(ts.UnixNano(), 10)
	msec = strings.TrimRight(strings.Replace(msec, sec, "", 1), "0")

	return []byte(`"` + sec + `.` + msec + `"`), nil
}

func (t *MessageTimestamp) UnmarshalJSON(b []byte) error {
	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	ts := strings.Split(str, ".")
	if len(ts) != 2 {
		return errors.New("could not unmarshal into MessageTimestamp")
	}

	sec, err := strconv.Atoi(ts[0])
	if err != nil {
		return err
	}

	msec, err := strconv.Atoi(ts[1])
	if err != nil {
		return err
	}

	*t = MessageTimestamp(time.Unix(int64(sec), int64(msec*1000)))

	return nil
}

// func (t *MessageTimestamp) String() string {
// 	return time.Time(*t).String()
// }
