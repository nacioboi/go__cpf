package main

import (
	"testing"
	"time"

	"github.com/nacioboi/go__cpf/cpf_options"
	cpf "github.com/nacioboi/go__cpf/cpf_release"
)

func TestRelease(t *testing.T) {
	cpf.Set(cpf_options.PREFIX_HANDLER, func() string {
		var pref string
		cpf.Formatted(&pref, "[%s] -- ", time.Now().Format("2006-01-02 15:04:05"))
		return pref
	})

	cpf.Log(INFO, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")

	var s string
	cpf.Formatted(&s, "FYI: %s", "golang")
	if s != "" {
		t.Errorf("Expected empty string, got %s", s)
	}
}
