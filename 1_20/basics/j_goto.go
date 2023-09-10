package basics

func GotoDemo() {
	//不推荐使用goto
	//goto模拟for
	//如果必须使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败。
	i := 0
HERE:
	println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
