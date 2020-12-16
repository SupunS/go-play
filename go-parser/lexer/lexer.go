package lexer

import (
	"testGo/readers"
)

// Lexer represents a lexer instance
type Lexer struct {
	reader *readers.CharReader
}

// NewLexer initialize a new lexer instance
func NewLexer(reader *readers.CharReader) (lexer Lexer) {
	lexer.reader = reader
	return
}

// NextToken retrieves the next token in the input stream
func (lexer *Lexer) NextToken() Token {
	for !lexer.reader.IsEOF() {
		char := lexer.reader.Peek()

		switch char {
		case '"':
			return lexer.processStringLiteral()
		}
		if isAlphaNumeric(char) {
			return lexer.processIdentifierToken()
		}

		lexer.reader.Consume()
	}

	return EOFToken{"EOF", eofTokenTag}
}

func isAlphaNumeric(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}

func (lexer *Lexer) processIdentifierToken() Token {
	var tokenText string = ""
	for !lexer.reader.IsEOF() {
		char := lexer.reader.Peek()

		if !isAlphaNumeric(char) {
			break
		}

		lexer.reader.Consume()
		tokenText += string(char)
	}

	return IdentifierToken{tokenText, identifierTokenTag}
}

func (lexer *Lexer) processStringLiteral() Token {
	// Starting double-quote has be read and verified before coming here.
	// Hence consume and append to the content
	char := lexer.reader.Consume()
	var content string = string(char)

	for !lexer.reader.IsEOF() {
		char = lexer.reader.Consume()
		content += string(char)
		switch char {
		case '"':
			break
		case '\\':
			// consume two tokens
			content += string(char)
			char = lexer.reader.Consume()
			continue
		default:
			continue
		}

		break
	}

	return StringLiteralToken{content, stringLiteralTokenTag}
}
