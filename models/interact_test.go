package models

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
)

func TestStartMode_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sm      StartMode
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			sm:   ModeCode,
			want: []byte(`"user_code"`),
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid",
			sm:      StartMode("invalid_mode"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sm.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("StartMode.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("StartMode.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartMode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    StartMode
		wantErr bool
	}{
		{
			name: "valid",
			in:   []byte(`"redirect"`),
			want: ModeRedirect,
		},
		{
			name:    "json",
			in:      []byte(`{}`),
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`"wrong_mode"`),
			wantErr: true,
		},
		{
			name:    "empty",
			in:      nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got StartMode
			err := json.Unmarshal(tt.in, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartMode.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("StartMode.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinishMethod_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		fm      FinishMethod
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			fm:   MethodPush,
			want: []byte(`"push"`),
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid",
			fm:      FinishMethod("wrong_method"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fm.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("FinishMethod.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("FinishMethod.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinishMethod_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    FinishMethod
		wantErr bool
	}{
		{
			name: "valid",
			in:   []byte(`"redirect"`),
			want: MethodRedirect,
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`"wrong_method"`),
			wantErr: true,
		},
		{
			name:    "json",
			in:      []byte(`{}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FinishMethod
			err := json.Unmarshal(tt.in, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("FinishMethod.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FinishMethod.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIAStart_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      IAStart
		want    []byte
		wantErr bool
	}{
		{
			name: "valid val",
			in:   IAStart{Mode: ModeRedirect},
			want: []byte(`{"mode":"redirect"}`),
		},
		{
			name: "valid ref",
			in:   IAStart{ModeCode, true},
			want: []byte(`"user_code"`),
		},
		{
			name:    "empty val",
			in:      IAStart{},
			wantErr: true,
		},
		{
			name:    "empty ref",
			in:      IAStart{IsRef: true},
			wantErr: true,
		},
		{
			name:    "invalid val",
			in:      IAStart{Mode: "wrong_mode"},
			wantErr: true,
		},
		{
			name:    "invalid ref",
			in:      IAStart{"wrong_mode", true},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("IAStart.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("IAStart.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIAStart_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    IAStart
		wantErr bool
	}{
		{
			name: "valid ref",
			in:   []byte(`"user_code_uri"`),
			want: IAStart{ModeCodeURI, true},
		},
		{
			name: "valid val",
			in:   []byte(`{"mode":"user_code"}`),
			want: IAStart{Mode: ModeCode},
		},
		{
			name:    "json ref",
			in:      []byte(`"something`),
			wantErr: true,
		},
		{
			name:    "json val",
			in:      []byte(`{"mode":"anything`),
			wantErr: true,
		},
		{
			name:    "empty ref",
			in:      []byte(`""`),
			wantErr: true,
		},
		{
			name:    "empty val",
			in:      []byte(`{"mode":""}`),
			wantErr: true,
		},
		{
			name:    "invalid ref",
			in:      []byte(`"wrong_mode"`),
			wantErr: true,
		},
		{
			name:    "invalid val",
			in:      []byte(`{"mode":"wrong_mode"}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got IAStart
			err := json.Unmarshal(tt.in, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("IAStart.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("IAStart.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIACallback_Encode(t *testing.T) {
	equal := func(a, b url.Values) bool {
		return a.Get("hash") == b.Get("hash") &&
			a.Get("interact_ref") == b.Get("interact_ref")
	}
	tests := []struct {
		name string
		in   IACallback
		want url.Values
	}{
		{
			name: "hash,ref",
			in:   IACallback{"g4kb5kj3b56kn234", "9128347"},
			want: url.Values{
				"hash":         []string{"g4kb5kj3b56kn234"},
				"interact_ref": []string{"9128347"},
			},
		},
		{
			name: "hash",
			in:   IACallback{Hash: "g4kb5kj3b56kn234"},
			want: url.Values{
				"hash":         []string{"g4kb5kj3b56kn234"},
				"interact_ref": []string{""},
			},
		},
		{
			name: "ref",
			in:   IACallback{InteractRef: "9128347"},
			want: url.Values{
				"hash":         []string{""},
				"interact_ref": []string{"9128347"},
			},
		},
		{
			name: "empty",
			in:   IACallback{},
			want: url.Values{
				"hash":         []string{""},
				"interact_ref": []string{""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.Encode()
			if !equal(got, tt.want) {
				t.Errorf("IACallback.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromQuery(t *testing.T) {
	tests := []struct {
		name string
		in   url.Values
		want IACallback
	}{
		{
			name: "hash,ref",
			in: url.Values{
				"hash":         []string{"k1l2j34hh0sd9f0y"},
				"interact_ref": []string{"92562837"},
			},
			want: IACallback{"k1l2j34hh0sd9f0y", "92562837"},
		},
		{
			name: "hash",
			in: url.Values{
				"hash":         []string{"k1l2j34hh0sd9f0y"},
				"interact_ref": []string{""},
			},
			want: IACallback{Hash: "k1l2j34hh0sd9f0y"},
		},
		{
			name: "ref",
			in: url.Values{
				"hash":         []string{""},
				"interact_ref": []string{"92562837"},
			},
			want: IACallback{InteractRef: "92562837"},
		},
		{
			name: "empty",
			in: url.Values{
				"hash":         []string{""},
				"interact_ref": []string{""},
			},
			want: IACallback{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromQuery(tt.in); got != tt.want {
				t.Errorf("FromQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
