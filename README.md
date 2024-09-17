# CPF for Golang (Conditionally Printf)

This is an EXTREMELY simple package that allows developers to `fmt.Printf()` only when importing `go_cpf/cpf_debug` but not when importing `go_cpf/cpf_release`.

Effectively, this allows changing an import for release and have zero performance impact.
And you still get to keep your debug prints for when you're developing.

## Usage

- For release mode, change the import to `go_cpf/cpf_release`:

```go
package main

import (
	cpf "go_cpf/cpf_debug"
)

func main() {
	cpf.Printf("Hello, World!\n")
}
```

## Installation

```bash
go get github.com/nacioboi/go_cpf
```

## Quick note on why this works

It works because the go compiler will realize that the `go_cpf/cpf_release.Printf` function does nothing and will remove it from the compiled binary.

The following code:

```go
package main

import (
	cpf "go_cpf/cpf_release"
)

func main() {
	cpf.Printf("Hello, %s!\n", "world")
}
```

Will compile to:

```text
TEXT main.main(SB) C:/Users/joel.DESKTOP-T6RK670/Documents/Programming-Projects/go_test/main.go
  main.go:9             0x469620                c3                      RET
  :-1                   0x469621                cc                      INT $0x3
  :-1                   0x469622                cc                      INT $0x3
  :-1                   0x469623                cc                      INT $0x3
  :-1                   0x469624                cc                      INT $0x3
  :-1                   0x469625                cc                      INT $0x3
  :-1                   0x469626                cc                      INT $0x3
  :-1                   0x469627                cc                      INT $0x3
  :-1                   0x469628                cc                      INT $0x3
  :-1                   0x469629                cc                      INT $0x3
  :-1                   0x46962a                cc                      INT $0x3
  :-1                   0x46962b                cc                      INT $0x3
  :-1                   0x46962c                cc                      INT $0x3
  :-1                   0x46962d                cc                      INT $0x3
  :-1                   0x46962e                cc                      INT $0x3
  :-1                   0x46962f                cc                      INT $0x3
  :-1                   0x469630                cc                      INT $0x3
  :-1                   0x469631                cc                      INT $0x3
  :-1                   0x469632                cc                      INT $0x3
  :-1                   0x469633                cc                      INT $0x3
  :-1                   0x469634                cc                      INT $0x3
  :-1                   0x469635                cc                      INT $0x3
  :-1                   0x469636                cc                      INT $0x3
  :-1                   0x469637                cc                      INT $0x3
  :-1                   0x469638                cc                      INT $0x3
  :-1                   0x469639                cc                      INT $0x3
  :-1                   0x46963a                cc                      INT $0x3
  :-1                   0x46963b                cc                      INT $0x3
  :-1                   0x46963c                cc                      INT $0x3
  :-1                   0x46963d                cc                      INT $0x3
  :-1                   0x46963e                cc                      INT $0x3
  :-1                   0x46963f                cc                      INT $0x3
```

See? The `.Printf` call is completely removed from the compiled binary.

## Contributing

Feel free to submit a pull request or open an issue.

Or, fork the project and make it your own.

## License

This software is covered under the MIT License.
