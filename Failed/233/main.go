package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func CountDigitOne(n int) int {
	q := []byte("1")
	var i int
	inChan := make(chan byte)
	var ans []int
	for i = 1; i <= n; i++ {
		go func(i int) {
			strBuf := bytes.NewBuffer([]byte(strconv.Itoa(i)))
			s := strBuf.Next(len(strconv.Itoa(i)))
			for _, v := range s {
				// fmt.Println(v)
				inChan <- v
			}
		}(i)
	}
	// var xx int
	// for tmp := range inChan {
	// 	go func(tmp byte) {
	// 		var a int
	// 		if tmp == q[0] {
	// 			fmt.Println(tmp)
	// 			a++
	// 			ans = append(ans, a)
	// 		}
	// 	}(tmp)
	// }
	// close(inChan)
Loop:
	for {
		select {
		case tmp := <-inChan:
			go func(tmp byte) {
				var a int
				if tmp == q[0] {
					// fmt.Println(tmp)
					a++
					ans = append(ans, a)
				}
			}(tmp)
		case <-time.After(10 * time.Millisecond):
			break Loop
		}
	}

	// for {
	// 	if len(ans)-xx == 0 {
	// 		break
	// 	}
	// 	temp := len(ans)
	// 	xx = temp
	// }
	fmt.Println(len(ans))
	return len(ans)
}
