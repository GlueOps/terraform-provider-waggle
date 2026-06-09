package main

// Generate Terraform Registry documentation from the provider schema and the
// examples/ directory. Run with `go generate ./...` (or `make docs`).
//
// This lives here rather than in main.go because main.go is OpenAPI-generated
// and carries a DO NOT EDIT header.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name waggle
