package main

import (
	"fmt"
	"niavka/compiler"
	"niavka/parser"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: niavka <filename>")
		os.Exit(1)
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	lexer := parser.NewNiavkaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewNiavkaParser(stream)
	p.BuildParseTrees = true
	tree := p.File()
	visitor := parser.NiavkaVisitor{}
	programNode := visitor.Visit(tree).(parser.ProgramNode)

	m, err := compiler.CompileProgramNode(programNode)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
