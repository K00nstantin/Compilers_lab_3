package main

import (
	"fmt"
)

func main() {
	input := `{ a := 5; b := a + 3; c := (a > b) and not (b = 0) }`
	lexer := NewLexer(input)
	tokens, err := lexer.Tokenize()
	if err != nil {
		fmt.Println("Lexical error:", err)
		return
	}
	for _, t := range tokens {
		fmt.Println(t)
	}
}
