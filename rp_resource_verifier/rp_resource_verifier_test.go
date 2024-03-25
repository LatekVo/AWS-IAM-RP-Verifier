package rp_resource_verifier_test

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/tests"
	"aws/iam/utilities"
	"testing"
)

func TestRpResourceVerifier(t *testing.T) {
	testcases := tests.GetTestCases()
	testsPath := "../tests/data/"

	for _, tc := range testcases {
		rpRaw, _ := utilities.LoadFile(testsPath + tc.In)
		rpJson, _ := rp_json_loader.RpJsonLoader(rpRaw)

		result, err := rp_resource_verifier.RpResourceVerifier(rpJson, "*")
		if err != nil && tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, err, tc.Want)
		}
		if result != tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, result, tc.Want)
		}
	}
}
