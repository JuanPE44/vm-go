package main

import (
	"fmt"
	"strings"
	"vm-go/pkg/assembler"
	"vm-go/pkg/ast"
	"vm-go/pkg/compiler"
	"vm-go/pkg/vm"
)

func CompileASTtoASM(node ast.Node) string {
	var sb strings.Builder
	generate(node, &sb)
	sb.WriteString("PRINT\n")
	sb.WriteString("HALT\n")
	return sb.String()
}

func generate(node ast.Node, sb *strings.Builder) {
	switch n := node.(type) {
	case *ast.NumberNode:
		sb.WriteString(fmt.Sprintf("PUSH %d\n", n.Value))

	case *ast.BinaryOpNode:
		generate(n.Left, sb)
		generate(n.Right, sb)

		switch n.Operator {
		case "+":
			sb.WriteString("ADD\n")
		case "-":
			sb.WriteString("SUB\n")
		case "*":
			sb.WriteString("MUL\n")
		case "/":
			sb.WriteString("DIV\n")
		}
	}
}

func main() {
	input := "2 * (2 + 3)"

	l := compiler.NewLexer(input)

	p := compiler.NewParser(l)
	ast := p.ParseExpression()

	programASM := CompileASTtoASM(ast)
	fmt.Println("--- Assembly Generado ---")
	fmt.Println(programASM)

	bytecode := assembler.CompileASM(programASM)

	virtualMachine := vm.NewVM(bytecode)
	engine := vm.NewEngine()
	engine.Execute(virtualMachine)

}
