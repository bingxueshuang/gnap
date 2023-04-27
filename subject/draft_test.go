package subject_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bingxueshuang/gnap/subject"
)

// testcase handles table driven testing
type testcase struct {
	in      []byte     // json data
	want    subject.ID // expected Subject Identifier
	wantErr bool       // if error is expected
}

// testvector is a list of test cases obtained from
// the examples in the draft document.
var testvector = []testcase{
	// Account Identifier Format
	{
		in: []byte(`{
			"format": "account",
			"uri": "acct:example.user@service.example.com"
		  }`),
		want: subject.ID{
			Format: subject.Account,
			URI:    "acct:example.user@service.example.com",
		},
	},
	// Email Identifier Format
	{
		in: []byte(`{
			"format": "email",
			"email": "user@example.com"
		  }`),
		want: subject.ID{
			Format: subject.Email,
			Email:  "user@example.com",
		},
	},
	// Issuer and Subject Identifier Format
	{
		in: []byte(`{
			"format": "iss_sub",
			"iss": "https://issuer.example.com/",
			"sub": "145234573"
		  }`),
		want: subject.ID{
			Format:  subject.IssuerSubject,
			Issuer:  "https://issuer.example.com/",
			Subject: "145234573",
		},
	},
	// Opaque Identifier Format
	{
		in: []byte(`{
			"format": "opaque",
			"id": "11112222333344445555"
		  }`),
		want: subject.ID{
			Format: subject.Opaque,
			ID:     "11112222333344445555",
		},
	},
	// Phone Number Identifier Format
	{
		in: []byte(`{
			"format": "phone_number",
			"phone_number": "+12065550100"
		  }`),
		want: subject.ID{
			Format: subject.PhoneNumber,
			Phone:  "+12065550100",
		},
	},
	// Decentralized Identifier Format with a bare DID
	{
		in: []byte(`{
			"format": "did",
			"url": "did:example:123456"
		  }`),
		want: subject.ID{
			Format: subject.DID,
			URL:    "did:example:123456",
		},
	},
	// Decentralized Identifier Format with a DID URL
	// with non-empty path and query components
	{
		in: []byte(`{
			"format": "did",
			"url": "did:example:123456/did/url/path?versionId=1"
		  }`),
		want: subject.ID{
			Format: subject.DID,
			URL:    "did:example:123456/did/url/path?versionId=1",
		},
	},
	// URI Format with a website URI
	{
		in: []byte(`{
			"format": "uri",
			"uri": "https://user.example.com/"
		  }`),
		want: subject.ID{
			Format: subject.URI,
			URI:    "https://user.example.com/",
		},
	},
	// URI Format with a random URN
	{
		in: []byte(`{
			"format": "uri",
			"uri": "urn:uuid:4e851e98-83c4-4743-a5da-150ecb53042f"
		  }`),
		want: subject.ID{
			Format: subject.URI,
			URI:    "urn:uuid:4e851e98-83c4-4743-a5da-150ecb53042f",
		},
	},
	// Aliases Identifier Format
	{
		in: []byte(`{
			"format": "aliases",
			"identifiers": [
			  {
				"format": "email",
				"email": "user@example.com"
			  },
			  {
				"format": "phone_number",
				"phone_number": "+12065550100"
			  },
			  {
				"format": "email",
				"email": "user+qualifier@example.com"
			  }
			]
		  }`),
		want: subject.ID{
			Format: subject.Aliases,
			Identifiers: []subject.NoAlias{
				{Format: subject.Email, Email: "user@example.com"},
				{Format: subject.PhoneNumber, Phone: "+12065550100"},
				{Format: subject.Email, Email: "user+qualifier@example.com"},
			},
		},
	},
}

func TestID(t *testing.T) {
	for i, tt := range testvector {
		var id subject.ID
		err := json.Unmarshal(tt.in, &id)
		if tt.wantErr != (err != nil) {
			t.Errorf("at %v: error = %v, wantErr = %v", i, err, tt.wantErr)
		}
		if !reflect.DeepEqual(id, tt.want) {
			t.Errorf("at %v: id = %v, want %v", i, id, tt.want)
		}
	}
}
