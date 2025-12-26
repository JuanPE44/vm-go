package vm

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

func (a *ALU) Equal(x, y int) int {
	if x == y {
		return 1
	}
	return 0
}

func (a *ALU) NotEqual(x, y int) int {
	if x != y {
		return 1
	}
	return 0
}

func (a *ALU) GreaterThan(x, y int) int {
	if x > y {
		return 1
	}
	return 0
}

func (a *ALU) LessThan(x, y int) int {
	if x < y {
		return 1
	}
	return 0
}

func (a *ALU) GreaterThanOrEqual(x, y int) int {
	if x >= y {
		return 1
	}
	return 0
}

func (a *ALU) LessThanOrEqual(x, y int) int {
	if x <= y {
		return 1
	}
	return 0
}

func (a *ALU) JumpIfTrue(x, y int) int {
	if x != 0 {
		return y
	}
	return 0
}
