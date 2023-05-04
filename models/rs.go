package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidTokenFormat is returned when a token format not
// defined in the registry is encountered.
var ErrInvalidTokenFormat = errors.New("invalid token format")

// TokenFormat represents a type-safe enum for defined token format types.
type TokenFormat struct {
	name string `json:"-"`
}

// Contents of Token Formats Registry
var (
	TFjwtSigned    = TokenFormat{"jwt-signed"}
	TFjwtEncrypted = TokenFormat{"jwt-encrypted"}
	TFMacaroon     = TokenFormat{"macaroon"}
	TFBiscuit      = TokenFormat{"biscuit"}
	TFzcap         = TokenFormat{"zcap"}
)

// tokenFormatRegistry is a quick mapping from strings
// to valid token formats.
var tokenFormatRegistry = map[string]TokenFormat{
	"jwt-signed":    TFjwtSigned,
	"jwt-encrypted": TFjwtEncrypted,
	"macaroon":      TFMacaroon,
	"biscuit":       TFBiscuit,
	"zcap":          TFzcap,
}

type RSDiscovery struct {
	// URL of the endpoint offering introspection. The location MUST be a URL [RFC3986]
	// with a scheme component that MUST be https, a host component, and optionally,
	// port, path and query components and no fragment components.
	//
	// [RFC3986]: https://www.rfc-editor.org/rfc/rfc3986
	Introspection URL `json:"introspection_endpoint,omitempty"` // OPTIONAL
	// list of token formats supported by this AS.
	TokenFormats []TokenFormat `json:"token_formats,omitempty"`
	// URL of the endpoint offering resource registration. The location MUST be a
	// URL [RFC3986] with a scheme component that MUST be https, a host component,
	// and optionally, port, path and query components and no fragment components.
	//
	// [RFC3986]: https://www.rfc-editor.org/rfc/rfc3986
	Registration URL `json:"registration,omitempty"` // OPTIONAL
	// location of the AS's grant request endpoint, used by the RS to derive
	// downstream access tokens. The location MUST be a URL [RFC3986] with a scheme
	// component that MUST be https, a host component, and optionally, port, path
	// and query components and no fragment components. This URL MUST be the same URL
	// used by client instances in support of GNAP requests.
	//
	// [RFC3986]: https://www.rfc-editor.org/rfc/rfc3986
	GrantRequest URL `json:"grant_request_endpoint"` // REQUIRED
	// list of the AS's supported key proofing mechanisms. The values of this list
	// correspond to possible values of the proof field of the "key" of the request.
	KeyProofs []ProofMethod `json:"key_proofs_supported,omitempty"` // OPTIONAL
}

// ResourceServer is the form of the resource server identifying itself using a key
// field or by passing an instance identifier directly.
type ResourceServer struct {
	// public key in a format similar to [Client]
	Key Key `json:"key"` // REQUIRED
	// RS MAY present its keys by reference by passing an instance identifier directly.
	Ref string `json:"-"`
}

// IOSReq represents token introspection request object to validate the access token
// and determine if the token is "active".
type IOSReq struct {
	// access token value presented to the RS by the client instance.
	Token string `json:"token"` // REQUIRED
	// proofing method used by the client instance to bind token to the RS request.
	Proof ProofMethod `json:"proof"` // RECOMMENDED
	// identification used to authenticate the resource server making this call,
	// either by value or by reference
	Server ResourceServer `json:"resource_server"` // REQUIRED
	// minimum access rights required to fulfill the request.
	Access []ARight `json:"access,omitempty"` // OPTIONAL
}

type IOSResponse struct {
	// if true, the access token presented is active. If any of the criteria for an
	// active token are not true, or if the AS is unable to make a determination
	// (such as the token is not found), the value is set to false and
	// other fields are omitted.
	Active bool `json:"active"` // REQUIRED
	// access rights associated with this access token. This array MAY be filtered or
	// otherwise limited for consumption by the identified RS, including being
	// an empty array.
	Access []ARight `json:"access,omitempty"` // REQUIRED
	// key bound to the access token, to allow the RS to validate the signature
	// of the request from the client instance.
	//
	// If the access token is a bearer token, this MUST NOT be included.
	// REQUIRED if the token is bound.
	Key Key `json:"key,omitempty"`
	// set of flags associated with the access token.
	Flags []TokenFlag `json:"flags,omitempty"` // OPTIONAL
	// timestamp after which this token is no longer valid. Expressed as a integer
	// seconds from UNIX Epoch.
	Expiry int64 `json:"exp,omitempty"` // OPTIONAL
	// timestamp at which this token was issued by the AS. Expressed as a integer
	// seconds from UNIX Epoch.
	IssuedAt int64 `json:"iat,omitempty"` // OPTIONAL
	// timestamp before which this token is not valid. Expressed as a integer
	// seconds from UNIX Epoch.
	NotBefore int64 `json:"nbf,omitempty"` // OPTIONAL
	// identifiers for the resource servers this token can be accepted at.
	Audience Audience `json:"aud,omitempty"` // OPTIONAL
	// identifier of the resource owner who authorized this token.
	Subject string `json:"sub,omitempty"` // OPTIONAL
	// grant endpoint URL of the AS that issued this token.
	Issuer URL `json:"iss,omitempty"` // OPTIONAL
	// instance identifier of the client instance that the token was issued to.
	InstanceID string `json:"instance_id,omitempty"` // OPTIONAL
}

