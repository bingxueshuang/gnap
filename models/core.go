package models

import (
	"fmt"

	"github.com/bingxueshuang/gnap/subject"
)

// Discovery represents the server's discovery information.
type Discovery struct {
	GrantRequest  URL               `json:"grant_request_endpoint"`
	StartModes    []StartMode       `json:"interaction_start_modes_supported,omitempty"`
	FinishMethods []FinishMethod    `json:"interaction_finish_methods_supported,omitempty"`
	KeyProofs     []ProofMethod     `json:"key_proofs_supported,omitempty"`
	SubFormats    []subject.Format  `json:"sub_id_formats_supported,omitempty"`
	AFormats      []AssertionFormat `json:"assertion_formats_supported,omitempty"`
	KeyRotation   bool              `json:"key_rotation_supported,omitempty"`
}

// GrantRequest represents the grant request for initiation
// of the gnap flow.
type GrantRequest struct {
	AccessToken ATRequest      `json:"access_token,omitempty"`
	Subject     SubRequest     `json:"subject,omitempty"`
	Client      ClientInstance `json:"client"`
	User        EndUser        `json:"user,omitempty"`
	Interact    IARequest      `json:"interact,omitempty"`
}

// NewRequest is constructor for [GrantRequest] with mandatory client
// and optional parameters.
func NewRequest(client ClientInstance, options ...requestOption) (req GrantRequest, err error) {
	g := &GrantRequest{Client: client}
	for _, setter := range options {
		err = setter(g)
		if err != nil {
			return
		}
	}
	return *g, nil
}

// requestOption is functional parameter for request constructor.
type requestOption func(*GrantRequest) error

// GrantResponse represents the AS response to a grant request.
type GrantResponse struct {
	Continue    ContinueResponse `json:"continue,omitempty"`
	AccessToken ATResponse       `json:"access_token,omitempty"`
	Interact    IAResponse       `json:"interact,omitempty"`
	Subject     SubResponse      `json:"subject,omitempty"`
	InstanceID  string           `json:"instance_id,omitempty"`
	Error       GNAPError        `json:"error,omitempty"`
}

// NewRequest is constructor for [GrantResponse] with optional parameters.
func NewResponse(options ...responseOption) (res GrantResponse, err error) {
	g := &GrantResponse{}
	for _, setter := range options {
		err = setter(g)
		if err != nil {
			return
		}
	}
	return *g, nil
}

// responseOption is functional parameter for response constructor.
type responseOption func(*GrantResponse) error

// ContinueRequest represents the continuation request
// sent by the client instance after successful interaction.
type ContinueRequest struct {
	InteractRef string `json:"interact_ref"`
}

// ContinueResponse represents the continuation object
// returned by the AS during the gnap request flow.
type ContinueResponse struct {
	URI   URL           `json:"uri"`
	Wait  int           `json:"wait"`
	Token ContinueToken `json:"access_token"`
}

// SingleToken is an optional parameter for [NewRequest]
// to request for single access token.
func SingleToken(token TokenRequest) requestOption {
	return func(req *GrantRequest) error {
		req.AccessToken = ATRequest{Single: token}
		return nil
	}
}

// MultiToken is an optional parameter for [NewRequest]
// to request for multiple access tokens.
func MultiToken(tokens ...TokenRequest) requestOption {
	return func(req *GrantRequest) error {
		set := make(map[string]struct{})
		for i := range tokens {
			label := tokens[i].Label
			if label == "" {
				return fmt.Errorf("missing label: %w", ErrInvalidTokenRequest)
			}
			_, ok := set[label]
			if ok { // duplicate
				return fmt.Errorf("duplicate label: %w", ErrInvalidTokenRequest)
			}
			set[label] = struct{}{}
		}
		req.AccessToken = ATRequest{Multiple: tokens}
		return nil
	}
}

// WithSubject is an optional parameter for [NewRequest]
// to request for subject information.
func WithSubject(sub SubRequest) requestOption {
	return func(req *GrantRequest) error {
		req.Subject = sub
		return nil
	}
}

// WithUser is an optional parameter for [NewRequest]
// to identify the RO to the AS.
func WithUser(user EndUser) requestOption {
	return func(req *GrantRequest) error {
		req.User = user
		return nil
	}
}

// WithInteract is an optional parameter for [NewRequest]
// to convey mode of interaction.
func WithInteract(ia IARequest) requestOption {
	return func(req *GrantRequest) error {
		req.Interact = ia
		return nil
	}
}

// WithSingleResponse is an optional parameter for [NewResponse]
// to grant single access token.
func WithSingleResponse(token TokenResponse) responseOption {
	return func(res *GrantResponse) error {
		res.AccessToken = ATResponse{Single: token}
		return nil
	}
}

// WithSingleResponse is an optional parameter for [NewResponse]
// to grant multiple access token.
func WithMultiResponse(tokens ...TokenResponse) responseOption {
	return func(res *GrantResponse) error {
		set := make(map[string]struct{})
		for i := range tokens {
			label := tokens[i].Label
			if label == "" {
				return fmt.Errorf("missing label: %w", ErrInvalidTokenResponse)
			}
			_, ok := set[label]
			if ok { // duplicate
				return fmt.Errorf("duplicate label: %w", ErrInvalidTokenResponse)
			}
		}
		res.AccessToken = ATResponse{Multiple: tokens}
		return nil
	}
}

// WithSubjectResponse is an optional parameter for [NewResponse]
// to convey subject information to the client.
func WithSubjectResponse(sub SubResponse) responseOption {
	return func(res *GrantResponse) error {
		res.Subject = sub
		return nil
	}
}

// WithContinue is an optional parameter for [NewResponse]
// to convey the means of flow continuation to the client.
func WithContinue(con ContinueResponse) responseOption {
	return func(res *GrantResponse) error {
		res.Continue = con
		return nil
	}
}

// WithInteractResponse is an optional parameter for [NewResponse]
// to convey the interaction urls to the client.
func WithInteractResponse(ia IAResponse) responseOption {
	return func(res *GrantResponse) error {
		res.Interact = ia
		return nil
	}
}

// WithError is an optional parameter for [NewResponse]
// to respond with an error to the client.
func WithError(err GNAPError) responseOption {
	return func(res *GrantResponse) error {
		res.Error = err
		return nil
	}
}

// WithInstanceID is an optional parameter for [NewResponse]
// to attribute a unique instance ID to the ongoing GNAP flow.
func WithInstanceID(id string) responseOption {
	return func(res *GrantResponse) error {
		res.InstanceID = id
		return nil
	}
}
