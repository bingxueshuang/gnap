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
	id := NoAlias{
		Format: Account,
		URI:    acc,
	}
	ok := validateAccount(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: Account, NoAlias: id}, nil
}

// NewIDEmail creates a new subject identifier of format [Email].
func NewIDEmail(email string) (ID, error) {
	id := NoAlias{
		Format: Email,
		Email:  email,
	}
	ok := validateEmail(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: Email, NoAlias: id}, nil
}

// NewIDIssuerSubject creates a new subject identifier of format [IssuerSubject].
func NewIDIssuerSubject(iss string, sub string) (ID, error) {
	id := NoAlias{
		Format:  IssuerSubject,
		Issuer:  iss,
		Subject: sub,
	}
	ok := validateIssuerSubject(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: IssuerSubject, NoAlias: id}, nil
}

// NewIDOpaque creates a new subject identifier of format [Opaque].
func NewIDOpaque(opaque string) (ID, error) {
	id := NoAlias{
		Format: Opaque,
		ID:     opaque,
	}
	ok := validateOpaque(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: Opaque, NoAlias: id}, nil
}

// NewIDPhoneNumber creates a new subject identifier of format [PhoneNumber].
func NewIDPhoneNumber(phone string) (ID, error) {
	id := NoAlias{
		Format: PhoneNumber,
		Phone:  phone,
	}
	ok := validatePhoneNumber(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: PhoneNumber, NoAlias: id}, nil
}

// NewIDdid creates a new subject identifier of format [DID].
func NewIDdid(did string) (ID, error) {
	id := NoAlias{
		Format: DID,
		URL:    did,
	}
	ok := validateDID(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: DID, NoAlias: id}, nil
}

// NewIDuri creates a new subject identifier of format [URI].
func NewIDuri(uri string) (ID, error) {
	id := NoAlias{
		Format: URI,
		URI:    uri,
	}
	ok := validateURI(id)
	if !ok {
		return ID{}, ErrInvalidSubjectID
	}
	return ID{Format: URI, NoAlias: id}, nil
}

// NewIDAliases creates a new subject identifier of format [Aliases].
func NewIDAliases(aliases []NoAlias) (ID, error) {
	id := ID{
		Format:      Aliases,
		Identifiers: aliases,
	}
	ok := validateAliases(id)
	if !ok {
		return ID{}, ErrInvalidAliases
	}
	return id, nil
}

// NoAlias is subject identifier with any valid format except
// aliases (to prevent nesting).
type NoAlias struct {
	Format  Format `json:"format"`
	URI     string `json:"uri,omitempty"`
	Email   string `json:"email,omitempty"`
	Issuer  string `json:"issuer,omitempty"`
	Subject string `json:"subject,omitempty"`
	ID      string `json:"id,omitempty"`
	Phone   string `json:"phone_number,omitempty"`
	URL     string `json:"url,omitempty"`
}

// ID is IETF Subject Identifier for Security Events.
type ID struct {
	Format      Format    `json:"format"`
	Identifiers []NoAlias `json:"identifiers,omitempty"`
	NoAlias
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *ID) UnmarshalJSON(data []byte) error {
	type Alias ID
	var alias Alias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	ok := validateAliases(ID(alias))
	if !ok {
		return ErrInvalidAliases
	}
	*id = ID(alias)
	id.NoAlias.Format = id.Format
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
	validator := formatRegistry[alias.Format]
	if validator == nil {
		return ErrInvalidFormat
	}
	ok := validator(NoAlias(alias))
	if !ok {
		return ErrInvalidSubjectID
	}
	*id = NoAlias(alias)
	return nil
}
