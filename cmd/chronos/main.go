package main

import (
	"fmt"
	"os"
	"time"

	"github.com/patrickmcnamara/chronos"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println(chronos.Now())
	case 2:
		// check if chronos
		c, err := chronos.Parse(os.Args[1])
		if err == nil {
			fmt.Println(c.ToTime().Format(time.RFC3339[11:19]))
			return
		}
		// otherwise, check if normal time
		t, err := time.Parse(time.RFC3339[11:19], os.Args[1])
		if err == nil {
			fmt.Println(chronos.FromTime(t))
			return
		}
		// otherwise, print error
		fmt.Fprintln(os.Stderr, "chronos: first argument should be of the form `CCCC` or `HH:MM:SS`")
		os.Exit(1)
	default:
		fmt.Fprintln(os.Stderr, "chronos: too many arguments")
		os.Exit(1)
	}
}
