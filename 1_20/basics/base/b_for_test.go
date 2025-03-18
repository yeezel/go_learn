package basics

import (
	"fmt"
	"testing"
)

func TestForDemo(t *testing.T) {
	//方式一
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//方式二 类似while
	sum = 0
	for sum <= 10 {
		sum += 1
	}
	fmt.Println(sum)

	//方式三 死循环
	sum = 0
	for {
		sum += 1
		if sum == 50 {
			break
		}
	}
	fmt.Println(sum)

	//方式四 range
	strings := []string{"google", "runoob"}
	for index, str := range strings {
		fmt.Println(index, str)
	}
	numbers := [6]int{1, 2, 3, 5}
	//读取key和value
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	// 读取 key
	for index := range numbers {
		fmt.Printf("key is: %d\n", index)
	}
	// 读取 value
	for _, value := range numbers {
		fmt.Printf("value is: %d\n", value)
	}

	//LABEL，建议使用全部大写字母
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				// break LABEL1
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