// RegReq represents the object format that the Resource Server will use to
// post a set of resources to the AS's resource registration endpoint along
// with information about what the RS will need to validate the request.
type RegReq struct {
	// list of access rights associated with the request
	Access []ARight `json:"access"` // REQUIRED
	// identification used to authenticate the resource server making this call,
	// either by value or by reference
	Server ResourceServer `json:"resource_server"` // REQUIRED
	// token format required to access the identified resource. If the field is omitted,
	// the token format is at the discretion of the AS. If the AS does not support the
	// requested token format, the AS MUST return an error to the RS.
	Format TokenFormat `json:"token_format_required"` // OPTIONAL
	// if present and set to true, the RS expects to make a token introspection request.
	// If absent or set to false, the RS does not anticipate needing to make an
	// introspection request for tokens relating to this resource set.
	IOS bool `json:"token_introspection_required"` // OPTIONAL
}

// RegResponse represents the object format with which the AS responds with a reference
// appropriate to represent the resources list that the RS presented in its request as
// well as any additional information the RS might need in future requests.
type RegResponse struct {
	// single string representing the list of resources registered in the request.
	// The RS MAY make this handle available to a client instance as part of a discovery
	// response as described in [GNAP] or as documentation to client software developers.
	ResourceRef string `json:"resource_reference"` // REQUIRED
	// instance identifier that the RS can use to refer to itself in future calls
	// to the AS, in lieu of sending its key by value.
	InstanceID string `json:"instance_id,omitempty"` // OPTIONAL
	// introspection endpoint of this AS, used to allow RS to perform token introspection.
	Endpoint URL `json:"endpoint,omitempty"` // OPTIONAL
}

// Audience represents the intended receivers of the access token. The access token
// is intended for use at one or more RS's. The AS can identify those RS's to allow
// each RS to ensure that the token is not receiving a token intended for someone else.
// The AS and RS have to agree on the nature of any audience identifiers represented
// by the token, but the URIs of the RS are a common pattern. It is either a string or
// an array of strings.
type Audience struct {
	Single string   `json:"-"`
	Many   []string `json:"-"`
}

// MarshalJSON implements the [json.Marshaler] interface.
func (tf TokenFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(tf.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (tf *TokenFormat) UnmarshalJSON(data []byte) error {
	var format string
	err := json.Unmarshal(data, &format)
	if err != nil {
		return err
	}
	f, ok := tokenFormatRegistry[format]
	if ok {
		*tf = f
		return nil
	}
	return ErrInvalidTokenFormat
}

// MarshalJSON implements the [json.Marshaler] interface.
func (rs ResourceServer) MarshalJSON() ([]byte, error) {
	if rs.Ref != "" {
		return json.Marshal(rs.Ref)
	}
	type Alias ResourceServer
	return json.Marshal(Alias(rs))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (rs *ResourceServer) UnmarshalJSON(data []byte) (err error) {
	var ref string
	err = json.Unmarshal(data, &ref)
	if err == nil {
		rs.Ref = ref
		return
	}
	type Alias ResourceServer
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err == nil {
		*rs = ResourceServer(alias)
		return
	}
	return
}

// MarshalJSON implements the [json.Marshaler] interface.
func (aud Audience) MarshalJSON() ([]byte, error) {
	if aud.Many == nil {
		return json.Marshal(aud.Single)
	}
	return json.Marshal(aud.Many)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (aud *Audience) UnmarshalJSON(data []byte) (err error) {
	var one string
	err = json.Unmarshal(data, &one)
	if err == nil {
		aud.Single = one
		return
	}
	var many []string
	err = json.Unmarshal(data, &many)
	if err == nil {
		aud.Many = many
		return
	}
	return
}
