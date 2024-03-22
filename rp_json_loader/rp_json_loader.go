package rp_json_loader

import (
	"aws/iam/rp_core"
	"encoding/json"
	"fmt"
	"log"
)

func RpJsonLoader(jsonBytes []byte) rp_core.RolePolicy {
	fmt.Println("Tried loading json file.")

	var newRolePolicy rp_core.RolePolicy
	err := json.Unmarshal(jsonBytes, &newRolePolicy)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return newRolePolicy
}
