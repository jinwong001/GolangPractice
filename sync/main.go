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
			time.Sleep(1 * time.Second)
			mux.Lock()
			sum += i
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main2() {
	// 必须新建map
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}

	time.Sleep(2 * time.Second)
	fmt.Println(c.Value("someKey"))
}

func main() {
	syncOnce()
}

func syncOnce() {
	// sync.Once 包含的程序只会执行一次
	var once sync.Once
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		//
		//go func() {
		//	// i 直接传入会有问题，会多次 i 相同,推荐下面
		//	fmt.Println("213 pre", i)
		//	once.Do(onced)
		//	fmt.Println("213")
		//}()

		go func(i int) {
			// i 直接传入会有问题，会多次 i 相同
			fmt.Println("213 pre", i)
			//onced 不会执行，once已经已经执行过一次
			once.Do(onced)
			fmt.Println("213")
		}(i)
	}
	time.Sleep(40000)
}

func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
