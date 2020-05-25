package main

import (
	"fmt"
)

func f(x int) int{
	return x
}

func foo() chan int {
	mychannel := make(chan int)
	go func() {
		for i := 0; ; i++ {
			mychannel <- i
		}
	}()
	return mychannel // returns the channel as returning argument
}


func Merge2Channels(f func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int){
	go func(in1 <-chan int, in2 <- chan int, out chan<- int, f func(int) int) {
		for i := 0; i < n; i++ {
			var f1 = make(chan int)
			var f2 = make(chan int)
			go func(f1 chan<- int, x1 int) {
				f1 <- f(x1)
			}(f1, <- in1)
			go func(f2 chan<- int, x2 int) {
				f2 <- f(x2)
			}(f2, <- in2)
			out <- <-f1 + <-f2
		}
	}(in1, in2, out, f)
}



func main() {
	const n = 100

	//wg := &sync.WaitGroup{}
	in1 := foo()
	in2 := foo()
	out := make(chan int)

	Merge2Channels(f, in1, in2, out, n)
	//wg.Wait()


	for i := 0; i < n; i++{
		x := <-out
		if x == 2 * f(i){
			fmt.Println(x, 2 * f(i))
		}
	}
	close(in1)
	close(in2)
	close(out)
}
