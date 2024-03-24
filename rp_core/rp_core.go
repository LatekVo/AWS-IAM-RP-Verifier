package rp_core

import (
	"encoding/json"
	"fmt"
)

// While default unmarshalling works with well-structured nested jsons,
// a custom function is required to catch any malformed jsons, as they won't be detected otherwise

type Statement struct {
	Sid      string   `json:"Sid" validate:"required"`
	Effect   string   `json:"Effect" validate:"required"`
	Action   []string `json:"Action" validate:"required"`
	Resource string   `json:"Resource" validate:"required"`
}

type PolicyDocument struct {
	Version   string      `json:"Version" validate:"required"`
	Statement []Statement `json:"Statement" validate:"required"`
}

type RolePolicy struct {
	PolicyName     string         `json:"PolicyName" validate:"required"`
	PolicyDocument PolicyDocument `json:"PolicyDocument" validate:"required"`
}

func (p *RolePolicy) UnmarshalJSON(data []byte) error {
	// this method is required to ensure all fields are present,
	// default unmarshaller doesn't validate required fields

	// Custom validation:
	var formlessObject map[string]interface{}
	err := json.Unmarshal(data, &formlessObject)

	if err != nil {
		return err
	}

	// check if each field exists
	requiredFields := []string{"PolicyName", "PolicyDocument"}
	for _, field := range requiredFields {
		value := formlessObject[field]
		if value == nil {
			return fmt.Errorf("missing required field: %v", field)
		}
	}

	// since type casting from map[string]interface{} struct-members to RP's struct-members is not possible in Go,
	// we have to rely on unmarshall to directly populate those struct fields

	// we create this nested version of RP to perform standard unmarshalling
	// without recursively calling the current function
	type NestedRolePolicy struct {
		PolicyName     string
		PolicyDocument PolicyDocument
	}

	var nrp NestedRolePolicy
	json.Unmarshal(data, &nrp)

	// there is no easy way to convert formlessObject
	rp := &RolePolicy{
		PolicyName:     nrp.PolicyName,
		PolicyDocument: nrp.PolicyDocument,
	}

	// replace empty rp with our new rp
	*p = *rp

	return nil
}
