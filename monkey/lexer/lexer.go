package lexer

import (
    "testing"
    "monkey/token"   
)

func TestNextToken(t *testing.T) {
    input := `+(){},:`
    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)
    for i, tt := range tests{
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType,
            tok.Type)
        }
        if tok.Literal != tt.expectedLiteral{
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral,
            tok.Literal)
        }
    }
}

type Lexer struct{
    input string
    position int // current postiion in inputs ( points to the current char)
    readPostion int // current reading position in input (after current char)
    ch byte // current char under examination
}

func New(input string) *Lexer{
    l := &Lexer{input: input}
    return l
}
