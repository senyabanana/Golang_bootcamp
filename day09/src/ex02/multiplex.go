package main

import (
	"fmt"
	"sync"
)

func multiplex(channels ...chan interface{}) chan interface{} {
	out := make(chan interface{})
	wg := sync.WaitGroup{}

	go func() {
		for _, ch := range channels {
			wg.Add(1)
			go func(ch chan interface{}) {
				defer wg.Done()
				for val := range ch {
					out <- val
				}
			}(ch)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 11; i <= 15; i++ {
			ch3 <- i
		}
	}()

	out := multiplex(ch1, ch2, ch3)

	for val := range out {
		fmt.Println(val)
	}
}
