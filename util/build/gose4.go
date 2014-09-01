package main

import (
	"github.com/ProductHealth/gose4/server"
	"github.com/ProductHealth/gose4/util"
	"os/exec"
	"os/user"
	"encoding/json"
	"fmt"
	"time"
	"strings"
	"runtime"
)

// Generate a minimum status json file, populating all build related fields
func main() {
	now := time.Now()
	// Extract variables

	// Write bare json file
	status := server.Status{}
	status.BuildMachine = util.GetCurrentHostName()

	status.SetBuildWhen(&now)
	status.GitSha = getCurrentGitRevision()
	status.BuildBy = getCurrentUser()

	r, _ := json.Marshal(status)
	print(fmt.Sprintf("%v", string(r)))
}

func getCurrentGitRevision() string {
	output, err := exec.Command("/usr/bin/env", "git", "rev-parse", "HEAD").Output()
	if err != nil {
		print(fmt.Sprintf("Could not determine git revision : %v", err))
		return "n/a"
	}
	return strings.TrimSpace(string(output))
}

func getCurrentUser() string {
	u, err := user.Current()
	if err != nil {
		print(fmt.Sprintf("Could not determine user : %v", err))
		return "n/a"
	}
	return u.Username
}

func getCompilerVersion() string {
	return runtime.NumCPU()
}
