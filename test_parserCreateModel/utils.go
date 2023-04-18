package main

import (
	"regexp"
	"strings"
)

func IsAIOperation(s string) bool {
	tokens := ExtractKeywords(s)
	if len(tokens) < 2 {
		return false
	}
	switch tokens[0] {
	case "create", "train":
		return len(tokens) >= 2 && tokens[1] == "model"
	case "predict":
		return true
	default:
		return false
	}
}

func ExtractKeywords(s string) []string {
	reg, err := regexp.Compile("[a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	matches := reg.FindAllString(s, -1)
	var keywords []string
	for _, match := range matches {
		if len(match) >= 1 { // 只保留长度超过2个字符的单词作为关键字;
			keyword := strings.ToLower(match) // 统一转换成小写
			keywords = append(keywords, keyword)
		}
	}
	return keywords
}
