package models

import "encoding/json"

// ProofMethod is a type-safe enum for GNAP key proofing methods.
type ProofMethod struct {
	name string `json:"-"`
}

// Contents of Key Proofing Methods Registry.
var (
	PMhttpSig = ProofMethod{"httpsig"}
	PMmtls    = ProofMethod{"mtls"}
	PMjwsd    = ProofMethod{"jwsd"}
	PMjws     = ProofMethod{"jws"}
)

// Key represents object format for GNAP key material.
type Key struct {
	// form of proof that client instance will use when presenting the key.
	Proof KeyProof `json:"proof"` // REQUIRED
	// public key and its properties represented as a JSON Web Key [RFC7517].
	// A JWK MUST contain the alg (Algorithm) and kid (Key ID) parameters.
	// The alg parameter MUST NOT be "none".
	//
	// [RFC7517]: https://www.rfc-editor.org/rfc/rfc7517
	JWK json.RawMessage `json:"jwk,omitempty"` // OPTIONAL
	// PEM serialized value of the certificate used to sign the request,
	// with optional internal whitespace per [RFC7468]. The PEM header and
	// footer are optionally removed.
	//
	// [RFC7468]: https://www.rfc-editor.org/rfc/rfc7468
	Cert string `json:"cert,omitempty"` // OPTIONAL
	// certificate thumbprint calculated as per [OAuth-MTLS] in base64 URL
	// encoding. Note that this format does not include the full public key.
	//
	// [OAUTH-MTLS]: https://www.rfc-editor.org/rfc/rfc8705
	CertS256 string `json:"cert#S256,omitempty"` // OPTIONAL
	// keys can also be passed by reference such that the party receiving the
	// reference will be able to determine the appropriate keying material for use
	// in that part of the protocol. Key references are a single opaque string.
	Ref string `json:"-"`
}

// ProofHTTPSig is object format of key proofing mechanism.
type KeyProof struct {
	// name of the key proofing method to be used.
	Method ProofMethod `json:"method"` // REQUIRED
	// HTTP signature algorithm, from the HTTP Signature Algorithm registry.
	// REQUIRED if "httpsig" proofing is declared in object form.
	SigAlg HTTPSigAlg `json:"alg"` // REQUIRED
	// algorithm used for the Content-Digest field, used
	// to protect the body when present in the message.
	// REQUIRED if "httpsig" proofing is declared in object form.
	DigAlg DigestAlg `json:"content-digest-alg"` // REQUIRED
	// determines if the proofing mechanism is declared as object (by value)
	// or as reference (by ref).
	IsRef bool `json:"-"`
}
