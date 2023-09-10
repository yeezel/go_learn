package basics

import "fmt"

func TypeChange() {
	var sum int = 17
	var count int = 5

	mean := float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
