package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var st string
	if n, err := strconv.Atoi(st); err == nil {
		fmt.Println(n)
	}
	strings.Split(st, "1")
}

var wg sync.WaitGroup

func RunTwoSeerviceAndWaitFinish(ctx context.Context, timeout time.Duration) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
		}
	}()
}
