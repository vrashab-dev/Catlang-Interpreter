package token

// TokenType represents the type of token
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	INDENT = "INDENT"
	INT    = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	MOD    = "%"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPEREN = "("
	RPEREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
)
