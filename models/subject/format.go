package subject

import (
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

// Format represents a subject identifier format as defined in
// Security Event Identifier Format Registry.
type Format string

// Subject identifier formats defined in the registery.
const (
	Account       Format = "account"
	Email         Format = "email"
	IssuerSubject Format = "iss_sub"
	Opaque        Format = "opaque"
	PhoneNumber   Format = "phone_number"
	DID           Format = "did"
	URI           Format = "uri"
	Aliases       Format = "aliases"
)

// Equal is simple equality comparison for [ID].
func Equal(a, b ID) bool {
	if a.Format != b.Format {
		return false
	}
	fmt := a.Format
	if fmt != Aliases {
		return EqualNoAlias(a.NoAlias(), b.NoAlias())
	}
	return slices.EqualFunc(a.Identifiers, b.Identifiers, EqualNoAlias)
}

// EqualNoAlias is a simple equality comparison for [NoAlias].
func EqualNoAlias(a, b NoAlias) bool {
	if a.Format != b.Format {
		return false
	}
	switch a.Format {
	case Account:
		return a.URI == b.URI
	case Email:
		return a.Email == b.Email
	case IssuerSubject:
		return a.Issuer == b.Issuer && a.Subject == b.Subject
	case Opaque:
		return a.ID == b.ID
	case PhoneNumber:
		return a.Phone == b.Phone
	case DID:
		return a.URL == b.URL
	case URI:
		return a.URI == b.URI
	}
	return false
}

// EmailRegex according to [RFC5322].
var EmailRegex = regexp.MustCompile(`^([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\[[\t -Z^-~]*])$`)

// PhoneRegex according to [E.164].
var PhoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// validator is used to validate subject identifer object according to format.
type validator func(id NoAlias) bool

// formatRegistry is a mapping from format identifiers to corresponding validator.
var formatRegistry = map[Format]validator{
	Account:       validateAccount,
	Email:         validateEmail,
	IssuerSubject: validateIssuerSubject,
	Opaque:        validateOpaque,
	PhoneNumber:   validatePhoneNumber,
	DID:           validateDID,
	URI:           validateURI,
	Aliases:       nil, // to prevent "aliases" format in NoAlias
}

// validateAccount verifies acct URI as defined in [RFC7565].
func validateAccount(id NoAlias) bool {
	if id.URI == "" {
		return false
	}
	acc, err := url.Parse(id.URI)
	if err != nil {
		return false
	}
	return acc.Scheme == "acct"
}

// validateEmail verifies email formatted as defined in [RFC5322].
func validateEmail(id NoAlias) bool {
	return EmailRegex.MatchString(id.Email)
}

// validateIssuerSubject verifies issuer and subject field as per
// "iss" and "sub" fields in JWT [RFC7519].
func validateIssuerSubject(id NoAlias) bool {
	return validStringOrURI(id.Issuer) && validStringOrURI(id.Subject)
}

// validStringOrURI verifies StringOrURI as defined in JWT [RFC7519].
func validStringOrURI(s string) bool {
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

// validateOpaque verifies [ID] formatted as [Opaque] format.
func validateOpaque(id NoAlias) bool {
	return id.ID != ""
}

// validatePhoneNumber verifies Phone number
// with country code.
func validatePhoneNumber(id NoAlias) bool {
	return PhoneRegex.MatchString(id.Phone)
}

// validateDID verifies DID url as defined in [w3-did-core].
func validateDID(id NoAlias) bool {
	if id.URL == "" {
		return false
	}
	did, err := url.Parse(id.URL)
	if err != nil {
		return false
	}
	return did.Scheme == "did"
}

// validateURI verifies URI as defined in [RFC3986].
func validateURI(id NoAlias) bool {
	if id.URI == "" {
		return false
	}
	_, err := url.Parse(id.URI)
	return err == nil
}

// validateAliases verifies subject identifier with aliases.
// If the list of identifiers containes exact duplicates, then
// it is considered as invalid.
func validateAliases(id ID) bool {
	visited := make([]NoAlias, 0, len(id.Identifiers))
	for _, v := range id.Identifiers {
		// exact duplicates
		if slices.ContainsFunc(visited, func(old NoAlias) bool {
			return EqualNoAlias(v, old)
		}) {
			return false
		}
		// invalid NoAlias
		if v.Validate() != nil {
			return false
		}
	}
	return true
}
