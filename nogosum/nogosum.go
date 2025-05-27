package nogosum

import (
	"fmt"

	"github.com/AVM-tiket/AVM-LAB-GOMOD/a"
)

func Echo(s string) string {
	return fmt.Sprintf("nogosum: %s", a.Echo(s))
}
