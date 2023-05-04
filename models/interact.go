package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidFinishMethod is returned when a finish method not defined
// in the registry is encountered.
var ErrInvalidFinishMethod = errors.New("invalid finish method")

// ErrInvalidStartMode is returned when a start mode not defined in the registry
// is encountered.
var ErrInvalidStartMode = errors.New("invalid start mode")

// StartMode is enum that defines the means of how the client
// allows the RS to communicate with the AS. Each interaction
// start modes has a unique identifying name.
type StartMode struct {
	name string `json:"-"`
}

// Contents of Interaction Start Modes Registry.
var (
	SMRedirect = StartMode{"redirect"}
	SMApp      = StartMode{"app"}
	SMCode     = StartMode{"user_code"}
	SMCodeURI  = StartMode{"user_code_uri"}
)

// startModeRegistry is a quick mapping from string to valid start modes.
var startModeRegistry = map[string]StartMode{
	"redirect":      SMRedirect,
	"app":           SMApp,
	"user_code":     SMCode,
	"user_code_uri": SMCodeURI,
}

// FinishMethod is enum that defines the means of how the
// client gets notified when the interaction is completed.
type FinishMethod struct {
	name string `json:"-"`
}

// Contents of the Interaction Finish Methods Registry.
var (
	FMRedirect = FinishMethod{"redirect"}
	FMPush     = FinishMethod{"push"}
)

// finishMethodRegistry is a quick mapping from string to valid finish methods.
var finishMethodRegistry = map[string]FinishMethod{
	"redirect": FMRedirect,
	"push":     FMPush,
}

// IAReq defines the object format for "interact" field of [Request].
// client instance declares the parameters for interaction methods that it
// can support using the interact field. It declares how the client can
// initiate and complete the request, as well as provide hints to the AS
// about user preferences such as locale.
type IAReq struct {
	// how the client instance can start an interaction.
	Start IAStart `json:"start"` // REQUIRED
	// how the client instance can receive an indication
	// that interaction has finished at the AS.
	Finish IAFinish `json:"finish,omitempty"` // OPTIONAL
	// additional info to inform the interaction process at the AS.
	Hints IAHints `json:"hints,omitempty"` // OPTIONAL
}

// IAResponse is included in "interact" field of [Response] if interaction is
// both supported and necessary. All supported interaction methods are included
// in the same interact object.
type IAResponse struct {
	// redirect to an arbitrary URI.
	// REQUIRED if the redirect start mode is possible for this request.
	Redirect URL `json:"redirect,omitempty"`
	// launch of an application URI.
	// REQUIRED if the app start mode is possible for this request.
	App URL `json:"app,omitempty"`
	// display a short user code.
	// REQUIRED if the user_code start mode is possible for this request.
	Code string `json:"code,omitempty"`
	// display a short user code and URI.
	// REQUIRED if the user_code_uri start mode is possible for this request.
	CodeURI CodeURI `json:"codeuri,omitempty"`
	// unique ASCII string value provided by the AS as a nonce.
	// REQUIRED if the interaction finish method requested by the client instance
	// is possible for this request.
	Finish string `json:"finish,omitempty"`
	// number of seconds after which this set of interaction responses will
	// expire and no longer be usable by the client instance.
	// If omitted, the interaction modes returned do not expire but MAY be
	// invalidated by the AS at any time.
	ExpiresIn int `json:"expires_in,omitempty"` // OPTIONAL
}

// IACallback defines the object format used by the AS to signal to the
// client instance that interaction is complete and the request can be continued
// at the client instance's callback URI.
type IACallback struct {
	// interaction hash value
	Hash string `json:"hash"` // REQUIRED
	// interaction reference generated for this interaction.
	IARef string `json:"interact_ref"` // REQUIRED
}

