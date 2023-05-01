package internal

import (
	"net/url"
	"regexp"
	"strings"
)

// EmailRegex according to [RFC5322].
//
// [RFC5322]: https://www.rfc-editor.org/info/rfc5322
var EmailRegex = regexp.MustCompile(`^([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\[[\t -Z^-~]*])$`)

// PhoneRegex according to [E.164].
//
// [E.164]: https://www.itu.int/rec/T-REC-E.164-201011-I/en
var PhoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// ID is the mirror of [subject.ID] to avoid circular dependencies.
type ID struct {
	URI     string
	Email   string
	Issuer  string
	Subject string
	ID      string
	Phone   string
	URL     string
}

// ValidateAccount verifies acct URI as defined in [RFC7565].
//
// [RFC7565]: https://www.rfc-editor.org/info/rfc7565
func ValidateAccount(id ID) error {
	if id.URI == "" {
		return Wrap("empty uri", ErrInvalidID)
	}
	u, err := url.Parse(id.URI)
	if err != nil {
		return Wrap("invalid uri", ErrInvalidID, err)
	}
	if u.Scheme != "acct" {
		return Wrap("acct scheme expected", ErrInvalidID)
	}
	return nil
}

// ValidateEmail verifies email formatted as defined in [RFC5322].
//
// [RFC5322]: https://www.rfc-editor.org/info/rfc5322
func ValidateEmail(id ID) error {
	ok := EmailRegex.MatchString(id.Email)
	if !ok {
		return Wrap("invalid email", ErrInvalidID)
	}
	return nil
}

// ValidateIssSub verifies issuer and subject field as per
// "iss" and "sub" fields in JWT [RFC7519].
//
// [RFC7519]: https://www.rfc-editor.org/info/rfc7519
func ValidateIssSub(id ID) error {
	ok := StringOrURI(id.Issuer) && StringOrURI(id.Subject)
	if !ok {
		return Wrap("invalid issuer-subject", ErrInvalidID)
	}
	return nil
}

// StringOrURI verifies StringOrURI as defined in JWT [RFC7519].
//
// [RFC7519]: https://www.rfc-editor.org/info/rfc7519
func StringOrURI(s string) bool {
	if s == "" {
		return false
	}
	// if ":" then url
	if strings.Contains(s, ":") {
		_, err := url.Parse(s)
		return err == nil
	}
	// else string
	return true
}

// ValidateOpaque verifies ID formatted as [Opaque] format.
func ValidateOpaque(id ID) error {
	if id.ID == "" {
		return Wrap("empty opaque id", ErrInvalidID)
	}
	return nil
}

// ValidatePhone verifies Phone number using regexp
// with country code.
func ValidatePhone(id ID) error {
	if !PhoneRegex.MatchString(id.Phone) {
		return Wrap("invalid phone number", ErrInvalidID)
	}
	return nil
}

// ValidateDID verifies DID url as defined in [w3-did-core].
//
// [w3-did-core]: https://www.w3.org/TR/did-core/
func ValidateDID(id ID) error {
	if id.URL == "" {
		return Wrap("empty DID URL", ErrInvalidID)
	}
	u, err := url.Parse(id.URL)
	if err != nil {
		return Wrap("invalid DID URL", ErrInvalidID, err)
	}
	if u.Scheme != "did" {
		return Wrap("did scheme expected", ErrInvalidID)
	}
	return nil
}

// ValidateURI verifies URI as defined in [RFC3986].
//
// [RFC3986]: https://www.rfc-editor.org/info/rfc3986
func ValidateURI(id ID) error {
	if id.URI == "" {
		return Wrap("empty URI", ErrInvalidID)
	}
	_, err := url.Parse(id.URI)
	if err != nil {
		return Wrap("invalid URI", ErrInvalidID)
	}
	return nil
}
