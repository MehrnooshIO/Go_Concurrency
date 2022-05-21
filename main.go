package main

import (
	"fmt"
)

func main()  {
	c := Generator()
	for i := 1; i < 5; i++ {
		fmt.Println("Channel says: ", <-c)
	}
}

// Generator returns a channel
func Generator() <- chan string {
	c:= make(chan string)
	go ChanWriter(c)
	return c
}
// ChanWriter fakes to do something and write the result in channel
func ChanWriter(c chan string) {
	func() {
		for i := 1; i < 5; i++ {
			c <- fmt.Sprintln("msg #", i)
		}
	}()
}

