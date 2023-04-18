package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Location struct {
	Line uint
	Col  uint
}

// for storing SQL reserved Keywords
type Keyword string

const (
	FromKeyword        Keyword = "from"
	AsKeyword          Keyword = "as"
	CreateModelKeyword Keyword = "create model"
	TrainModelKeyword  Keyword = "train model"
	PredictKeyword     Keyword = "predict"
	WithKeyword        Keyword = "with"
)

// for storing SQL syntax
type Symbol string

const (
	SemicolonSymbol  Symbol = ";"
	AsteriskSymbol   Symbol = "*"
	CommaSymbol      Symbol = ","
	LeftParenSymbol  Symbol = "("
	RightParenSymbol Symbol = ")"
	EqSymbol         Symbol = "="
	NeqSymbol        Symbol = "<>"
	NeqSymbol2       Symbol = "!="
	ConcatSymbol     Symbol = "||"
	PlusSymbol       Symbol = "+"
	LtSymbol         Symbol = "<"
	LteSymbol        Symbol = "<="
	GtSymbol         Symbol = ">"
	GteSymbol        Symbol = ">="
)

type TokenKind uint

const (
	KeywordKind TokenKind = iota
	SymbolKind
	IdentifierKind
	StringKind
	NumericKind
	BoolKind
	NullKind
)

type Token struct {
	Value string
	Kind  TokenKind
	Loc   Location
}

// isKeyword检查给定的字符串是否是一个SQL关键字
func isKeyword(s string) bool {
	switch Keyword(s) {
	case FromKeyword, AsKeyword, CreateModelKeyword, TrainModelKeyword, PredictKeyword, WithKeyword:
		return true
	default:
		return false
	}
}

// isSymbol检查给定的字符串是否是一个SQL符号
func isSymbol(s string) bool {
	switch Symbol(s) {
	case SemicolonSymbol, AsteriskSymbol, CommaSymbol, LeftParenSymbol, RightParenSymbol, EqSymbol, NeqSymbol, NeqSymbol2, ConcatSymbol, PlusSymbol, LtSymbol, LteSymbol, GtSymbol, GteSymbol:
		return true
	default:
		return false
	}
}

func tokenize(input string) ([]*Token, error) {
	var tokens []*Token
	var currToken string
	var currTokenKind TokenKind
	var inString bool
	var currLoc Location
	var inPredict bool // 新增一个变量，用于判断是否在predict语句中
	var inModel bool   // 新增一个变量，用于判断是否在create model或train model语句中
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	for i, r := range input {
		if r == '"' {
			if inString {
				tokens = append(tokens, &Token{Value: currToken, Kind: StringKind, Loc: currLoc})
				currToken = ""
				inString = false
			} else {
				currLoc = Location{Line: currLoc.Line, Col: uint(i)}
				inString = true
			}
		} else if inString {
			currToken += string(r)
		} else if unicode.IsSpace(r) {
			if currToken != "" {
				if ok := isKeyword(currToken); ok {
					currTokenKind = KeywordKind
					tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
				} else {
					currTokenKind = IdentifierKind
					tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
				}
				currToken = ""
			}
			if r == '\n' {
				currLoc = Location{Line: currLoc.Line + 1, Col: 0}
			} else {
				currLoc = Location{Line: currLoc.Line, Col: currLoc.Col + 1}
			}
		} else if isSymbol(string(r)) && string(r) != "," { // 忽略所有逗号
			if currToken != "" {
				if ok := isKeyword(currToken); ok {
					currTokenKind = KeywordKind
					tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
				} else {
					currTokenKind = IdentifierKind
					tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
				}
				currToken = ""
			}
			currTokenKind = SymbolKind
			if string(r) == "[" { // 新增判断，如果是"["，则将inPredict设为true
				inPredict = true
			}
			if string(r) == "]" { // 新增判断，如果是"]"，则将inPredict设为false
				inPredict = false
			}
			if inPredict || inModel { // 新增判断，如果在create model或train model语句中，则将整个"create model iris_Knn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;"或"train model iris_knn WITH euclidean;"作为一个整体
				currToken += string(r)
			} else if string(r) != "(" && string(r) != ")" && string(r) != "," { // 忽略所有括号和逗号
				tokens = append(tokens, &Token{Value: string(r), Kind: currTokenKind, Loc: currLoc})
			}
			if r == '\n' {
				currLoc = Location{Line: currLoc.Line + 1, Col: 0}
			} else {
				currLoc = Location{Line: currLoc.Line, Col: currLoc.Col + 1}
			}
		} else {
			currToken += string(r)
		}
		if currToken == "create" || currToken == "train" { // 新增判断，如果当前token是"create"或"train"，则将inModel设为true
			inModel = true
		}
		if currToken == "model" { // 新增判断，如果当前token是"model"，则将inModel设为false
			inModel = false
		}
	}
	if currToken != "" {
		if ok := isKeyword(currToken); ok {
			currTokenKind = KeywordKind
			tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
		} else {
			currTokenKind = IdentifierKind
			tokens = append(tokens, &Token{Value: currToken, Kind: currTokenKind, Loc: currLoc})
		}
	}
	return tokens, nil

}

func main() {
	source := "create model iris_Knn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;"
	tokens, err := tokenize(source)
	if err != nil {
		// handle error
	}

	for _, token := range tokens {
		fmt.Println(token.Value)
	}
	fmt.Println()
	source = "train model iris_knn WITH euclidean;"

	tokens, err = tokenize(source)
	if err != nil {
		// handle error
	}

	for _, token := range tokens {
		fmt.Println(token.Value)
	}
	fmt.Println()
	source = "PREDICT(iris_knn, [5.1, 3.5, 1.4, 0.2]) FROM iris_table;"
	tokens, err = tokenize(source)
	if err != nil {
		// handle error
	}

	for _, token := range tokens {
		fmt.Println(token.Value)
	}

}
