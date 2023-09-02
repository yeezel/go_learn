package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var b int = 200
	var ip *int /* 声明指针变量 */
	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。一个指针变量通常缩写为 ptr
	var ptr *int

	if ptr == nil {
		println("ptr is null pointer")
	}
	println(ptr)

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap(&a, &b)
	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

/* 定义交换值函数*/
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
