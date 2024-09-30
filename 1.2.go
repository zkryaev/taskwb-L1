package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	ans := make(chan int, len(nums))

	wg := sync.WaitGroup{}
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ans <- n * n
		}(nums[i])
	}
	wg.Wait()
	close(ans)
	for v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
