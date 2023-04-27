package subject_test

import (
	"testing"

	"github.com/bingxueshuang/gnap/subject"
)

func TestEqualNoAlias(t *testing.T) {
	tests := []struct {
		name string
		a    subject.NoAlias
		b    subject.NoAlias
		want bool
	}{
		{
			name: "equal",
			a:    subject.NoAlias{Format: subject.Email, Email: "test@example.com"},
			b:    subject.NoAlias{Format: subject.Email, Email: "test@example.com"},
			want: true,
		},
		{
			name: "format",
			a:    subject.NoAlias{Format: subject.Account},
			b:    subject.NoAlias{Format: subject.Opaque},
		},
		{
			name: "account",
			a:    subject.NoAlias{Format: subject.Account, URI: "acct:example.com/user"},
			b:    subject.NoAlias{Format: subject.Account, URI: "acct:example.com/admin"},
		},
		{
			name: "email",
			a:    subject.NoAlias{Format: subject.Email, Email: "user@example.com"},
			b:    subject.NoAlias{Format: subject.Email, Email: "admin@example.com"},
		},
		{
			name: "iss-sub",
			a: subject.NoAlias{
				Format:  subject.IssuerSubject,
				Issuer:  "https://example.com/",
				Subject: "f7e9153",
			},
			b: subject.NoAlias{
				Format:  subject.IssuerSubject,
				Issuer:  "https://example.com/",
				Subject: "k3jl4h",
			},
		},
		{
			name: "opaque",
			a:    subject.NoAlias{Format: subject.Opaque, ID: "lkjlk56j7"},
			b:    subject.NoAlias{Format: subject.Opaque, ID: "lkjl4576j7"},
		},
		{
			name: "phone",
			a:    subject.NoAlias{Format: subject.PhoneNumber, Phone: "9184749907"},
			b:    subject.NoAlias{Format: subject.PhoneNumber, Phone: "918567907"},
		},
		{
			name: "did",
			a:    subject.NoAlias{Format: subject.DID, URL: "did:example.com/user"},
			b:    subject.NoAlias{Format: subject.DID, URL: "did:example.com/admin"},
		},
		{
			name: "uri",
			a:    subject.NoAlias{Format: subject.URI, URI: "https://example.com/user"},
			b:    subject.NoAlias{Format: subject.URI, URI: "https://example.com/admin"},
		},
		{
			name: "invalid",
			a:    subject.NoAlias{Format: "wrong"},
			b:    subject.NoAlias{Format: "wrong"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.EqualNoAlias(tt.a, tt.b); got != tt.want {
				t.Errorf("EqualNoAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    subject.ID
		b    subject.ID
		want bool
	}{
		{
			name: "noalias",
			a:    subject.ID{Format: subject.Email, Email: "foo@bar.com"},
			b:    subject.ID{Format: subject.Email, Email: "foo@bar.com"},
			want: true,
		},
		{
			name: "alias",
			a: subject.ID{
				Format: subject.Aliases,
				Identifiers: []subject.NoAlias{
					{Format: subject.Email, Email: "foo@bar.com"},
				},
			},
			b: subject.ID{
				Format: subject.Aliases,
				Identifiers: []subject.NoAlias{
					{Format: subject.Email, Email: "foo@bar.com"},
				},
			},
			want: true,
		},
		{
			name: "invalid",
			a:    subject.ID{Format: "wrong"},
			b:    subject.ID{Format: "invalid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.Equal(tt.a, tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
