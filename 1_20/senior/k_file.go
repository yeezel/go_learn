package senior

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
)

func FileDemo() {
	// openFile()
	// readAndWrite()

	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func openFile() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n') //读取一行
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func readByte() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		buf := make([]byte, 1024)
		n, _ := inputReader.Read(buf)
		if n == 0 {
			break
		}
		fmt.Printf("%s\n", string(buf))
	}
}

func writeFile() {

	/*
		os.O_RDONLY：只读
		os.O_WRONLY：只写
		os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
		os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0 。
	*/
	// 如果文件不存在则自动创建
	// 在写文件时，不管是 Unix 还是 Windows，都需要使用 0666
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	// 如果写入的东西很简单，可以只用这一行代码即可
	// fmt.Fprintf(outputFile, "Some test data.\n")

	// 或者不使用缓冲区直接写
	// outputFile.WriteString("hello world!\n")

	// 使用缓冲区
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

func readAndWrite() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = os.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

func readColumn() {

	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		// 前提数据是按列排列并用空格分隔的，否则fmt.Fscanf()自定义分隔符
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// 支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib，还有tar归档
func compress() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

// 用 buffer 读取命令行参数中的文件
func flagFile() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		// for {
		// 	buf, err := bufio.NewReader(f).ReadBytes('\n')
		// 	fmt.Fprintf(os.Stdout, "%s", buf)
		// 	if err == io.EOF {
		// 		break
		// 	}
		// }

		//用slice读写文件
		const NBUF = 512
		var buf [NBUF]byte
	FOROUT:
		for {
			switch nr, err := f.Read(buf[:]); true {
			case nr < 0:
				fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
				os.Exit(1)
			case nr == 0: // EOF
				break FOROUT
			case nr > 0:
				if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
					fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
				}
			}
		}
		f.Close()
	}
}
