package main

import (
	"fmt"

	"github.com/NeowayLabs/nash/scanner"
)

func main() {
	l := scanner.Lex("example", `shells = (nash rc csh zsh sh "MS PowerShell" bash)`)

	for tok := range l.Tokens {
		fmt.Println(tok)
	}
}
