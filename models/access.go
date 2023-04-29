package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidTokenFlag is returned when a token flag not defined in
// the registry is encountered.
var ErrInvalidTokenFlag = errors.New("invalid token flag")

// TokenFlag represents GNAP access token flags.
type TokenFlag string

// Registry of access token flags.
const (
	FlagBearer  TokenFlag = "bearer"
	FlagDurable TokenFlag = "durable"
)

// flagRegistry is a quick mapping from allowed values to token flags.
var flagRegistry = map[string]TokenFlag{
	"bearer":  FlagBearer,
	"durable": FlagDurable,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (tf TokenFlag) MarshalJSON() ([]byte, error) {
	_, ok := flagRegistry[string(tf)]
	if ok {
		return json.Marshal(string(tf))
	}
	return nil, ErrInvalidTokenFlag
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (tf *TokenFlag) UnmarshalJSON(data []byte) error {
	var flag string
	err := json.Unmarshal(data, &flag)
	if err == nil {
		*tf = TokenFlag(flag)
		return nil
	}
	return ErrInvalidTokenFlag
}

// TokenRequest represents access token request object for requesting
// access to resources.
type TokenRequest struct {
	Access []AccessRight `json:"access"`
	Label  string        `json:"label,omitempty"`
	Flags  []TokenFlag   `json:"flags,omitempty"`
}

// TokenResponse represents access token grant response by the AS.
type TokenResponse struct {
	URI         URL         `json:"uri"`
	Wait        int         `json:"wait"`
	AccessToken AccessToken `json:"access_token"`
}

// AccessToken represents the access token granted by the AS.
type AccessToken struct {
	Value     string        `json:"value"`
	Label     string        `json:"label,omitempty"`
	Manage    URL           `json:"manage,omitempty"`
	Access    []AccessRight `json:"access"`
	ExpiresIn int           `json:"expires_in,omitempty"`
	Key       ClientKey     `json:"key,omitempty"`
	Flags     []TokenFlag   `json:"flags,omitempty"`
}

// ContinueToken represents continuation access token to
// be presented for continuation request.
type ContinueToken struct {
	Value     string      `json:"value"`
	Label     string      `json:"label,omitempty"`
	Manage    URL         `json:"manage,omitempty"`
	ExpiresIn int         `json:"expires_in,omitempty"`
	Flags     []TokenFlag `json:"flags,omitempty"`
}
