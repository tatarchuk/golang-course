package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = os.Args[1]
	fmt.Println("Hello, " + name)
}
