package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	critical := color.New(color.Bold, color.FgHiRed).SprintFunc()
	poor := color.New(color.Bold, color.FgGreen).SprintFunc()
	a := 8.988765678
	b := 1.1
	aString := fmt.Sprintf("%5.2f", a)
	bString := fmt.Sprintf("%5.2f", b)
	x := critical(aString)
	fmt.Printf("%2s %2s dadsf\n", x, poor(bString))
}
