package main

import (
	"fmt"
)

type VM struct {
	running   bool
	stack     []int
	callStack []int
	pc        int
	alu       ALU
	program   []byte
	memory    map[int]int
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

func (vm *VM) add() {
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

func (vm *VM) sub() {
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

	target := vm.program[vm.pc+1]
	if int(target) >= len(vm.program) {
		panic("invalid jump target")
	}
	vm.pc = int(target)
}

func (vm *VM) jump_if_false() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}
	cond := vm.pop()
	target := vm.program[vm.pc+1]

	if cond == 0 {
		if int(target) >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = int(target)
		return
	}

	vm.pc += 2
}

func (vm *VM) jump_if_true() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}
	cond := vm.pop()
	target := vm.program[vm.pc+1]

	if cond != 0 {
		if int(target) >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = int(target)
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

func (vm *VM) store(key int) {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}

	value := vm.pop()
	vm.memory[key] = value
}

func (vm *VM) load(key int) {
	value := vm.memory[key]
	vm.push(value)
}

func (vm *VM) call() {
	if len(vm.stack) < 1 {
		panic("not enough operands")
	}

	address := vm.program[vm.pc+1]
	vm.callStack = append(vm.callStack, vm.pc+2)
	vm.pc = int(address)
}

func (vm *VM) ret() {
	if len(vm.callStack) == 0 {
		panic("no function to return from")
	}

	vm.pc = vm.callStack[len(vm.callStack)-1]
	vm.callStack = vm.callStack[:len(vm.callStack)-1]
}
