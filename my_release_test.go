package main

import (
	"os"
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

	cpf.Set(cpf_options.PRINT_IN_INTERVALS, 1)
	cpf.Log(INFO, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")
	cpf.Log(INFO, "Hello, %s!\n", "World")
	cpf.Log(INFO, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")
	cpf.Log(DETAIL, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")
	cpf.Log(INFO, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")
	cpf.Set(cpf_options.LOG_LEVEL, DETAIL)
	cpf.Log(DETAIL, "This is detail level\n")
	cpf.Set(cpf_options.LOG_LEVEL, INFO)
	cpf.Log(DETAIL, "IF THIS IS PRINTED THEN THE TEST IS A FAILURE EVEN IF IT SAYS IT PASSED!\n")
	var s string
	cpf.Formatted(&s, "FYI: %s", "golang")
	cpf.Log(INFO, "%s %s, %d, %v\n",
		s,
		"is a great language",
		55,
		[]int{1, 2, 3},
	)

	file, err := os.Create("test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer file.Sync()

	cpf.Set(cpf_options.PRINT_IN_INTERVALS, 0)
	cpf.Add(1, file)

	cpf.Log(INFO, "This should be printed to the file\n")

	cpf.Del(1)

	cpf.Log(INFO, "This should not be printed to the file\n")
}
