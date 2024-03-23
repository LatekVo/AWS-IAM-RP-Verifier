package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/utilities"
	"testing"
)

func TestMain(t *testing.T) {
	testcases := []struct {
		in   string
		want bool
	}{
		{"valid_payload_clean.json", true},
		{"invalid_payload_clean.json", false},
		{"bad_filename.json", false},
		{"wrong_format.json", false},
		{"wrong_filetype.txt", false},
	}
	for _, tc := range testcases {
		rp_raw := utilities.LoadFile("tests/data/" + tc.in)
		rp_json := rp_json_loader.RpJsonLoader(rp_raw)
		result := rp_resource_verifier.RpResourceVerifier(rp_json, "*")
		if result != tc.want {
			t.Errorf("Got: %v, want: %v", result, tc.want)
		}
	}

}
