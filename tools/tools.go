//go:build tools

// Package tools tracks build-time tool dependencies so `go mod tidy` keeps them
// in go.mod. It is never compiled into the provider binary.
package tools

import (
	// Documentation generation for the Terraform Registry (`go generate ./...`).
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
