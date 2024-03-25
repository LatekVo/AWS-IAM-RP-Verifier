package rp_json_loader_test

import (
	"aws/iam/rp_json_loader"
	"aws/iam/tests"
	"aws/iam/utilities"
	"testing"
)

func TestRpJsonLoader(t *testing.T) {
	testcases := tests.GetTestCases()
	testsPath := "../tests/data/"

	for _, tc := range testcases {
		rpRaw, _ := utilities.LoadFile(testsPath + tc.In)
		_, err := rp_json_loader.RpJsonLoader(rpRaw)

		if err != nil && tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, err, tc.Want)
		}
	}
}
