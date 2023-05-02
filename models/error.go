package models

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
