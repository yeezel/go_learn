package basics

import (
	"fmt"
	"reflect"
)

func ArrayDemo() {

	//区别
	var arr1 = new([5]int)            // 使用 new 函数创建一个指向数组的指针
	var arr2 [5]int                   // 直接声明一个数组
	var arr3 []int                    // 直接声明一个切片
	fmt.Println(reflect.TypeOf(arr1)) // *[5]int
	fmt.Println(reflect.TypeOf(arr2)) // [5]int
	fmt.Println(reflect.TypeOf(arr3)) // []int

	// 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, x := range balance {
		fmt.Printf("第 %d 位 x 的值 = %f\n", i, x)
	}
	// 声明数组的同时快速初始化数组
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, x := range balance1 {
		fmt.Printf("第 %d 位 x 的值 = %f\n", i, x)
	}

	//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小，实际上这个是一个切片
	balance2 := []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, x := range balance2 {
		fmt.Printf("第 %d 位 x 的值 = %f\n", i, x)
	}

	//只为个别初始化
	balance3 := []string{3: "Chris", 4: "Ron"}
	for i, x := range balance3 {
		fmt.Printf("第 %d 位 x 的值 = %v\n", i, x)
	}

	/* 声明 n 是一个长度为 10 的数组 */
	var n [10]int

	/* 为数组 n 初始化元素 */
	for i := 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}
	/* 输出每个数组元素的值 */
	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	//多维数组
	var screen [1920][1080]int
	for y := 0; y < 1080; y++ {
		for x := 0; x < 1920; x++ {
			screen[x][y] = 0
		}
	}

	//把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种情况：
	// 传递数组的指针
	// 使用数组的切片 (常用)
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f\n", x)
}
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
