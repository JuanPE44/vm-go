package main

func main() {
	vm := VM{
		alu: ALU{},
		program: []any{
			"PUSH", 1,
			"PUSH", 2,
			// condicion del jump
			"PUSH", 4,
			"STORE", "x",
			"PUSH", 3,
			"PUSH", 5,
			"PUSH", 6,
			"PUSH", 7,
			"LOAD", "x",
			"PRINT",
		},
		memory: make(map[string]int),
	}

	vm.Run()
}
