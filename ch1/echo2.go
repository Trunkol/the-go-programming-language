package main

import (
	"fmt"
	"os"
)

func echo2() {
	for key, arg := range os.Args[0:] {
		fmt.Println(key, " - ", arg)
	}
}
