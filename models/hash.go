package models

import (
	"encoding/json"
	"errors"

	"github.com/bingxueshuang/gnap/models/internal"
)

var ErrInvalidHashMethod = errors.New("invalid hash method")

// HashMethod is a type-safe enum for hash methods used for interaction completion.
type HashMethod struct {
	name string              `json:"-"`
	sum  func([]byte) []byte `json:"-"`
}

// Contents of [IANA Named Information Hash Algorithms Registry].
//
// [IANA Named Information Hash Algorithms Registry]: https://www.iana.org/assignments/named-information/named-information.xhtml#hash-alg
var (
	HM_SHA_256  = HashMethod{"sha-256", internal.SumSHA_256}
	HM_SHA_384  = HashMethod{"sha-384", internal.SumSHA_384}
	HM_SHA_512  = HashMethod{"sha-512", internal.SumSHA_512}
	HM_SHA3_224 = HashMethod{"sha3-224", internal.SumSHA3_224}
	HM_SHA3_384 = HashMethod{"sha3-384", internal.SumSHA3_384}
	HM_SHA3_512 = HashMethod{"sha3-512", internal.SumSHA3_512}

	HM_BLAKE2s_256 = HashMethod{"blake2s-256", internal.SumBlake2s_256}
	HM_BLAKE2b_256 = HashMethod{"blake2b-256", internal.SumBlake2b_256}
	HM_BLAKE2b_512 = HashMethod{"blake2b-512", internal.SumBlake2b_512}
)

// hashMethodRegistry is a quick mapping from string values to valid hash methods.
var hashMethodRegistry = map[string]HashMethod{
	"sha-256":     HM_SHA_256,
	"sha-384":     HM_SHA_384,
	"sha-512":     HM_SHA_512,
	"sha3-224":    HM_SHA3_224,
	"sha3-384":    HM_SHA3_384,
	"sha3-512":    HM_SHA3_512,
	"blake2s-256": HM_BLAKE2s_256,
	"blake2b-256": HM_BLAKE2b_256,
	"blake2b-512": HM_BLAKE2b_512,
}

// MarshalJSON implements the [json.Marshaler] interface.
func (hm HashMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(hm.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (hm *HashMethod) UnmarshalJSON(data []byte) error {
	var hash string
	err := json.Unmarshal(data, &hash)
	if err != nil {
		return err
	}
	hmethod, ok := hashMethodRegistry[hash]
	if ok {
		*hm = hmethod
		return nil
	}
	return ErrInvalidHashMethod
}

// Sum is a helper function to call the hash algorithm of the
// hash method.
func (hm HashMethod) Sum(data []byte) []byte {
	if hm.sum == nil {
		return nil
	}
	return hm.sum(data)
}
