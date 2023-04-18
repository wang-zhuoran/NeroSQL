// package main

// // import (
// // 	"bytes"
// // 	"fmt"
// // 	"strings"
// // )

// // type TokenType int
// // type State int

// // const (
// // 	Keyword TokenType = iota
// // 	Identifier
// // 	Symbol
// // )

// // const (
// // 	Normal State = iota
// // 	InArray
// // )

// // type Token struct {
// // 	Type  TokenType
// // 	Value string
// // }

// // func Lex(input string) []Token {
// // 	var tokens []Token
// // 	state := Normal
// // 	var buffer bytes.Buffer
// // 	for _, word := range strings.FieldsFunc(input, func(char rune) bool {
// // 		return char == ' ' || char == '(' || char == ')' || char == ';' || char == '[' || char == ']'
// // 	}) {
// // 		switch strings.ToUpper(word) {
// // 		case "CREATE":
// // 			tokens = append(tokens, Token{Keyword, "CREATE"})
// // 		case "MODEL":
// // 			tokens = append(tokens, Token{Keyword, "MODEL"})
// // 		case "AS":
// // 			tokens = append(tokens, Token{Keyword, "AS"})
// // 		case "SELECT":
// // 			tokens = append(tokens, Token{Keyword, "SELECT"})
// // 		case "FROM":
// // 			tokens = append(tokens, Token{Keyword, "FROM"})
// // 		case "TRAIN":
// // 			tokens = append(tokens, Token{Keyword, "TRAIN"})
// // 		case "WITH":
// // 			tokens = append(tokens, Token{Keyword, "WITH"})
// // 		case "PREDICT":
// // 			tokens = append(tokens, Token{Keyword, "PREDICT"})
// // 		case "ARRAY":
// // 			tokens = append(tokens, Token{Keyword, "ARRAY"})
// // 		default:
// // 			if len(word) > 0 {
// // 				tokens = append(tokens, Token{Identifier, word})
// // 			}
// // 		}
// // 		// for _, char := range word {
// // 		// 	switch char {
// // 		// 	case '(':
// // 		// 		tokens = append(tokens, Token{Symbol, "("})
// // 		// 	case ')':
// // 		// 		tokens = append(tokens, Token{Symbol, ")"})
// // 		// 	case ';':
// // 		// 		tokens = append(tokens, Token{Symbol, ";"})
// // 		// 	case '[':
// // 		// 		tokens = append(tokens, Token{Symbol, "["})
// // 		// 	case ']':
// // 		// 		tokens = append(tokens, Token{Symbol, "]"})
// // 		// 	}
// // 		// }

// // 		for _, char := range word {
// // 			switch state {
// // 			case Normal:
// // 				if char == '[' {
// // 					state = InArray
// // 					buffer.WriteRune(char)
// // 				} else if char == ' ' || char == '(' || char == ')' || char == ';' {
// // 					switch char {
// // 					case '(':
// // 						tokens = append(tokens, Token{Symbol, "("})
// // 					case ')':
// // 						tokens = append(tokens, Token{Symbol, ")"})
// // 					case ';':
// // 						tokens = append(tokens, Token{Symbol, ";"})
// // 					}
// // 				} else {
// // 					buffer.WriteRune(char)
// // 				}
// // 			case InArray:
// // 				buffer.WriteRune(char)
// // 				if char == ']' {
// // 					tokens = append(tokens, Token{Symbol, buffer.String()})
// // 					buffer.Reset()
// // 					state = Normal
// // 				}
// // 			}
// // 		}
// // 	}
// // 	return tokens
// // }

// import (
// 	"fmt"
// 	"strings"
// )

// // location of the token in source code
// type Location struct {
// 	Line uint
// 	Col  uint
// }

// // for storing SQL reserved Keywords
// type Keyword string

