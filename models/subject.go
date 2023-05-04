package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/bingxueshuang/gnap/subject"
)

// ErrInvalidAFormat is returned when an assertion format
// not defined in the registry is encountered.
var ErrInvalidAFormat = errors.New("invalid assertion format")

// AFormat is type-safe enum for gnap assertion formats.
// It defines a means to pass identity assertions between the
// AS and the client instance.
type AFormat struct {
	name string `json:"-"`
}

// Contents of the Assertion Formats Registry.
var (
	AFidToken = AFormat{"id_token"}
	AFsaml2   = AFormat{"saml2"}
)

// aFormatRegistry is a quick mapping from strings to valid AFormat values.
var aFormatRegistry = map[string]AFormat{
	"id_token": AFidToken,
	"saml2":    AFsaml2,
}

// SubReq defines the subject field of [Request] as a JSON object if the
// client instance is requesting information about the RO from AS.
type SubReq struct {
	// array of subject identifier subject formats requested for the RO.
	// REQUIRED if subject identifiers are requested.
	SFormats []subject.Format `json:"sub_id_formats,omitempty"`
	// array of requested assertion formats.
	// REQUIRED if assertions are requested.
	AFormats []AFormat `json:"assertion_formats,omitempty"`
	// array of subject identifiers representing the subject that info is being
	// requested for. All identifiers in the array MUST identify the same subject.
	SubIDs []subject.ID `json:"sub_ids,omitempty"` // OPTIONAL
	// a boolean field to differentiate between zero value and set value.
	// MUST be set to true if SubReq is not zero value.
	NonZero bool `json:"-"`
}

// SubResponse represents the "subject" field of [Response] object. It is returned
// by the AS if info about the RO is requested and the AS grants the client instance
// access to that data.
type SubResponse struct {
	// array of subject identifiers for the RO.
	// REQUIRED if returning subject identifiers.
	SubIDs []subject.ID `json:"sub_ids,omitempty"`
	// array containing identity assertions as [Assertion] objects.
	// REQUIRED if returning assertions.
	Assertions []Assertion `json:"assertions,omitempty"`
	// timestamp as an [RFC3339] date string, indicating when the identified
	// account was last updated.
	//
	// [RFC3339]: https://www.rfc-editor.org/rfc/rfc3339
	UpdatedAt time.Time `json:"updated_at"` // RECOMMENDED
}

// Assertion defines a JSON object for representing an identity assertion
// with respect to a subject.
type Assertion struct {
	// assertion format (such as OPEN ID Connect ID Token or SAML2).
	Format AFormat `json:"format"` // REQUIRED
	// assertion value as the JSON string serialization of the assertion.
	Value string `json:"value"` // REQUIRED
}

// MarshalJSON implements the [json.Marshaler] interface.
func (af AFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(af.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (af *AFormat) UnmarshalJSON(data []byte) error {
	var format string
	err := json.Unmarshal(data, &format)
	if err != nil {
		return err
	}
	f, ok := aFormatRegistry[format]
	if ok {
		*af = f
		return nil
	}
	return ErrInvalidAFormat
}
