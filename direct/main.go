package direct

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("direct: Hello %s", s)
}
