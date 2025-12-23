package main

func main() {
	vm := VM{
		running: true,
		alu:     ALU{},
		program: []any{
			"PUSH", 1,
			"PUSH", 2,
			"PUSH", 3,
			"CALL", 15,
			"PRINT",
			"HALT",

			"PUSH", 4,
			"PUSH", 5,
			"RET",

			"PUSH", 6,
			"RET",
		},
		memory: make(map[string]int),
	}

	vm.Run()
}
