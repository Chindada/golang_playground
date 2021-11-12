package main

import (
	"errors"
	"fmt"
)

func main() {
	// var a []int
	// go com(&a)
	// for {
	// 	a = append(a, 1)
	// }
	err1 := errors.New("Cancel already")
	// err2 := errors.New("Cancel already")
	fmt.Println(err1.Error() == "Cancel already")
}

func com(arr *[]int) {
	for {
		tmp := *arr
		fmt.Println(len(tmp))
		// fmt.Println(tmp[len(tmp)-1])
	}
}
