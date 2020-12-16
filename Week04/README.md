# 学习笔记

## 2020-12-10

Standard Go Project Layout

https://github.com/golang-standards/project-layout/blob/master/README_zh.md

请参阅Go 1.4 release notes。注意，你并不局限于

顶级 internal 目录。

### 依赖注入，控制反转

https://blog.golang.org/wire

### 生命周期

https://github.com/go-kratos/kratos/blob/v2/app.go

orm

https://github.com/facebook/ent

bilibili v2 代码参考
https://github.com/go-kratos/service-layout
https://github.com/go-kratos/kratos/blob/v2/transport/http/service.go

---------------------------------------------------------------------------------------------------

## 2020-12-12

gRPC，基于文件定义服务，服务而非对象、消息而非引用

restful 时间长了，接口实现和文档可能会不一致

区分proto message 必选和可选

Protobuf v3 中，建议使用：https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto

### API Errors

使用一小组标准错误配合大量资源

业务层只处理自己关心的错误，也只识别自己关心的错误；不关心的错误统一返回 500 （unkonwn）错误

在PB文件中定义错误，并生成IsXXX()判断具体错误的方法

kratos错误代码处理示例：https://github.com/go-kratos/kratos/blob/v2/examples/kratos-demo/api/kratos/demo/v1/greeter_grpc.pb.go

### API Design

google.protobuf.FieldMask

必看->谷歌API设计指南：https://www.bookstack.cn/read/API-design-guide/API-design-guide-01-%E7%AE%80%E4%BB%8B.md

### 配置管理

环境变量，静态配置，动态配置，全局配置

动态配置，可以结合：https://pkg.go.dev/expvar 使用

### Functional options

https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

the-site-reliability-workbook2

margin note3 工具 看电子书

### Configuration & APIs

yaml+proto

用DTO->DO的deep copy思路，将配置初始化，和对象加载，分开

### Configuration Best Pratice

不建议将二进制文件的发布和配置文件发布做为两个流程处理，因为回滚二进制文件的时候，往往忘记配置文件的回滚

### 测试

mock/fake stub

Google测试之道

微软测试之道

subtests

https://blog.golang.org/subtests

gomock

利用go官方提供的：Subtests + Gomock完成整个单元测试

API测试框架：例如yapi

业务层biz层，利用gomock模拟interface的实现，来进行业务单元测试

https://golang.org/pkg/testing/#T.Cleanup

https://golang.google.cn/pkg/testing/#T.Cleanup