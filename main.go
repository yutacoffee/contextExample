package main

import (
	"context"
	"fmt"
	"time"

)

func longProcess(ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main ()  {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
	go longProcess(ch)

	CTXLOOP:
		for {
			select {
			case <- ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP

			case <- ch:
				fmt.Println("success")
				break CTXLOOP

			}
		}
	fmt.Println("###############")
}