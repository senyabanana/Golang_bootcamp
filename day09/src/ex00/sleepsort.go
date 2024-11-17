package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(s []int) <-chan int {
	lenS := len(s)
	out := make(chan int, lenS)
	wg := sync.WaitGroup{}

	wg.Add(lenS)
	for i := range s {
		n := s[i]
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			out <- n
		}()
	}
	wg.Wait()
	close(out)

	return out
}

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

	result := sleepSort(nums)

	for num := range result {
		fmt.Println(num)
	}
}
