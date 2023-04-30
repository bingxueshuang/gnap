package internal

import "encoding/json"

// Format is a typed enum for subject identifier format.
// Identifier Formats define how to encode identity info for a subject.
type Format struct {
	Name string `json:"-"`
}

// Registry of defined subject id formats. Aliases format is not
// implemented (since not necessary for GNAP).
var (
	FormatAccount = Format{"account"}      // acct URI
	FormatEmail   = Format{"email"}        // email id
	FormatIssSub  = Format{"iss_sub"}      // issuer and subject
	FormatOpaque  = Format{"opaque"}       // opaque string
	FormatPhone   = Format{"phone_number"} // phone number
	FormatDID     = Format{"did"}          // w3 did
	FormatURI     = Format{"uri"}          // uri
	// aliases format not implemented
)

// String implements Stringer interface for fmt.Print* functions.
func (f Format) String() string {
	return string(f.Name)
}

// MarshalJSON implements [json.Marshaler] interface.
func (f Format) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Name)
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (f *Format) UnmarshalJSON(data []byte) error {
	var fmt string
	err := json.Unmarshal(data, &fmt)
	if err != nil {
		return err
	}
	_, ok := Validators[Format{fmt}]
	if ok {
		*f = Format{fmt}
		return nil
	}
	return ErrInvalidFormat
}
