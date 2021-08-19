package main

import "fmt"

func main() {
	fmt.Printf("Alice: %v\nBob: %v", parsearPuntaje(2), parsearPuntaje(1))
}

func parsearPuntaje(valor int) string {
	switch valor {
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	default:
		return "0"
	}
}
