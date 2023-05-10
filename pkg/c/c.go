package c

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("pkg/c: Hello %s", s)
}
