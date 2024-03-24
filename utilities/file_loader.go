package utilities

import (
	"os"
)

func LoadFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)

	// todo: add edge-case file checking here

	return content, err
}
