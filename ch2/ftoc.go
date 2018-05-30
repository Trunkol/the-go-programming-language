package main

import "fmt"

func ftoc() {
	const freezingF, boilingF = 32.1, 212.12981928
	fmt.Printf("%gºF = %.2gºC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gºF = %.2fºC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
