package concurrency

import (
	"fmt"
	"time"
)

// ========================= Goroutines ===================
func sayConcurrent(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func ConcurrentLauncher() {
	go sayConcurrent("world") // launch a lightweight thread called "goroutine"
	sayConcurrent("hello")
}

// ============================ Channels ====================
func FirstChanUser() {
  myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

  myChannel := make(chan int) // How to build a channel, it will allow communication with goroutines

  go functionPublishingInAChannel(myArray[:len(myArray)/2], myChannel)
  go functionPublishingInAChannel(myArray[len(myArray)/2:], myChannel)

  x, y := <-myChannel, <-myChannel // block until at least 1 value is received

  fmt.Println(x, y)

}

func functionPublishingInAChannel(intArray []int, outputChan chan int){ // chan is a new "type"
  sum := 0
  for _, v := range intArray {
    sum += v
  }

  outputChan <- sum // publish "sum" into the channel, block if a value is already in the channel has not been ridden
}

// ===================== Buffered channels ==================

func FirstBufferedChanUser() {
    myChannel := make(chan int, 2) // 1 value can stay in the channel without being read without blocking the other publisher

  myChannel <- 1
  myChannel <- 2 // will have created a deadlock if not buffered.

  close(myChannel) // Close the channel, it is not mandatory !

  x, isChannelClosed := <-myChannel // x = 1, the second indicate if the channel has been closed (can be closed without being empty !)

  fmt.Println(x, isChannelClosed)
}


//============================ Use range on channels
// Publish on pipe
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func FirstRangeOnPipe() {
	c := make(chan int, 10)

  go fibonacci(cap(c), c)

  for i := range c { // read all the values into the pipe until pipe closure
		fmt.Println(i)
	}
}

// ======================= SELECT ==============
func infiniteFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // Block until one of the two action is possible
		case c <- x:                        // Publishing something on channel c
			x, y = y, x+y
		case <-quit:                        //OR receiving anything on channel quit
			fmt.Println("quit")
			return
    //add a "default :"" if we don't want to block
		}
	}
}

func FirstExitChannel() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // consume channel
		}
		quit <- 0
	}() // this is how to declare an "anonymous function"
	infiniteFibonacci(c, quit)
}


// And there is many ways to
