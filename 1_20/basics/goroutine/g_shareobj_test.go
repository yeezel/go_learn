package goroutine

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

// 这个结构在构造函数 NewPerson() 中初始化的同时会启动一个后台协程 backend()。
// backend() 方法会在一个无限循环中执行 chF 中放置的所有函数，有效地将它们序列化从而提供了安全的并发访问
func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// 更改和读取 salary 的方法会通过将一个匿名函数写入 chF 通道中，然后让 backend() 按顺序执行以达到其目的。
// Set salary.
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// Retrieve salary.
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

// 使用协程实现同步互斥锁，顺序安全的执行对象共享修改值
// 这是一个简化的例子，它不应该被用在这种案例下。但是它却向我们展示了在更复杂的场景中该如何解决这种问题
func TestShareObjDemo() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	for i := 1; i < 100; i++ {
		bs.SetSalary(4000.25 * float64(i))
		fmt.Println("Salary changed:")
		fmt.Println(bs)
	}
}
