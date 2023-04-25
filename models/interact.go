package models

import (
	"encoding/json"
	"errors"
	"net/url"
)

// ErrInvalidStartMode is returned when a start mode
// not defined in the registry is encountered.
var ErrInvalidStartMode = errors.New("invalid start mode")

// ErrInvalidFinishMethod is returned when a finish method
// not defined in the registry is encountered.
var ErrInvalidFinishMethod = errors.New("invalid finish method")

// StartMode indicates how the client instance can start an interaction.
type StartMode string

// Contents of Interaction Start Modes Registry.
const (
	ModeRedirect StartMode = "redirect"
	ModeApp      StartMode = "app"
	ModeCode     StartMode = "user_code"
	ModeCodeURI  StartMode = "user_code_uri"
)

// startModesRegistry is a quick mapping of registry values
// to verify that a string is valid start mode.
var startModesRegistry = map[string]StartMode{
	"redirect":      ModeRedirect,
	"app":           ModeApp,
	"user_code":     ModeCode,
	"user_code_uri": ModeCodeURI,
}

// MarshalJSON implements [json.Marshaler] interface.
func (sm StartMode) MarshalJSON() ([]byte, error) {
	_, ok := startModesRegistry[string(sm)]
	if !ok {
		return nil, ErrInvalidStartMode
	}
	return json.Marshal(string(sm))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (sm *StartMode) UnmarshalJSON(data []byte) error {
	var mode string
	err := json.Unmarshal(data, &mode)
	if err != nil {
		return err
	}
	_, ok := startModesRegistry[mode]
	if !ok {
		return ErrInvalidStartMode
	}
	*sm = StartMode(mode)
	return nil
}

// FinishMethod indicates how the client instance can
// receive an indication that interaction has finished
// at the AS.
type FinishMethod string

// Contents of Interaction Finish Methods Registry.
const (
	MethodPush     FinishMethod = "push"
	MethodRedirect FinishMethod = "redirect"
)

// finishMethodRegistry is a quick mapping to
// check if a string is a valid FinishMethod.
var finishMethodsRegistry = map[string]FinishMethod{
	"push":     MethodPush,
	"redirect": MethodRedirect,
}

// MarshalJSON implements [json.Marshaler] interface.
func (fm FinishMethod) MarshalJSON() ([]byte, error) {
	_, ok := finishMethodsRegistry[string(fm)]
	if !ok {
		return nil, ErrInvalidFinishMethod
	}
	return json.Marshal(string(fm))
}

// UnmarshalJSON implements [json.Unmarshaler] interface.
func (fm *FinishMethod) UnmarshalJSON(data []byte) error {
	var method string
	err := json.Unmarshal(data, &method)
	if err != nil {
		return err
	}
	_, ok := finishMethodsRegistry[method]
	if !ok {
		return ErrInvalidFinishMethod
	}
	*fm = FinishMethod(method)
	return nil
}

// IARequest describes the modes that the client instance supports
// for allowing the RO to interact with the AS and modes for the
// client instance to receive updates when interaction is complete.
type IARequest struct {
	Start  []IAStart `json:"start"`
	Finish *IAFinish `json:"finish,omitempty"`
	Hints  *IAHints  `json:"hints,omitempty"`
}

// IAResponse indicates that interaction through some set of
// defined mechanisms needs to take place.
type IAResponse struct {
	Redirect  *URL       `json:"redirect,omitempty"`
	App       *URL       `json:"app,omitempty"`
	UserCode  string     `json:"user_code,omitempty"`
	CodeURI   *IACodeURI `json:"user_code_uri,omitempty"`
	Finish    string     `json:"finish,omitempty"`
	ExpiresIn int        `json:"expires_in,omitempty"`
}

// IACallback represents the information conveyed to the
// client through the interaction callback.
type IACallback struct {
	Hash        string `json:"hash"`
	InteractRef string `json:"interact_ref"`
}

// Encode creates query parameters from the
// IACallback object.
func (c IACallback) Encode() url.Values {
	v := url.Values{}
	v.Set("hash", c.Hash)
	v.Set("interact_ref", c.InteractRef)
	return v
}

// FromQuery constructs a new IACallback from the
// given query string parameters.
func FromQuery(params url.Values) IACallback {
	return IACallback{
		Hash:        params.Get("hash"),
		InteractRef: params.Get("interact_ref"),
	}
}

// IAStart indicates how the client instance can start an interaction.
type IAStart struct {
	Mode  StartMode `json:"mode"`
	IsRef bool      `json:"-"`
}

// MarshalJSON implements the [json.Marshaler] interface. Encodes to
// json string or json object with mode property.
func (ias IAStart) MarshalJSON() ([]byte, error) {
	if ias.IsRef {
		return json.Marshal(ias.Mode)
	}
	type Alias IAStart
	alias := Alias(ias)
	return json.Marshal(alias)
}

// UnmarshalJSON implements [json.Unmarshaler] interface. Decodes
// from a start mode string or a start mode object.
func (ias *IAStart) UnmarshalJSON(data []byte) error {
	var mode StartMode
	err := json.Unmarshal(data, &mode)
	if err == nil { // start mode by reference
		*ias = IAStart{mode, true}
		return nil
	}
	type Alias IAStart
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	*ias = IAStart(alias)
	return nil
}

// IAFinish indicates how the client instance can receive an
// indication that interaction has finished at the AS.
type IAFinish struct {
	Method     FinishMethod `json:"method"`
	URI        *URL         `json:"uri"`
	Nonce      string       `json:"nonce"`
	HashMethod HashMethod   `json:"hash_method,omitempty"`
}

// IAHints provides additional information to inform the
// interaction process at the AS.
type IAHints struct {
	UILocales []string `json:"uilocales,omitempty"`
}

// IACodeURI represents a User Code or URI object that indicates
// a short user-typable code and a short URI.
type IACodeURI struct {
	Code string `json:"code"`
	URI  URL    `json:"uri"`
}
