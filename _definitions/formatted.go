package definitions

import "fmt"

func FormattedImplementation(out *string, format string, args ...interface{}) {
	*out = fmt.Sprintf(format, args...)
}
