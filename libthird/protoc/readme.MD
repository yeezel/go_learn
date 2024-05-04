# 简介
protobuf 即 Protocol Buffers，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。protobuf 性能和效率大幅度优于 JSON、XML 等其他的结构化数据格式。protobuf 是以二进制方式存储的，占用空间小，但也带来了可读性差的缺点。protobuf 在通信协议和数据存储等领域应用广泛。例如著名的分布式缓存工具 Memcached 的 Go 语言版本groupcache 就使用了 protobuf 作为其 RPC 数据格式。

# 安装
## 下载
[下载地址-单个二进制文件](https://github.com/protocolbuffers/protobuf/releases)
验证：protoc --version

## 安装插件
这个工具用来将 .proto 文件转换为 Golang 代码
`go install github.com/golang/protobuf/protoc-gen-go`


# 步骤
1. 编写.proto文件
2. 翻译成go代码：protoc --go_out=. *.proto
3. 测试：go run main.go


# 字段类型
## 标量类型（Scalar）
|proto类型|	go类型|	备注|	proto类型|	go类型|	备注|
|-|-|-|-|-|
|double|	float64	|	float|	float32	|
|int32|	int32|		int64|	int64|	
|uint32|	uint32|		uint64|	uint64|	
|sint32|	int32|	适合负数|	sint64|	int64|	适合负数|
|fixed32|	uint32|	固长编码，适合大于2^28的值|	fixed64|	uint64|	固长编码，适合大于2^56的值|
|sfixed32|	int32|	固长编码|	sfixed64|	int64|	固长编码|
|bool|	bool|		|string|	string|	UTF8 编码，长度不超过 2^32|
|bytes|	[]byte|	任意字节序列，长度不超过 2^32|		|	||

- 标量类型如果没有被赋值，则不会被序列化，解析时，会赋予默认值。
    - strings：空字符串
    - bytes：空序列
    - bools：false
    - 数值类型：0

## 枚举(Enumerations)
枚举类型适用于提供一组预定义的值，选择其中一个。例如我们将性别定义为枚举类型。
```
message Student {
  string name = 1;
  enum Gender {
    FEMALE = 0;
    MALE = 1;
  }
  Gender gender = 2;
  repeated int32 scores = 3;
}
```
枚举类型的第一个选项的标识符必须是0，这也是枚举类型的默认值。
别名（Alias），允许为不同的枚举值赋予相同的标识符，称之为别名，需要打开allow_alias选项。
```
message EnumAllowAlias {
  enum Status {
    option allow_alias = true;
    UNKOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
```

## 使用其他消息类型
Result是另一个消息类型，在 SearchReponse 作为一个消息字段类型使用。
```
message SearchResponse {
  repeated Result results = 1; 
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```
嵌套写也是支持的：
```
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```
如果定义在其他文件中，可以导入其他消息类型来使用：
`import "myproject/other_protos.proto";`

## 任意类型(Any)

Any 可以表示不在 .proto 中定义任意的内置类型
```
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

## oneof
```
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```
## map
```
message MapRequest {
  map<string, int32> points = 1;
}
```

# 定义服务(Services)
如果消息类型是用来远程通信的(Remote Procedure Call, RPC)，可以在 .proto 文件中定义 RPC 服务接口。例如我们定义了一个名为 SearchService 的 RPC 服务，提供了 Search 接口，入参是 SearchRequest 类型，返回类型是 SearchResponse
```
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```
官方仓库也提供了一个[插件列表](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)，帮助开发基于 Protocol Buffer 的 RPC 服务。

# protoc 其他参数
命令行使用方法: `protoc --proto_path=IMPORT_PATH --<lang>_out=DST_DIR path/to/file.proto`

- `--proto_path=IMPORT_PATH`：可以在 .proto 文件中 import 其他的 .proto 文件，proto_path 即用来指定其他 .proto 文件的查找目录。如果没有引入其他的 .proto 文件，该参数可以省略。
- `--<lang>_out=DST_DIR`：指定生成代码的目标文件夹，例如 –go_out=. 即生成 GO 代码在当前文件夹，另外支持 cpp/java/python/ruby/objc/csharp/php 等语言

# 推荐风格

- 文件(Files)
    - 文件名使用小写下划线的命名风格，例如 lower_snake_case.proto
    - 每行不超过 80 字符
    - 使用 2 个空格缩进

- 包(Packages)
    - 包名应该和目录结构对应，例如文件在my/package/目录下，包名应为 my.package

- 消息和字段(Messages & Fields)
    - 消息名使用首字母大写驼峰风格(CamelCase)，例如message StudentRequest { ... }
    - 字段名使用小写下划线的风格，例如 string status_code = 1
    - 枚举类型，枚举名使用首字母大写驼峰风格，例如 enum FooBar，枚举值使用全大写下划线隔开的风格(CAPITALS_WITH_UNDERSCORES )，例- 如 FOO_DEFAULT=1

- 服务(Services)
    - RPC 服务名和方法名，均使用首字母大写驼峰风格，例如service FooService{ rpc GetSomething() }


# 参考
[Go Protobuf 简明教程](https://geektutu.com/post/quick-go-protobuf.html)