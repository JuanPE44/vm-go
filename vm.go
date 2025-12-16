package main

import "fmt"

type VM struct {
	stack   []int
	pc      int
	alu     ALU
	program []any
}

func (vm *VM) push(value int) {
	vm.stack = append(vm.stack, value)
}

func (vm *VM) pop() int {
	if len(vm.stack) == 0 {
		panic("stack underflow")
	}

	value := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return value
}

func (vm *VM) sum() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.Sum(x, y))
}

func (vm *VM) print() {
	cont := vm.stack
	for len(cont) > 0 {
		fmt.Println(cont[len(cont)-1])
		cont = cont[:len(cont)-1]
	}

}

func (vm *VM) res() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.Res(x, y))
}

func (vm *VM) mul() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.Mul(x, y))
}

func (vm *VM) div() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.Div(x, y))
}

func (vm *VM) Run() {
	for vm.pc < len(vm.program) {
		instr := vm.program[vm.pc].(string)
		vm.pc++

		switch instr {
		case "PUSH":
			value := vm.program[vm.pc].(int)
			vm.pc++
			vm.push(value)

		case "POP":
			vm.pop()

		case "SUM":
			vm.sum()

		case "RES":
			vm.res()

		case "MUL":
			vm.mul()

		case "DIV":
			vm.div()

		case "PRINT":
			vm.print()

		default:
			panic("unknown instruction: " + instr)
		}
	}
}
