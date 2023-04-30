package internal

import (
	"encoding/json"
	"errors"
)

// ErrInvalidFormat is returned when subject id format encountered
// is not defined in the registry or invalid.
var ErrInvalidFormat = errors.New("invalid subject id format")

// ErrInvalidID is returned when the subject identifier object
// is found to be malformed.
var ErrInvalidID = errors.New("invalid subject identifier")

// ID is the Subject Identifier object which is a JSON [RFC8259] object
// whose contents may be used to identify a subject within some context.
type ID struct {
	Format  Format `json:"format"`
	URI     string `json:"uri,omitempty"`
	Email   string `json:"email,omitempty"`
	Issuer  string `json:"iss,omitempty"`
	Subject string `json:"sub,omitempty"`
	ID      string `json:"id,omitempty"`
	Phone   string `json:"phone_number,omitempty"`
	URL     string `json:"url,omitempty"`
}

// Validate checks if the object is valid or not.
func (id ID) Validate() error {
	validator := Validators[id.Format]
	if validator == nil {
		return ErrInvalidFormat
	}
	return validator(id)
}

// UnmarshalJSON implements the [json.UnmarshalJSON.] interface
func (id *ID) UnmarshalJSON(data []byte) error {
	type Alias ID
	var alias Alias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	err = id.Validate()
	if err != nil {
		return err
	}
	*id = ID(alias)
	return nil
}

// NewAccount constructs a subject identifier using an account at a
// service provider, identified with an "acct" URI as defined in [RFC7565].
func NewAccount(uri string) (ID, error) {
	id := ID{Format: FormatAccount, URI: uri}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewEmail constructs a subject identifier using an email address
// defined in [RFC5322].
func NewEmail(email string) (ID, error) {
	id := ID{Format: FormatAccount, Email: email}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIssSub constructs a subject identifier using a pair of "iss" and "sub" members,
// analogous to how subjects are identified using the "iss" and "sub" claims in
// OpenID Connect [OpenID.Core] ID Tokens and JWT [RFC7519].
func NewIssSub(issuer, subject string) (ID, error) {
	id := ID{
		Format:  FormatAccount,
		Issuer:  issuer,
		Subject: subject,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewOpaque constructs a subject identifier with a string with no semantics asserted
// beyond its usage as an identifier for the subject, such as a UUID or hash.
func NewOpaque(opaque string) (ID, error) {
	id := ID{Format: FormatAccount, ID: opaque}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewPhone constructs a subject identifier using a telephone number,
// including international dialing prefix, formatted according to [E.164].
func NewPhone(phone string) (ID, error) {
	id := ID{Format: FormatAccount, Phone: phone}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewDID constructs a subject identifier using a Decentralized Identifier
// (DID) URL as defined in [DID].
func NewDID(url string) (ID, error) {
	id := ID{Format: FormatAccount, URL: url}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewURI constructs a subject identifier using a URI as defined in [RFC3986].
func NewURI(uri string) (ID, error) {
	id := ID{Format: FormatAccount, URI: uri}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}
