package basics

import "fmt"

func LabelDemo() {
	//LABEL，建议使用全部大写字母
	println("continue label==================")
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	println("break label==================")
LABEL2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
