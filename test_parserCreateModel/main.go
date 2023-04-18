package main

import (
	"fmt"
)

func main() {
	input := "create model irisKnn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;"
	var p Parser
	tokens, _ := lex(input)
	stmt, _, ok := p.parseCreateModelStatement(tokens, 0, Token{})
	if ok {
		fmt.Println(stmt)
	}

}
