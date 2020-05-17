package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	for {
		select {
		case i := <-ch1:
			time.Sleep(time.Second * 2)
			fmt.Println("ch1:", i)

		}
	}
}
