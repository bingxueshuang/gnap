package subject

import (
	"encoding/json"

	"github.com/bingxueshuang/gnap/subject/internal"
)

// Errors during creating subject id object.
var (
	ErrInvalidFormat = internal.ErrInvalidFormat
	ErrInvalidID     = internal.ErrInvalidID
)

// ID is the Subject Identifier object which is a JSON [RFC8259] object
// whose contents may be used to identify a subject within some context.
//
// [RFC8259]: https://www.rfc-editor.org/info/rfc8259
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

// validators is a quick lookup for validate function for every
// defined [Format].
var validators = map[Format]func(internal.ID) error{
	Account:     internal.ValidateAccount,
	Email:       internal.ValidateEmail,
	IssSub:      internal.ValidateIssSub,
	Opaque:      internal.ValidateOpaque,
	PhoneNumber: internal.ValidatePhone,
	DID:         internal.ValidateDID,
	URI:         internal.ValidateURI,
}

// Validate checks if the subject id object is valid or not.
func (id ID) Validate() error {
	validator := validators[id.Format]
	if validator == nil {
		return ErrInvalidFormat
	}
	iid := internal.ID{
		URI:     id.URI,
		Email:   id.Email,
		Issuer:  id.Issuer,
		Subject: id.Subject,
		ID:      id.ID,
		Phone:   id.Phone,
		URL:     id.URL,
	}
	return validator(iid)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
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

// NewIDAccount constructs a subject identifier using an account at a
// service provider, identified with an "acct" URI as defined in [RFC7565].
//
// [RFC7565]: https://www.rfc-editor.org/info/rfc7565
func NewIDAccount(uri string) (ID, error) {
	id := ID{Format: Account, URI: uri}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDEmail constructs a subject identifier using an email address
// defined in [RFC5322].
//
// [RFC5322]: https://www.rfc-editor.org/info/rfc5322
func NewIDEmail(email string) (ID, error) {
	id := ID{Format: Email, Email: email}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDIssSub constructs a subject identifier using a pair of "iss" and "sub" members,
// analogous to how subjects are identified using the "iss" and "sub" claims in
// OpenID Connect [OpenID.Core] ID Tokens and JWT [RFC7519].
//
// [OpenID.Core]: https://openid.net/specs/openid-connect-core-1_0.html
// [RFC7519]: https://www.rfc-editor.org/info/rfc7519
func NewIDIssSub(issuer, subject string) (ID, error) {
	id := ID{
		Format:  IssSub,
		Issuer:  issuer,
		Subject: subject,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDOpaque constructs a subject identifier with a string with no semantics asserted
// beyond its usage as an identifier for the subject, such as a UUID or hash.
func NewIDOpaque(opaque string) (ID, error) {
	id := ID{Format: Opaque, ID: opaque}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDPhone constructs a subject identifier using a telephone number,
// including international dialing prefix, formatted according to [E.164].
//
// [E.164]: https://www.itu.int/rec/T-REC-E.164-201011-I/en
func NewIDPhone(phone string) (ID, error) {
	id := ID{Format: PhoneNumber, Phone: phone}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDdid constructs a subject identifier using a Decentralized Identifier
// (DID) URL as defined in [W3DID].
//
// [W3DID]: https://www.w3.org/TR/did-core/
func NewIDdid(url string) (ID, error) {
	id := ID{Format: DID, URL: url}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDuri constructs a subject identifier using a URI as defined in [RFC3986].
//
// [RFC3986]: https://www.rfc-editor.org/info/rfc3986
func NewIDuri(uri string) (ID, error) {
	id := ID{Format: URI, URI: uri}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}
