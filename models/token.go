package models

// TokenFlag is a typed enum for access token flags.
// Flags define attributes or behavior associated with the access token.
type TokenFlag struct {
	name string `json:"-"`
}

// Contents of Access Token Flags Registry.
var (
	TFBearer  = TokenFlag{"bearer"}  // Bearer Access Token (default: bound access token)
	TFDurable = TokenFlag{"durable"} // Access Token durable even after rotation
)

// TokenReq is an object used for describing the requested access rights
// and attributes associated with the access token.
type TokenReq struct {
	// access rights that the client instance is requesting for the access tokens
	// to be used at the RS.
	Access []ARight `json:"access"` // REQUIRED
	// unique name chosen by the client instance to refer to the resulting access token.
	// The value of this field is opaque to the AS. If this field is included in the
	// request, the AS MUST include the same label in the token response.
	// REQUIRED if used as part of a multiple access token request.
	// OPTIONAL otherwise.
	Label string `json:"label,omitempty"`
	// set of flags that indicate desired attributes or behavior to be attached
	// to the access token by the AS.
	Flags []TokenFlag `json:"flags,omitempty"` // OPTIONAL
}

// TokenResponse is an object representing the access token grant by the AS.
// It describes the token attributes and token management info.
type TokenResponse struct {
	// value of the access token as a string. The value MUST be limited to the
	// `token68` character set defined in [RFC9110] to facilitate transmission over
	// HTTP headers and within other protocols without requiring additional encoding.
	//
	// [RFC9110]: https://www.rfc-editor.org/rfc/rfc9110
	Value string `json:"value"` // REQUIRED
	// value of the label the client instance provided in the associated
	// token request, if present.
	// REQUIRED for multiple access tokens or if a label was included in the single
	// access token request.
	// OPTIONAL for a single access token where no label was included in the request.
	Label string `json:"label,omitempty"`
	// management URI for this access token. This URI MUST be an absolute URI.
	Manage URL `json:"manage,omitempty"` // OPTIONAL
	// description of the rights associated with this access token.
	// If included, this MUST reflect the rights associated with the issued access
	// token. These rights MAY vary from what was requested by the client instance.
	Access []ARight `json:"access"` // REQUIRED
	// number of seconds in which the access will expire.
	// The client instance MUST NOT use the access token past this time.
	ExpiresIn int `json:"expires_in,omitempty"` // OPTIONAL
	// key that the token is bound to, if different from the client instance's
	// presented key. The client instance MUST be able to dereference or process the key
	// information in order to be able to sign subsequent requests using the access token.
	// It is RECOMMENDED that keys returned for use with access tokens be key references
	// that the client instance can correlate to its known keys. OPTIONAL.
	Key Key `json:"key,omitempty"` // REQUIRED
	// set of flags that represent attributes or behaviors of the access token
	// issued by the AS.
	Flags []TokenFlag `json:"flags,omitempty"` // OPTIONAL
}

// ContToken defines unique access token for continuing the request,
// called the "continuation access token".
type ContToken struct {
	// value of the access token as a string. Similar to [TokenResponse.Value].
	Value string `json:"value"` // REQUIRED
	// management URI for the continuation token. MUST be absolute URI.
	Manage URL `json:"manage,omitempty"` // OPTIONAL
	// number of seconds in which the continuation token will expire.
	// The client instance MUST NOT use the token past this time.
	ExpiresIn int `json:"expires_in,omitempty"` // OPTIONAL
	// set of flags that represent attributes or behaviors of the
	// continuation token. MUST NOT contain the flag "bearer".
	Flags []TokenFlag `json:"flags,omitempty"` // OPTIONAL
}
