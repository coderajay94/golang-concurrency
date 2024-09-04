package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	defer func() {
		fmt.Println("time taken", time.Since(now))
	}()

	evilNinjas := []string{"raghu", "raghav", "ram", "raju"}

	for _, ninja := range evilNinjas {
		go attack(ninja)
	}

	time.Sleep(time.Second * 1)
}

func attack(ninja string) {
	fmt.Println("attacked by ", ninja)
}
