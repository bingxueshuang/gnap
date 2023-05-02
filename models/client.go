package models

import (
	"github.com/bingxueshuang/gnap/subject"
)

// Client represents the object used by the client instance to identify
// itself by including its client information in the client field of the
// request and by signing the request with its unique key (during new grant request).
type Client struct {
	// public key of the client instance to be used in this request
	// or a reference to a key.
	Key Key `json:"key"` // REQUIRED
	// identifier string that the AS can use to identify the
	// client software comprising this client instance.
	// The contents and format of this field are up to the AS.
	ClassID string `json:"class_id"` // OPTIONAL
	// object containing additional information that the AS MAY display
	// to the RO during interaction, authorization, and management.
	Display Display `json:"display"` // OPTIONAL
	// if the client instance has an instance identifier that the AS can use to
	// determine appropriate key information, the client instance can send this
	// instance identifier as a direct reference value in lieu of the client object.
	// The instance identifier MAY be assigned to a client instance at runtime
	// through a grant response or MAY be obtained in another fashion,
	// such as a static registration process at the AS.
	// REQUIRED if client is identified by reference.
	Ref string `json:"-"`
}

// Display represents the "display" field of the [Client] object sent by the
// client instance if it has additional information to display to the RO during
// any interactions at the AS.
type Display struct {
	// display name of the client software.
	Name string `json:"name"` // RECOMMENDED
	// user-facing info about the client software, such as a web page.
	// This URI MUST be an absolute URI.
	URI URL `json:"uri,omitempty"` // OPTIONAL
	// display image of the client software. This URI MUST be an absolute URI.
	// The logo MAY be passed by value by using a data: URI [RFC2397]
	// referencing an image mediatype.
	//
	// [RFC2397]: https://www.rfc-editor.org/rfc/rfc2397
	Logo URL `json:"logo_uri,omitempty"` // OPTIONAL
}

// User defines the object for "user" field of [Request] to the AS. This field
// is sent if the client instance knows the identity of the end user through
// one or more identifiers or assertions.
type User struct {
	// array of subject identifiers for the end user.
	SubIDs []subject.ID `json:"sub_ids,omitempty"` // OPTIONAL
	// array containing identity assertions as objects regarding the end user.
	Assertions []Assertion `json:"assertions,omitempty"` // OPTIONAL
	// a boolean field to differentiate between zero value and set value.
	// MUST be set to true if User is not zero value.
	NonZero bool `json:"-"`
	// If AS can identify the current end user to the client instance with
	// a reference, it can be used by the client instance to refer to the
	// end user across multiple requests. If the client instance has a reference
	// for the end user at this AS, the client instance MAY pass that reference
	// as a string. The format of this string is opaque to the client instance.
	// REQUIRED if user is identified by reference.
	Ref string `json:"-"`
}
