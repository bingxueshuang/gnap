package models

import (
	"encoding/json"
	"errors"

	"github.com/yaronf/httpsign"
)

// ErrInvalidSigAlg is returned when a http signature algorithm
// not defined in the registry is encountered.
var ErrInvalidSigAlg = errors.New("invalid http signature algorithm")

// ErrInvalidDigestAlg is returned when a content digest algorithm
// not defined in the registry is encountered.
var ErrInvalidDigestAlg = errors.New("invalid content digest algorithm")

// ErrInvalidProofMethod is returned when a key proof method not
// defined in the registry is encountered.
var ErrInvalidProofMethod = errors.New("invalid proof method")

// ErrInvalidKeyFormat is returned when a key format not defined
// in the registry is encountered.
var ErrInvalidKeyFormat = errors.New("invalid key format")

// HTTPSigAlg represents http signature algorithm.
type HTTPSigAlg string

// Registry of http signature algorithms.
const (
	RSA_PSS_SHA512    HTTPSigAlg = "rsa-pss-sha512"
	RSA_SHA256        HTTPSigAlg = "rsa-v1_5-sha256"
	HMAC_SHA256       HTTPSigAlg = "hmac-sha256"
	ECDSA_P256_SHA256 HTTPSigAlg = "ecdsa-p256-sha256"
	ECDSA_P384_SHA384 HTTPSigAlg = "ecdsa-p384-sha384"
	ED25519           HTTPSigAlg = "ed25519"
)

