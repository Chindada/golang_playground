package main

import "fmt"

func main() {
	// var historyClose []int
	// for {
	// 	tmp := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// 	historyClose = append(historyClose, tmp...)
	// 	if len(historyClose) < 5 {
	// 		continue
	// 	} else {
	// 		historyClose = historyClose[len(historyClose)-5:]
	// 	}
	// 	fmt.Println(historyClose)
	// }
	ch := make(chan int)
	ch2 := make(chan int)
	chanMap := make(map[int]*chan int)
	chanMap[1] = &ch
	chanMap[2] = &ch2
	go func() {
		for {
			num := <-ch
			num2 := <-ch2
			fmt.Println(num, num2)
		}
	}()

	go func() {
		for {
			ch := chanMap[1]
			*ch <- 2
			ch2 := chanMap[2]
			*ch2 <- 3
		}
	}()
	for {

	}
}
