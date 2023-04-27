package subject

import (
	"encoding/json"
	"errors"
)

// Errors during creating subject id object.
var (
	ErrInvalidFormat    = errors.New("format not defined in the registry")
	ErrInvalidSubjectID = errors.New("malformed subject identifier")
	ErrInvalidAliases   = errors.New("malformed subject id aliases")
)

// NewIDAccount creates a new subject identifier of format [Account].
func NewIDAccount(acc string) (ID, error) {
	id := ID{
		Format: Account,
		URI:    acc,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDEmail creates a new subject identifier of format [Email].
func NewIDEmail(email string) (ID, error) {
	id := ID{
		Format: Email,
		Email:  email,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDIssuerSubject creates a new subject identifier of format [IssuerSubject].
func NewIDIssuerSubject(iss string, sub string) (ID, error) {
	id := ID{
		Format:  IssuerSubject,
		Issuer:  iss,
		Subject: sub,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDOpaque creates a new subject identifier of format [Opaque].
func NewIDOpaque(opaque string) (ID, error) {
	id := ID{
		Format: Opaque,
		ID:     opaque,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDPhoneNumber creates a new subject identifier of format [PhoneNumber].
func NewIDPhoneNumber(phone string) (ID, error) {
	id := ID{
		Format: PhoneNumber,
		Phone:  phone,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDdid creates a new subject identifier of format [DID].
func NewIDdid(did string) (ID, error) {
	id := ID{
		Format: DID,
		URL:    did,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDuri creates a new subject identifier of format [URI].
func NewIDuri(uri string) (ID, error) {
	id := ID{
		Format: URI,
		URI:    uri,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NewIDAliases creates a new subject identifier of format [Aliases].
func NewIDAliases(aliases []NoAlias) (ID, error) {
	id := ID{
		Format:      Aliases,
		Identifiers: aliases,
	}
	err := id.Validate()
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

// NoAlias is subject identifier with any valid format except
// aliases (to prevent nesting).
type NoAlias struct {
	Format  Format `json:"format"`
	URI     string `json:"uri,omitempty"`
	Email   string `json:"email,omitempty"`
	Issuer  string `json:"iss,omitempty"`
	Subject string `json:"sub,omitempty"`
	ID      string `json:"id,omitempty"`
	Phone   string `json:"phone_number,omitempty"`
	URL     string `json:"url,omitempty"`
}

// ID is IETF Subject Identifier for Security Events.
type ID struct {
	Format      Format    `json:"format"`
	URI         string    `json:"uri,omitempty"`
	Email       string    `json:"email,omitempty"`
	Issuer      string    `json:"iss,omitempty"`
	Subject     string    `json:"sub,omitempty"`
	ID          string    `json:"id,omitempty"`
	Phone       string    `json:"phone_number,omitempty"`
	URL         string    `json:"url,omitempty"`
	Identifiers []NoAlias `json:"identifiers,omitempty"`
}

// Validate checks if all fields of id are valid.
func (id NoAlias) Validate() error {
	validate := formatRegistry[id.Format]
	if validate == nil {
		return ErrInvalidFormat
	}
	if validate(id) {
		return nil
	}
	return ErrInvalidSubjectID
}

// Validate checks if all fields of id are valid.
func (id ID) Validate() error {
	v, ok := formatRegistry[id.Format]
	if !ok {
		return ErrInvalidFormat
	}
	if v != nil {
		return id.NoAlias().Validate()
	}
	ok = validateAliases(id)
	if !ok {
		return ErrInvalidSubjectID
	}
	return nil
}

// NoAlias is a helper to convert [ID] to [NoAlias].
func (id ID) NoAlias() (n NoAlias) {
	if id.Format == Aliases {
		return
	}
	return NoAlias{
		Format:  id.Format,
		URI:     id.URI,
		Email:   id.Email,
		Issuer:  id.Issuer,
		Subject: id.Subject,
		ID:      id.ID,
		Phone:   id.Phone,
		URL:     id.URL,
	}
}

// SubjectID is a helper to convert [NoAlias] to [ID].
func (id NoAlias) SubjectID() (s ID) {
	return ID{
		Format:  id.Format,
		URI:     id.URI,
		Email:   id.Email,
		Issuer:  id.Issuer,
		Subject: id.Subject,
		ID:      id.ID,
		Phone:   id.Phone,
		URL:     id.URL,
	}
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *ID) UnmarshalJSON(data []byte) error {
	type Alias ID
	var alias Alias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	err = ID(alias).Validate()
	if err != nil {
		return err
	}
	*id = ID(alias)
	return nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *NoAlias) UnmarshalJSON(data []byte) error {
	type Alias NoAlias
	var alias Alias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	err = NoAlias(alias).Validate()
	if err != nil {
		return err
	}
	*id = NoAlias(alias)
	return nil
}
