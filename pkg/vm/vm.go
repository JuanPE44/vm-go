package vm

import (
	"fmt"
	"vm-go/pkg/canvas"
)

type ValueType int

const (
	IntType ValueType = iota
	BoolType
	RefType // heap
)

const (
	VAL_INT ValueType = iota
	VAL_BOOL
	VAL_STRING
	VAL_NULL
)

type Value struct {
	Type ValueType
	Int  int
	Bool bool
	Ref  int
}

type Scope struct {
	vars map[int]Value
}

type Frame struct {
	returnPC int
	locals   []Value
	stack    []Value
}

type VM struct {
	running   bool
	callStack []Frame
	pc        int
	alu       ALU
	program   []byte
	scopes    []*Scope
	canvas    *canvas.Canvas
}

func NewVM(program []byte) *VM {
	return &VM{
		alu:     ALU{},
		program: program,
		scopes: []*Scope{
			{vars: make(map[int]Value)},
		},
		callStack: []Frame{
			{
				returnPC: -1,
				locals:   make([]Value, 0),
				stack:    make([]Value, 0),
			},
		},
		canvas: canvas.NewCanvas(20, 10),
	}
}

func (vm *VM) currentScope() *Scope {
	return vm.scopes[len(vm.scopes)-1]
}

func (vm *VM) currentFrame() *Frame {
	return &vm.callStack[len(vm.callStack)-1]
}

func (vm *VM) push(value Value) {
	frame := vm.currentFrame()
	frame.stack = append(frame.stack, value)
}

func (vm *VM) pop() Value {
	frame := vm.currentFrame()
	if len(frame.stack) == 0 {
		panic("stack underflow")
	}

	value := frame.stack[len(frame.stack)-1]
	frame.stack = frame.stack[:len(frame.stack)-1]
	return value
}

func (vm *VM) add() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: ADD expects int + int")
	}

	vm.push(vm.alu.Add(a, b))
}

func (vm *VM) sub() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: SUB expects int - int")
	}

	vm.push(vm.alu.Sub(a, b))
}

func (vm *VM) mul() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: MUL expects int * int")
	}

	vm.push(vm.alu.Mul(a, b))
}

func (vm *VM) div() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: DIV expects int / int")
	}

	vm.push(vm.alu.Div(a, b))
}

func (vm *VM) jump() {
	target := vm.program[vm.pc+1]
	if int(target) >= len(vm.program) {
		panic("invalid jump target")
	}
	vm.pc = int(target)
}

func (vm *VM) jump_if_false() {
	cond := vm.pop()

	if cond.Type != BoolType {
		panic("type error: JUMP_IF_FALSE expects bool")
	}

	frame := vm.currentFrame()

	if len(frame.stack) < 1 {
		panic("not enough operands")
	}
	target := vm.program[vm.pc+1]

	if !cond.Bool {
		if int(target) >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = int(target)
		return
	}

	vm.pc += 2
}

func (vm *VM) jump_if_true() {
	cond := vm.pop()
	if cond.Type != BoolType {
		panic("type error: JUMP_IF_TRUE expects bool")
	}

	frame := vm.currentFrame()

	if len(frame.stack) < 1 {
		panic("not enough operands")
	}
	target := vm.program[vm.pc+1]

	if cond.Bool {
		if int(target) >= len(vm.program) {
			panic("invalid jump target")
		}
		vm.pc = int(target)
		return
	}

	vm.pc += 2
}

func (vm *VM) eq() {

	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: EQ expects int + int")
	}

	vm.push(vm.alu.Equal(a, b))
}

func (vm *VM) neq() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: NEQ expects int + int")
	}

	vm.push(vm.alu.NotEqual(a, b))
}

func (vm *VM) gt() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: GT expects int + int")
	}

	vm.push(vm.alu.GreaterThan(a, b))
}

func (vm *VM) lt() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: LT expects int + int")
	}

	vm.push(vm.alu.LessThan(a, b))
}

func (vm *VM) ge() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: GE expects int + int")
	}

	vm.push(vm.alu.GreaterThanOrEqual(a, b))
}

func (vm *VM) le() {
	frame := vm.currentFrame()
	if len(frame.stack) < 2 {
		panic("not enough operands")
	}

	b := vm.pop()
	a := vm.pop()

	if a.Type != IntType || b.Type != IntType {
		panic("type error: LE expects int + int")
	}

	vm.push(vm.alu.LessThanOrEqual(a, b))
}

func (vm *VM) store(key int) {
	frame := vm.currentFrame()
	if len(frame.stack) < 1 {
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

	vm.callStack = append(vm.callStack, Frame{
		returnPC: vm.pc + 2,
		locals:   make([]Value, 0),
		stack:    make([]Value, 0),
	})
	// crear nuevo scope
	vm.scopes = append(vm.scopes, &Scope{
		vars: make(map[int]Value),
	})

	vm.pc = int(address)
}

func (vm *VM) ret() {
	curr := vm.currentFrame()
	retPC := curr.returnPC
	vm.callStack = vm.callStack[:len(vm.callStack)-1]
	vm.pc = retPC

}

func (vm *VM) dup() {
	frame := vm.currentFrame()
	if len(frame.stack) < 1 {
		panic("stack underflow")
	}
	v := frame.stack[len(frame.stack)-1]
	frame.stack = append(frame.stack, v)
}

func (vm *VM) sysDrawPixel() {
	b := vm.pop()
	a := vm.pop()
	vm.canvas.DrawPixel(a.Int, b.Int)
}

func (vm *VM) sysPresent() {
	vm.canvas.Present()
}

func (vm *VM) print() {
	frame := vm.currentFrame()
	if len(frame.stack) == 0 {
		panic("empty stack")
	}
	fmt.Println(frame.stack[len(frame.stack)-1].Int)
}

func (vm *VM) dump() {
	frame := vm.currentFrame()
	cont := frame.stack
	for len(cont) > 0 {
		fmt.Println(cont[len(cont)-1].Int)
		cont = cont[:len(cont)-1]
	}
}
