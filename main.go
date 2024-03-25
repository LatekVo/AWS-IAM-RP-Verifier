package main

import (
	"aws/iam/rp_json_loader"
	"aws/iam/rp_resource_verifier"
	"aws/iam/utilities"
	"flag"
	"fmt"
	"log"
)

func main() {
	// load filepath from command line if available
	filepath := flag.String("f", "tests/data/valid_payload_clean.json", "A path to the tested json")
	
	flag.Parse()

	fmt.Println(*filepath)

	// loading from files is a separate, optional functionality
	rp_raw, _ := utilities.LoadFile(*filepath)
	rp_json, _ := rp_json_loader.RpJsonLoader(rp_raw)

	result, err := rp_resource_verifier.RpResourceVerifier(rp_json, "*")

	if err != nil {
		log.Fatalf("RpResourceVerifier returned error while parsing: %v", err)
	}

	// Results: false - "*", true - [anything_else]
	fmt.Println(result)
}
