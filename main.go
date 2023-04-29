package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"niavka/compiler"
	"niavka/parser"
	"os"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewNiavkaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewNiavkaParser(stream)
	p.BuildParseTrees = true
	tree := p.File()
	visitor := parser.NiavkaVisitor{}
	programNode := visitor.Visit(tree).(parser.ProgramNode)

	m := compiler.CompileProgramNode(programNode)

	fmt.Println(m)
}
