package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if strings.Contains(input.Text(), os.Args[1]) {
			fmt.Println(input.Text())		
		}
	}
}
