package main

import (
	"fmt"
	"math/rand"
	"time"
)

func interface1() {
	for i := 1; i <= 4; i++ {
		fmt.Println("[bisa1 bisa2 bisa3]", i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func interface2() {
	for i := 1; i <= 4; i++ {
		fmt.Println("[coba1 coba2 coba3]", i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 4; i++ {
		if i%2 == 0 {
			go interface2()
		} else {
			go interface1()
		}
	}

	time.Sleep(5 * time.Second)
}
