package main

import (
	"fmt"
	"os"
)

func main() {
	for key, arg := range os.Args[0:] {
		fmt.Println(key, " - ", arg)
	}
}
