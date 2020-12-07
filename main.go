package main

import (
	"adventofcode/day1"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var dayAnswer int
	color.Green("Welcome to my 2020 Advent of Code!")
	color.Blue("Choose a day challenge: (1-25)")

	_, err := fmt.Scan(&dayAnswer)
	if err != nil {
		color.Red("could not read your input")
		color.Yellow("run this program again...")
		return
	}

	switch dayAnswer {
	case 1:
		day1.DisplayExpenses()
		return
	default:
		color.Yellow(fmt.Sprintf("bad input...%v", dayAnswer))
	}
}
