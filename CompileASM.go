package main

import (
	"strconv"
	"strings"
)

func CompileASM(source string) []byte {
	lines := strings.Split(source, "\n")

	labels := map[string]int{}
	pc := 0

	// resolve labels
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Label
		if strings.HasSuffix(line, ":") {
			name := strings.TrimSuffix(line, ":")
			labels[name] = pc
			continue
		}

		parts := strings.Fields(line)
		op := parts[0]

		switch op {
		case "PUSH", "LOAD", "STORE":
			pc += 2

		case "JUMP", "JUMP_IF_TRUE", "JUMP_IF_FALSE", "CALL":
			pc += 2

		case "POP", "ADD", "SUB", "MUL", "DIV",
			"EQ", "NEQ", "GT", "GE", "LT", "LE",
			"PRINT", "RET", "HALT":
			pc += 1

		default:
			panic("unknown instruction: " + op)
		}
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
			bytecode = append(bytecode, OP_PUSH, byte(v))

		case "POP":
			bytecode = append(bytecode, OP_POP)

		case "ADD":
			bytecode = append(bytecode, OP_ADD)

		case "SUB":
			bytecode = append(bytecode, OP_SUB)

		case "MUL":
			bytecode = append(bytecode, OP_MUL)

		case "DIV":
			bytecode = append(bytecode, OP_DIV)

		case "PRINT":
			bytecode = append(bytecode, OP_PRINT)

		case "EQ":
			bytecode = append(bytecode, OP_EQ)

		case "NEQ":
			bytecode = append(bytecode, OP_NEQ)

		case "GT":
			bytecode = append(bytecode, OP_GT)

		case "GE":
			bytecode = append(bytecode, OP_GE)

		case "LT":
			bytecode = append(bytecode, OP_LT)

		case "LE":
			bytecode = append(bytecode, OP_LE)

		case "JUMP":
			target := labels[parts[1]]
			bytecode = append(bytecode, OP_JUMP, byte(target))

		case "JUMP_IF_TRUE":
			target := labels[parts[1]]
			bytecode = append(bytecode, OP_JUMP_IF_TRUE, byte(target))

		case "JUMP_IF_FALSE":
			target := labels[parts[1]]
			bytecode = append(bytecode, OP_JUMP_IF_FALSE, byte(target))

		case "HALT":
			bytecode = append(bytecode, OP_HALT)

		case "LOAD":
			slot, _ := strconv.Atoi(parts[1])
			bytecode = append(bytecode, OP_LOAD, byte(slot))

		case "STORE":
			slot, _ := strconv.Atoi(parts[1])
			bytecode = append(bytecode, OP_STORE, byte(slot))

		case "CALL":
			target := labels[parts[1]]
			bytecode = append(bytecode, OP_CALL, byte(target))

		case "RET":
			bytecode = append(bytecode, OP_RET)

		default:
			panic("unknown instruction: " + op)
		}
	}

	return bytecode
}
