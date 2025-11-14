1. [并发 VS 并行](#并发-vs-并行)
2. [协程 goroutine](#协程-goroutine)
3. [通道 channel](#通道-channel)
4. [串联的Channels（管道 Pipeline）](#串联的channels管道-pipeline)
5. [带缓存的Channel](#带缓存的channel)


## 并发 VS 并行

- 并发：如何你只有一个处理器，但有多个任务，可以分配时间片，从而同时处理所有任务。
- 并行：多核处理器同时执行多个任务

## 协程 goroutine

Go 使用工作共享和工作窃取双重模型来管理 goroutine 

- 工作共享：调度器将工作分配给可用的进程
- 工作窃取：调度器中未充分利用的进程会尝试从另一个进程中窃取工作

## 通道 channel

channel 是一种通过信号在 goroutine 之间进行通信的机制。信号可以有数据也可以没数据

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels。
当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine之前（译注：happens before，这是Go语言并发内存模型的一个关键术语！）。
在讨论并发编程时，当我们说x事件在y事件之前发生（happens before），我们并不是说x事件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了


## 串联的Channels（管道 Pipeline）

Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的管道（pipeline）。

使用管道 pipeline 的一个好处是程序中有不变的数据流，因此 goroutine 和 通道不必等所有都就绪才开始执行。另外，因为您不必把所有内容都保存为变量，就节省了变量和内存空间的使用。最后，管道简化了程序设计并提升了维护性。

使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。

```go
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	
	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

```

其实你并不需要关闭每一个channel。只要当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。不管一个channel是否被关闭，当它没有被引用时将会被Go语言
的垃圾自动回收器回收。（不要将关闭一个打开文件的操作和关闭一个channel操作混淆。对于每个打开的文件，都需要在不使用的使用调用对应的Close方法来关闭文件。）
试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制

为了表明这种意图并防止被滥用，Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。

因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将是一个编译错误。

```go
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
```

## 带缓存的Channel

带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。`make(chan int, 5)`

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个 goroutine 执行发送操作而向队列插入元素。

那么channel的缓存队列将不是满的也不是空的（图8.4），因此对该channel执行的发送或接收操作都不会发生阻塞。通过这种方式，channel的缓存队列解耦了接收和发送的goroutine。

泄漏的 goroutines 并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。

无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的。

## 并发的循环

