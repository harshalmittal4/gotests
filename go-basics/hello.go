package main

// import path of main package is <module path>(in go.mod)/go-basics

import (
	"fmt"

	morestrings "learn/testing/gotests/revstringstesting" // dependency on local module

	"github.com/google/go-cmp/cmp" // dependency on external module
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go")) // fmt.Fprintln() - Spaces are always added between
}
