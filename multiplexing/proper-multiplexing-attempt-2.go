// multiplexing.go

package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan int, 3)
	done := make(chan bool)
	completed := make(chan bool)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("Received data from ch1")
			case val := <-ch2:
				fmt.Println(val)
			case c := <-ch3:
				fmt.Println(c)
			case <-done: // when the channel is closed, this case is executed now
				fmt.Println("exiting...")
				completed <- true
				return
			}
		}
	}()

	ch1 <- 100
	ch2 <- "ch2 msg"

	// added to avoid leaking the 'select' goroutine
	close(done)
	<-completed
}
