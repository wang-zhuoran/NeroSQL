package main

import (
	"nerosql/parser"
	"os"
)

func main() {
	if err := parser.ParseSQL(os.Stdin); err != nil {
		panic(err)
	}
}
