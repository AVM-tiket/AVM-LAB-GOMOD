package b

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("b: Hello %s", s)
}
