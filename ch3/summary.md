### 格式化代码工具

* ``gofmt -w test.go`` 格式化并重写 test.go,不加 -w 只打印格式化的代码, 不重写到 test.go
* ``gofmt -w *.go`` 格式化并重写所有 go文件
* ``gofmt dir1`` 格式化并重写 dir1目录及其子目录下所有的 go文件


### 依赖管理

``go install`` 安装非标准库的包文件


### 单元测试
``go test`` 是一个轻量级的单元测试框架