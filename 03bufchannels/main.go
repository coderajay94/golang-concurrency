package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("buffered channel demo")
	channel := make(chan string)

	//check 1
	//counter := 5
	// go attack(channel, counter)

	// for i := 0; i < counter; i++ {
	// 	fmt.Println(<-channel)
	// }
	// fmt.Println("-------------")

	//check 2
	go attack2(channel)

	for {
		scoreCard, hasData := <-channel
		if !hasData {
			break
		}
		fmt.Println(scoreCard)
	}
	fmt.Println("-------------")

}

func attack(buff chan string, counter int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < counter; i++ {
		score := rand.Intn(10)
		buff <- fmt.Sprintf("hit a score of %d", score)
	}
	fmt.Println("this statement will not execute,as goroutine will not push anything into channel and main will finish")
}

func attack2(buff chan string) {
	rand.Seed(time.Now().UnixNano())
	randome := rand.Intn(10)
	fmt.Println("generated random no is:", randome)
	for i := 0; i < randome; i++ {
		score := rand.Intn(10)
		buff <- fmt.Sprintf("got a score of %d", score)
	}
	fmt.Println("unblocking main function as not putting any data into channel now")
	close(buff)
}
