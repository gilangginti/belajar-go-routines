package belajar_go_routines

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channnel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channnel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channnel)
	}()
	for data := range channnel {
		fmt.Println("menerima data ", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRes(channel1)
	go GiveMeRes(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
func TestSelectChannelDefault(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRes(channel1)
	go GiveMeRes(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
