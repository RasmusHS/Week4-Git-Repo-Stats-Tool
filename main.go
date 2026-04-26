package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	repoFlag := flag.String("path", ".", "Path to your local repo")
	flag.Parse()
}
