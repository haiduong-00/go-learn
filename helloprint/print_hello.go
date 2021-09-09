package hello

import (
	"fmt"
)

func print_hello(name string) string {
	return fmt.Sprintf("Hello %v, thanks for using this application", name)
}
