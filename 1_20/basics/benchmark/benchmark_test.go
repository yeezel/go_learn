package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

/*
Benchmark 基准测试只有在所有的测试通过后才能运行！
函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
参数为 b *testing.B。
执行基准测试时，需要添加 -bench 参数。

运行所有的基准测试函数: go test –test.bench=.*
也可以从main函数中执行：testing.Benchmark()

基准测试报告每一列值对应的含义如下：
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间	单位为 ns（纳秒，ns/op）
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}
*/

func main() {
	fmt.Println("time", testing.Benchmark(BenchmarkHello).String())
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
	fmt.Println("Parallel", testing.Benchmark(BenchmarkParallel).String())
}

func BenchmarkHello(b *testing.B) {
	//... // 耗时操作
	//如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprint("hello")
	}
}

// 测试channel异步协程
func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

// 测试channel缓冲
func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

// 测试并发
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	//使用 RunParallel 测试并发性能
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
