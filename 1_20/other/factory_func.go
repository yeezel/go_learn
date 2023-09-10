package other

import "strings"

// 工厂函数
func FactoryDemo() {
	addBmp := factoryPic(".bmp")
	addJpeg := factoryPic(".jpeg")

	println(addBmp("file"))  // returns: file.bmp
	println(addJpeg("file")) // returns: file.jpeg
}

func factoryPic(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
