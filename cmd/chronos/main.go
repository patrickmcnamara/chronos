// Chronos tells the current chronos time.
//
package main

import (
	"fmt"

	"github.com/patrickmcnamara/chronos"
)

func main() {
	now := chronos.Now()
	fmt.Println(now)
}
