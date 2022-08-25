package channel

import "fmt"

func count(ch chan int64, endCh chan<- int64) {
	for {
		count := <-ch
		if count == 100 {
			break
		}
		fmt.Println(count)
		ch <- count + 1
	}
	endCh <- 1
}

//func main() {
//	ch := make(chan int64)
//	endCh := make(chan int64)
//	go count(ch, endCh)
//	go count(ch, endCh)
//	ch <- 0
//	<-endCh
//}
