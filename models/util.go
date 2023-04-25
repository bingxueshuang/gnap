package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

// ErrInvalidURL is the error returned in case of invalid URL
// during json marshaling, unmarshaling or parsing of [URL].
var ErrInvalidURL = errors.New("invalid url")

// URL is wrapper around [url.URL] with json marshaling and unmarshaling.
type URL struct {
	*url.URL
}

// UnmarshalJSON implements [json.Unmarshaler].
func (u *URL) UnmarshalJSON(data []byte) error {
	var stringURL string
	err := json.Unmarshal(data, &stringURL)
	if err != nil {
		return err
	}
	if stringURL != "" {
		uu, err := url.Parse(stringURL)
		if err != nil {
			return err
		}
		u.URL = uu
		return nil
	}
	return ErrInvalidURL
}

// MarshalJSON implements [json.Marshaler].
func (u URL) MarshalJSON() ([]byte, error) {
	if u.URL == nil || u.String() == "" {
		return nil, ErrInvalidURL
	}
	return []byte(u.String()), nil
}

// ParseURL is a helper function for creating [URL].
// It is perfectly alright to directly create [URL]
// from struct literals.
func ParseURL(raw string) (URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return URL{}, fmt.Errorf("%w: %w", ErrInvalidURL, err)
	}
	return URL{u}, nil
}

// Seconds a wrapper around [time.Duration] to store
// the number of seconds only.
type Seconds time.Duration

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (s *Seconds) UnmarshalJSON(data []byte) error {
	var seconds int64
	err := json.Unmarshal(data, &seconds)
	if err != nil {
		return err
	}
	*s = Seconds(seconds * int64(time.Second))
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (s Seconds) MarshalJSON() ([]byte, error) {
	seconds := int(time.Duration(s).Seconds())
	return json.Marshal(seconds)
}

// Seconds is a convenience wrapper around [time.Duration.Seconds].
func (s Seconds) Seconds() int {
	return int(time.Duration(s).Seconds())
}

// NewSeconds is a convenience constructor for creating
// [Seconds] from numerical seconds.
func NewSeconds(seconds int) Seconds {
	return Seconds(int64(seconds) * int64(time.Second))
}
