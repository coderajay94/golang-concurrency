package main

import (
	"fmt"
	"sync"
)

func main() {

	var beeper sync.WaitGroup
	ninja := "ajay"
	beeper.Add(1)
	go runNinja(ninja, &beeper)
	beeper.Wait()
	fmt.Println("Starting")

	var neener sync.WaitGroup
	ninjas := []string{"raghu", "rajeev"}
	neener.Add(len(ninjas))
	for _, ninja := range ninjas {
		hitNinjas(ninja, &neener)
	}
	neener.Wait()
}

func hitNinjas(ninja string, neener *sync.WaitGroup) {
	fmt.Println("hitting Ninja -------<>-----", ninja)
	neener.Done()
}

func runNinja(ninja string, beeper *sync.WaitGroup) {
	fmt.Println("Ninja--------------->called Ninja", ninja)
	beeper.Done()
}
