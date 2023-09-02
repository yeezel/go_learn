package main

func InverseProduct(a string, b string) string {
	a_inv_future := InverseFuture(a) // start as a goroutine
	b_inv_future := InverseFuture(b) // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}

func Product(a_inv, b_inv string) string {
	return a_inv + b_inv
}

func InverseFuture(a string) chan string {
	future := make(chan string)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

func Inverse(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}

// 让A和B字符串反转后拼接，A和B字符串可以同时进行反转计算
func main() {
	a := "wo"
	b := "you"
	s := InverseProduct(a, b)
	println(s)
}
