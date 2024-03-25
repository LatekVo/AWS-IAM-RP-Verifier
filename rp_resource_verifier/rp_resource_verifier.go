package rp_resource_verifier

import (
	"aws/iam/rp_core"
	"aws/iam/rp_structure_verifier"
	"errors"
)

func RpResourceVerifier(testedRolePolicy rp_core.RolePolicy, expected string) (bool, error) {
	// Task: "Method shall return logical false if an input JSON Resource field contains a single asterisk and true in any other case."


	statements := testedRolePolicy.PolicyDocument.Statement

	// expecting exactly one statement
	if len(statements) != 1 {
		return true, errors.New("multiple statements present in role policy")
	}

	isStructureValid := rp_structure_verifier.RpStructureVerifier(testedRolePolicy)
	if !isStructureValid {
		return true, errors.New("role policy json structure is malformed")
	}

	var resourceField = statements[0].Resource

	return resourceField != expected, nil
}
