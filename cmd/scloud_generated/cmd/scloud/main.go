package main

//go:generate go run gen_version.go

import (
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd"
)

func main() {
	cmd.Execute()
}
