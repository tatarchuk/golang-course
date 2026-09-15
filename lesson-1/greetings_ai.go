package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// os.Args[0] is the program itself; the first real argument is os.Args[1].
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	username := os.Args[1]
	fmt.Printf("Hello, %s!\n", username)
}
