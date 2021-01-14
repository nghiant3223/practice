package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}