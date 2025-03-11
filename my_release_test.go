package main

import (
	"testing"

	cpf "github.com/nacioboi/go__cpf/cpf_release"
)

func TestRelease(t *testing.T) {
	cpf.Set(0, nil)
	cpf.Add(0, nil)
	cpf.Del(0)
	cpf.Log(0, "This should not be printed\n")
	cpf.Formatted(nil, "This should not be printed\n")
}
