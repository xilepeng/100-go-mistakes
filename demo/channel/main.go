package main

// fatal error: all goroutines are asleep - deadlock!
// 对没有初始化的 channel (nil channel)进行读写操作都会发生阻塞
func deadlock() {
	var ch chan int
	ch <- 1
	//	goroutine 1 [chan send (nil chan)]:
	//<-ch
	//goroutine 1 [chan receive (nil chan)]:
}
func main() {
	deadlock()
}
