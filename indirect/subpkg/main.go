package subpkg

import (
	"fmt"

	"github.com/AVM-tiket/AVM-LAB-GOMOD/direct"
)

func Echo(s string) string {
	return fmt.Sprintf("indirect: %s", direct.Echo(s))
}
