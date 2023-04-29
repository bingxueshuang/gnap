package models

import (
	"encoding/json"
	"errors"

	"golang.org/x/exp/slices"
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

// NewTokenRequest is a constructor for TokenRequest.
func NewTokenRequest(rights []AccessRight, options ...tokenRequestOption) (req TokenRequest, err error) {
	tr := &TokenRequest{Access: rights}
	for _, setter := range options {
		err = setter(tr)
		if err != nil {
			return
		}
	}
	return *tr, nil
}

// tokenRequestOption is functional parameter for TokenResponse constructor.
type tokenRequestOption func(*TokenRequest) error

// TokenResponse represents the access token granted by the AS.
type TokenResponse struct {
	Value     string        `json:"value"`
	Label     string        `json:"label,omitempty"`
	Manage    URL           `json:"manage,omitempty"`
	Access    []AccessRight `json:"access"`
	ExpiresIn int           `json:"expires_in,omitempty"`
	Key       ClientKey     `json:"key,omitempty"`
	Flags     []TokenFlag   `json:"flags,omitempty"`
}

// NewTokenResponse is constructor for TokenResponse.
func NewTokenResponse(value string, access []AccessRight, options ...tokenResponseOption) (res TokenResponse, err error) {
	tr := &TokenResponse{Value: value, Access: access}
	for _, setter := range options {
		err = setter(tr)
		if err != nil {
			return
		}
	}
	return *tr, nil
}

// tokenResponseOption is functional parameter for TokenResponse constructor.
type tokenResponseOption func(*TokenResponse) error

// ContinueToken represents continuation access token to
// be presented for continuation request.
type ContinueToken struct {
	Value     string      `json:"value"`
	Label     string      `json:"label,omitempty"`
	Manage    URL         `json:"manage,omitempty"`
	ExpiresIn int         `json:"expires_in,omitempty"`
	Flags     []TokenFlag `json:"flags,omitempty"`
}

// WithLabel is optional parameter for [NewTokenRequest]
// to request a label for the token.
func WithLabel(label string) tokenRequestOption {
	return func(req *TokenRequest) error {
		req.Label = label
		return nil
	}
}

// WithFlag is optional parameter for [NewTokenRequest]
// to mention the flags associated with the token.
func WithFlag(flag TokenFlag) tokenRequestOption {
	return func(req *TokenRequest) error {
		flags := req.Flags
		if slices.Contains(flags, flag) { // avoid duplicates
			return nil
		}
		req.Flags = append(flags, flag)
		return nil
	}
}

// WithLabelResponse is optional parameter for [NewTokenResponse]
// to provide the label with the access token.
func WithLabelResponse(label string) tokenResponseOption {
	return func(res *TokenResponse) error {
		res.Label = label
		return nil
	}
}

// WithManage is optional parameter for [NewTokenResponse]
// to provide the token management URI.
func WithManage(manage URL) tokenResponseOption {
	return func(res *TokenResponse) error {
		res.Manage = manage
		return nil
	}
}

// WithExpiry is optional parameter for [NewTokenResponse]
// to provide expiry duration for the access token.
func WithExpiry(seconds int) tokenResponseOption {
	return func(res *TokenResponse) error {
		res.ExpiresIn = seconds
		return nil
	}
}

// WithKey is optional parameter for [NewTokenResponse]
// to provide the key to be presented with the access token.
func WithKey(key ClientKey) tokenResponseOption {
	return func(res *TokenResponse) error {
		res.Key = key
		return nil
	}
}

// WithFlags is optional parameter for [NewTokenResponse]
// to mention the flags associated with the access token.
func WithFlags(flags []TokenFlag) tokenResponseOption {
	return func(res *TokenResponse) error {
		res.Flags = flags
		return nil
	}
}
