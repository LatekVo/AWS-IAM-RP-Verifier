package json_utils

import (
	// "encoding/json"
	"fmt"
)

func LoadJson() {
	fmt.Println("Tried loading json file.")
}

func GetField(fieldname string) string {
	fmt.Println("Tried getting json field.")

	// todo: allow null
	return fieldname
}
