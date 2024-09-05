package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int = 0
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func main() {
	//read basic mutex lock
	//runner()

	//read and write mutex lock
	readAndWrite()

}

// ----------------------
func readAndWrite() {
	//use it for read intensive operation
	//lock can be held by multiple readers and 1 writer
	go read()
	go read()
	go read()
	go read()
	go write()
	time.Sleep(5 * time.Second)
	fmt.Println("done------------------------")
}

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Println("Aquiring read lock...")
	time.Sleep(1 * time.Second)
	fmt.Println("releasing read lock...")

}

func write() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	fmt.Println("Aquired write lock...")
	time.Sleep(1 * time.Second)
	fmt.Println("released write lock...")
}

//----------------------

func runner() {
	runner := 1000
	fmt.Println("hello world")
	for i := 0; i < runner; i++ {
		go increment()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("count is:", counter)
}

func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}
