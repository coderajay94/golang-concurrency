package main

import (
	"fmt"
	"time"
)

func main() {

	//example of unbuffered channel
	now := time.Now()
	defer func() {
		fmt.Println("time elapsed: ", time.Since(now))
	}()

	evilNinja := "Raghav"
	smokeSignal := make(chan bool)
	go attack(evilNinja, smokeSignal)
	fmt.Println("smokeSignal: ", <-smokeSignal)
	//fmt.Println("smokeSignal2: ", <-smokeSignal)

}

func attack(ninja string, channel chan bool) {
	fmt.Println("attacked by ", ninja)
	time.Sleep(time.Second * 1)
	channel <- true
}
