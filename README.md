# Goalng 练习代码

## FAQ

###

numeric-constants.go 中有一行代码是

    Big     = 1 << 100

这里面Big 是什么？

### 声明变量与赋值

struct 字段遇见的问题

```python

type
Vertex struct {
    X int
    Y int
}
```

v = Vertex{1, 3} 会有语法错误
-   是因为 v 是在此时声明变量 \*

v.X := 4 会有错误
-   是因为 v.X 仅仅是赋值 \*

### 结构体声明方法

```golang
func (v *Vertex) Abs() float64 {
    return 0
}

func (f MyFloat) Abs() float64 {
    return 0
}
```

在为结构体声明变量的时候用的是指针，在为自定义的类型 MyFloat 声明变量时，用
的是变量

### interface 使用方法

### goroutines 使用方法

go tour 中为什么`say(hello)`能够正确的执行`0..4`五次，而`go say(hello)`只能够正
确的执行`0..3`四次。

`go` 会将say('world')丢掉 goroutines 里面运行，进程退出的时候不会等到 goroutine
结束。

### select 的使用方法

选择准备好的goroutines 来执行，多个 goroutines 都可以执行的时候，则随机选择。


### defer 使用

一个 defer 的过程要处理缓存对象，参数拷贝，多次函数调用。这样性能会有损耗，在平
时代码中要合理的使用 defer。

### range

range 会复制对象，见 range_copy.go
