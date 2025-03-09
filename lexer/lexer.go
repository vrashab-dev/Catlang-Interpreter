package lexer

type Lexer struct {
	input string
	postion int
	readPosition int
     ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}