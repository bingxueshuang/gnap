package models

import (
	"encoding/json"
	"errors"
)

// ErrInvalidErrorCode is returned when a code not defined in the
// error codes registry is encountered.
var ErrInvalidErrorCode = errors.New("invalid error code")

// ErrorCode defines a type-safe enum for GNAP error codes.
type ErrorCode struct {
	name string `json:"-"`
}

// Contents of Error Code Registry.
var (
	ECInvalidRequest      = ErrorCode{"invalid_request"}
	ECInvalidClient       = ErrorCode{"invalid_client"}
	ECInvalidInteraction  = ErrorCode{"invalid_interaction"}
	ECInvalidFlag         = ErrorCode{"invalid_flag"}
	ECInvalidRotation     = ErrorCode{"invalid_rotation"}
	ECkrNotSupported      = ErrorCode{"key_rotation_not_supported"}
	ECInvalidContinuation = ErrorCode{"invalid_continuation"}
	ECUserDenied          = ErrorCode{"user_denied"}
	ECRequestDenied       = ErrorCode{"request_denied"}
	ECUnknownUser         = ErrorCode{"unknown_user"}
	ECUnknownInteraction  = ErrorCode{"unknown_interaction"}
	ECTooFast             = ErrorCode{"too_fast"}
	ECTooManyAttempts     = ErrorCode{"too_many_attempts"}
)

// errorCodeRegistry is a quick mapping from string to valid error codes.
var errorCodeRegistry = map[string]ErrorCode{
	"invalid_request":            ECInvalidRequest,
	"invalid_client":             ECInvalidClient,
	"invalid_interaction":        ECInvalidInteraction,
	"invalid_flag":               ECInvalidFlag,
	"invalid_rotation":           ECInvalidRotation,
	"key_rotation_not_supported": ECkrNotSupported,
	"invalid_continuation":       ECInvalidContinuation,
	"user_denied":                ECUserDenied,
	"request_denied":             ECRequestDenied,
	"unknown_user":               ECUnknownUser,
	"unknown_interaction":        ECUnknownInteraction,
	"too_fast":                   ECTooFast,
	"too_many_attempts":          ECTooManyAttempts,
}

// ErrorDescription maps error codes to default description as given in the draft.
var ErrorDescription = map[ErrorCode]string{
	ECInvalidRequest:      "The request is missing a required parameter, includes an invalid parameter value or is otherwise malformed.",
	ECInvalidClient:       "The request was made from a client that was not recognized or allowed by the AS, or the client's signature validation failed.",
	ECInvalidInteraction:  "The client instance has provided an interaction reference that is incorrect for this request or the interaction modes in use have expired.",
	ECInvalidFlag:         "The flag configuration is not valid.",
	ECInvalidRotation:     "The token rotation request is not valid.",
	ECkrNotSupported:      "The AS does not allow rotation of this access token's key.",
	ECInvalidContinuation: "The continuation of the referenced grant could not be processed.",
	ECUserDenied:          "The RO denied the request.",
	ECRequestDenied:       "The request was denied for an unspecified reason.",
	ECUnknownUser:         "The user presented in the request is not known to the AS or does not match the user present during interaction.",
	ECUnknownInteraction:  "The interaction integrity could not be established.",
	ECTooFast:             "The client instance did not respect the timeout in the wait response before the next call.",
	ECTooManyAttempts:     "A limit has been reached in the total number of reasonable attempts.",
}

// Error represents the "error" field of [Response] object that is returned
// if the AS determines that the request cannot be completed for any reason.
type Error struct {
	// single ASCII error code defining the error.
	Code ErrorCode `json:"code"` // REQUIRED
	// human-readable string description of the error intended for the
	// developer of the client.
	Desc string `json:"description,omitempty"` // OPTIONAL
	// a boolean field indicating whether the error code is represented
	// by reference as an JSON string.
	IsRef bool `json:"-"`
}

// MarshalJSON implements the [json.Marshaler] interface.
func (ec ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(ec.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (ec *ErrorCode) UnmarshalJSON(data []byte) error {
	var code string
	err := json.Unmarshal(data, &code)
	if err != nil {
		return err
	}
	ecode, ok := errorCodeRegistry[code]
	if ok {
		*ec = ecode
		return nil
	}
	return ErrInvalidErrorCode
}

// MarshalJSON implements the [json.Marshaler] interface.
func (e Error) MarshalJSON() ([]byte, error) {
	if e.IsRef {
		return json.Marshal(e.Code)
	}
	type Alias Error
	return json.Marshal(Alias(e))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (e *Error) UnmarshalJSON(data []byte) (err error) {
	var ref ErrorCode
	err = json.Unmarshal(data, &ref)
	if err == nil {
		e.Code = ref
		e.IsRef = true
		return
	}
	type Alias Error
	var alias Alias
	err = json.Unmarshal(data, &alias)
	if err == nil {
		*e = Error(alias)
		return
	}
	return
}
