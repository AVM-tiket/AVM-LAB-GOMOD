package d

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("pkg/d: Hello %s", s)
}
