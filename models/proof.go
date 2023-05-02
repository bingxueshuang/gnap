package models

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
