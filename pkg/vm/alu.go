package vm

type ALU struct{}

func (alu *ALU) Add(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.Sum")
	}

	return Value{
		Type: IntType,
		Int:  a.Int + b.Int,
	}
}

func (alu *ALU) Sub(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.Sub")
	}

	return Value{
		Type: IntType,
		Int:  a.Int - b.Int,
	}
}

func (alu *ALU) Mul(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.Sub")
	}

	return Value{
		Type: IntType,
		Int:  a.Int * b.Int,
	}
}

func (alu *ALU) Div(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.Div")
	}

	return Value{
		Type: IntType,
		Int:  a.Int / b.Int,
	}
}

func (alu *ALU) Equal(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.Equal")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int == b.Int,
	}
}

func (alu *ALU) NotEqual(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.NotEqual")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int != b.Int,
	}
}

func (alu *ALU) GreaterThan(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.GreaterThan")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int > b.Int,
	}
}

func (alu *ALU) LessThan(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.LessThan")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int < b.Int,
	}
}

func (alu *ALU) GreaterThanOrEqual(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.GreaterThanOrEqual")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int >= b.Int,
	}
}

func (alu *ALU) LessThanOrEqual(a, b Value) Value {
	if a.Type != IntType || b.Type != IntType {
		panic("type error in ALU.LessThanOrEqual")
	}

	return Value{
		Type: BoolType,
		Bool: a.Int <= b.Int,
	}
}
