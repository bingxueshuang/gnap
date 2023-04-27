package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/bingxueshuang/gnap/subject"
)

// ErrInvalidAFormat is returned when an assertion format not defined
// in the registry is encountered.
var ErrInvalidAFormat = errors.New("invalid assertion format")

// AssertionFormat is a valid assertion format as defined in the draft.
type AssertionFormat string

// Registry of valid assertion formats.
const (
	AFidToken AssertionFormat = "id_token"
	AFsaml2   AssertionFormat = "saml2"
)

// assertionFormatsRegistry provides quick lookup for valid or invalid check.
var assertionFormatsRegistry = map[string]AssertionFormat{
	"id_token": AFidToken,
	"saml2":    AFsaml2,
}

// MarshalJSON implements [json.Marshaler] interface.
func (af AssertionFormat) MarshalJSON() ([]byte, error) {
	_, ok := assertionFormatsRegistry[string(af)]
	if !ok {
		return nil, ErrInvalidAFormat
	}
	return json.Marshal(string(af))
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (af *AssertionFormat) UnmarshalJSON(data []byte) error {
	var format string
	err := json.Unmarshal(data, &format)
	if err != nil {
		return err
	}
	_, ok := assertionFormatsRegistry[format]
	if ok {
		*af = AssertionFormat(format)
		return nil
	}
	return ErrInvalidAFormat
}

// SubRequest describes the information about the RO that the client instance
// is requesting to be returned directly in the response from the AS.
type SubRequest struct {
	SFormats []subject.Format  `json:"sub_id_formats,omitempty"`
	AFormats []AssertionFormat `json:"assertion_formats,omitempty"`
	SubIDs   []subject.ID      `json:"sub_ids,omitempty"`
}

// SubResponse contains claims about the RO as known and declared by the AS.
type SubResponse struct {
	SubIDs     []subject.ID `json:"sub_ids,omitempty"`
	Assertions []Assertion  `json:"assertions,omitempty"`
	UpdatedAt  time.Time    `json:"updated_at,omitempty"`
}

// EndUser identifies the end user to the AS in a manner that the AS can verify
// (by value or by reference), either directly or by interacting with the
// end user to determine their status as the RO.
type EndUser struct {
	SubIDs     []subject.ID `json:"sub_ids,omitempty"`
	Assertions []Assertion  `json:"assertions,omitempty"`
	Ref        string
}

// MarshalJSON implements the [json.Marshaler] interface.
func (u EndUser) MarshalJSON() ([]byte, error) {
	if u.Ref != "" {
		return json.Marshal(u.Ref)
	}
	type Alias EndUser
	return json.Marshal(Alias(u))
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (u *EndUser) UnmarshalJSON(data []byte) error {
	var ref string
	err := json.Unmarshal(data, &ref)
	if err == nil { // valid string
		u.Ref = ref
		return nil
	}
	type Alias EndUser
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	*u = EndUser(alias)
	return nil
}

// Assertion represents identity assertions used to convey subject information.
type Assertion struct {
	Value  string          `json:"value"`
	Format AssertionFormat `json:"format"`
}
