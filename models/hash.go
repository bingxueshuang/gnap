package models

import (
	"crypto"
	"encoding/json"
	"errors"
	"hash"
)

// HashMethod is a hash name string from IANA Named
// Information Hash Algorithm Registry.
type HashMethod string

// Contents of IANA Named Information Hash Algorithm Registry.
const (
	SHA_256  HashMethod = "sha-256"  // [RFC6920]
	SHA_384  HashMethod = "sha-384"  // [FIPS 180-4]
	SHA_512  HashMethod = "sha-512"  // [FIPS 180-4]
	SHA3_224 HashMethod = "sha3-224" // [FIPS 202]
	SHA3_384 HashMethod = "sha3-384" // [FIPS 202]
	SHA3_512 HashMethod = "sha3-512" // [FIPS 202]

	BLAKE2s_256 HashMethod = "blake2s-256" // [RFC7693]
	BLAKE2b_256 HashMethod = "blake2b-256" // [RFC7693]
)

// hashFuncRegistry maps HashMethod to corresponding cryptographic
// hash algorithm. The imports mentioned in the comments need to be
// included before usage.
var HashFuncRegistry = map[HashMethod]crypto.Hash{
	SHA_256:     crypto.SHA256,      // crypto/sha256
	SHA_384:     crypto.SHA384,      // crypto/sha512
	SHA_512:     crypto.SHA512,      // crypto/sha512
	SHA3_224:    crypto.SHA3_224,    // golang.org/x/crypto/sha3
	SHA3_384:    crypto.SHA3_384,    // golang.org/x/crypto/sha3
	SHA3_512:    crypto.SHA3_512,    // golang.org/x/crypto/sha3
	BLAKE2s_256: crypto.BLAKE2s_256, // golang.org/x/crypto/blake2s
	BLAKE2b_256: crypto.BLAKE2b_256, // golang.org/x/crypto/blake2b
}

// ErrInvalidHashMethod is returned when a hash method that is not
// defined in the registry is encountered.
var ErrInvalidHashMethod = errors.New("invalid hash method")

// MarshalJSON implements [json.Marshaler] interface.
func (hm HashMethod) MarshalJSON() ([]byte, error) {
	_, ok := HashFuncRegistry[hm]
	if !ok {
		return nil, ErrInvalidHashMethod
	}
	return json.Marshal(string(hm))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (hm *HashMethod) UnmarshalJSON(data []byte) error {
	var hash string
	err := json.Unmarshal(data, &hash)
	if err != nil {
		return err
	}
	_, ok := HashFuncRegistry[HashMethod(hash)]
	if !ok {
		return ErrInvalidHashMethod
	}
	*hm = HashMethod(hash)
	return nil
}

// HashFunc returns the hash algorithm corresponding to the
// HashMethod hm. If hash method is empty, then [SHA_256] is
// returned by default. In case of invalid hash method, nil
// is returned.
func (hm HashMethod) HashFunc() hash.Hash {
	if hm == "" {
		// SHA_256 is default
		hm = SHA_256
	}
	hash, ok := HashFuncRegistry[hm]
	if !ok {
		return nil
	}
	return hash.New()
}
