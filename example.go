package example

import (
	"fmt"
)

const version = "no-version"

func Exec() string {

	return fmt.Sprintf("This is go-module-example-5 %s", version)
}
