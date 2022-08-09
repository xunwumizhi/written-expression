# Go 语言学习大纲

## 资料汇总

基础知识点：https://www.runoob.com/go/go-data-types.html
Go 入门教程 | Laravel 学院  https://laravelacademy.org/books/golang-tutorials
Go语言入门教程，Golang入门教程（非常详细）  http://c.biancheng.net/golang/

官网API文档：https://pkg.go.dev/

golang-standards/project-layout: Standard Go Project Layout  https://github.com/golang-standards/project-layout

## 准备工作

了解Go程序运行部署的情况。静态语言，动态语言；可执行文件，脚本解释器Python执行，编译成字节码在JVM上执行（Kotlin、Groovy、JRuby、Jython、Scala ）。GO代码最终编译成二进制，执行时不依赖语言的运行环境

关于变量几个概念：声明、定义（不显示初始化则默认值）、初始化

声明一个变量只是将变量名标识符的有关信息告诉编译器，使编译器“认识”该标识符

在 go 里面
定义（声明）、初始化
定义并初始化



go工程的特点：包依赖，首字母大小写设置包级别的权限，测试文件编译时忽略

学会写测试用例：test

错误处理：函数显示返回 error 变量 

## 基础

Go 语言数据类型 | 菜鸟教程  https://www.runoob.com/go/go-data-types.html
    指针、结构体

常量：iota, nil 

变量: 声明并定义 := ；一次赋值多个复制注意变量的作用域

```go
var a, b int
var (
    name string
    age int
)

a, b := 1, 2
a, b := b, a
```

函数：多返回值，匿名函数、值传递
```go
type AddFn func(a, b)

var myfn AddFn = func(a, b int){
    fmt.Println(a + b)
}

// 匿名函数
myfn := func(a, b int){
    fmt.Println(a + b)
}
```


数组、切片：底层是数组，传参的时候需要注意底层数组是否发生变化

结构体：匿名结构体，匿名成员、tag
```go
tyep Book struct {
    Name string
    Price int
}

type Store struct {
    Book // 匿名成员
    ChildrenBook Book `json:"children_book"`
    Name string
}

mybook := Book{} // 声明初始化
var mybook Book // 声明
mybookPtr := &Book{}
var book1 *B
```
可用用结构体实现类

## 抽象封装

- 接口 interface

## 并发

go 函数名( 参数列表 )

# JavaScript、Vue
