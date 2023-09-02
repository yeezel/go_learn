package main

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to returned channel
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

// 函数 sieve()、generate() 和 filter() 都是工厂；
// 它们创建通道并返回，而且使用了协程的 lambda 函数。
// main() 函数现在短小清晰：它调用 sieve() 返回了包含素数的通道，然后通过 fmt.Println(<-primes) 打印出来。
func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
