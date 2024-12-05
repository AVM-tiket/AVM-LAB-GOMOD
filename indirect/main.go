package indirect

import (
	"fmt"

	"github.com/AVM-tiket/AVM-LAB-GOMOD/direct"
)

func Echo(s string) string {
	return fmt.Sprintf("indirect: Hello %s", s)
}

func EchoIndirect(s string) string {
	return fmt.Sprintf("indirect: %s", direct.Echo(s))
}
