package main

import "fmt"

// reciver
func rcv (ch1 <-chan int, ch2 chan int){
	for {
	    i := <-ch1
	    fmt.Printf("ch1:%d ", i)
		if i < 0{
			ch2 <- i
		}
    }
}

// reciver done
func done (ch2 <-chan int){ 
    fmt.Printf("Start \n")
	for {
	    i := <-ch2
		if i < 0{
			fmt.Printf("\nreciver done:%d \n", i)
		}
    }
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go rcv(ch1, ch2)
	go done(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	ch1 <- -1
    fmt.Printf("main done \n")
}
