package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidDigestAlg is returned when the digest algorithm is not defined
// in the registry.
var ErrInvalidDigestAlg = errors.New("invalid digest algorithm")

// ErrInvalidSigAlg is returned when the signature is not defined in the registry.
var ErrInvalidSigAlg = errors.New("invalid signature algorithm")

// HTTPSigAlg represents type-safe enum for HTTP Signature Algorithm.
type HTTPSigAlg struct {
	name string `json:"-"`
}

// Contents of HTTP Signature Algorithms Registry.
var (
	// RSASSA-PSS using SHA-512
	RSA_PSS_SHA512 = HTTPSigAlg{"rsa-pss-sha512"}
	// RSASSA-PKCS1-v1_5 using SHA-256
	RSA_V1_5_SHA256 = HTTPSigAlg{"rsa-v1_5-sha256"}
	// HMAC using SHA-256
	HMAC_SHA256 = HTTPSigAlg{"hmac-sha256"}
	// ECDSA using curve P-256 DSS and SHA-256
	ECDSA_P256_SHA256 = HTTPSigAlg{"ecdsa-p256-sha256"}
	// ECDSA using curve P-384 DSS and SHA-384
	ECDSA_P384_SHA384 = HTTPSigAlg{"ecdsa-p384-sha384"}
	// Edwards Curve DSA using curve edwards25519
	ED25519 = HTTPSigAlg{"ed25519"}
)

// httpSigAlgRegistry is a quick mapping from string to
// valid HTTP Signature Algorithm.
var httpSigAlgRegistry = map[string]HTTPSigAlg{
	"rsa-pss-sha512":    RSA_PSS_SHA512,
	"rsa-v1_5-sha256":   RSA_V1_5_SHA256,
	"hmac-sha256":       HMAC_SHA256,
	"ecdsa-p256-sha256": ECDSA_P256_SHA256,
	"ecdsa-p384-sha384": ECDSA_P384_SHA384,
	"ed25519":           ED25519,
}

// DigestAlg represents type-safe enum for HTTP Content Digest Algorithm.
type DigestAlg struct {
	name string `json:"-"`
}

// Contents of Hash Algorithms for HTTP Digest Fields Registry.
var (
	// SHA-256 algorithm (standard)
	CDsha256 = DigestAlg{"sha-256"}
	// SHA-512 algorithm (standard)
	CDsha512 = DigestAlg{"sha-512"}
)

// digestAlgRegistry is a quick mapping from strings to valid digest algorithms.
var digestAlgRegistry = map[string]DigestAlg{
	"sha-256": CDsha256,
	"sha-512": CDsha512,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (alg HTTPSigAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(alg.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (alg *HTTPSigAlg) UnmarshalJSON(data []byte) error {
	var sigalg string
	err := json.Unmarshal(data, &sigalg)
	if err != nil {
		return err
	}
	a, ok := httpSigAlgRegistry[sigalg]
	if ok {
		*alg = a
		return nil
	}
	return ErrInvalidSigAlg
}

// MarshalJSON implements the [json.Marshaler] interface.
func (alg DigestAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(alg.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (alg *DigestAlg) UnmarshalJSON(data []byte) error {
	var digalg string
	err := json.Unmarshal(data, &digalg)
	if err != nil {
		return err
	}
	a, ok := digestAlgRegistry[digalg]
	if ok {
		*alg = a
		return nil
	}
	return ErrInvalidDigestAlg
}
