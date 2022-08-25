package channel

var x int64
var endCh = make(chan int64)

func add(ch chan int64) {
	for i := 0; i < 50000; i++ {
		// 发送一个值 挂起自身 等待其他协程唤醒
		ch <- 1
		// 自增操作
		x++
		// 操作结束，挂起自身，唤醒其他协程
		<-ch
	}
	// 操作完成唤醒主协程
	endCh <- 1
}

//func main() {
//	ch := make(chan int64)
//
//	go add(ch)
//	go add(ch)
//
//	// 唤醒其中一个协程
//	<-ch
//	// 挂起主协程
//	<-endCh
//
//	fmt.Println(x)
//}
