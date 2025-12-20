package main

func main() {
	vm := VM{
		alu: ALU{},
		program: []any{
			"PUSH", 4,
			// condicion del jump
			"PUSH", 2,
			"JUMP_IF_TRUE", 8,
			"PUSH", 2,
			"PUSH", 2,
			"PUSH", 2,
			"PRINT",
		},
	}

	vm.Run()
}
