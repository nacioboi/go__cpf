package main

import (
	"os"
	"testing"
	"time"

	cpf "github.com/nacioboi/go__cpf/cpf_debug"
	"github.com/nacioboi/go__cpf/cpf_options"
)

func TestDebug(t *testing.T) {
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

	cpf.Del(cpf.DEFAULT_HANDLER_ID)

	cpf.Log(INFO, "This should be printed to the file only\n")

	cpf.Del(1)
	cpf.Add(cpf.DEFAULT_HANDLER_ID, nil)

	cpf.Log(INFO, "This should not be printed to the file\n")

	cpf.Set(cpf_options.PRINT_IN_INTERVALS, 1)
	cpf.Log(INFO, "Only printed once 1.0\n")
	cpf.Log(INFO, "Only printed once 1.0\n")

	cpf.Log(INFO, "Only printed once 2.0\n")
	cpf.Log(INFO, "Only printed once 2.0\n")
}
