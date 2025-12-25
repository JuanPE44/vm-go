package main

func main() {
	vm := VM{
		alu: ALU{},
		program: []byte{
			OP_PUSH, 1,
			OP_PUSH, 2,
			OP_PUSH, 3,
			OP_CALL, 15,
			OP_PRINT,
			OP_HALT,

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
