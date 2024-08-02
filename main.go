package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: esptool <filename>")
	}
}