// sigAlgRegistry is a quick mapping from string values
// to valid http signature algorithms.
var sigAlgRegistry = map[string]HTTPSigAlg{
	"rsa-pss-sha512":    RSA_PSS_SHA512,
	"rsa-v1_5-sha256":   RSA_SHA256,
	"hmac-sha256":       HMAC_SHA256,
	"ecdsa-p256-sha256": ECDSA_P256_SHA256,
	"ecdsa-p384-sha384": ECDSA_P384_SHA384,
	"ed25519":           ED25519,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (alg HTTPSigAlg) MarshalJSON() ([]byte, error) {
	_, ok := sigAlgRegistry[string(alg)]
	if ok {
		return json.Marshal(string(alg))
	}
	return nil, ErrInvalidSigAlg
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (alg *HTTPSigAlg) UnmarshalJSON(data []byte) error {
	var sigalg string
	err := json.Unmarshal(data, &sigalg)
	if err != nil {
		return err
	}
	_, ok := sigAlgRegistry[sigalg]
	if ok {
		*alg = HTTPSigAlg(sigalg)
		return nil
	}
	return ErrInvalidSigAlg
}

// DigestAlg represents HTTP Content Digest algorithm.
type DigestAlg string

// Registry of http content digest algorithms.
const (
	DigestSha256 DigestAlg = httpsign.DigestSha256
	DigestSha512 DigestAlg = httpsign.DigestSha512
)

// digestAlgRegistry is a quick mapping from string
// values to valid digest algorithms.
var digestAlgRegistry = map[string]DigestAlg{
	httpsign.DigestSha256: DigestSha256,
	httpsign.DigestSha512: DigestSha512,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (alg DigestAlg) MarshalJSON() ([]byte, error) {
	_, ok := digestAlgRegistry[string(alg)]
	if ok {
		return json.Marshal(string(alg))
	}
	return nil, ErrInvalidDigestAlg
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (alg *DigestAlg) UnmarshalJSON(data []byte) error {
	var digalg string
	err := json.Unmarshal(data, &digalg)
	if err != nil {
		return err
	}
	_, ok := digestAlgRegistry[digalg]
	if ok {
		*alg = DigestAlg(digalg)
		return nil
	}
	return ErrInvalidDigestAlg
}

// ProofMethod represents the proofing method for
// presenting the key.
type ProofMethod string

// Registry of permitted ProofMethods.
const (
	ProofHTTPSig ProofMethod = "httpsig"
	ProofMTLS    ProofMethod = "mtls"
	ProofJWSD    ProofMethod = "jwsd"
	ProofJWS     ProofMethod = "jws"
)

// proofRegistry is a mapping from valid string values to
// corresponding proof methods.
var proofRegistry = map[string]ProofMethod{
	"httpsig": ProofHTTPSig,
	"mtls":    ProofMTLS,
	"jwsd":    ProofJWSD,
	"jws":     ProofJWS,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (p ProofMethod) MarshalJSON() ([]byte, error) {
	_, ok := proofRegistry[string(p)]
	if ok {
		return json.Marshal(string(p))
	}
	return nil, ErrInvalidProofMethod
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (p *ProofMethod) UnmarshalJSON(data []byte) error {
	var proof string
	err := json.Unmarshal(data, &proof)
	if err != nil {
		return err
	}
	_, ok := proofRegistry[string(proof)]
	if ok {
		*p = ProofMethod(proof)
		return nil
	}
	return ErrInvalidProofMethod
}

// Proof implements the [Proof] interface.
func (p ProofMethod) Proof() ProofMethod {
	return p
}

// KeyFormat defines enum of permitted key formats.
type KeyFormat string

// Registry of KeyFormat values.
const (
	FormatJWK      KeyFormat = "jwk"
	FormatCert     KeyFormat = "cert"
	FormatCertS256 KeyFormat = "cert#S256"
)

// keyFormatRegistry is a simple mapping from valid string values
// to the corresponding key format.
var keyFormatRegistry = map[string]KeyFormat{
	"jwk":       FormatJWK,
	"cert":      FormatCert,
	"cert#S256": FormatCertS256,
}

// MarshalJSON implements the [json.MarshalJSON] interface.
func (kf KeyFormat) MarshalJSON() ([]byte, error) {
	_, ok := keyFormatRegistry[string(kf)]
	if ok {
		return json.Marshal(string(kf))
	}
	return nil, ErrInvalidKeyFormat
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (kf *KeyFormat) UnmarshalJSON(data []byte) error {
	var format string
	err := json.Unmarshal(data, &format)
	if err != nil {
		return err
	}
	_, ok := keyFormatRegistry[string(format)]
	if ok {
		*kf = KeyFormat(format)
		return nil
	}
	return ErrInvalidKeyFormat
}

// ClientInstance defines a client by reference (string) or
// by value (object).
type ClientInstance struct {
	Key     ClientKey     `json:"key"`
	ClassID string        `json:"class_id,omitempty"`
	Display ClientDisplay `json:"display,omitempty"`
	Ref     string
}

// NewClient is the constructor for client instance by value (object).
func NewClient(key ClientKey, options ...clientOption) (client ClientInstance, err error) {
	c := &ClientInstance{Key: key}
	for _, setter := range options {
		err = setter(c)
		if err != nil {
			return
		}
	}
	return *c, nil
}

// clientOption is a functional parameter for client constructor.
type clientOption func(c *ClientInstance) error

// ClientDisplay presents information regarding the client
// instance for displaying to the user.
type ClientDisplay struct {
	Name string `json:"name"`
	URI  URL    `json:"uri,omitempty"`
	Logo URL    `json:"logo_uri,omitempty"`
}

// ClientKey the key object of the client. It is used as
// either a key object by value (object) or by reference (string).
type ClientKey struct {
	Proof    Proofer         `json:"proof"`
	JWK      json.RawMessage `json:"jwk,omitempty"`
	Cert     string          `json:"cert,omitempty"`
	CertS256 string          `json:"cert#S256,omitempty"`
	Ref      string          `json:"-"`
}

// TODO: constructor for ClientKey to facilitate with std.

// Proofer describes any object that conveys the proofing information.
type Proofer interface {
	Proof() ProofMethod
}

// HTTPSig represents HTTP signature proofing method.
type HTTPSig struct {
	Method    ProofMethod `json:"method"` // == "httpsig"
	SigAlg    HTTPSigAlg  `json:"alg"`
	DigestAlg DigestAlg   `json:"content-digest"`
}

// Proof implements [Proofer] interface.
func (sig HTTPSig) Proof() ProofMethod {
	return ProofHTTPSig
}
