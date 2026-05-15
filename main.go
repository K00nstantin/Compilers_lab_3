package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "program.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", filename, err)
		return
	}

	input := string(data)

	lexer := NewLexer(input)
	tokens, err := lexer.Tokenize()
	if err != nil {
		fmt.Println("Лексическая ошибка:", err)
		return
	}

	parser := NewParser(tokens)
	if err := parser.parseProgram(); err != nil {
		fmt.Println("Синтаксическая ошибка:", err)
		return
	}

	if !parser.isAtEnd() {
		cur := parser.current()
		fmt.Printf("Синтаксическая ошибка: лишние токены после конца программы (строка %d, колонка %d): %v\n",
			cur.Line, cur.Col, cur.Type)
		return
	}

	fmt.Println("Синтаксический анализ успешно завершён. Программа принадлежит грамматике.")
}
