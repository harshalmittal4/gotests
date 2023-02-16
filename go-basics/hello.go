package main 
// import path of main package is <module path>(in go.mod)/go-basics

import (
    "fmt"
    "github.com/google/go-cmp/cmp" // dependency on external module
    "learn/testing/gotests/morestrings"
)

func main() {
    fmt.Println("Hello, world.")
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
    // fmt.Fprintln() - Spaces are always added between 
    // the specified operands and a newline is appended at the end.
}