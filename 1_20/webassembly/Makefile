all: weba.wasm wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./

weba.wasm : webassembly.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o weba.wasm