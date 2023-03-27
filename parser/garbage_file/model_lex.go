// package parser

// import (
// 	"fmt"
// 	"strings"
// 	"text/scanner"
// )

// type TokenType int

// const (
// 	_ TokenType = iota
// 	CREATE
// 	MODEL
// 	SELECT
// 	FROM
// 	ID
// 	STRING
// 	COMMA
// )

// func (t TokenType) String() string {
// 	switch t {
// 	case CREATE:
// 		return "CREATE"
// 	case MODEL:
// 		return "MODEL"
// 	case SELECT:
// 		return "SELECT"
// 	case FROM:
// 		return "FROM"
// 	case ID:
// 		return "ID"
// 	case STRING:
// 		return "STRING"
// 	case COMMA:
// 		return "COMMA"
// 	default:
// 		panic(fmt.Sprintf("unknown token type: %d", t))
// 	}
// }

// type Token struct {
// 	Type  TokenType
// 	Value string
// }

// type Lexer struct {
// 	input   string
// 	scanner scanner.Scanner
// 	tokens  []Token
// }

// func NewLexer(input string) *Lexer {
// 	var lex Lexer

// 	lex.input = input
// 	lex.scanner.Init(strings.NewReader(input))
// 	lex.scanner.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanChars | scanner.SkipComments
// 	lex.scanner.Whitespace ^= 1 << '\n' // ignore newlines

// 	for tok := lex.nextToken(); tok.Type != scanner.EOF; {
// 		lex.tokens = append(lex.tokens, tok)
// 	}
// 	lex.tokens = append(lex.tokens, Token{Type: scanner.EOF})

// 	return &lex
// }

// func (l *Lexer) nextToken() Token {
// 	if l.scanner.Error != nil {
// 		panic(fmt.Sprintf("scanner error: %s", l.scanner.Error))
// 	}

// 	pos, tok, lit := l.scanner.Scan()
// 	switch tok {
// 	case scanner.EOF:
// 		return Token{Type: scanner.EOF}
// 	case scanner.Ident:
// 		if pos.Line == 1 && (lit == "CREATE" || lit == "MODEL" || lit == "SELECT" || lit == "FROM") {
// 			return Token{Type: TokenType(lit), Value: lit}
// 		} else {
// 			return Token{Type: ID, Value: lit}
// 		}
// 	case scanner.String:
// 		return Token{Type: STRING, Value: lit}
// 	case ',':
// 		return Token{Type: COMMA, Value: string(tok)}
// 	default:
// 		panic(fmt.Sprintf("unexpected token %q", tok))
// 	}
// }

// func lex(input string) []Token {
// 	lexer := NewLexer(input)
// 	tokens := make([]Token, len(lexer.tokens))
// 	for i, tok := range lexer.tokens {
// 		tokens[i] = tok
// 	}

// 	return tokens
// }
