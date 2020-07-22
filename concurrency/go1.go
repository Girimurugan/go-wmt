package main
 
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
 
// This assign a WaitGroup, to be used later
var wait sync.WaitGroup 
 
func Printer(a int) {
 
	// decrement semaphore counter at end
	defer wait.Done() 
 
	// sleep for a while
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Println("a is: ", a)
}
 
func main() {
	max_routine := 5
 
	// number of goroutines to wait for
	wait.Add(max_routine) 
 
	fmt.Println("At start")
	for i := 0; i < max_routine; i++ {
		fmt.Println("Sending to routine: ", i)
		go Printer(i)
	}
 
	// block until sempaphore counter decrements to 0. 
	wait.Wait() 
	fmt.Println("At end!")
}