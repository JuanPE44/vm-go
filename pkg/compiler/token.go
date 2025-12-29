package compiler

type TokenType int

const (
	TokenNumber TokenType = iota // 123
	TokenPlus                    // +
	TokenMinus                   // -
	TokenStar                    // *
	TokenSlash                   // /
	TokenLParen                  // (
	TokenRParen                  // )
	TokenEOF                     // Fin del archivo
)

type Token struct {
	Type    TokenType
	Literal string
}
