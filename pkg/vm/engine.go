package vm

import (
	"fmt"
	"vm-go/pkg/opcodes"
)

type ExecutionEngine struct {
}

func NewEngine() *ExecutionEngine {
	return &ExecutionEngine{}
}

func (ee *ExecutionEngine) Execute(vm *VM) {
	vm.running = true
	for vm.running && vm.pc < len(vm.program) {
		instr := vm.program[vm.pc]
		switch instr {
		case opcodes.OP_PUSH:
			value := vm.program[vm.pc+1]
			vm.push(int(value))
			vm.pc += 2

		case opcodes.OP_ADD, opcodes.OP_SUB, opcodes.OP_MUL, opcodes.OP_DIV, opcodes.OP_POP, opcodes.OP_PRINT, opcodes.OP_EQ, opcodes.OP_NEQ, opcodes.OP_GT, opcodes.OP_LT, opcodes.OP_GE, opcodes.OP_LE, opcodes.OP_DUMP, opcodes.OP_DUP:
			switch instr {
			case opcodes.OP_ADD:
				vm.add()
			case opcodes.OP_SUB:
				vm.sub()
			case opcodes.OP_MUL:
				vm.mul()
			case opcodes.OP_DIV:
				vm.div()
			case opcodes.OP_POP:
				vm.pop()
			case opcodes.OP_PRINT:
				vm.print()
			case opcodes.OP_DUMP:
				vm.dump()
			case opcodes.OP_DUP:
				vm.dup()
			case opcodes.OP_EQ:
				vm.eq()
			case opcodes.OP_NEQ:
				vm.neq()
			case opcodes.OP_GT:
				vm.gt()
			case opcodes.OP_LT:
				vm.lt()
			case opcodes.OP_GE:
				vm.ge()
			case opcodes.OP_LE:
				vm.le()
			}
			vm.pc++

		case opcodes.OP_STORE:
			addr := vm.program[vm.pc+1]
			vm.store(int(addr))
			vm.pc += 2

		case opcodes.OP_LOAD:
			addr := vm.program[vm.pc+1]
			vm.load(int(addr))
			vm.pc += 2

		case opcodes.OP_JUMP:
			vm.jump()

		case opcodes.OP_JUMP_IF_FALSE:
			vm.jump_if_false()

		case opcodes.OP_JUMP_IF_TRUE:
			vm.jump_if_true()

		case opcodes.OP_HALT:
			vm.running = false

		case opcodes.OP_CALL:
			vm.call()

		case opcodes.OP_RET:
			vm.ret()

		default:
			panic(fmt.Sprintf("unknown instruction: 0x%X at PC %d", instr, vm.pc))
		}
	}
}
