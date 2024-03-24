package rp_structure_verifier

import (
	"aws/iam/rp_core"
	"reflect"
)

func RpStructureVerifier(rpObject rp_core.RolePolicy) bool {
	if (reflect.TypeOf(rpObject) == reflect.TypeOf(rp_core.RolePolicy{})) {
		return true
	}

	return false
}
