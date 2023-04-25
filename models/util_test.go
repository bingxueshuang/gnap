package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func TestURL_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      URL
		out     []byte
		wantErr bool
	}{
		{
			name: "valid",
			in: URL{&url.URL{
				Scheme:   "http",
				Host:     "foo.com:8000",
				Path:     "/path",
				RawQuery: "query=value",
				Fragment: "hash",
			}},
			out: []byte(`"http://foo.com:8000/path?query=value#hash"`),
		},
		{
			name:    "nil",
			in:      URL{},
			wantErr: true,
		},
		{
			name:    "empty",
			in:      URL{&url.URL{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("URL.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.out) {
				t.Errorf("URL.MarshalJSON() = %v, want %v", got, tt.out)
				return
			}
		})
	}
}

func TestURL_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		out     *URL
		wantErr bool
	}{
		{
			name: "valid",
			in:   []byte(`"https://example.com/long/path/name/"`),
			out: &URL{&url.URL{
				Scheme: "https",
				Host:   "example.com",
				Path:   "/long/path/name/",
			}},
		},
		{
			name:    "json",
			in:      []byte(`"https://example.com`),
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`""`),
			wantErr: true,
		},
		{
			name:    "empty",
			in:      []byte(`"http://192.168.0.%31/"`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new(URL)
			err := got.UnmarshalJSON(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.out != nil && (got == nil || tt.out.String() != got.String()) {
				t.Errorf("URL.UnmarshalJSON() = %v, want %v", got, tt.out)
				return
			}
		})
	}
}

func TestParseURL(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{
			name: "valid",
			in:   "http://example.com/path?query=value#hash",
		},
		{
			name:    "invalid",
			in:      `http://[fe80::%31%25en0]/`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.in != got.String() {
				t.Errorf("ParseURL() = %v, want %v", got, tt.in)
			}
		})
	}
}

func ExampleURL() {
	type request struct {
		URL  URL    `json:"url"`
		Name string `json:"name"`
	}
	data := `{"name":"GitHub","url":"https://github.com/"}`
	var req request
	_ = json.Unmarshal([]byte(data), &req)
	fmt.Println(req.Name, req.URL)
	// Output: GitHub https://github.com/
}
