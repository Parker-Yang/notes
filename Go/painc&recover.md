```go
func main() {
      fmt.Println("c")
   defer func() { // 必须要先声明defer，否则不能捕获到panic异常
      fmt.Println("d")
      if err := recover(); err != nil {
         fmt.Println(err) // 这里的err其实就是panic传入的内容
      }
      fmt.Println("e")
   }()
   f() //开始调用f
   fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
   fmt.Println("a")
   panic("异常信息")
   fmt.Println("b") //这里开始下面代码不会再执行
}
-------output-------
c
a
d
异常信息
e
```

#### panic：
 1、内建函数
 2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
 3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行，这里的defer 有点类似 try-catch-finally 中的 finally
 4、直到goroutine整个退出，并报告错误

#### recover：
 1、内建函数
 2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
 3、一般的调用建议
 a). 在defer函数中，通过recever来终止一个gojroutine的panicking过程，从而恢复正常代码的执行
 b). 可以获取通过panic传递的error

## 简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。