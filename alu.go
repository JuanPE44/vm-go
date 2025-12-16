package main

type ALU struct{}

func (a *ALU) Sum(x, y int) int {
	return x + y
}

func (a *ALU) Res(x, y int) int {
	return x - y
}

func (a *ALU) Mul(x, y int) int {
	return x * y
}

func (a *ALU) Div(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}
