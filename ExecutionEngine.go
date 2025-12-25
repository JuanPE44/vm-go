package main

import "fmt"

type ExecutionEngine struct {
}

func (ee *ExecutionEngine) Execute(vm *VM) {
	vm.running = true
	for vm.running && vm.pc < len(vm.program) {
		instr := vm.program[vm.pc]
		switch instr {
		case OP_PUSH:
			value := vm.program[vm.pc+1]
			vm.push(int(value))
			vm.pc += 2

		case OP_ADD, OP_SUB, OP_MUL, OP_DIV, OP_POP, OP_PRINT, OP_EQ, OP_NEQ, OP_GT, OP_LT, OP_GE, OP_LE:
			switch instr {
			case OP_ADD:
				vm.add()
			case OP_SUB:
				vm.sub()
			case OP_MUL:
				vm.mul()
			case OP_DIV:
				vm.div()
			case OP_POP:
				vm.pop()
			case OP_PRINT:
				vm.print()
			case OP_EQ:
				vm.eq()
			case OP_NEQ:
				vm.neq()
			case OP_GT:
				vm.gt()
			case OP_LT:
				vm.lt()
			case OP_GE:
				vm.ge()
			case OP_LE:
				vm.le()
			}
			vm.pc++

		case OP_STORE:
			addr := vm.program[vm.pc+1]
			vm.store(int(addr))
			vm.pc += 2

		case OP_LOAD:
			addr := vm.program[vm.pc+1]
			vm.load(int(addr))
			vm.pc += 2

		case OP_JUMP:
			vm.jump()

		case OP_JUMP_IF_FALSE:
			vm.jump_if_false()

		case OP_JUMP_IF_TRUE:
			vm.jump_if_true()

		case OP_HALT:
			vm.running = false

		case OP_CALL:
			vm.call()

		case OP_RET:
			vm.ret()

		default:
			panic(fmt.Sprintf("unknown instruction: 0x%X at PC %d", instr, vm.pc))
		}
	}
}
