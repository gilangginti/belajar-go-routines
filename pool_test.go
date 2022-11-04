package belajar_go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Empty"
		},
	}
	pool.Put("Gilang")
	pool.Put("Fauzi")
	pool.Put("Muhammad")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data ", data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
