package main


import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var i = 10
	mutex := &sync.RWMutex{}
	mutex.Lock()
	i++
	fmt.Println("Before RLock: ",i)
	mutex.Unlock()
	mutex.RLock()
	i++
	mutex.RUnlock()
	fmt.Println("After RLock: ", i)
	time.Sleep(time.Duration(1) * time.Second)
}