// const (
// 	SelectKeyword      Keyword = "select"
// 	FromKeyword        Keyword = "from"
// 	AsKeyword          Keyword = "as"
// 	TableKeyword       Keyword = "table"
// 	CreateKeyword      Keyword = "create"
// 	DropKeyword        Keyword = "drop"
// 	InsertKeyword      Keyword = "insert"
// 	IntoKeyword        Keyword = "into"
// 	ValuesKeyword      Keyword = "values"
// 	IntKeyword         Keyword = "int"
// 	TextKeyword        Keyword = "text"
// 	BoolKeyword        Keyword = "boolean"
// 	WhereKeyword       Keyword = "where"
// 	AndKeyword         Keyword = "and"
// 	OrKeyword          Keyword = "or"
// 	TrueKeyword        Keyword = "true"
// 	FalseKeyword       Keyword = "false"
// 	UniqueKeyword      Keyword = "unique"
// 	IndexKeyword       Keyword = "index"
// 	OnKeyword          Keyword = "on"
// 	PrimarykeyKeyword  Keyword = "primary key"
// 	NullKeyword        Keyword = "null"
// 	LimitKeyword       Keyword = "limit"
// 	OffsetKeyword      Keyword = "offset"
// 	DeleteKeyword      Keyword = "delete"
// 	CreateModelKeyword Keyword = "create model"
// 	TrainModelKeyword  Keyword = "train model"
// 	PredictKeyword     Keyword = "predict"
// 	WithKeyword        Keyword = "with"
// )

// // for storing SQL syntax
// type Symbol string

// const (
// 	SemicolonSymbol  Symbol = ";"
// 	AsteriskSymbol   Symbol = "*"
// 	CommaSymbol      Symbol = ","
// 	LeftParenSymbol  Symbol = "("
// 	RightParenSymbol Symbol = ")"
// 	EqSymbol         Symbol = "="
// 	NeqSymbol        Symbol = "<>"
// 	NeqSymbol2       Symbol = "!="
// 	ConcatSymbol     Symbol = "||"
// 	PlusSymbol       Symbol = "+"
// 	LtSymbol         Symbol = "<"
// 	LteSymbol        Symbol = "<="
// 	GtSymbol         Symbol = ">"
// 	GteSymbol        Symbol = ">="
// )

// type TokenKind uint

// const (
// 	KeywordKind TokenKind = iota
// 	SymbolKind
// 	IdentifierKind
// 	StringKind
// 	NumericKind
// 	BoolKind
// 	NullKind
// )

// type Token struct {
// 	Value string
// 	Kind  TokenKind
// 	Loc   Location
// }

// func (t Token) bindingPower() uint {
// 	switch t.Kind {
// 	case KeywordKind:
// 		switch Keyword(t.Value) {
// 		case AndKeyword:
// 			fallthrough
// 		case OrKeyword:
// 			return 1
// 		}
// 	case SymbolKind:
// 		switch Symbol(t.Value) {
// 		case EqSymbol:
// 			fallthrough
// 		case NeqSymbol:
// 			return 2

// 		case LtSymbol:
// 			fallthrough
// 		case GtSymbol:
// 			return 3

// 		// For some reason these are grouped separately
// 		case LteSymbol:
// 			fallthrough
// 		case GteSymbol:
// 			return 4

// 		case ConcatSymbol:
// 			fallthrough
// 		case PlusSymbol:
// 			return 5
// 		}
// 	}

// 	return 0
// }

// func (t *Token) equals(other *Token) bool {
// 	return t.Value == other.Value && t.Kind == other.Kind
// }

// // cursor indicates the current position of the lexer
// type cursor struct {
// 	pointer uint
// 	loc     Location
// }

// // longestMatch iterates through a source string starting at the given
// // cursor to find the longest matching substring among the provided
// // options
// func longestMatch(source string, ic cursor, options []string) string {
// 	var value []byte
// 	var skipList []int
// 	var match string

// 	cur := ic

// 	for cur.pointer < uint(len(source)) {

// 		value = append(value, strings.ToLower(string(source[cur.pointer]))...)
// 		cur.pointer++

// 	match:
// 		for i, option := range options {
// 			for _, skip := range skipList {
// 				if i == skip {
// 					continue match
// 				}
// 			}

