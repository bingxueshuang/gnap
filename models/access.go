package models

// ATReq represents the "access_token" field of the [Request], which MUST be included
// if the client instance is requesting one or more access tokens for the purpose
// of accessing a protected resource (such as API).
type ATReq struct {
	Single TokenReq   `json:"-"`
	Multi  []TokenReq `json:"-"`
}

// ATResponse represents the "access_token" field of the [Response], which MUST be
// included if the AS has successfully granted one or more access tokens to the
// client instance.
type ATResponse struct {
	Single TokenResponse   `json:"-"`
	Multi  []TokenResponse `json:"-"`
}

// ARight represent rights of access that are associated with the the access token.
// Rights of access can be defined by the RS as either an object or a string.
type ARight struct {
	// type of resource request as a string. This field MAY define which
	// other fields are allowed in the request object.
	Type string `json:"type"` // REQUIRED
	// types of actions the client instance will take at the RS
	Actions []string `json:"actions,omitempty"` // OPTIONAL
	// typically URIs identifying the location of the RS.
	Locations []string `json:"locations,omitempty"` // OPTIONAL
	// kinds of data available to the client instance at the RS's API
	Datatypes []string `json:"datatypes,omitempty"` // OPTIONAL
	// string identifier indicating a specific resource at the RS.
	Identifier string `json:"identifier,omitempty"` // OPTIONAL
	// types or levels of privilege being requested at the resource.
	Privileges []string `json:"privileges,omitempty"` // OPTIONAL
	// access rights MAY be communicated as a string known to the AS
	// representing the access being requested. such refs indicate a
	// specific access at a protected resource.
	Ref string `json:"-"`
}
