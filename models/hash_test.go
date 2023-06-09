package models

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestHashMethod_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		hm      HashMethod
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			hm:   SHA_512,
			want: []byte(`"sha-512"`),
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "invalid",
			hm:      HashMethod("wrong-hash"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.hm.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("HashMethod.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("HashMethod.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMethod_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		want    HashMethod
		in      []byte
		wantErr bool
	}{
		{
			name: "valid",
			in:   []byte(`"sha3-512"`),
			want: SHA3_512,
		},
		{
			name:    "json",
			in:      []byte(`{}`),
			wantErr: true,
		},
		{
			name:    "invalid",
			in:      []byte(`"sha-48"`),
			wantErr: true,
		},
		{
			name:    "empty",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got HashMethod
			err := json.Unmarshal(tt.in, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashMethod.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != got {
				t.Errorf("HashMethod.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMethod_Sum(t *testing.T) {
	tests := []struct {
		name string
		hm   HashMethod
		in   []byte
		want []byte
	}{
		{
			name: "default",
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x3d, 0x75, 0x38, 0x8, 0x8f, 0x4a, 0x26, 0xd9, 0xc8, 0xc5, 0x16, 0xdf, 0x15, 0x88, 0x1e, 0xa1, 0x63, 0x19, 0xa7, 0xea, 0xce, 0xd6, 0x97, 0x2d, 0x24, 0xea, 0xfc, 0xe6, 0x7, 0x9, 0xd5, 0x2f},
		},
		{
			name: "sha-256",
			hm:   SHA_256,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x3d, 0x75, 0x38, 0x8, 0x8f, 0x4a, 0x26, 0xd9, 0xc8, 0xc5, 0x16, 0xdf, 0x15, 0x88, 0x1e, 0xa1, 0x63, 0x19, 0xa7, 0xea, 0xce, 0xd6, 0x97, 0x2d, 0x24, 0xea, 0xfc, 0xe6, 0x7, 0x9, 0xd5, 0x2f},
		},
		{
			name: "sha-384",
			hm:   SHA_384,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x10, 0x7d, 0xa7, 0xa1, 0xe, 0xf0, 0xc3, 0x30, 0xe8, 0xa1, 0x34, 0xa5, 0x32, 0xab, 0xc1, 0xb8, 0x6d, 0x90, 0xf0, 0x95, 0xa, 0x76, 0x87, 0xea, 0xe0, 0xec, 0xd5, 0x60, 0x2b, 0x23, 0x53, 0x47, 0x67, 0xb4, 0xf9, 0xb0, 0x1e, 0x77, 0xdb, 0xff, 0x1c, 0xe, 0xad, 0xa6, 0x16, 0xd, 0x0, 0x5b},
		},
		{
			name: "sha-512",
			hm:   SHA_512,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x29, 0xdd, 0x85, 0xb1, 0x64, 0xbd, 0xc8, 0x8d, 0x56, 0x33, 0x31, 0xbd, 0x1e, 0xe3, 0xad, 0x48, 0xdb, 0xbf, 0x6d, 0x40, 0x86, 0xb9, 0xc6, 0xae, 0x79, 0x1a, 0xa1, 0x1b, 0xb5, 0xa, 0x1d, 0x81, 0x8a, 0x5e, 0x67, 0xa9, 0x69, 0x9a, 0xf9, 0x1d, 0x9c, 0x8b, 0xb8, 0x2d, 0x81, 0x20, 0xaf, 0x3e, 0x90, 0xf9, 0x7c, 0xb2, 0xcd, 0xdb, 0xb8, 0x35, 0x7a, 0xc4, 0x9b, 0xb, 0x7a, 0x85, 0x62, 0x7b},
		},
		{
			name: "sha3-224",
			hm:   SHA3_224,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x71, 0xd7, 0x6e, 0x8c, 0x4, 0x77, 0xff, 0xcb, 0xa8, 0x8b, 0x43, 0x43, 0x2, 0xeb, 0xb0, 0xe4, 0xc7, 0xb3, 0x99, 0x5d, 0x15, 0x6b, 0x1, 0xec, 0x8c, 0x52, 0x9f, 0x54},
		},
		{
			name: "sha3-384",
			hm:   SHA3_384,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0xdc, 0xac, 0xf8, 0xdf, 0x36, 0x1, 0x7a, 0x49, 0xed, 0x7, 0xa9, 0x70, 0x0, 0x8a, 0x7f, 0x84, 0x56, 0xdd, 0x31, 0xb3, 0x60, 0xc2, 0xee, 0x5d, 0x53, 0xc, 0x4, 0xae, 0x1c, 0x6b, 0x66, 0x4f, 0xb8, 0xc2, 0xdc, 0xf3, 0xa0, 0x92, 0x22, 0x37, 0x4b, 0x71, 0xb, 0xeb, 0x77, 0x71, 0xe2, 0x7d},
		},
		{
			name: "sha3-512",
			hm:   SHA3_512,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0xd5, 0x9c, 0x5f, 0xb8, 0x13, 0x7f, 0x2, 0x24, 0x4c, 0x69, 0x48, 0xc3, 0x22, 0x2, 0x10, 0xbc, 0xcb, 0x22, 0xaa, 0x48, 0x20, 0x15, 0x59, 0x70, 0xe4, 0x42, 0x8f, 0x8b, 0x61, 0x64, 0x30, 0x31, 0xf9, 0x90, 0xdb, 0x63, 0x34, 0x88, 0x18, 0x8b, 0x36, 0x91, 0x9f, 0x24, 0x8, 0x79, 0x1b, 0x81, 0xf4, 0x90, 0x5d, 0x65, 0xc0, 0xa8, 0xdb, 0xe2, 0x56, 0xbb, 0x84, 0x0, 0x4a, 0xf4, 0x65, 0xd9},
		},
		{
			name: "blake2s-256",
			hm:   BLAKE2s_256,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0x4, 0xff, 0x47, 0x60, 0x63, 0x2a, 0x40, 0xfc, 0x1d, 0x2f, 0x3d, 0x6a, 0x3, 0xb0, 0xc5, 0xe4, 0x68, 0x54, 0x99, 0x56, 0x5f, 0x75, 0x76, 0xde, 0xe9, 0x95, 0xfb, 0xaf, 0x2c, 0x3c, 0x95, 0x92},
		},
		{
			name: "blake2b-256",
			hm:   BLAKE2b_256,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0xcc, 0xc, 0x33, 0x45, 0x53, 0xf7, 0x77, 0xfb, 0x81, 0x48, 0xa, 0x1d, 0xa0, 0xd4, 0x56, 0x51, 0x47, 0xaa, 0x97, 0xed, 0x71, 0x51, 0xb1, 0x70, 0x7d, 0x54, 0xf9, 0xa2, 0x89, 0x42, 0xbb, 0x49},
		},
		{
			name: "blake2b-512",
			hm:   BLAKE2b_512,
			in:   []byte("gnap-core-protocol"),
			want: []byte{0xb9, 0x36, 0x2f, 0xb9, 0x42, 0x8d, 0x53, 0x6e, 0x29, 0xb2, 0x9a, 0x15, 0xeb, 0x69, 0xf, 0xa8, 0xa5, 0xc1, 0x34, 0xc7, 0x7, 0x55, 0xda, 0x2f, 0x8f, 0x10, 0xbe, 0x9, 0x2e, 0xa9, 0x19, 0x4f, 0xe0, 0x1c, 0xe0, 0x63, 0xe5, 0xdc, 0x2, 0xcf, 0xf5, 0xad, 0x20, 0xaf, 0x2e, 0x4e, 0x4f, 0xce, 0xc5, 0x1b, 0xc, 0xe3, 0xe1, 0xc3, 0x3e, 0x22, 0xf5, 0xa1, 0xd7, 0xba, 0x87, 0xe3, 0x74, 0x4e},
		},
		{
			name: "invalid",
			hm:   "wrong-hash",
			in:   []byte("gnap-core-protocol"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.Sum(tt.in); !bytes.Equal(got, tt.want) {
				t.Errorf("HashMethod.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
