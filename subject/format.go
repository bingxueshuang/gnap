package subject

import (
	"encoding/json"
)

// Format is a typed enum for subject identifier format.
// Identifier Formats define how to encode identity info for a subject.
type Format struct {
	name string `json:"-"`
}

// Registry of defined subject id formats. Aliases format is not
// implemented (since not necessary for GNAP).
var (
	Account     = Format{"account"}      // acct URI
	Email       = Format{"email"}        // email id
	IssSub      = Format{"iss_sub"}      // issuer and subject
	Opaque      = Format{"opaque"}       // opaque string
	PhoneNumber = Format{"phone_number"} // telephone number
	DID         = Format{"did"}          // w3 did
	URI         = Format{"uri"}          // uri
	// aliases format not implemented
)

// String implements Stringer interface for fmt.Print* functions.
func (f Format) String() string {
	return string(f.name)
}

// MarshalJSON implements [json.Marshaler] interface.
func (f Format) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.name)
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (f *Format) UnmarshalJSON(data []byte) error {
	var fmt string
	err := json.Unmarshal(data, &fmt)
	if err != nil {
		return err
	}
	_, ok := validators[Format{fmt}]
	if ok {
		*f = Format{fmt}
		return nil
	}
	return ErrInvalidFormat
}
