package rp_structure_verifier_test

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_structure_verifier"
	"aws/iam/tests"
	"aws/iam/utilities"
	"testing"
)

func TestRpStructureVerifier(t *testing.T) {
	testcases := tests.GetTestCases()
	testsPath := "../tests/data/"

	for _, tc := range testcases {
		rpRaw, _ := utilities.LoadFile(testsPath + tc.In)
		rpJson, _ := rp_json_loader.RpJsonLoader(rpRaw)

		isStructureValid := rp_structure_verifier.RpStructureVerifier(rpJson)
		if !isStructureValid {
			t.Errorf("Case: %v, got: %v, want: %v", tc.In, isStructureValid, tc.Want)
		}
	}
}
