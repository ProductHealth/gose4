package main

import (
	"flag"
	"fmt"
	"github.com/ProductHealth/gose4"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// Generate a minimum status json file, populating all build related fields
func main() {
	now := time.Now()
	// Extract variables

	// Populate bare json file
	status := gose4.Status{}
	flag.StringVar(&status.ArtifactId, "artifactid", "unknown", "Artifact Id")
	flag.StringVar(&status.BuildNumber, "buildnumber", "unknown", "Build number")
	flag.Parse()
	status.BuildMachine = gose4.GetCurrentHostName()
	status.SetBuildWhen(&now)
	status.GitSha = getCurrentGitRevision()
	status.BuildBy = getCurrentUser()
	status.CompilerVersion = runtime.Version()
	write(status)
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

func write(status gose4.Status) {
	// Write to temp dir
	filename := "gose4_initialization.go"
	t, _ := template.New("gose4").Parse(statusTemplate)
	f, _ := os.Create(filename)
	defer f.Close()
	err := t.Execute(f, status)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Write %v", filename)
	}
}