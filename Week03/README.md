# 学习笔记

2020-12-03

MESI 内存屏障 CPU屏障 CAS Effective Go

go ，什么时候结束，如何让它结束 Never start a goroutine without knowning when it will stop

log.fatal ,调用os.exit ,会无条件终止程序；defer不会被调用到

https://learnku.com/docs/effective-go/2020/introduction/6236，高效的 Go 编程 Effective Go

https://colobu.com/，鸟窝

http://docscn.studygolang.com/ref/mem，Go 内存模型

https://www.jianshu.com/p/5e44168f47a3，Go的内存模型

https://github.com/cch123/golang-notes，主要是源码分析

https://chai2010.cn/advanced-go-programming-book/，曹大佬


---------------------------------------------------------------------------------------------

2020-12-05

Do not communicate by sharing memory; instead, share memory by communicating.

Data race ,向量时钟

最晚加锁，最早释放

Redis COW BGsave

空转优化，pause,intel pause