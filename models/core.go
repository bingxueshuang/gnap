package models

import "github.com/bingxueshuang/gnap/subject"

// Request represents the JSON document that the client instance sends
// to start a GNAP request at the grant endpoint of the AS. The document
// is a JSON object where each field represents a different aspect of the
// client instance's request.
type Request struct {
	// rights and properties associated with the requested access token.
	// REQUIRED if requesting an access token.
	Access ATReq `json:"access_token,omitempty"`
	// info about the RO  to be returned directly in the response
	// from the AS.
	// REQUIRED if requesting subject information.
	Subject SubReq `json:"subject,omitempty"`
	// client instance indentification, including the key that the client
	// instance will use to protect requests at the AS and any user-facing
	// information about the client instance used in interactions.
	Client Client `json:"client"` // REQUIRED
	// end user identification to the AS in a manner that the AS can verify.
	User User `json:"user,omitempty"` // OPTIONAL
	// interaction modes that the client instance supports for allowing the RO
	// to interact with the AS and modes for the client instance to receive
	// updates when interaction is complete.
	// REQUIRED if interaction is supported.
	Interact IAReq `json:"interact,omitempty"`
}

// Response represents the JSON object that the AS sends to the client
// instance in response to a GNAP request.
type Response struct {
	// indicates that the client instance can continue the request by making
	// one or more continuation requests.
	// REQUIRED if continuation calls are allowed for this client instance
	// on this grant request.
	Continue ContResponse `json:"continue,omitempty"`
	// single access token or set of access tokens that the client instance
	// can use to call the RS on behalf of the RO.
	// REQUIRED if an access token is included.
	Access ATResponse `json:"access_token,omitempty"`
	// indicates that interaction through some set of defined mechanisms
	// needs to take place.
	// REQUIRED if interaction is expected.
	Interact IAResponse `json:"interact,omitempty"`
	// claims about the RO as known and declared by the AS.
	// REQUIRED if subject information is included.
	Subject SubResponse `json:"subject,omitempty"`
	// identifier this client instance can use to identify itself when
	// making future requests.
	InstanceID string `json:"instance_id,omitempty"` // OPTIONAL
	// error code indicating that something has gone wrong.
	// REQUIRED for an error condition.
	Error Error `json:"error,omitempty"`
}

// ContReq defines the continuation request at the AS after
// a successful interaction with the RO.
type ContReq struct {
	// the interaction reference obtained during interaction finish.
	IARef string `json:"interact_ref"` // REQUIRED
}

// ContResponse defines the "continue" field of the [Response] object.
// If the AS determines that the grant request can be continued by
// the client instance, the AS responds with the continue field.
type ContResponse struct {
	// URI at which the client instance can make continuation requests.
	// This URI MUST be an an absolute URI.
	URI URL `json:"uri"` // REQUIRED
	// The amount of time in integer seconds the client instance MUST wait
	// after receiving this response and calling the continuation URI.
	Wait int `json:"wait"` // RECOMMENDED
	// A unique access token for continuing the request, called the
	// "continuation access token".
	Token ContToken `json:"access_token"` // REQUIRED
}

// Discovery defines the json object format with which the client MAY
// send an HTTP OPTIONS request to the grant request endpoint to retrieve
// server's discovery information.
type Discovery struct {
	// location of the AS's grant request endpoint. The location MUST be an absolute
	// URL [RFC3986] with a scheme component (which MUST be "https"), a host component,
	// and optionally, port, path and query components and no fragment components.
	// This URL MUST match the URL the client instance used to make the discovery request.
	//
	// [RFC3986]: https://www.rfc-editor.org/rfc/rfc3986
	Endpoint URL `json:"grant_request_endpoint"` // REQUIRED
	// list of the AS's interaction start methods.
	StartModes []StartMode `json:"interaction_start_modes_supported,omitempty"` // OPTIONAL
	// list of the AS's interaction finish methods.
	FinishMethods []FinishMethod `json:"interaction_finish_methods_supported,omitempty"` // OPTIONAL
	// list of the AS's supported key proofing mechanisms.
	KeyProofs []ProofMethod `json:"key_proofs_supported,omitempty"` // OPTIONAL
	// list of the AS's supported subject identifier formats.
	SFormats []subject.Format `json:"sub_id_formats_supported,omitempty"` // OPTIONAL
	// list of the AS's supported assertion formats.
	AFormats []AFormat `json:"assertion_formats_supported,omitempty"` // OPTIONAL
	// The boolean "true" indicates that rotation of access token bound keys by
	// the client is supported by the AS. The absence of this field or a boolean
	// "false" value indicates that this feature is not supported.
	KeyRotation bool `json:"key_rotation_supported,omitempty"` // OPTIONAL
}
