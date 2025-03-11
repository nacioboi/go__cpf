# CPF for Golang (Conditionally Printf)

`go__cpf` is a lightweight logging package that allows developers to conditionally log messages based on the import.

This enables zero-cost logging removal in release builds by simply switching imports from `cpf_debug` to `cpf_release`.

## Features

### 1. Log levels

Use `Log(level, format, args...)` to print messages based on the current log level.

```go
cpf.Log(INFO, "Hello, %s!\n", "World")
```

Again, switching to `cpf_release` removes all logging at compile time.

### 1.5. Defining custom log levels

`go__cpf` does not include log levels by default, so you need to define them yourself.

To define log levels, create a `LogLevel` (the name doesn't matter) type alias and declare different levels:

```go
type LogLevel = int // MUST be an `int`

const (
	INFO LogLevel = iota  // General information
	DETAIL                // More detailed logging
	DEBUG                 // Debug-level logging
	ERROR                 // Error messages
)
```

A few things to note:

- Make sure its a type alias and not a type on its own.
  - Notice that there is an equal `=` sign in the type definition.
- It must equal to an `int` in order to be compatible with `go__cpf`

Then, you may use them with `cpf.Log`:

```go
cpf.Set(cpf_options.LOG_LEVEL, DETAIL)     // Set logging level to DETAIL
cpf.Log(INFO, "Application started\n")     // This will be printed
cpf.Log(DETAIL, "Detailed log message\n")  // Will print only if DETAIL level is enabled
```

### 2. Log intervals

Restrict how often messages are printed:

```go
cpf.Set(cpf_options.PRINT_IN_INTERVALS, 1)
```

This will only print every second log message.

This is useful when calling `cpf.Log` in a loop since printing to screen is incredibly slow.

### 3. Custom Output Handlers

You can override the default output behavior and redirect logs to a custom handler. The handler function must have the signature func(level int, msg string).

#### Example 1: Writing logs to a file

```go
import (
	"os"
	"log"

	cpf "github.com/nacioboi/go__cpf/cpf_debug"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file")
	}
	defer file.Close()
	defer file.Sync()

	// Remember what key you used, in this case, `1`.
	// You might need it to delete the handler later on.
	cpf.Add(1, file) 

	cpf.Log(0, "Hi inside stdout and the file :D")

	cpf.Del(1)

	cpf.Log(0, "Only from within stdout now")
}
```

#### Example: Logging over the network.

Obviously this should be done with caution.

```go
import (
	"net"
	"fmt"
	
	cpf "github.com/nacioboi/go__cpf/cpf_debug"
)

func main() {
	conn, err := net.Dial("udp", "192.168.1.100:9999")
	if err != nil {
		panic("Failed to connect to log server")
	}
	defer conn.Close()

	cpf.Add(1337, func(level int, msg string) {
		fmt.Fprintf(conn, "[Level %d] %s", level, msg)
	})

	cpf.Log("NEVER SHARE THE USERS PRIVATE INFO OVER A NETWORK USING `go__cpf` since that would be mean.")
}
```

#### Deleting and Reinserting the default handler

```go
cpf.Del(cpf.DEFAULT_HANDLER_ID)
cpf.Add(cpf.DEFAULT_HANDLER_ID, nil) // it knows about `DEFAULT_HANDLER_ID`.
```

### 4. Log Prefixes

Customize log message prefixes dynamically:

```go
cpf.Set(cpf_options.PREFIX_HANDLER, func() string {
	var pref string
	cpf.Formatted(&pref, "[%s] -- ", time.Now().Format("2006-01-02 15:04:05"))
	return pref
})
```

### 5. String Formatting

Format strings without printing:

```go
var s string
cpf.Formatted(&s, "Processed %d items", 42)
fmt.Println(s) // Output: Processed 42 items
```

## Usage Example

```go
package main

import (
	cpf "github.com/nacioboi/go__cpf/cpf_debug"
	"github.com/nacioboi/go__cpf/cpf_options"
)

type LogLevel = int // MUST be an `int`

const (
	INFO LogLevel = iota  // General information
	// ...
)

func main() {
	cpf.Set(cpf_options.LOG_LEVEL, INFO)
	cpf.Log(INFO, "Application started\n")

	var msg string
	cpf.Formatted(&msg, "User %s logged in", "Alice")
	cpf.Log(INFO, msg)
}
```

## Installation

```bash
go get github.com/nacioboi/go__cpf
```

## Quick note on why this works

It works because the go compiler will realize that the `go__cpf/cpf_release.Printf` function does nothing and will remove it from the compiled binary.

The following code:

```go
package main

import (
	cpf "github.com/nacioboi/go__cpf/cpf_release"
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
