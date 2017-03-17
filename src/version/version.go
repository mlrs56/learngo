package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s", runtime.Version())
	fmt.Printf("%s", runtime.GOOS)
}