// IAStart defines the means of initiation of interaction by the client,
// if the client instance is capable of starting interaction with the end user.
type IAStart struct {
	// interaction start mode.
	Mode StartMode `json:"mode"`
	// a boolean field indicating whether the start mode is represented
	// by reference as an opaque string.
	IsRef bool `json:"-"`
}

// IAFinish represents the the object under the "finish" key in [IAReq].
// It is sent to the AS if the client is capable of receiving a message from
// the AS indicating that the RO has completed the interaction.
type IAFinish struct {
	// callback method that the AS will use to contact the client instance.
	Method FinishMethod `json:"method"` // REQUIRED
	// the URI that the AS will either send the RO to after interaction
	// or send an HTTP POST request. This URI MUST be an absolute URI,
	// and MUST NOT contain any fragment component.
	// REQUIRED for redirect and push methods.
	URI URL `json:"uri,omitempty"`
	// unique ASCII string value to be used in the calculation of the
	// "hash" query parameter sent to the callback URI, must be sufficiently
	// random to be unguessable by an attacker. MUST be generated by the
	// client instance as a unique value for this request.
	Nonce string `json:"nonce"` // REQUIRED
	// An identifier of a hash calculation mechanism to be used for the
	// callback hash. If absent, the default value is sha-256.
	Hash HashMethod `json:"hash,omitempty"` // OPTIONAL
}

// IAHints is an object describing one or more suggestions from the client
// instance that the AS can use to help drive user interaction.
type IAHints struct {
	// end user's preferred locales that the AS can use during interaction,
	// particularly before the RO has authenticated.
	UILocales []string `json:"ui_locales,omitempty"` // OPTIONAL
}

// CodeURI represents the "user_code_uri" field of the [IAResponse] object.
// AS responds with this field if the client instance indicates that it can
// display a short user-typeable code and AS supports this mode.
type CodeURI struct {
	// unique short code that the end user can type into a provided URI.
	// To facilitate usability, this string MUST be case-insensitive,
	// MUST consist of only easily typeable characters (such as letters or numbers).
	// The string MUST be randomly generated so as to be unguessable by an attacker
	// within the time it is accepted. The time in which this code will be accepted
	// SHOULD be short lived, such as several minutes. It is RECOMMENDED that this
	// code be no more than eight characters in length.
	Code string `json:"code"` // REQUIRED
	// interaction URI that the client instance will direct the RO to.
	// This URI MUST be short enough to be communicated to the end user by
	// the client instance. It is RECOMMENDED that this URI be short enough for
	// an end user to type in manually. The URI MUST NOT contain the code value.
	// This URI MUST be an absolute URI.
	URI URL `json:"uri"` // REQUIRED
}

// MarshalJSON implements the [json.Marshaler] interface.
func (sm StartMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(sm.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (sm *StartMode) UnmarshalJSON(data []byte) error {
	var mode string
	err := json.Unmarshal(data, &mode)
	if err != nil {
		return err
	}
	m, ok := startModeRegistry[mode]
	if ok {
		*sm = m
		return nil
	}
	return ErrInvalidStartMode
}

// MarshalJSON implements the [json.Marshaler] interface.
func (fm FinishMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(fm.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (fm *FinishMethod) UnmarshalJSON(data []byte) error {
	var method string
	err := json.Unmarshal(data, &method)
	if err != nil {
		return err
	}
	m, ok := finishMethodRegistry[method]
	if ok {
		*fm = m
		return nil
	}
	return ErrInvalidFinishMethod
}

// MarshalJSON implements the [json.Marshaler] interface.
func (start IAStart) MarshalJSON() ([]byte, error) {
	if start.IsRef {
		return json.Marshal(start.Mode)
	}
	type Alias IAStart
	return json.Marshal(Alias(start))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (start *IAStart) UnmarshalJSON(data []byte) (err error) {
	var ref StartMode
	err = json.Unmarshal(data, &ref)
	if err == nil {
		start.Mode = ref
		return
	}
	type Alias IAStart
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err == nil {
		*start = IAStart(alias)
		return
	}
	return
}
