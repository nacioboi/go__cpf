//go:generate go run generate_cpf.go
package main

import (
	"fmt"
	"os"
	"strings"
)

const debugTemplate = `// Code generated by go generate; DO NOT EDIT.

package cpf_debug

import (
	definitions "github.com/nacioboi/go__cpf/_definitions"
	"github.com/nacioboi/go__cpf/cpf_options"
)

const c__MAX_INTERVAL_AMOUNT = 1<<16 - 1

var (
	interval_amount     uint16
	interval_count_down uint16
	log_level           int

	output_handlers = make(map[int]definitions.CustomOutputHandlerEntry, 0)

	prefix_handler func() string
)

func init() {
	default_output_handler := func(message string) {
		print(message)
	}
	entry := definitions.CustomOutputHandlerEntry{
		I: 0,
		H: default_output_handler,
	}
	output_handlers[1<<32-1] = entry
}

func Add(key int, value interface{}) {
	definitions.AddImplementation(
		key,
		value,
		output_handlers,
	)
}

func Del(key int) {
	definitions.DelImplementation(
		key,
		output_handlers,
	)
}

func Set(key cpf_options.T__Option, value interface{}) {
	definitions.SetImplementation(
		key,
		value,
		&interval_amount,
		&interval_count_down,
		c__MAX_INTERVAL_AMOUNT,
		&prefix_handler,
		&log_level,
	)
}

func Log(level int, format string, args ...interface{}) {
	definitions.LogImplementation(
		level,
		log_level,
		output_handlers,
		prefix_handler,
		&interval_amount,
		&interval_count_down,
		format,
		args...,
	)
}

func Formatted(out *string, format string, args ...interface{}) {
	definitions.FormattedImplementation(out, format, args...)
}
`

const releaseTemplate = `//go:build !debug
// +build !debug

// Code generated by go generate; DO NOT EDIT.

package cpf_release

import "github.com/nacioboi/go__cpf/cpf_options"

type CustomOutputHandler = func(int, string)

func Add(key int, value interface{}) {}

func Del(key int) {}

func Set(key cpf_options.T__Option, value interface{}) {}

func Log(level int, format string, args ...interface{}) {}

func Formatted(out *string, format string, args ...interface{}) {}
`

func writeFile(path, content string) {
	err := os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0644)
	if err != nil {
		fmt.Println("Failed to generate", path, ":", err)
		os.Exit(1)
	}
	fmt.Println("Generated", path)
}

func main() {
	writeFile("cpf_debug/cpf_debug.go", debugTemplate)
	writeFile("cpf_release/cpf_release.go", releaseTemplate)
}
