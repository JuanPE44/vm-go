package vm

import (
	"fmt"
)

type Scope struct {
	vars map[int]int
}

type VM struct {
	running   bool
	stack     []int
	callStack []int
	pc        int
	alu       ALU
	program   []byte
	scopes    []*Scope
}

func NewVM(program []byte) *VM {
	return &VM{
		alu:     ALU{},
		program: program,
		scopes: []*Scope{
			{vars: make(map[int]int)},
		},
	}
}

func (vm *VM) currentScope() *Scope {
	return vm.scopes[len(vm.scopes)-1]
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

func (vm *VM) dump() {
	cont := vm.stack
	for len(cont) > 0 {
		fmt.Println(cont[len(cont)-1])
		cont = cont[:len(cont)-1]
	}
}

func (vm *VM) print() {
	if len(vm.stack) == 0 {
		panic("empty stack")
	}
	fmt.Println(vm.stack[len(vm.stack)-1])
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
	vm.currentScope().vars[key] = value
}

func (vm *VM) load(key int) {
	for i := len(vm.scopes) - 1; i >= 0; i-- {
		if v, ok := vm.scopes[i].vars[key]; ok {
			vm.push(v)
			return
		}
	}
	panic("variable not found")
}

func (vm *VM) call() {
	address := vm.program[vm.pc+1]
	vm.callStack = append(vm.callStack, vm.pc+2)

	// crear nuevo scope
	vm.scopes = append(vm.scopes, &Scope{
		vars: make(map[int]int),
	})

	vm.pc = int(address)
}

func (vm *VM) ret() {
	if len(vm.callStack) == 0 {
		panic("no function to return from")
	}

	vm.pc = vm.callStack[len(vm.callStack)-1]
	vm.callStack = vm.callStack[:len(vm.callStack)-1]

	// destruir scope
	vm.scopes = vm.scopes[:len(vm.scopes)-1]
}