// 			// Deal with cases like INT vs INTO
// 			if option == string(value) {
// 				skipList = append(skipList, i)
// 				if len(option) > len(match) {
// 					match = option
// 				}

// 				continue
// 			}

// 			sharesPrefix := string(value) == option[:cur.pointer-ic.pointer]
// 			tooLong := len(value) > len(option)
// 			if tooLong || !sharesPrefix {
// 				skipList = append(skipList, i)
// 			}
// 		}

// 		if len(skipList) == len(options) {
// 			break
// 		}
// 	}

// 	return match
// }

// func lexSymbol(source string, ic cursor) (*Token, cursor, bool) {
// 	c := source[ic.pointer]
// 	cur := ic
// 	// Will get overwritten later if not an ignored syntax
// 	cur.pointer++
// 	cur.loc.Col++

// 	switch c {
// 	// Syntax that should be thrown away
// 	case '\n':
// 		cur.loc.Line++
// 		cur.loc.Col = 0
// 		fallthrough
// 	case '\t':
// 		fallthrough
// 	case ' ':
// 		return nil, cur, true
// 	}

// 	// Syntax that should be kept
// 	Symbols := []Symbol{
// 		EqSymbol,
// 		NeqSymbol,
// 		NeqSymbol2,
// 		LtSymbol,
// 		LteSymbol,
// 		GtSymbol,
// 		GteSymbol,
// 		ConcatSymbol,
// 		PlusSymbol,
// 		CommaSymbol,
// 		LeftParenSymbol,
// 		RightParenSymbol,
// 		SemicolonSymbol,
// 		AsteriskSymbol,
// 	}

// 	var options []string
// 	for _, s := range Symbols {
// 		options = append(options, string(s))
// 	}

// 	// Use `ic`, not `cur`
// 	match := longestMatch(source, ic, options)
// 	// Unknown character
// 	if match == "" {
// 		return nil, ic, false
// 	}

// 	cur.pointer = ic.pointer + uint(len(match))
// 	cur.loc.Col = ic.loc.Col + uint(len(match))

// 	// != is rewritten as <>: https://www.postgresql.org/docs/9.5/functions-comparison.html
// 	if match == string(NeqSymbol2) {
// 		match = string(NeqSymbol)
// 	}

// 	return &Token{
// 		Value: match,
// 		Loc:   ic.loc,
// 		Kind:  SymbolKind,
// 	}, cur, true
// }

// func lexKeyword(source string, ic cursor) (*Token, cursor, bool) {
// 	cur := ic
// 	Keywords := []Keyword{
// 		SelectKeyword,
// 		InsertKeyword,
// 		ValuesKeyword,
// 		TableKeyword,
// 		CreateKeyword,
// 		DropKeyword,
// 		WhereKeyword,
// 		FromKeyword,
// 		IntoKeyword,
// 		TextKeyword,
// 		BoolKeyword,
// 		IntKeyword,
// 		AndKeyword,
// 		OrKeyword,
// 		AsKeyword,
// 		TrueKeyword,
// 		FalseKeyword,
// 		UniqueKeyword,
// 		IndexKeyword,
// 		OnKeyword,
// 		PrimarykeyKeyword,
// 		NullKeyword,
// 		LimitKeyword,
// 		OffsetKeyword,
// 		DeleteKeyword,
// 		CreateModelKeyword,
// 		TrainModelKeyword,
// 		PredictKeyword,
// 		WithKeyword,
// 	}

// 	var options []string
// 	for _, k := range Keywords {
// 		options = append(options, string(k))
// 	}

// 	match := longestMatch(source, ic, options)
// 	if match == "" {
// 		return nil, ic, false
// 	}

// 	cur.pointer = ic.pointer + uint(len(match))
// 	cur.loc.Col = ic.loc.Col + uint(len(match))

// 	Kind := KeywordKind
// 	if match == string(TrueKeyword) || match == string(FalseKeyword) {
// 		Kind = BoolKind
// 	}

// 	if match == string(NullKeyword) {
// 		Kind = NullKind
// 	}

