package lexer

const eofTokenTag byte = 0
const identifierTokenTag byte = 1
const stringLiteralTokenTag byte = 2

// Token represents any token emitted by the lexer
type Token interface {
}

// IdentifierToken represents an identifier
type IdentifierToken struct {
	text string
	tag  byte
}

// StringLiteralToken represents a quoted string literal
type StringLiteralToken struct {
	text string
	tag  byte
}

// EOFToken represents the last token of a file. It marks the end of a file.
type EOFToken struct {
	text string
	tag  byte
}
