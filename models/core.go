package models

import "github.com/bingxueshuang/gnap/subject"

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

// GrantResponse represents the AS response to a grant request.
type GrantResponse struct {
	Continue    ContinueResponse `json:"continue,omitempty"`
	AccessToken ATResponse       `json:"access_token,omitempty"`
	Interact    IAResponse       `json:"interact,omitempty"`
	Subject     SubResponse      `json:"subject,omitempty"`
	InstanceID  string           `json:"instance_id,omitempty"`
	Error       GNAPError        `json:"error,omitempty"`
}

// ContinueRequest represents the continuation request
// sent by the client instance after successful interaction.
type ContinueRequest struct {
	InteractRef string `json:"interact_ref"`
}

// ContinueResponse represents the continuation object
// returned by the AS during the gnap request flow.
type ContinueResponse struct {
	URI         URL           `json:"uri"`
	Wait        int           `json:"wait"`
	AccessToken ContinueToken `json:"access_token"`
}
