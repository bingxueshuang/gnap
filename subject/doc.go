// Package subject provides models for **Subject Identifiers for Security Event Tokens**
// as defined in draft-ietf-secevent-subject-identifiers-16. See
// https://datatracker.ietf.org/doc/draft-ietf-secevent-subject-indentifiers/ for
// the latest draft version.

// For creating a new Subject Identifier, use constructor functions matching:
//
//	NewID(format)?
//
// where `format` is a subject id format. For more complicated usages, you can directly
// use [ID] as struct literal. For including list of identifiers in aliases format,
// ID.NoAlias helper function can be used. See the example of [ID] type for more details.
package subject // import "github.com/bingxueshuang/gnap/subject"
