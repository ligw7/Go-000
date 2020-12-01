学习笔记

### 什么时候用panic，什么时候用error?

- 使用**panic**：对于真正意外的情况，那些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出。
- 使用**error**：对于其他的错误情况，我们应该是期望使用error来进行判定。

> *You only need to check the error value if you care about the result. -- Dave*
>
> *Error are values.*

### error的几种实现

#### 1. 预定义错误：Sentinel Error
常常会定义成包级别变量：
```go
var Err123 = errors.New("123")
```

类似： io.EOF,更底层的syscall.ENOENT

Lowlight:
- Sentinel errors 成为你API公共部分
- Sentinel errors 在两个包之间创建依赖
- 结论：尽可能避免 Sentinel errors

#### 2. 自定义错误：Error types
```go
package myError

type MyError struct {
    Msg     string
    File    string
    Line    int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%s:%d: %s",e.File,e.Line,e.Msg)
}
```
用法：
```go
switch errType:= err.(type){
    case nil:
        // call succeeded,nothing to do
    case *MyError:
        fmt.Println("error occurred on line:",err.line)
    default:
        // unknown error
}
```
或者
```go
ok,myError := err.(*MyError)
```

Highlight: 
- 可以携带更多的上下文

Lowlight: 
- 同 Sentinel Error

#### 3. 非透明错误处理：Opaque errors

不输出错误内容，只输出成功或失败
```go
func fn() error {
    x,err := bar.Foo()
    if err != nil {
        return err
    }
    // use x
    return nil
}
```

只判断是否为error或nil

Highlight: 
- 代码和调用者之间的耦合最少

Lowlight: 
- 更细粒度的信息获取不到

##### 进一步：

Assert errors for behaviour, not type

断言错误实现了特定的行为，而不是断言错误是特定的类型或值
```go
type temporary interface {
    Temporary() bool
}
// IsTemporary returns true if err is temporary.
fun IsTemporary(err error) bool {
    te,ok := err.(temporary)
    return ok && te.Temporary()
}
```

Highlight: 
- 逻辑可以在不导入定义错误的包或者实际上不了解err的底层类型的情况下实现--我们只对它的行为感兴趣。

### Coding时，如何优雅的处理error？