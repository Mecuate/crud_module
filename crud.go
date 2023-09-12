package crud_module

import (
	"fmt"
)

var Version versionType = "0.0.1"

func Log(vers string) {
	fmt.Println("[DEBUG] " + vers)
}
