package rp_resource_verifier

import (
	"aws/iam/rp_core"
	"fmt"
)

func RpResourceVerifier(TestedRolePolicy rp_core.RolePolicy, expected string) bool {
	fmt.Println("Tried loading json file.")

	// todo: iterate through every statement
	var resourceField = TestedRolePolicy.PolicyDocument.Statement[0].Resource

	return resourceField == expected
}
