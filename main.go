package main

func main() {
	vm := VM{
		alu: ALU{},
		program: []any{
			"PUSH", 1,
			"PUSH", 2,
			"PUSH", 3,
			"PUSH", 4,
			"MUL",
			"MUL",
			"MUL",
			"PRINT",
		},
	}

	vm.Run()
}
