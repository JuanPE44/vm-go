package main

import "os"

const (
	OP_PUSH byte = iota
	OP_POP
	OP_ADD
	OP_SUB
	OP_PRINT
	OP_MUL
	OP_DIV
	OP_JUMP
	OP_JUMP_IF_TRUE
	OP_JUMP_IF_FALSE
	OP_EQ
	OP_NEQ
	OP_QT
	OP_LT
	OP_GE
	OP_GT
	OP_LE
	OP_LOAD
	OP_STORE
	OP_CALL
	OP_RET
	OP_HALT
)

func main() {
	source, err := os.ReadFile("program.asm")
	if err != nil {
		panic(err)
	}
	bytecode := CompileASM(string(source))
	vm := &VM{
		alu:     ALU{},
		program: bytecode,
		memory:  make(map[int]int),
	}

	engine := &ExecutionEngine{}
	engine.Execute(vm)
}
