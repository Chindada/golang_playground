package main

import (
	"fmt"
	"time"
)

func main() {
	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)
	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	// prompt := promptui.Prompt{
	// 	Label: "Number",
	// }
	// result, err := prompt.Run()
	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("You choose %q\n", result)
	a := time.Unix(0, 1634605208032373000)
	last := time.Date(a.Year(), a.Month(), a.Day(), 13, 0, 0, 0, time.Local)
	fmt.Println(a)
	fmt.Println(last)
	fmt.Println(a.After(last))
}
