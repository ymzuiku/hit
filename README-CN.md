# 在Golang中使用三元运算

> [查看英文文档](./README.md)

Golang的设计理念是一个事物只有一种解法, 所以连三元运算也没有提供, 官方的思路是使用 if / else 代替三元运算

如果你平时习惯使用概括性更强的语句进行编写代码, 你可能希望可以这样去编写代码:

```js
var a = 20 > 50 ? true: false;  // a = false
var b = a || 100  // b = 100
var c = b && 50  // c = 50
var fn = c > 40 && func(){ return "ok" } // fn = "ok"
```

使用`hit`之后,以上的代码我们可以这样写:

```go
var a = If(20 > 50, true, false)  // a = false
var b = Or(a, 100)  // b = 100
var c = If(b, 50)  // c = 50
var fn = If(c > 40, func(){ return "ok" }) // fn = "ok"
```

抛弃教条,享受 `hit` 吧

## 安装

```
$ go get github.com/ymzuiku/hit
```

hit只有两个API

```go
func If(args ...interface{}) interface{}
func Or(args ...interface{}) interface{}
```

## 使用例子

> 当需返回取某个参数时, 并且该参数类型是函数时, hit会运行函数并且使用函数返回值作为参数的值

以下条件都会识别为`不成立`

- nil
- 0
- error
- ""
- "f", "F", "false", "FALSE", "False"
- "0", "0.0", "-0"



### `If(a, b, c)` 类似 `a ? b : c`

if `a` == `成立`, return `b`, else return `c`

```go
import . "github.com/ymzuiku/hit"

func main(){
    value1 := If(20 > 5, "ok", "cancel")
    log.Println(value1) // ok
    
    value2 := If("test", "ok", "cancel")
    log.Println(value2) // ok
    
    value3 := If("", "ok", "cancel")
    log.Println(value3) // cancel
    
    value4 := If("false", "ok", "cancel")
    log.Println(value4) // cancel
    
    value5 := If(5, "ok", "cancel")
    log.Println(value5) // ok
    
    value6 := If(0, "ok", "cancel")
    log.Println(value6) // cancel
    
    value7 := If(nil, "ok", "cancel")
    log.Println(value7) // cancel
    
    value8 := If(errors.New("test-error"), func() { log.Println("if err != nil, 这个参数不会进行读取" }) // value8 is error
    
    //如果参数的类型是 `func()`, hit 读取参数时会运行它并且返回nil
    value9 := If(20 > 5, func() { log.Println("ok") }, func() { log.Println("cancel") }) // run: log.Println("ok"), value = nil
    
    // 如果参数的类型是 `func()interface{}`, hit 读取参数时会运行它并且返回函数的返回值.
    value10 := If(func() interface{} { return true }, func()interface{} { return "ok"  }, func()interface{} { return "cancel" })
    log.Println(value9) // ok

}
```

### `If(a, b)` 类似 `a && b`

if `a` == `成立`, return `b`, else return `nil`

```go
import . "github.com/ymzuiku/hit"

func main(){
    value1 := If("test", "ok")
    log.Println(value1) // "ok"
    
    value2 := If(500, "ok")
    log.Println(value2) // "ok"
    
    value3 := If(func() interface{} { return false }, func() { log.Println("this func no run") })
    log.Println(value3) // func no run, and value = nil
    
    value4 := If(func() interface{} { return true }, func() interface{} { log.Println("func is run"); return 50 })
    log.Println(value4) // func is run, and value = 50
}
```

### `Or(a, b)` 类似 `a || b`

if `a` == `成立`, return `a`, else return `b`

```go
import . "github.com/ymzuiku/hit"

func main(){
    value1 := Or("test", "ok")
    log.Println(value1) // "test"
    
    value2 := Or(500, "ok")
    log.Println(value2) // 500
    
    value3 := Or(func() interface{} { return 100 }, func() { log.Println("this func no run") })
    log.Println(value3) // func no run, and value = 100
    
    value4 := Or(func() interface{} { return false }, func() interface{} { log.Println("func is run"); return 50 })
    log.Println(value4) // func is run, and value = 50
}
```

## 测试用例:

![](https://user-gold-cdn.xitu.io/2018/9/22/165fe3c7fac459c3?w=1242&h=84&f=png&s=22863)

hit代码的测试覆盖率为 100%, 查看具体的测试代码: [hit_test.go](./hit_test.go)


### 开原协议

```
MIT License

Copyright (c) 2013-present, Facebook, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
