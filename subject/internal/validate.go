package internal

import (
	"net/url"
	"regexp"
	"strings"
)

// EmailRegex according to [RFC5322].
var EmailRegex = regexp.MustCompile(`^([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\[[\t -Z^-~]*])$`)

// PhoneRegex according to [E.164].
var PhoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// Validators is a quick lookup for validate function for every
// defined [Format].
var Validators = map[Format]func(ID) error{
	FormatAccount: validateAccount,
	FormatEmail:   validateEmail,
	FormatIssSub:  validateIssSub,
	FormatOpaque:  validateOpaque,
	FormatPhone:   validatePhone,
	FormatDID:     validateDID,
	FormatURI:     validateURI,
}

// validateAccount verifies acct URI as defined in [RFC7565].
func validateAccount(id ID) error {
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

// validateEmail verifies email formatted as defined in [RFC5322].
func validateEmail(id ID) error {
	ok := EmailRegex.MatchString(id.Email)
	if !ok {
		return Wrap("invalid email", ErrInvalidID)
	}
	return nil
}

// validateIssSub verifies issuer and subject field as per
// "iss" and "sub" fields in JWT [RFC7519].
func validateIssSub(id ID) error {
	ok := stringOrURI(id.Issuer) && stringOrURI(id.Subject)
	if !ok {
		return Wrap("invalid issuer-subject", ErrInvalidID)
	}
	return nil
}

// stringOrURI verifies StringOrURI as defined in JWT [RFC7519].
func stringOrURI(s string) bool {
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

// validateOpaque verifies ID formatted as [Opaque] format.
func validateOpaque(id ID) error {
	if id.ID == "" {
		return Wrap("empty opaque id", ErrInvalidID)
	}
	return nil
}

// validatePhone verifies Phone number using regexp
// with country code.
func validatePhone(id ID) error {
	if !PhoneRegex.MatchString(id.Phone) {
		return Wrap("invalid phone number", ErrInvalidID)
	}
	return nil
}

// validateDID verifies DID url as defined in [w3-did-core].
func validateDID(id ID) error {
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

// validateURI verifies URI as defined in [RFC3986].
func validateURI(id ID) error {
	if id.URI == "" {
		return Wrap("empty URI", ErrInvalidID)
	}
	_, err := url.Parse(id.URI)
	if err != nil {
		return Wrap("invalid URI", ErrInvalidID)
	}
	return nil
}
