package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Printf("The operating system is: %s\n", runtime.GOOS)
	fmt.Printf("Path is %s\n", os.Getenv("PATH"))
}
