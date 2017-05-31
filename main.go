package main

import (
	"fmt"

	"github.com/johnsudaar/testpkg1"
)

func main() {
	fmt.Println("PKG1: " + testpkg1.Version())
}
