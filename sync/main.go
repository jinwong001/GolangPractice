package main

import (
	"sync"
	"time"
	"fmt"
)

func main1() {
	var wg sync.WaitGroup
	mux := sync.Mutex{}
	sum := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1*time.Second)
			mux.Lock()
			sum +=i
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

type SafeCounter struct {
	v map[string] int
	mux sync.Mutex
}

func (c *SafeCounter)Inc(key string){
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter)Value(key string)int{
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main(){
	// 必须新建map
	c:=SafeCounter{v:make(map[string]int)}
	for i:=0;i<1000;i++{
		go c.Inc("someKey")
	}

	time.Sleep(2*time.Second)
	fmt.Println(c.Value("someKey"))




}


