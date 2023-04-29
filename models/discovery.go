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
