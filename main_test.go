package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/rp_structure_verifier"
	"aws/iam/tests"
	"aws/iam/utilities"
	"testing"
)

func TestMain(t *testing.T) {
	testcases := tests.GetTestCases()
	testsPath := "tests/data/"

	for _, tc := range testcases {
		rpRaw, err := utilities.LoadFile(testsPath + tc.In)
		if err != nil && !tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, err, tc.Want)
		}

		rpJson, err := rp_json_loader.RpJsonLoader(rpRaw)
		if err != nil && !tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, err, tc.Want)
		}

		isStructureValid := rp_structure_verifier.RpStructureVerifier(rpJson)
		if !isStructureValid {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, isStructureValid, tc.Want)
		}

		// this err var was already handled in previous statements
		result, err := rp_resource_verifier.RpResourceVerifier(rpJson, "*")
		if err != nil && !tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, err, tc.Want)
		}

		if result != tc.Want {
			t.Errorf("Case: %v, got: %v, Want: %v", tc.In, result, tc.Want)
		}
	}
}

func FuzzMain(f *testing.F) {
	clean_payload_path := "tests/data/valid_payload_clean.json"
	clean_payload_bytes, _ := utilities.LoadFile(clean_payload_path)
	f.Add(clean_payload_bytes)
	f.Fuzz(func(t *testing.T, fuzzedData []byte) {
		rpJson, err := rp_json_loader.RpJsonLoader(fuzzedData)
		if err != nil {
			t.Skipf("RpJsonLoader rejected file: %v", fuzzedData)
		}

		isStructureValid := rp_structure_verifier.RpStructureVerifier(rpJson)
		if !isStructureValid {
			t.Skipf("RpStructureVerifier rejected object: %v", rpJson)
		}

		_, err = rp_resource_verifier.RpResourceVerifier(rpJson, "*")
		if err != nil {
			// this error should never occur
			t.Errorf("Error while verifying: %v", err)
		}

	})
}
