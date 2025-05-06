package main

import (
	"fmt"
	"os"
)

func getFlagValue(flag string) (string, error) {

	for i, arg := range os.Args {
		if arg == flag && i+1 < len(os.Args) {
			return os.Args[i+1], nil
		}
	}

	return "", fmt.Errorf("no flag %s was found", flag)
}
