package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lines []string
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if strings.Contains(input.Text(), os.Args[1]) {
			lines = append(lines, input.Text())
		}
	}

	for _, line := range lines{
		fmt.Println(line)
	}
}
