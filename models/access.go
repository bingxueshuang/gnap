package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidAccessRight is returned when an access right object
// is found to be malformed.
var ErrInvalidAccessRight = errors.New("invalid access right")

// ErrInvalidTokenRequest is returned when a token request is
// found to be malformed.
var ErrInvalidTokenRequest = errors.New("invalid token request")

// ErrInvalidTokenResponse is returned when a token response
// is found to be malformed.
var ErrInvalidTokenResponse = errors.New("invalid token response")

// AccessRight represents the rights and privileges requested
// or granted during a gnap request flow.
type AccessRight struct {
	Type       string   `json:"type"`
	Actions    []string `json:"actions,omitempty"`
	Locations  []string `json:"locations,omitempty"`
	Datatypes  []string `json:"datatypes,omitempty"`
	Identifier string   `json:"identifier,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	Ref        string   `json:"-"`
}

// MarshalJSON implements the [json.Marshaler] interface.
func (r AccessRight) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return json.Marshal(r.Ref)
	}
	type Alias AccessRight
	return json.Marshal(Alias(r))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (r *AccessRight) UnmarshalJSON(data []byte) error {
	var ref string
	err := json.Unmarshal(data, &ref)
	if err == nil { // by reference
		*r = AccessRight{Ref: ref}
		return nil
	}
	type Alias AccessRight
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err == nil {
		*r = AccessRight(alias)
		return nil
	}
	return ErrInvalidAccessRight
}

// ATRequest is a wrapper aroung TokenRequest for
// managing single and multiple access token requests.
type ATRequest struct {
	Single   TokenRequest
	Multiple []TokenRequest
}

// ATResponse is a wrapper around TokenResponse for
// managing single and multiple access token responses.
type ATResponse struct {
	Single   TokenResponse
	Multiple []TokenResponse
}

// MarshalJSON implements the [json.Marshaler] interface.
func (req ATRequest) MarshalJSON() ([]byte, error) {
	if req.Multiple == nil {
		return json.Marshal(req.Single)
	}
	return json.Marshal(req.Multiple)
}

// MarshalJSON implements the [json.Marshaler] interface.
func (req ATResponse) MarshalJSON() ([]byte, error) {
	if req.Multiple == nil {
		return json.Marshal(req.Single)
	}
	return json.Marshal(req.Multiple)
}

// UnmarshalJSON implements [json.UnmarshalJSON] interface.
func (req *ATRequest) UnmarshalJSON(data []byte) error {
	var one TokenRequest
	err := json.Unmarshal(data, &one)
	if err == nil { // valid single access token request
		req.Single = one
		return nil
	}
	var many []TokenRequest
	err = json.Unmarshal(data, &many)
	if err == nil { // valid multiple access token request
		req.Multiple = many
		return nil
	}
	return ErrInvalidTokenRequest
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (req *ATResponse) UnmarshalJSON(data []byte) error {
	var one TokenResponse
	err := json.Unmarshal(data, &one)
	if err == nil { // valid single access token response
		req.Single = one
		return nil
	}
	var many []TokenResponse
	err = json.Unmarshal(data, &many)
	if err == nil { // valid multiple access token response
		req.Multiple = many
		return nil
	}
	return ErrInvalidTokenResponse
}
