package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/utilities"
	"fmt"
)

func main() {
	rp_raw := utilities.LoadFile("tests/data/valid_payload_clean.json")
	rp_json := rp_json_loader.RpJsonLoader(rp_raw)
	result := rp_resource_verifier.RpResourceVerifier(rp_json, "*")

	fmt.Println("Results:", result)
}
