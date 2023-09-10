package basics

// 结构体不暴露给外部
type file struct {
	fd   int    // 文件描述符
	name string // 文件名
}

func StructFactory() {
	f := NewFile(10, "./test.txt")
	println(f.name)
}

// 让外部包强制使用该工厂方法
func NewFile(fd int, name string) *file {
	if fd < 0 {
		return nil
	}

	//如果是结构体，new(file) 和 &file{} 是等价的
	return &file{fd, name}
}
