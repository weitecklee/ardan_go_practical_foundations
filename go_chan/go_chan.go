package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		/*
			// BUG: all goroutines use the same "i" due to scope
			go func() {
				fmt.Println(i)
			}()
		*/
		/*
			// Fix 1: use parameter
			go func(n int) {
				fmt.Println(n)
			}(i)
		*/
		// Fix 2: Use loop body variable
		i := i // "i" shadows "i" from the for loop
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)
	ch := make(chan string)
	// ch <- "hi" // will cause deadlock
	go func() {
		ch <- "hi"
	}()
	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch) // must close channel otherwise deadlock
	}()

	for msg := range ch {
		fmt.Println("got: ", msg)
	}

	/* for/range on ch does this
	for msg := range ch {
		msg, ok := <-ch
		if !ok {
			break
		}
		...
	}
	*/

	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed -> panic

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))

}

func sleepSort(values []int) []int {
	ch := make(chan int)
	sorted := make([]int, 0, len(values))
	// max := 0
	for _, v := range values {
		v := v
		// if v > max {
		// 	max = v
		// }
		go func() {
			time.Sleep(time.Millisecond * time.Duration(v))
			ch <- v
		}()
	}
	// go func() {
	// 	time.Sleep(time.Millisecond * time.Duration(max*2))
	// 	close(ch)
	// }()
	// for v := range ch {
	// 	sorted = append(sorted, v)
	// }

	for range values { // for i := 0; i < len(values); i++ {
		sorted = append(sorted, <-ch)
	}
	return sorted
}

/* Channel semantics:
- send & receive will block until opposite directions (with exceptions)
- receive from closed channel will return zero value without blocking
- send to closed channel will panic
- closing closed channel will panic
- send/receive to nil channel will block forever
*/
