package main

import (
	"os"
)

func main() {
	os.Exit(start(os.Args[1:]...))
}
