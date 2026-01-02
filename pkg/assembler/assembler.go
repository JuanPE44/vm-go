package assembler

import (
	"strconv"
	"strings"
	"vm-go/pkg/opcodes"
)

func instrSize(op string) int {
	switch op {
	case "PUSH", "LOAD", "STORE",
		"JUMP", "JUMP_IF_TRUE", "JUMP_IF_FALSE", "CALL":
		return 2

	case "POP", "ADD", "SUB", "MUL", "DIV",
		"EQ", "NEQ", "GT", "GE", "LT", "LE",
		"PRINT", "RET", "HALT", "DUP", "DUMP", "SYS_DRAW_PIXEL", "SYS_PRESENT":
		return 1

	default:
		panic("unknown instruction: " + op)
	}
}

func CompileASM(source string) []byte {
	lines := strings.Split(source, "\n")

	labels := map[string]int{}
	pc := 0
	// labels
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			name := strings.TrimSuffix(line, ":")
			labels[name] = pc
			continue
		}

		parts := strings.Fields(line)
		op := parts[0]

		pc += instrSize(op)
	}

	// emit bytecode
	bytecode := []byte{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasSuffix(line, ":") {
			continue
		}

		parts := strings.Fields(line)
		op := parts[0]

		switch op {

		case "PUSH":
			v, _ := strconv.Atoi(parts[1])
			bytecode = append(bytecode, opcodes.OP_PUSH_INT, byte(v))

		case "POP":
			bytecode = append(bytecode, opcodes.OP_POP)

		case "ADD":
			bytecode = append(bytecode, opcodes.OP_ADD)

		case "SUB":
			bytecode = append(bytecode, opcodes.OP_SUB)

		case "MUL":
			bytecode = append(bytecode, opcodes.OP_MUL)

		case "DIV":
			bytecode = append(bytecode, opcodes.OP_DIV)

		case "PRINT":
			bytecode = append(bytecode, opcodes.OP_PRINT)

		case "DUMP":
			bytecode = append(bytecode, opcodes.OP_DUMP)

		case "DUP":
			bytecode = append(bytecode, opcodes.OP_DUP)

		case "EQ":
			bytecode = append(bytecode, opcodes.OP_EQ)

		case "NEQ":
			bytecode = append(bytecode, opcodes.OP_NEQ)

		case "GT":
			bytecode = append(bytecode, opcodes.OP_GT)

		case "GE":
			bytecode = append(bytecode, opcodes.OP_GE)

		case "LT":
			bytecode = append(bytecode, opcodes.OP_LT)

		case "LE":
			bytecode = append(bytecode, opcodes.OP_LE)

		case "JUMP":
			target := labels[parts[1]]
			bytecode = append(bytecode, opcodes.OP_JUMP, byte(target))

		case "JUMP_IF_TRUE":
			target := labels[parts[1]]
			bytecode = append(bytecode, opcodes.OP_JUMP_IF_TRUE, byte(target))

		case "JUMP_IF_FALSE":
			target := labels[parts[1]]
			bytecode = append(bytecode, opcodes.OP_JUMP_IF_FALSE, byte(target))

		case "HALT":
			bytecode = append(bytecode, opcodes.OP_HALT)

		case "LOAD":
			slot, _ := strconv.Atoi(parts[1])
			bytecode = append(bytecode, opcodes.OP_LOAD, byte(slot))

		case "STORE":
			slot, _ := strconv.Atoi(parts[1])
			bytecode = append(bytecode, opcodes.OP_STORE, byte(slot))

		case "CALL":
			target := labels[parts[1]]
			bytecode = append(bytecode, opcodes.OP_CALL, byte(target))

		case "RET":
			bytecode = append(bytecode, opcodes.OP_RET)

		case "SYS_DRAW_PIXEL":
			bytecode = append(bytecode, opcodes.OP_SYS_DRAW_PIXEL)

		case "SYS_PRESENT":
			bytecode = append(bytecode, opcodes.OP_SYS_PRESENT)

		default:
			panic("unknown instruction: " + op)
		}
	}

	return bytecode
}
