package basics

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringDemo(t *testing.T) {
	var name string = "lili"

	//使用+进行字符串拼接，这不是最高效的做法，更好的办法是使用函数 strings.Join()，最好的办法使用字节缓冲（bytes.Buffer）拼接
	println(name + " hello!!")

	//多行字符串
	var des = `
Cannot proceed, the divider is zero.
dividee: %d
divider: 0
`
	fmt.Printf(des, 5)
	println()

	//字符串的比较，一般的比较运算符（==、!=、<、<=、>=、>）通过在内存中按字节比较来实现字符串的对比
	println("字符串的比较：", "hello" == "hello")

	//len()函数获取长度
	println("len()长度：", len("fjwieofjwe"))

	//获取字符
	println("获取字符", "helloworld"[5-3])

	//获取前缀和后缀
	psfix := "This is an example of a string"
	println("前缀Thi开头：", strings.HasPrefix(psfix, "Thi"))
	println("后缀ing结尾：", strings.HasSuffix(psfix, "ing"))

	//包含
	a, b := "hello", "el"
	println("a变量是否包含el字符串：", strings.Contains(a, b))

	//替换，最后一个参数n = -1 则替换所有字符串
	newstr := strings.Replace("goGogogo", "go", "biu", 2)
	println("替换字符串：", newstr)

	//统计
	var manyG = "gg,gg,Gg,gg"
	println("有多少个gg：", strings.Count(manyG, "gg"))

	//重复字符串
	println("重复3次go：", strings.Repeat("go", 3))

	//大小写转换
	println("转小写：", strings.ToLower("HELLO"))
	println("转大写：", strings.ToUpper("hello"))

	//删除开头和结尾的字符串
	println("删除指定cut字符串:", strings.Trim("cuthellocutcut", "cut"))
	println("删除空格字符串:", strings.TrimSpace("	 hello    "))

	//切割字符串，返回slice数据类型
	for i, data := range strings.Fields("aa bb cc") {
		println("fields切割单词字符串：", i, "=", data)
	}
	for i, data := range strings.Split("aa,bb,cc", ",") {
		println("split切割单词字符串：", i, "=", data)
	}

	//字符串转基础类型
	liu := "666"
	num, err := strconv.Atoi(liu)
	println("转数字类型：", num, err)
	println("数字转字符串：", strconv.Itoa(12398))
}