// 	return &Token{
// 		Value: match,
// 		Kind:  Kind,
// 		Loc:   ic.loc,
// 	}, cur, true
// }

// func lexNumeric(source string, ic cursor) (*Token, cursor, bool) {
// 	cur := ic

// 	periodFound := false
// 	expMarkerFound := false

// 	for ; cur.pointer < uint(len(source)); cur.pointer++ {
// 		c := source[cur.pointer]
// 		cur.loc.Col++

// 		isDigit := c >= '0' && c <= '9'
// 		isPeriod := c == '.'
// 		isExpMarker := c == 'e'

// 		// Must start with a digit or period
// 		if cur.pointer == ic.pointer {
// 			if !isDigit && !isPeriod {
// 				return nil, ic, false
// 			}

// 			periodFound = isPeriod
// 			continue
// 		}

// 		if isPeriod {
// 			if periodFound {
// 				return nil, ic, false
// 			}

// 			periodFound = true
// 			continue
// 		}

// 		if isExpMarker {
// 			if expMarkerFound {
// 				return nil, ic, false
// 			}

// 			// No periods allowed after expMarker
// 			periodFound = true
// 			expMarkerFound = true

// 			// expMarker must be followed by digits
// 			if cur.pointer == uint(len(source)-1) {
// 				return nil, ic, false
// 			}

// 			cNext := source[cur.pointer+1]
// 			if cNext == '-' || cNext == '+' {
// 				cur.pointer++
// 				cur.loc.Col++
// 			}
// 			continue
// 		}

// 		if !isDigit {
// 			break
// 		}
// 	}

// 	// No characters accumulated
// 	if cur.pointer == ic.pointer {
// 		return nil, ic, false
// 	}

// 	return &Token{
// 		Value: source[ic.pointer:cur.pointer],
// 		Loc:   ic.loc,
// 		Kind:  NumericKind,
// 	}, cur, true
// }

// // lexCharacterDelimited looks through a source string starting at the
// // given cursor to find a start- and end- delimiter. The delimiter can
// // be escaped be preceeding the delimiter with itself.
// func lexCharacterDelimited(source string, ic cursor, delimiter byte) (*Token, cursor, bool) {
// 	cur := ic

// 	if len(source[cur.pointer:]) == 0 {
// 		return nil, ic, false
// 	}

// 	if source[cur.pointer] != delimiter {
// 		return nil, ic, false
// 	}

// 	cur.loc.Col++
// 	cur.pointer++

// 	var value []byte
// 	for ; cur.pointer < uint(len(source)); cur.pointer++ {
// 		c := source[cur.pointer]

// 		if c == delimiter {
// 			// SQL escapes are via double characters, not backslash.
// 			if cur.pointer+1 >= uint(len(source)) || source[cur.pointer+1] != delimiter {
// 				cur.pointer++
// 				cur.loc.Col++
// 				return &Token{
// 					Value: string(value),
// 					Loc:   ic.loc,
// 					Kind:  StringKind,
// 				}, cur, true
// 			}
// 			value = append(value, delimiter)
// 			cur.pointer++
// 			cur.loc.Col++
// 		}

// 		value = append(value, c)
// 		cur.loc.Col++
// 	}

// 	return nil, ic, false
// }

// func lexIdentifier(source string, ic cursor) (*Token, cursor, bool) {
// 	// Handle separately if is a double-quoted identifier
// 	if token, newCursor, ok := lexCharacterDelimited(source, ic, '"'); ok {
// 		// Overwrite from stringkind to identifierkind
// 		token.Kind = IdentifierKind
// 		return token, newCursor, true
// 	}

// 	cur := ic

// 	c := source[cur.pointer]
// 	// Other characters count too, big ignoring non-ascii for now
// 	isAlphabetical := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
// 	if !isAlphabetical {
// 		return nil, ic, false
// 	}
// 	cur.pointer++
// 	cur.loc.Col++

// 	value := []byte{c}
// 	for ; cur.pointer < uint(len(source)); cur.pointer++ {
// 		c = source[cur.pointer]

