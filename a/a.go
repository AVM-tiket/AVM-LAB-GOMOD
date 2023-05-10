package a

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("a: Hello %s", s)
}
