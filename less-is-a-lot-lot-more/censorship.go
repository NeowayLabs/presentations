package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	src := `
// HelloWorldService
package main

const hello = "Hello World!"

// This is a forbidden phrase!
const evil = "Where is my daily paçoca???"

func main() {
    fmt.Println(hello)
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Change the package comment
	f.Comments[0].List[0].Text = f.Comments[0].List[0].Text + " - CENSORED BY NEOWAY"

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	// Remove the evil variable declaration from the list of declarations.
	f.Decls = append(f.Decls[:1], f.Decls[2:]...)

	// Use the comment map to filter comments that don't belong anymore
	// (the comments associated with the variable declaration), and create
	// the new comments list.
	f.Comments = cmap.Filter(f).Comments()

	// Print the modified AST.
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.Bytes())
}
