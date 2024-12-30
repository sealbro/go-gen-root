package main

import (
	"fmt"
)

func main() {
	app := NewApp()
	result := app.calculatorService.Calculate()
	fmt.Println(result)
}
