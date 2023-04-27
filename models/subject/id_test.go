package subject_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/bingxueshuang/gnap/models/subject"
	"golang.org/x/exp/slices"
)

func ExampleID() {
	var testid subject.ID
	data := []byte(`{
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
			"email": "user+word@example.com"
		  }
		]
	  }`)
	_ = json.Unmarshal(data, &testid)
	personal, _ := subject.NewIDEmail("user@example.com")
	number, _ := subject.NewIDPhoneNumber("+12065550100")
	workmail, _ := subject.NewIDEmail("user+word@example.com")
	newid, _ := subject.NewIDAliases([]subject.NoAlias{
		personal.NoAlias(),
		number.NoAlias(),
		workmail.NoAlias(),
	})
	fmt.Println(subject.Equal(newid, testid))
	// Output:
	// true
}

func TestNoAlias_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    subject.NoAlias
		wantErr bool
	}{
		{
			name: "valid",
			in: []byte(`{
				"format": "email",
				"email": "user@example.com"
			  }`),
			want: subject.NoAlias{
				Format: subject.Email,
				Email:  "user@example.com",
			},
		},
		{
			name:    "invalid json",
			in:      []byte(`}`),
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`{}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new(subject.NoAlias)
			err := got.UnmarshalJSON(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoAlias.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if *got != tt.want {
				t.Errorf("NoAlias.UnmarshalJSON() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    subject.ID
		wantErr bool
	}{
		{
			name: "valid",
			in: []byte(`{
				"format": "email",
				"email": "user@example.com"
			  }`),
			want: subject.ID{
				Format: subject.Email,
				Email:  "user@example.com",
			},
		},
		{
			name: "aliases",
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
		{
			name:    "invalid json",
			in:      []byte(`}`),
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`{}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new(subject.ID)
			err := got.UnmarshalJSON(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoAlias.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NoAlias.UnmarshalJSON() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestNoAlias_SubjectID(t *testing.T) {
	tests := []struct {
		name string
		in   subject.NoAlias
		want subject.ID
	}{
		{
			name: "valid",
			in: subject.NoAlias{
				Format: subject.Email,
				Email:  "valid@example.com",
			},
			want: subject.ID{
				Format: subject.Email,
				Email:  "valid@example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.SubjectID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoAlias.SubjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_NoAlias(t *testing.T) {
	tests := []struct {
		name string
		in   subject.ID
		want subject.NoAlias
	}{
		{
			name: "valid",
			in: subject.ID{
				Format: subject.Email,
				Email:  "valid@example.com",
			},
			want: subject.NoAlias{
				Format: subject.Email,
				Email:  "valid@example.com",
			},
		},
		{
			name: "alias",
			in: subject.ID{
				Format: subject.Aliases,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.NoAlias(); got != tt.want {
				t.Errorf("ID.NoAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoAlias_Validate(t *testing.T) {
	tests := []struct {
		name    string
		in      subject.NoAlias
		wantErr bool
	}{
		{
			name: "valid",
			in: subject.NoAlias{
				Format: subject.Email,
				Email:  "test@example.com",
			},
		},
		{
			name: "fmt",
			in: subject.NoAlias{
				Format: "wrong",
			},
			wantErr: true,
		},
		{
			name: "invalid",
			in: subject.NoAlias{
				Format: subject.PhoneNumber,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.in.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("NoAlias.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestID_Validate(t *testing.T) {
	tests := []struct {
		name    string
		in      subject.ID
		wantErr bool
	}{
		{
			name: "valid NoAlias",
			in: subject.ID{
				Format: subject.Email,
				Email:  "test@example.com",
			},
		},
		{
			name: "fmt",
			in: subject.ID{
				Format: "wrong",
			},
			wantErr: true,
		},
		{
			name: "valid Aliases",
			in: subject.ID{
				Format: subject.Aliases,
				Identifiers: []subject.NoAlias{
					{Format: subject.Email, Email: "valid@example.com"},
				},
			},
		},
		{
			name: "invalid aliases",
			in: subject.ID{
				Format: subject.Aliases,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.in.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("NoAlias.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewIDAccount(t *testing.T) {
	tests := []struct {
		name    string
		acc     string
		wantErr bool
	}{
		{
			name: "valid",
			acc:  "acct:example.user@service.example.com",
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid uri",
			acc:     "https://google%30.com/path",
			wantErr: true,
		},
		{
			name:    "invalid scheme",
			acc:     "http://example.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDAccount(tt.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.URI != tt.acc {
				t.Errorf("NewIDAccount() URI = %v, want %v", got.URI, tt.acc)
			}
		})
	}
}

func TestNewIDEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:  "valid",
			email: "valid@example.com",
		},
		{
			name:    "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.Email != tt.email {
				t.Errorf("NewIDEmail() = %v, want %v", got.Email, tt.email)
			}
		})
	}
}

func TestNewIDIssuerSubject(t *testing.T) {
	tests := []struct {
		name    string
		iss     string
		sub     string
		wantErr bool
	}{
		{
			name: "valid",
			iss:  "https://issuer.example.com/",
			sub:  "145234573",
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid uri",
			iss:     "https://exa%32.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDIssuerSubject(tt.iss, tt.sub)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDIssuerSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.Issuer != tt.iss || got.Subject != tt.sub {
				type isssub struct {
					iss string
					sub string
				}
				t.Errorf(
					"NewIDIssuerSubject() = %v, want %v",
					isssub{got.Issuer, got.Subject},
					isssub{tt.iss, tt.sub},
				)
			}
		})
	}
}

func TestNewIDOpaque(t *testing.T) {
	tests := []struct {
		name    string
		opaque  string
		wantErr bool
	}{
		{
			name:   "valid",
			opaque: "11112222333344445555",
		},
		{
			name:    "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDOpaque(tt.opaque)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDOpaque() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.ID != tt.opaque {
				t.Errorf("NewIDOpaque() = %v, want %v", got.ID, tt.opaque)
			}
		})
	}
}

func TestNewIDPhoneNumber(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{
			name:  "valid",
			phone: "+12065550100",
		},
		{
			name:    "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDPhoneNumber(tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.Phone != tt.phone {
				t.Errorf("NewIDPhoneNumber() = %v, want %v", got.Phone, tt.phone)
			}
		})
	}
}

func TestNewIDdid(t *testing.T) {
	tests := []struct {
		name    string
		did     string
		wantErr bool
	}{
		{
			name: "valid bare",
			did:  "did:example:123456",
		},
		{
			name: "valid params",
			did:  "did:example:123456/did/url/path?versionId=1",
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid url",
			did:     "https://exampl%30.com/profile",
			wantErr: true,
		},
		{
			name:    "invalid scheme",
			did:     "https://exampl.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDdid(tt.did)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDdid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.URL != tt.did {
				t.Errorf("NewIDdid() = %v, want %v", got.URL, tt.did)
			}
		})
	}
}

func TestNewIDuri(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		wantErr bool
	}{
		{
			name: "valid uri",
			uri:  "https://user.example.com/",
		},
		{
			name: "valid urn",
			uri:  "urn:uuid:4e851e98-83c4-4743-a5da-150ecb53042f",
		},
		{
			name:    "empty",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDuri(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDuri() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.URI != tt.uri {
				t.Errorf("NewIDuri() = %v, want %v", got.URI, tt.uri)
			}
		})
	}
}

func TestNewIDAliases(t *testing.T) {
	tests := []struct {
		name    string
		aliases []subject.NoAlias
		wantErr bool
	}{
		{
			name: "valid",
			aliases: []subject.NoAlias{
				{Format: subject.Email, Email: "user@example.com"},
				{Format: subject.PhoneNumber, Phone: "+12065550100"},
				{Format: subject.Email, Email: "user+qualifier@example.com"},
			},
		},
		{
			name: "duplicates",
			aliases: []subject.NoAlias{
				{Format: subject.Email, Email: "user@example.com"},
				{Format: subject.Email, Email: "user@example.com"},
			},
			wantErr: true,
		},
		{
			name: "invalid",
			aliases: []subject.NoAlias{
				{},
			},
			wantErr: true,
		},
		{
			name:    "empty",
			aliases: []subject.NoAlias{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subject.NewIDAliases(tt.aliases)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDAliases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !slices.Equal(got.Identifiers, tt.aliases) {
				t.Errorf("NewIDAliases() = %v, want %v", got.Identifiers, tt.aliases)
			}
		})
	}
}