// 		// Other characters count too, big ignoring non-ascii for now
// 		isAlphabetical := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
// 		isNumeric := c >= '0' && c <= '9'
// 		if isAlphabetical || isNumeric || c == '$' || c == '_' {
// 			value = append(value, c)
// 			cur.loc.Col++
// 			continue
// 		}

// 		break
// 	}

// 	return &Token{
// 		// Unquoted identifiers are case-insensitive
// 		Value: strings.ToLower(string(value)),
// 		Loc:   ic.loc,
// 		Kind:  IdentifierKind,
// 	}, cur, true
// }

// func lexString(source string, ic cursor) (*Token, cursor, bool) {
// 	return lexCharacterDelimited(source, ic, '\'')
// }

// func lexCreateModel(source string, ic cursor) (*Token, cursor, bool) {
// 	cur := ic
// 	// CreateModelSymbols := []Symbol{
// 	// 	// list of symbols that can appear after CREATE MODEL keyword...
// 	// 	Symbol{Value: "(", Kind: KeywordKind},
// 	// 	Symbol{Value: "AS", Kind: KeywordKind},
// 	// 	Symbol{Value: "USING", Kind: KeywordKind},
// 	// 	Symbol{Value: "WITH", Kind: KeywordKind},
// 	// }

// 	// check if source starts with CREATE MODEL keyword
// 	if !strings.HasPrefix(source[ic.pointer:], string(CreateModelKeyword)) {
// 		return nil, ic, false
// 	}

// 	// move cursor past CREATE MODEL keyword
// 	cur.pointer += uint(len(CreateModelKeyword))
// 	cur.loc.Col += uint(len(CreateModelKeyword))

// 	// extract model_name
// 	modelName, newCursor, ok := lexIdentifier(source, cur)
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	// extract AS keyword
// 	if !strings.HasPrefix(source[cur.pointer:], " AS ") {
// 		return nil, ic, false
// 	}
// 	cur.pointer += 4
// 	cur.loc.Col += 4

// 	// extract model type
// 	modelType, newCursor, ok := lexIdentifier(source, cur)
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	// extract arguments
// 	if !strings.HasPrefix(source[cur.pointer:], "(") {
// 		return nil, ic, false
// 	}
// 	cur.pointer++
// 	cur.loc.Col++

