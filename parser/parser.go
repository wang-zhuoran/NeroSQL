package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/marianogappa/sqlparser"
)

func prompt() {
	fmt.Print("nerosql>")
}

func ParseSQL(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for {
		prompt()
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("Error reading input: %v", err)
			}
			// End of input
			break
		}

		sql := strings.TrimSpace(scanner.Text())
		if sql == "" {
			continue
		}
		if strings.ToLower(sql) == "exit()" {
			break
		}

		query, err := sqlparser.Parse(sql)
		if err != nil {
			log.Printf("Error: %s", err)
			continue
		}

		fmt.Printf("Query: %v\n", query)
	}

	return nil
}
