1. WebAssembly 是一种二进制编码格式，而不是一门新的语言。
2. WebAssembly 不是为了取代 JavaScript，而是一种补充（至少现阶段是这样），结合 WebAssembly 的性能优势，很大可能集中在对性能要求高（例如游戏，AI），或是对交互体验要求高（例如移动端）的场景。
3. C/C++ 等语言可以编译 WebAssembly 的目标文件，也就是说，其他语言可以通过编译器支持，而写出能够在浏览器前端运行的代码。

执行步骤：
1. 新建webassembly.go添加代码
2. 将 webassembly.go 编译为 weba.wasm，
    - `GOOS=js GOARCH=wasm go build -o weba.wasm`
3. 拷贝 wasm_exec.js
    - `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./`
4. 创建 index.html
5. 安装goexec，在命令行下执行go代码：`go install github.com/shurcooL/goexec`
6. 启动web服务：`goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'`
6. 使用Makefile自动化上述步骤，执行命令：`make`


进一步尝试：
- WebAssembly 的二进制分析工具 [WebAssembly Code Explorer](https://wasdk.github.io/wasmcodeexplorer/)
- 借鉴 Vue 实现的 Golang WebAssembly 前端框架 [Vugu](https://github.com/vugu/vugu)，完全使用 Go，不用写任何的 JavaScript 代码

- 使用 Go Assembly 前端渲染的一些[例子](https://stdiopt.github.io/gowasm-experiments/)
- [jsgo](https://github.com/dave/jsgo) 这个项目汇聚一些小而精的项目，包括 2048，俄罗斯方块等游戏
- [TodoMVC](https://jsgo.io/dave/todomvc)使用Go完整开发前端的项目

[syscall/js文档](https://golang.org/pkg/syscall/js)
[Go WebAssembly 官方文档 - github.com](https://github.com/golang/go/wiki/WebAssembly)