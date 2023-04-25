package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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