// 	args, newCursor, ok := lexCharacterDelimited(source, cur, ')')
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	return &Token{
// 		Value: string(CreateModelKeyword) + " " + modelName.Value + " AS " + modelType.Value + args.Value,
// 		Kind:  KeywordKind,
// 		Loc:   ic.loc,
// 	}, cur, true
// }

// func lexTrainModel(source string, ic cursor) (*Token, cursor, bool) {
// 	cur := ic
// 	// TrainModelSymbols := []Symbol{
// 	// 	// list of symbols that can appear after TRAIN MODEL keyword...
// 	// 	Symbol{Value: "WITH", Kind: KeywordKind},
// 	// }

// 	// check if source starts with TRAIN MODEL keyword
// 	if !strings.HasPrefix(source[ic.pointer:], string(TrainModelKeyword)) {
// 		return nil, ic, false
// 	}

// 	// move cursor past TRAIN MODEL keyword
// 	cur.pointer += uint(len(TrainModelKeyword))
// 	cur.loc.Col += uint(len(TrainModelKeyword))

// 	// extract model_name
// 	modelName, newCursor, ok := lexIdentifier(source, cur)
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	// extract WITH keyword
// 	if !strings.HasPrefix(source[cur.pointer:], " with ") {
// 		return nil, ic, false
// 	}
// 	cur.pointer += 6
// 	cur.loc.Col += 6

// 	// extract arguments
// 	if !strings.HasPrefix(source[cur.pointer:], "(") {
// 		return nil, ic, false
// 	}
// 	cur.pointer++
// 	cur.loc.Col++

// 	args, newCursor, ok := lexCharacterDelimited(source, cur, ')')
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	return &Token{
// 		Value: string(TrainModelKeyword) + " " + modelName.Value + " with " + args.Value,
// 		Kind:  KeywordKind,
// 		Loc:   ic.loc,
// 	}, cur, true
// }

// func lexPredict(source string, ic cursor) (*Token, cursor, bool) {
// 	cur := ic

// 	// check if source starts with PREDICT keyword
// 	if !strings.HasPrefix(source[ic.pointer:], string(PredictKeyword)) {
// 		return nil, ic, false
// 	}

// 	// move cursor past PREDICT keyword
// 	cur.pointer += uint(len(PredictKeyword))
// 	cur.loc.Col += uint(len(PredictKeyword))

// 	// extract model_name
// 	modelName, newCursor, ok := lexIdentifier(source, cur)
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	// extract input
// 	if !strings.HasPrefix(source[cur.pointer:], ", ") {
// 		return nil, ic, false
// 	}
// 	cur.pointer += 2
// 	cur.loc.Col += 2

// 	input, newCursor, ok := lexCharacterDelimited(source, cur, ')')
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	// extract new_table
// 	if !strings.HasPrefix(source[cur.pointer:], " from ") {
// 		return nil, ic, false
// 	}
// 	cur.pointer += 6
// 	cur.loc.Col += 6

// 	newTable, newCursor, ok := lexIdentifier(source, cur)
// 	if !ok {
// 		return nil, ic, false
// 	}
// 	cur = newCursor

// 	return &Token{
// 		Value: string(PredictKeyword) + " " + modelName.Value + ", " + input.Value + " FROM " + newTable.Value,
// 		Kind:  KeywordKind,
// 		Loc:   ic.loc,
// 	}, cur, true
// }

// type lexer func(string, cursor) (*Token, cursor, bool)

// // lex splits an input string into a list of Tokens. This process
// // can be divided into following tasks:
// //
// // 1. Instantiating a cursor with pointing to the start of the string
// //
// // 2. Execute all the lexers in series.
// //
// // 3. If any of the lexer generate a Token then add the Token to the
// // Token slice, update the cursor and restart the process from the new
// // cursor location.
// func lex(source string) ([]*Token, error) {
// 	var tokens []*Token
// 	cur := cursor{}

// lex:
// 	for cur.pointer < uint(len(source)) {
// 		lexers := []lexer{lexCreateModel, lexTrainModel, lexPredict}
// 		// var isAIOperation bool
// 		for _, l := range lexers {
// 			if token, newCursor, ok := l(source, cur); ok {
// 				cur = newCursor

// 				// Omit nil tokens for valid, but empty syntax like newlines
// 				if token != nil {
// 					tokens = append(tokens, token)
// 				}

// 				// if operation lexCreateModel, lexTrainModel, lexPredict success, then return tokens directly and do not do other lex
// 				return tokens, nil
// 			}
// 		}

// 		lexers = []lexer{lexKeyword, lexSymbol, lexString, lexNumeric, lexIdentifier}
// 		for _, l := range lexers {
// 			if token, newCursor, ok := l(source, cur); ok {
// 				cur = newCursor

// 				// Omit nil tokens for valid, but empty syntax like newlines
// 				if token != nil {
// 					tokens = append(tokens, token)
// 				}

// 				continue lex
// 			}
// 		}

// 		hint := ""
// 		if len(tokens) > 0 {
// 			hint = " after " + tokens[len(tokens)-1].Value
// 		}
// 		for _, t := range tokens {
// 			fmt.Println(t.Value)
// 		}
// 		return nil, fmt.Errorf("unable to lex token%s, at %d:%d", hint, cur.loc.Line, cur.loc.Col)
// 	}

// 	return tokens, nil
// }

// func main() {
// 	// create_model_statement := "CREATE MODEL iris_knn AS kNN('features', 'label', 3) AS output FROM iris_table;"
// 	// train_model_statement := "TRAIN MODEL iris_knn WITH distance_function='euclidean';"
// 	pred_statement := "PREDICT(iris_knn, input) FROM iris_table;"
// 	// tokens := Lex(create_model_statement)
// 	// for _, token := range tokens {
// 	// 	fmt.Println(token)
// 	// }
// 	// fmt.Println()
// 	// for _, token := range Lex(train_model_statement) {
// 	// 	fmt.Println(token)
// 	// }
// 	// for _, token := range Lex(pred_statement) {
// 	// 	fmt.Println(token)
// 	// }

// 	// tokens, err := lex(create_model_statement)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// for _, token := range tokens {
// 	// 	fmt.Println(token.Value)
// 	// }

// 	// tokens, err := lex(train_model_statement)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// for _, token := range tokens {
// 	// 	fmt.Println(token.Value)
// 	// }

// 	tokens, err := lex(pred_statement)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, token := range tokens {
// 		fmt.Println(token.Value)
// 	}

// }

package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

// LinearRegression predicts the output based on the input features using linear regression
func LinearRegression(features [][]float64, label []float64, optimizer string, learning_rate float64, epochs int) []float64 {
	numFeatures := len(features[0])
	numSamples := len(features)

	// Initialize weights with zeros
	weights := make([]float64, numFeatures)

	// Gradient descent to update weights
	alpha := learning_rate
	beta := 0.9 // used for momentum optimizer, 不想费劲让用户自定义了，就先这样吧。。。
	numIterations := epochs
	var prevGradient []float64
	for iter := 0; iter < numIterations; iter++ {
		for i := 0; i < numSamples; i++ {
			prediction := 0.0
			for j := 0; j < numFeatures; j++ {
				prediction += weights[j] * features[i][j]
			}
			error := label[i] - prediction
			gradient := make([]float64, numFeatures)
			for j := 0; j < numFeatures; j++ {
				gradient[j] = error * features[i][j]
				if optimizer == "sgd" {
					weights[j] += alpha * gradient[j]
				} else if optimizer == "momentum" {
					if iter == 0 {
						prevGradient = make([]float64, numFeatures)
					}
					prevGradient[j] = alpha*gradient[j] + beta*prevGradient[j]
					weights[j] += prevGradient[j]
				} else if optimizer == "adam" {
					beta1 := 0.9                                   // Exponential decay rate for the first moment estimates
					beta2 := 0.999                                 // Exponential decay rate for the second moment estimates
					epsilon := 1e-8                                // A small constant for numerical stability
					var m []float64 = make([]float64, numFeatures) // First moment vector
					var v []float64 = make([]float64, numFeatures) // Second moment vector
					t := iter + 1                                  // Bias correction term
					for j := 1; j < numFeatures; j++ {
						m[j-1] = beta1*m[j-1] + (1-beta1)*gradient[j]             // Update first moment estimate
						v[j-1] = beta2*v[j-1] + (1-beta2)*gradient[j]*gradient[j] // Update second moment estimate
						mHat := m[j-1] / (1 - math.Pow(beta1, float64(t)))        // Compute bias-corrected first moment estimate
						vHat := v[j-1] / (1 - math.Pow(beta2, float64(t)))        // Compute bias-corrected second moment estimate
						weights[j] += alpha * mHat / (math.Sqrt(vHat) + epsilon)  // Update weights
					}
				} else {
					panic("Invalid optimizer")
				}
			}
		}
	}

	return weights
}

func main() {
	// Open the CSV file
	file, err := os.Open("iris.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Convert the records to [][]float64
	var data [][]float64
	for _, record := range records {
		var row []float64
		for _, value := range record {
			// Parse the value as a float
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				panic(err)
			}
			row = append(row, floatValue)
		}
		data = append(data, row)
	}
	fmt.Print(data)

	features := data[:len(data)-1]
	label := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		label[i] = data[i][len(data[i])-1]
	}
	fmt.Print(features)
	fmt.Print(label)
	optimizer := "sgd"
	learning_rate := 0.01
	epochs := 100

	weights := LinearRegression(features, label, optimizer, learning_rate, epochs)
	fmt.Println(weights)

	input := []float64{5.1, 3.5, 1.4, 0.2} // example input features

	prediction := 0.0
	for i := 0; i < len(weights); i++ {
		prediction += weights[i] * input[i]
	}

	fmt.Println(prediction)
}
