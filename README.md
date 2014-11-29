# Goalng 练习代码

## FAQ

###

numeric-constants.go 中有一行代码是

    Big     = 1 << 100

这里面Big 是什么？

### 声明变量与赋值<a id="sec-1-1-2" name="sec-1-1-2"></a>

struct 字段遇见的问题

    type Vertex struct {
        X int
        Y int
    }

v = Vertex{1, 3} 会有语法错误
-   是因为 v 是在此时声明变量 \*

v.X := 4 会有错误
-   是因为 v.X 仅仅是赋值 \*

### 结构体声明方法<a id="sec-1-1-3" name="sec-1-1-3"></a>

    func (v \*Vertex) Abs() float64 {
        return 0
    }

    func (f MyFloat) Abs() float64 {
        return 0
    }

在为结构体声明变量的时候用的是指针，在为自定义的类型 MyFloat 声明变量时，用
的是变量

### interface 使用方法

### goroutines 使用方法

go tour 中为什么`say(hello)`能够正确的执行`0..4`五次，而`go say(hello)`只能够正
确的执行`0..3`四次。

### select 的使用方法

选择准备好的goroutines 来执行，多个 goroutines 都可以执行的时候，则随机选择。