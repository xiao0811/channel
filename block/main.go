package main

import (
	"fmt"
	"time"
)

const (
	max     = 50000000
	block   = 500
	bufSize = 100
)

func test() {
	done := make(chan struct{})
	c := make(chan int, bufSize)

	go func() {
		count := 0
		for x := range c {
			count += x
		}
		fmt.Println(count)
		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func testBlock() {
	done := make(chan struct{})
	c := make(chan [block]int, bufSize)

	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
		fmt.Println(count)
		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if b[n] == max-1 {
				break
			}
		}
		c <- b
	}

	close(c)
	<-done
}

func main() {
	start := time.Now()
	test()
	// testBlock()
	fmt.Println(time.Now().Sub(start))
}
