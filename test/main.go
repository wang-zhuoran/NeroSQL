package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isAIOperation(s string) bool {
	tokens := extractKeywords(s)
	if len(tokens) < 2 {
		return false
	}
	switch tokens[0] {
	case "create", "train":
		return len(tokens) >= 2 && tokens[1] == "model"
	case "select":
		return len(tokens) >= 2 && tokens[1] == "predict"
	default:
		return false
	}
}

func extractKeywords(s string) []string {
	reg, err := regexp.Compile("[a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	matches := reg.FindAllString(s, -1)
	var keywords []string
	for _, match := range matches {
		if len(match) >= 1 { // 只保留长度超过2个字符的单词作为关键字
			keyword := strings.ToLower(match) // 统一转换成小写
			keywords = append(keywords, keyword)
		}
	}
	return keywords
}

func readUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	if scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
	}

	if isAIOperation(line) {
		for {
			if scanner.Scan() {
				newLine := strings.TrimSpace(scanner.Text())
				line = line + " " + newLine
				// fmt.Println(line)
				if strings.LastIndex(newLine, ";") == len(newLine)-1 {
					// fmt.Println(line)
					// fmt.Println("detected")
					break
				}
			}
		}

	}
	return line, nil

}

func main() {
	// fmt.Println(extractKeywords("select predict()"))
	// fmt.Println(extractKeywords("select * from table"))
	// fmt.Println(isAIOperation("create model"))
	// fmt.Println(isAIOperation("train model"))
	// fmt.Println(isAIOperation("select * from"))
	// fmt.Println(isAIOperation("select predict()"))

	var text string
	text, _ = readUserInput()
	// text := "create model\n SELECT kNN('features', 'label', 3) AS prediction\n FROM iris_table;\n"
	// fmt.Println(strings.LastIndex(text, ";\n") == len(text)-2)
	fmt.Println(text)

	// for {
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	if scanner.Scan() {
	// 		line := strings.TrimSpace(scanner.Text())
	// 		fmt.Println(line)
	// 		if strings.LastIndex(line, ";") == len(line)-1 {
	// 			fmt.Println(line)
	// 			fmt.Println("detected")
	// 			break
	// 		}
	// 	}

	// }

}
