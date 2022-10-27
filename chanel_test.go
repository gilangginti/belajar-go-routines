package belajar_go_routines

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Gilang"
	}()
	data := <-channel
	fmt.Println(data)
}

func GiveMeRes(channel chan string) {
	channel <- "Fauzi"
}

func TestResponse(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeRes(channel)
	data := <-channel
	fmt.Println(data)
}

func TestBuffer(t *testing.T) {
	channel := make(chan int, 3)
	defer close(channel)
	// fmt.Println(cap(channel))
	for i := 0; i < cap(channel); i++ {
		channel <- i
		fmt.Println(<-channel)
	}
}
