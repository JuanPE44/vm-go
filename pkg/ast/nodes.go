package ast

type Node interface {
	isNode()
}

type NumberNode struct {
	Value int
}

func (n *NumberNode) isNode() {}

type BinaryOpNode struct {
	Left     Node
	Operator string // "+", "-", "*", "/"
	Right    Node
}

func (b *BinaryOpNode) isNode() {}
