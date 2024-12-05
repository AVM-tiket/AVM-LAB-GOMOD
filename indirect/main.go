package indirect

import (
	"fmt"
)

func Echo(s string) string {
	return fmt.Sprintf("indirect: Hello %s", s)
}
