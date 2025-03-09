package lexer

import (
	"testing"
	
	"Catlang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+-{}(),;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MUL, "*"},
		{token.DIV, "/"},
		{token.MOD, "%%"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPEREN, "("},
		{token.RPEREN, ")"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d]-tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]-literal wrong. expected=%q,got=%q", i, tt.expectedLiteral, tok.Literal)

		}
	}
}
