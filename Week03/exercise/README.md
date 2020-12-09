# 作业：

## 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

### 代码运行结果

```bash

$ go build
$ ./week03.exe
http server start
listenOsSignal: http server happened error
http server exit
errgroup err:listen tcp: address 8080: missing port in address
exit
-------------------------------------------------------------------
$ go build
$ ./week03.exe
http server start
listenOsSignal: signal is coming
http server exit
errgroup err:http: Server closed
exit


```
