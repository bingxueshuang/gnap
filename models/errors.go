package models

import (
	"errors"
	"fmt"
)

// ErrInvalidErrorCode is returned when the GNAPError.Code field is not
// defined in the registry of gnap error codes.
var ErrInvalidErrorCode = errors.New("invalid error code")

// Errors corresponding to GNAP error codes
var (
	ErrGInvalidClient       = errors.New("invalid_client")
	ErrGInvalidInteraction  = errors.New("invalid_interaction")
	ErrGInvalidFlag         = errors.New("invalid_flag")
	ErrGInvalidRotation     = errors.New("invalid_rotation")
	ErrGKRNotSupported      = errors.New("key_rotation_not_supported")
	ErrGInvalidContinuation = errors.New("invalid_continuation")
	ErrGUserDenied          = errors.New("user_denied")
	ErrGRequestDenied       = errors.New("request_denied")
	ErrGUnknownUser         = errors.New("unknown_user")
	ErrGUnknownInteraction  = errors.New("unknown_interaction")
	ErrGTooFast             = errors.New("too_fast")
	ErrGTooManyAttempts     = errors.New("too_many_attempts")
)

// DefaultDescription is the default description for error codes.
var DefaultDescription = map[string]string{
	"invalid_request":            "The request is missing a required parameter, includes an invalid parameter value or is otherwise malformed.",
	"invalid_client":             "The request was made from a client that was not recognized or allowed by the AS, or the client's signature validation failed.",
	"invalid_interaction":        "The client instance has provided an interaction reference that is incorrect for this request or the interaction modes in use have expired.",
	"invalid_flag":               "The flag configuration is not valid.",
	"invalid_rotation":           "The token rotation request is not valid.",
	"key_rotation_not_supported": "The AS does not allow rotation of this access token's key.",
	"invalid_continuation":       "The continuation of the referenced grant could not be processed.",
	"user_denied":                "The RO denied the request.",
	"request_denied":             "The request was denied for an unspecified reason.",
	"unknown_user":               "The user presented in the request is not known to the AS or does not match the user present during interaction.",
	"unknown_interaction":        "The interaction integrity could not be established.",
	"too_fast":                   "The client instance did not respect the timeout in the wait response before the next call.",
	"too_many_attempts":          "A limit has been reached in the total number of reasonable attempts.",
}

// errorRegistry denotes the IANA registry for GNAP error codes.
var errorRegistry = map[string]error{
	"invalid_client":             ErrGInvalidClient,
	"invalid_interaction":        ErrGInvalidInteraction,
	"invalid_flag":               ErrGInvalidFlag,
	"invalid_rotation":           ErrGInvalidRotation,
	"key_rotation_not_supported": ErrGKRNotSupported,
	"invalid_continuation":       ErrGInvalidContinuation,
	"user_denied":                ErrGUserDenied,
	"request_denied":             ErrGRequestDenied,
	"unknown_user":               ErrGUnknownUser,
	"unknown_interaction":        ErrGUnknownInteraction,
	"too_fast":                   ErrGTooFast,
	"too_many_attempts":          ErrGTooManyAttempts,
}

// GNAPError is the error occurred during the GNAP protocol.
type GNAPError struct {
	Code string `json:"code"`
	Desc string `json:"description,omitempty"`
}

// Error implements error interface.
func (e GNAPError) Error() string {
	_, ok := errorRegistry[e.Code]
	if !ok {
		return ErrInvalidErrorCode.Error()
	}
	desc := e.Desc
	if desc == "" {
		desc = DefaultDescription[e.Code]
	}
	return fmt.Sprintf("%s: %v", e.Code, desc)
}

// Is implements [errors.Is].
func (e GNAPError) Is(target error) bool {
	_, valid := errorRegistry[e.Code]
	if !valid {
		return errors.Is(target, ErrInvalidErrorCode)
	}
	switch v := target.(type) {
	case GNAPError:
		return e.Code == v.Code
	case *GNAPError:
		return e.Code == v.Code
	}
	return false
}

// Unwrap returns the underlying error corresponding to the
// status code, or [ErrInvalidErrorCode] if status code is
// invalid.
func (e GNAPError) Unwrap() error {
	err, ok := errorRegistry[e.Code]
	if !ok {
		err = ErrInvalidErrorCode
	}
	return err
}
