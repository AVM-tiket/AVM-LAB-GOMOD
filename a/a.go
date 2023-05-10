package a

import (
	"fmt"
)

func Echo(s string) {
	return fmt.Sprintf("a: Hello %s", s)
}
