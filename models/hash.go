package models

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
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
	BLAKE2b_512 HashMethod = "blake2b-512" // [RFC7693]
)

// hashFuncRegistry maps HashMethod to corresponding cryptographic
// hash algorithm. The imports mentioned in the comments need to be
// included before usage.
var hashFuncRegistry = map[HashMethod]func([]byte) []byte{
	SHA_256: func(data []byte) []byte { x := sha256.Sum256(data); return x[:] },
	SHA_384: func(data []byte) []byte { x := sha512.Sum384(data); return x[:] },
	SHA_512: func(data []byte) []byte { x := sha512.Sum512(data); return x[:] },

	SHA3_224: func(data []byte) []byte { x := sha3.Sum224(data); return x[:] },
	SHA3_384: func(data []byte) []byte { x := sha3.Sum384(data); return x[:] },
	SHA3_512: func(data []byte) []byte { x := sha3.Sum512(data); return x[:] },

	BLAKE2s_256: func(data []byte) []byte { x := blake2s.Sum256(data); return x[:] },
	BLAKE2b_256: func(data []byte) []byte { x := blake2b.Sum256(data); return x[:] },
	BLAKE2b_512: func(data []byte) []byte { x := blake2b.Sum512(data); return x[:] },
}

// ErrInvalidHashMethod is returned when a hash method that is not
// defined in the registry is encountered.
var ErrInvalidHashMethod = errors.New("invalid hash method")

// MarshalJSON implements [json.Marshaler] interface.
func (hm HashMethod) MarshalJSON() ([]byte, error) {
	_, ok := hashFuncRegistry[hm]
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
	_, ok := hashFuncRegistry[HashMethod(hash)]
	if !ok {
		return ErrInvalidHashMethod
	}
	*hm = HashMethod(hash)
	return nil
}

// Sum calculates cryptographic hash digest
// using hm HashMethod. Returns nil if the hash algorithm
// is not present in the registry.
func (hm HashMethod) Sum(data []byte) []byte {
	if hm == "" {
		// SHA_256 is default
		hm = SHA_256
	}
	hash, ok := hashFuncRegistry[hm]
	if ok {
		return hash(data)
	}
	return nil
}
