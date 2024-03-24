package rp_resource_verifier

import (
	"aws/iam/rp_core"
	"aws/iam/rp_structure_verifier"
	"errors"
)

func RpResourceVerifier(testedRolePolicy rp_core.RolePolicy, expected string) (bool, error) {

	statements := testedRolePolicy.PolicyDocument.Statement

	// expecting exactly one statement
	if len(statements) != 1 {
		return false, errors.New("multiple statements present in role policy")
	}

	isStructureValid := rp_structure_verifier.RpStructureVerifier(testedRolePolicy)
	if !isStructureValid {
		return false, errors.New("role policy json structure is malformed")
	}

	var resourceField = statements[0].Resource

	return resourceField == expected, nil
}
