package main

const statusTemplate = `package main

import (
	"github.com/ProductHealth/gose4"
)

func init() {
	gose4.ServiceStatus = gose4.Status{
		ArtifactId: "{{.ArtifactId}}",
		BuildNumber: "{{.BuildNumber}}",
		BuildMachine: "{{.BuildMachine}}",
		BuildBy: "{{.BuildBy}}",
		BuildWhen: "{{.BuildWhen}}",
		CompilerVersion: "{{.CompilerVersion}}",
		GitSha: "{{.GitSha}}",
	}
}
`
