package main

import (
	"fmt"
)

type VM struct {
	stack   []int
	pc      int
	alu     ALU
	program []any
	memory  map[string]int
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

func (vm *VM) jump() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}

	target := vm.program[vm.pc+1].(int)
	if target < 0 || target >= len(vm.program) {
		panic("invalid jump target")
	}

	vm.pc = target
}

func (vm *VM) jump_if_false() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}
	cond := vm.pop()
	target := vm.program[vm.pc+1].(int)

	if cond == 0 {
		if target < 0 || target >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = target
		return
	}

	vm.pc += 2
}

func (vm *VM) jump_if_true() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}
	cond := vm.pop()
	target := vm.program[vm.pc+1].(int)

	if cond != 0 {
		if target < 0 || target >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = target
		return
	}

	vm.pc += 2
}

func (vm *VM) eq() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.Equal(x, y))
}

func (vm *VM) neq() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.NotEqual(x, y))
}

func (vm *VM) gt() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.GreaterThan(x, y))
}

func (vm *VM) lt() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.LessThan(x, y))
}

func (vm *VM) ge() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.GreaterThanOrEqual(x, y))
}

func (vm *VM) le() {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	x := vm.pop()
	y := vm.pop()
	vm.push(vm.alu.LessThanOrEqual(x, y))
}

func (vm *VM) store(key string) {
	if len(vm.stack) < 2 {
		panic("not enough operands")
	}

	value := vm.pop()
	vm.memory[key] = value
}

func (vm *VM) load(key string) {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}

	value := vm.memory[key]
	vm.push(value)
}

func (vm *VM) Run() {
	for vm.pc < len(vm.program) {
		switch instr := vm.program[vm.pc].(type) {
		case string:
			switch instr {
			case "PUSH":
				value := vm.program[vm.pc+1].(int)
				vm.push(value)
				vm.pc += 2

			case "POP":
				vm.pop()
				vm.pc++

			case "SUM":
				vm.sum()
				vm.pc++

			case "RES":
				vm.res()
				vm.pc++

			case "MUL":
				vm.mul()
				vm.pc++

			case "DIV":
				vm.div()
				vm.pc++

			case "PRINT":
				vm.print()
				vm.pc++

			case "JUMP":
				vm.jump()

			case "JUMP_IF_FALSE":
				vm.jump_if_false()

			case "JUMP_IF_TRUE":
				vm.jump_if_true()

			case "EQ":
				vm.eq()
				vm.pc++

			case "NEQ":
				vm.neq()
				vm.pc++

			case "GT":
				vm.gt()
				vm.pc++

			case "LT":
				vm.lt()
				vm.pc++

			case "GE":
				vm.ge()
				vm.pc++

			case "LE":
				vm.le()
				vm.pc++

			case "STORE":
				vm.store(vm.program[vm.pc+1].(string))
				vm.pc += 2

			case "LOAD":
				vm.load(vm.program[vm.pc+1].(string))
				vm.pc += 2

			default:
				panic("unknown instruction: " + instr)
			}
		default:
			panic("expected opcode, got operand")
		}
	}
}
