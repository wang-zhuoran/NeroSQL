// Client code
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	client_sending()
}

func prompt() {
	fmt.Print("nerosql>")
}

func client_sending() {
	for {

		prompt()

		conn, err := net.DialTimeout("tcp", "localhost:8080", 5*time.Second)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		// defer conn.Close()

		message, err := readUserInput()
		if err == nil && message == "" {
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = sendMessage(conn, message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		response, err := receiveMessage(conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Received response:", response)

		time.Sleep(1 * time.Second) // Wait for 1 second before sending the next message
	}
}

func receiveMessage(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func sendMessage(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}

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
		if line == "" {
			return "", nil
		}
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
