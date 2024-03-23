package rp_json_loader

import (
	"aws/iam/rp_core"
	"encoding/json"
	"log"
)

func RpJsonLoader(jsonBytes []byte) rp_core.RolePolicy {
	var newRolePolicy rp_core.RolePolicy
	err := json.Unmarshal(jsonBytes, &newRolePolicy)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return newRolePolicy
}
