package main

import (
	"fmt"
	lex "testGo/lexer"
	"testGo/readers"
)

func main() {
	fmt.Println("Hello")
	reader := readers.NewCharReader("/Users/supun/Supun/testGo/main.go")
	lexer := lex.NewLexer(&reader)

	for !reader.IsEOF() {
		token := lexer.NextToken()
		fmt.Println(token)

		switch token.(type) {
		case lex.IdentifierToken:
			fmt.Println("An identifier")
		default:
			fmt.Println("Something else")
		}

		fmt.Println()
	}

	// var f foo = foo{}
	f := new(foo)
	b := f
	fmt.Println(b)

}

type foo struct {
	a int
	b string
}

type bar struct {
	a int
}
