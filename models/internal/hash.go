package internal

import (
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

// SumSHA_256 obtains the cryptographic hash using SHA2 algorithm
// with 256 bits hash length as per [RFC6920].
//
// [RFC6920]: https://www.iana.org/go/rfc6920
func SumSHA_256(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

// SumSHA_384 obtains the cryptographic hash using SHA2 algorithm
// with 384 bits hash length as per [FIPS 180-4].
//
// [FIPS 180-4]: https://dx.doi.org/10.6028/NIST.FIPS.180-4
func SumSHA_384(data []byte) []byte {
	sum := sha512.Sum384(data)
	return sum[:]
}

// SumSHA_512 obtains the cryptographic hash using SHA2 algorithm
// with 512 bits hash length as per [FIPS 180-4].
//
// [FIPS 180-4]: https://dx.doi.org/10.6028/NIST.FIPS.180-4
func SumSHA_512(data []byte) []byte {
	sum := sha512.Sum512(data)
	return sum[:]
}

// SumSHA3_224 obtains the cryptographic hash using SHA3 algorithm
// with 224 bits hash length as per [FIPS 202].
//
// [FIPS 202]: https://dx.doi.org/10.6028/NIST.FIPS.202
func SumSHA3_224(data []byte) []byte {
	sum := sha3.Sum224(data)
	return sum[:]
}

// SumSHA3_384 obtains the cryptographic hash using SHA3 algorithm
// with 384 bits hash length as per [FIPS 202].
//
// [FIPS 202]: https://dx.doi.org/10.6028/NIST.FIPS.202
func SumSHA3_384(data []byte) []byte {
	sum := sha3.Sum384(data)
	return sum[:]
}

// SumSHA3_224 obtains the cryptographic hash using SHA3 algorithm
// with 512 bits hash length as per [FIPS 202].
//
// [FIPS 202]: https://dx.doi.org/10.6028/NIST.FIPS.202
func SumSHA3_512(data []byte) []byte {
	sum := sha3.Sum512(data)
	return sum[:]
}

// SumBlake2s_256 obtains the cryptographic hash using Blake-2s algorithm
// with 256 bits hash length as per [RFC7693].
//
// [RFC7693]: https://www.iana.org/go/rfc7693
func SumBlake2s_256(data []byte) []byte {
	sum := blake2s.Sum256(data)
	return sum[:]
}

// SumBlake2b_256 obtains the cryptographic hash using Blake-2b algorithm
// with 256 bits hash length as per [RFC7693].
//
// [RFC7693]: https://www.iana.org/go/rfc7693
func SumBlake2b_256(data []byte) []byte {
	sum := blake2b.Sum256(data)
	return sum[:]
}

// SumBlake2b_512 obtains the cryptographic hash using Blake-2b algorithm
// with 512 bits hash length as per [RFC7693].
//
// [RFC7693]: https://www.iana.org/go/rfc7693
func SumBlake2b_512(data []byte) []byte {
	sum := blake2b.Sum512(data)
	return sum[:]
}
