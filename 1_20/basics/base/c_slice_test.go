package basics

import (
	"bytes"
	"fmt"
	"testing"
)

//切片是一个 长度可变的数组

func TestSliceDemo(t *testing.T) {

	//声明切片的格式是： var identifier []type（不需要说明长度）
	// var x = []int{2, 3, 5, 7, 11} //直接初始化

	var arr1 [6]int
	//默认下限为 0，默认上限为 len(arr1)
	var slice1 []int = arr1[2:5] // item at index 5 not included!
	//var slice1 []type = arr1[:] 等同于 = arr1[0:len(arr1)]  等同于 slice1 = &arr1

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	print(slice1)
	//cap() 可扩容的最大值：len(arr)-start_length
	//0 <= len(s) <= cap(s)
	// print the slice
	printSlice(slice1)
	println("重新分片==================================")
	// 重新分片
	slice1 = slice1[0:4]
	printSlice(slice1)
	// grow the slice beyond capacity
	//slice1 = slice1[0:7] // panic: runtime error: slice bound out of range

	println("利用切片对字符串高效拼接==================================")
	//使用切片进行字符串拼接，这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok && buffer.Len() < 20 { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")

	println("make创建切片==================================")

	//当数组还没定义的情况下，可以使用make创建切片
	//其中第二个参数作为切片初始长度而第三个参数作为相关数组的长度。这么做的好处是我们的切片在达到容量上限后可以扩容。改变切片长度的过程称之为切片重组 reslicing
	//其中第三个参数可以省略，省略后和第二个参数长度一样
	//分配一个有10个int值得数组，并创建一个长度为3，容量为10为切片，该切片指向数组得前3个元素
	slice2 := make([]int, 3, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice2); i++ {
		slice2 = slice2[0 : i+1]
		slice2[i] = i
	}

	printSlice(slice2)

	println("切片的复制与追加==================================")

	slSrc := []int{1, 2, 3}
	slDes := make([]int, 10)

	n := copy(slDes, slSrc)
	fmt.Println(slDes)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6) //若切片容量不够，会自动创建新切片存储元素，方法总是返回成功，除非系统内存耗尽了
	fmt.Println(sl3)
}

func getNextString() (str string, err bool) {
	return "gogogo==", true
}

func print(data []int) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("Slice at %d is %d\n", i, data[i])
	}

	for i, d := range data {
		fmt.Printf("range Slice at %d is %d\n", i, d)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) //len=3 cap=5 slice=[0 0 0]
}
