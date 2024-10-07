package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10}

	res := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(a))
		for i := 0; i < len(a); i++ {
			go func(i int) {
				defer wg.Done()
				res <- a[i] * a[i]
			}(i)
		}
		wg.Wait()
		close(res) 
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() { 
		var sum int
		defer wg.Done()
		for r := range res { 
			sum += r
		}
		fmt.Println(sum)
	}()
	wg.Wait()
}
