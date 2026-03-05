package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("World")
// 	say("hello")
// }

// Channel Example
// import "fmt"

// func sum(s []int, c chan int) {
// 	sum:=0
// 	for _, v := range s {
// 		sum += v
// 	}

// 	c <- sum
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c

// 	fmt.Println(x, y, x+y)
// }

// Ex 2 : Buffered Channels
// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// Ex 3 : range function to receive channels data
// import "fmt"
// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for range n {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i) // instead og doing x <-c, we get the values directly into i because the range function handles it
// 	}
// }

//  Ex 4
// import "fmt"
// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y

// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func () {
// 		for range 10 {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// Ex 5
import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)

// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("boom.")
// 			return
// 		default:
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// }

// Ex 6: this program will panic because 1000s of go rutines will try to manipulate the map in an in-ordered manner so to fix it will use mutex in ex 7
// type Counter struct {
// 	v map[string]int
// }

// func (c *Counter) Inc(key string){
// 	c.v[key]++
// }

// func (c *Counter) Value(key string) int {
// 	return c.v[key]
// }

// func main() {
// 	c := Counter{v: make(map[string]int)}
// 	for range 1000 {
// 		go c.Inc("somekey")
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }

// Ex 7
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Inc(key string){
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()

	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}