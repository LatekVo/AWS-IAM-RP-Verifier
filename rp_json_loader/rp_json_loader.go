package rp_json_loader

import (
	"aws/iam/rp_core"
	"encoding/json"
)

func RpJsonLoader(jsonBytes []byte) (rp_core.RolePolicy, error) {
	var newRolePolicy rp_core.RolePolicy
	err := json.Unmarshal(jsonBytes, &newRolePolicy)

	return newRolePolicy, err
}
