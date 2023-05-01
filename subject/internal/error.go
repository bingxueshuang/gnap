package internal

import "errors"

// ErrInvalidFormat is returned when subject id format encountered
// is not defined in the registry or invalid.
var ErrInvalidFormat = errors.New("invalid subject id format")

// ErrInvalidID is returned when the subject identifier object
// is found to be malformed.
var ErrInvalidID = errors.New("invalid subject identifier")

// wrapped is error wrapper that wraps multiple errors into
// single one. Useful alternative to %w in [fmt.Errorf].
type wrapped struct {
	Message string
	Errors  []error
}

// Error implements [error] interface.
func (w *wrapped) Error() string {
	return w.Message
}

// Unwrap implements [errors.Unwrap].
func (w *wrapped) Unwrap() []error {
	return w.Errors
}

// Wrap is utility function to create wrapped errors.
func Wrap(msg string, errors ...error) error {
	return &wrapped{Message: msg, Errors: errors}
}
