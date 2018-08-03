package main

import (
	"fmt"

	"github.com/patrickmcnamara/chronos/chronos"
)

func main() {
	now := chronos.Now()
	fmt.Println(now)
}
