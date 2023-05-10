package b

import (
	"fmt"
)

func Echo(s string) {
	return fmt.Sprintf("b: Hello %s", s)
}
