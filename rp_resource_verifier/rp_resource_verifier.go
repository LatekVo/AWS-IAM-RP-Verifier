package rp_resource_verifier

import (
	"aws/iam/rp_core"
)

func RpResourceVerifier(testedRolePolicy rp_core.RolePolicy, expected string) bool {

	statements := testedRolePolicy.PolicyDocument.Statement

	// expecting exactly one statement
	if len(statements) != 1 {
		return false
	}

	var resourceField = statements[0].Resource

	return resourceField == expected
}
