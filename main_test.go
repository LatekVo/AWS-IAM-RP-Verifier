package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/rp_structure_verifier"
	"aws/iam/utilities"
	"testing"
)

func TestMain(t *testing.T) {
	testcases := []struct {
		in   string
		want bool
	}{
		{"valid_payload_clean.json", true},
		{"valid_misordered.json", true},
		{"valid_missing_extension", true}, // make sure we don't care about externalities
		{"invalid_payload_clean.json", false},
		{"invalid_wrong_format.json", false},
		{"invalid_field_types.json", false},
		{"bad_filename.json", false},
	}
	for _, tc := range testcases {
		rpRaw, err := utilities.LoadFile("tests/data/" + tc.in)
		if err != nil && tc.want {
			t.Errorf("Case: %v, got: %v, want: %v", tc.in, err, tc.want)
		}

		rpJson, err := rp_json_loader.RpJsonLoader(rpRaw)
		if err != nil && tc.want {
			t.Errorf("Case: %v, got: %v, want: %v", tc.in, err, tc.want)
		}

		isStructureValid := rp_structure_verifier.RpStructureVerifier(rpJson)
		if !isStructureValid {
			t.Errorf("Case: %v, got: %v, want: %v", tc.in, isStructureValid, tc.want)
		}

		result := rp_resource_verifier.RpResourceVerifier(rpJson, "*")
		if result != tc.want {
			t.Errorf("Case: %v, got: %v, want: %v", tc.in, result, tc.want)
		}
	}
}
