//go:build testgen
// +build testgen

package main

import "github.com/ezraisw/test-gen/internal/example"

func main() {
	example.Generate()
}
