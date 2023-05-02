package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// ErrInvalidURL is returned when URL encountered is not valid.
var ErrInvalidURL = errors.New("invalid url")

// URL is a helper type that guards json unmarshaling only valid URLs.
type URL string

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidURL, err)
	}
	if s == "" {
		return fmt.Errorf("%w: empty string", ErrInvalidURL)
	}
	_, err = url.Parse(s)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidURL, err)
	}
	*u = URL(s)
	return nil
}

// NewURL is utility function to convert [url.URL] into [URL].
func NewURL(u *url.URL) URL {
	return URL(u.String())
}
