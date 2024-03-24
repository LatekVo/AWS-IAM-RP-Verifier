package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/utilities"
	"fmt"
	"log"
)

func main() {
	// loading from files is a separate, optional functionality
	rp_raw, _ := utilities.LoadFile("tests/data/valid_payload_clean.json")
	rp_json, _ := rp_json_loader.RpJsonLoader(rp_raw)

	result, err := rp_resource_verifier.RpResourceVerifier(rp_json, "*")

	if err != nil {
		log.Fatalf("RpResourceVerifier returned error while parsing: %v", err)
	}

	fmt.Println("Results:", result)
}
