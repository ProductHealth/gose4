package main


const statusTemplate = `package main

import (
	"github.com/ProductHealth/gose4/server"
)

func init() {
	server.ServiceStatus = server.Status{
		ArtifactId: "{{.ArtifactId}}",
		BuildNumber: "{{.ArtifactId}}",
		BuildMachine: "{{.BuildMachine}}",
		BuildBy: "{{.BuildBy}}",
		BuildWhen: "{{.BuildWhen}}",
		CompilerVersion: "{{.CompilerVersion}}",
		GitSha: "{{.GitSha}}",
	}
}
`

