package models

import "github.com/bingxueshuang/gnap/models/internal"

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
