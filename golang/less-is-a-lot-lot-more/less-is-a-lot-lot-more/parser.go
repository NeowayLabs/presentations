package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, "./less-is-a-lot-lot-more/parser.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Decls {
		fmt.Println(s)
	}
} // END
