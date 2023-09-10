package cruds

import (
	"fmt"
)

var Version="0.0.0"

func Log(vers string) {
	fmt.Println("[DEBUG] " + vers)
   }