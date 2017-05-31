package main

import (
	"fmt"

	"github.com/johnsudaar/testpkg1"
	"github.com/johnsudaar/testpkg2"
)

func main() {
	fmt.Println("PKG1: " + testpkg1.Version())
	fmt.Println("PKG2: " + testpkg2.Version())
}
