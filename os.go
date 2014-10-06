package gose4

import (
	"fmt"
	"os"
)

func GetCurrentHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		print(fmt.Sprintf("Could not determine hostname : %v", err))
		return "n/a"
	}
	return hostname
}
