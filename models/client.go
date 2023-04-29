package models

import (
	"encoding/json"
	"errors"

	"github.com/yaronf/httpsign"
)

var ErrInvalidSigAlg = errors.New("invalid http signature algorithm")
var ErrInvalidDigestAlg = errors.New("invalid content digest algorithm")
var ErrInvalidProofMethod = errors.New("invalid proof method")
var ErrInvalidKeyFormat = errors.New("invalid key format")

type HTTPSigAlg string

const (
	RSA_PSS_SHA512    HTTPSigAlg = "rsa-pss-sha512"
	RSA_SHA256        HTTPSigAlg = "rsa-v1_5-sha256"
	HMAC_SHA256       HTTPSigAlg = "hmac-sha256"
	ECDSA_P256_SHA256 HTTPSigAlg = "ecdsa-p256-sha256"
	ECDSA_P384_SHA384 HTTPSigAlg = "ecdsa-p384-sha384"
	ED25519           HTTPSigAlg = "ed25519"
)

var sigAlgRegistry = map[string]HTTPSigAlg{
	"rsa-pss-sha512":    RSA_PSS_SHA512,
	"rsa-v1_5-sha256":   RSA_SHA256,
	"hmac-sha256":       HMAC_SHA256,
	"ecdsa-p256-sha256": ECDSA_P256_SHA256,
	"ecdsa-p384-sha384": ECDSA_P384_SHA384,
	"ed25519":           ED25519,
}

func (alg HTTPSigAlg) MarshalJSON() ([]byte, error) {
	_, ok := sigAlgRegistry[string(alg)]
	if ok {
		return json.Marshal(string(alg))
	}
	return nil, ErrInvalidSigAlg
}

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

type DigestAlg string

const (
	DigestSha256 DigestAlg = httpsign.DigestSha256
	DigestSha512 DigestAlg = httpsign.DigestSha512
)

var digestAlgRegistry = map[string]DigestAlg{
	httpsign.DigestSha256: DigestSha256,
	httpsign.DigestSha512: DigestSha512,
}

func (alg DigestAlg) MarshalJSON() ([]byte, error) {
	_, ok := digestAlgRegistry[string(alg)]
	if ok {
		return json.Marshal(string(alg))
	}
	return nil, ErrInvalidDigestAlg
}

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

type ProofMethod string

const (
	ProofHTTPSig ProofMethod = "httpsig"
	ProofMTLS    ProofMethod = "mtls"
	ProofJWSD    ProofMethod = "jwsd"
	ProofJWS     ProofMethod = "jws"
)

var proofRegistry = map[string]ProofMethod{
	"httpsig": ProofHTTPSig,
	"mtls":    ProofMTLS,
	"jwsd":    ProofJWSD,
	"jws":     ProofJWS,
}

func (p ProofMethod) MarshalJSON() ([]byte, error) {
	_, ok := proofRegistry[string(p)]
	if ok {
		return json.Marshal(string(p))
	}
	return nil, ErrInvalidProofMethod
}

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

func (p ProofMethod) Proof() ProofMethod {
	return p
}

type KeyFormat string

const (
	FormatJWK      KeyFormat = "jwk"
	FormatCert     KeyFormat = "cert"
	FormatCertS256 KeyFormat = "cert#S256"
)

var keyFormatRegistry = map[string]KeyFormat{
	"jwk":       FormatJWK,
	"cert":      FormatCert,
	"cert#S256": FormatCertS256,
}

func (kf KeyFormat) MarshalJSON() ([]byte, error) {
	_, ok := keyFormatRegistry[string(kf)]
	if ok {
		return json.Marshal(string(kf))
	}
	return nil, ErrInvalidKeyFormat
}

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

type ClientInstance struct {
	Key     ClientKey     `json:"key"`
	ClassID string        `json:"class_id,omitempty"`
	Display ClientDisplay `json:"display,omitempty"`
	Ref     string
}

type ClientDisplay struct {
	Name string `json:"name"`
	URI  URL    `json:"uri,omitempty"`
	Logo URL    `json:"logo_uri,omitempty"`
}

type ClientKey struct {
	Proof    Proofer         `json:"proof"`
	JWK      json.RawMessage `json:"jwk,omitempty"`
	Cert     string          `json:"cert,omitempty"`
	CertS256 string          `json:"cert#S256,omitempty"`
	Ref      string          `json:"-"`
}

type Proofer interface {
	Proof() ProofMethod
}

type HTTPSig struct {
	Method           ProofMethod `json:"method"` // == "httpsig"
	Alg              HTTPSigAlg  `json:"alg"`
	ContentDigestAlg DigestAlg   `json:"content-digest"`
}

func (sig HTTPSig) Proof() ProofMethod {
	return ProofHTTPSig
}
