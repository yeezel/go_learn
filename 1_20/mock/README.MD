gomock 是官方提供的 mock 框架，同时还提供了 mockgen 工具用来辅助生成测试代码。

安装：
go get -u github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

使用步骤：
1. 使用 mockgen 生成 db_mock.go：`mockgen -source=db.go -destination=db_mock.go -package=main`
2. 新建 db_test.go，写测试用例
3. 执行测试：`go test . -cover -v`