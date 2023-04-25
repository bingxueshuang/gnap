package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestGNAPError_Error(t *testing.T) {
	tests := []struct {
		name string
		in   GNAPError
		want string
	}{
		{
			name: "valid",
			in:   GNAPError{"invalid_interaction", "malformed interaction request"},
			want: "invalid_interaction: malformed interaction request",
		},
		{
			name: "default",
			in:   GNAPError{Code: "too_fast"},
			want: "too_fast: The client instance did not respect the timeout in the wait response before the next call.",
		},
		{
			name: "invalid",
			in:   GNAPError{Code: "wrong_code"},
			want: "invalid error code",
		},
		{
			name: "empty",
			in:   GNAPError{},
			want: "invalid error code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.Error(); got != tt.want {
				t.Errorf("GNAPError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGNAPError_Unwrap(t *testing.T) {
	tests := []struct {
		name string
		in   GNAPError
		want error
	}{
		{
			name: "valid",
			in:   GNAPError{Code: "invalid_flag"},
			want: ErrGInvalidFlag,
		},
		{
			name: "invalid",
			in:   GNAPError{Code: "wrong_code"},
			want: ErrInvalidErrorCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errors.Unwrap(tt.in)
			if got != tt.want {
				t.Errorf("GNAPError.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGNAPError_Is(t *testing.T) {
	tests := []struct {
		name   string
		in     GNAPError
		target error
		want   bool
	}{
		{
			name:   "valid",
			in:     GNAPError{"invalid_rotation", "rotation needs to be signed"},
			target: ErrGInvalidRotation,
			want:   true,
		},
		{
			name:   "invalid self",
			in:     GNAPError{Code: "wrong_code"},
			target: ErrInvalidErrorCode,
			want:   true,
		},
		{
			name:   "invalid other",
			in:     GNAPError{Code: "too_many_attempts"},
			target: GNAPError{Code: "invalid_code"},
			want:   false,
		},
		{
			name:   "other",
			in:     GNAPError{Code: "key_rotation_not_supported"},
			target: errors.New("unknown error"),
			want:   false,
		},
		{
			name:   "struct",
			in:     GNAPError{Code: "invalid_continuation"},
			target: GNAPError{"invalid_continuation", "continuation request is expected to have interact_ref"},
			want:   true,
		},
		{
			name:   "pointer",
			in:     GNAPError{Code: "unknown_interaction"},
			target: &GNAPError{"unknown_interaction", "interaction not specified in the registry"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errors.Is(tt.in, tt.target)
			if got != tt.want {
				t.Errorf("GNAPError.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ExampleGNAPError decodes a json data into GNAPError and
// checks that error is equivalent to ErrGInvalidClient.
func ExampleGNAPError() {
	type response struct {
		Token string    `json:"token,omitempty"`
		Error GNAPError `json:"error,omitempty"`
	}
	data := []byte(`{"error":{"code": "invalid_client","description":"client not recognized"}}`)
	var res response
	_ = json.Unmarshal(data, &res)
	fmt.Println(res.Error)
	fmt.Println(errors.Is(res.Error, ErrGInvalidClient))
	// Output:
	// invalid_client: client not recognized
	// true
}
