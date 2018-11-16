package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

func main() {
	//c := make(chan int)
	//a := []int{1, 8, 9, 4, 5}
	////sum(a, c)
	//go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)
	//x, y := <-c, <-c // 从 c 中获取
	//fmt.Println(x, y, x+y)

	//sendChan()
	////fmt.Scanln()
	//fmt.Println("done!")

	//threadPool1()
	// threadPool2()
	threadPool3()
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
	close(c)
}

func sendChan() {
	c := make(chan int)
	//go func() {
	//	c <- 1
	//}()
	c <- 1
	v := <-c
	fmt.Println(v)
}

func threadPool1() {
	// 原理：
	//  同时启动 3个 go routine
	//  job range 等待 过来的消息，进行处理
	// 三个工作线程同时从一个job Channel中取数据

	//两个channel，一个用来放置工作项，一个用来存放处理结果。
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// 添加9个任务后关闭Channel
	// channel to indicate that's all the work we have.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	// 添加9个任务后关闭Channel
	// channel to indicate that's all the work we have.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	//获取所有的处理结果
	for a := 1; a <= 9; a++ {
		<-result
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("work", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func threadPool2() {
	var p Pool
	url := []string{"11111", "22222", "33333", "444444", "55555", "66666", "77777", "88888", "999999"}
	p.Init(9, len(url))

	for i := range url {
		u := url[i]
		p.AddTask(func() error {
			return Download(u)
		})
	}

	p.SetFinishCallback(DownloadFinish)
	p.Start()
	p.Stop()
}

func Download(url string) error {
	time.Sleep(1 * time.Second)
	fmt.Println("Download " + url)
	return nil
}

func DownloadFinish() {
	fmt.Println("Download finsh")
}

type Pool struct {
	Queue         chan func() error
	RuntineNumber int
	Total         int

	Result         chan error
	FinishCallback func()
}

//初始化
func (self *Pool) Init(runtineNumber int, total int) {
	self.RuntineNumber = runtineNumber
	self.Total = total
	self.Queue = make(chan func() error, total)
	self.Result = make(chan error, total)
}

func (self *Pool) Start() {
	//开启 number 个goruntine
	// for 循环，一直等待<-self.Queue 任务过来，执行任务

	for i := 0; i < self.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task()
				self.Result <- err
			}
		}()
	}

	//获取每个任务的处理结果
	for j := 0; j < self.RuntineNumber; j++ {
		res, ok := <-self.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	//结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback()
	}
}

//关闭
func (self *Pool) Stop() {
	close(self.Queue)
	close(self.Result)
}

func (self *Pool) AddTask(task func() error) {
	self.Queue <- task
}

func (self *Pool) SetFinishCallback(fun func()) {
	self.FinishCallback = fun
}

func threadPool3() {
	// 原理
	// 通过 缓存通道 channel 超过容量 阻塞特性，
	// 一个工作线程完成，缓存通道，取出一个元素，然后 不阻塞，接收消息
	var wg sync.WaitGroup
	ch := make(chan struct{}, 100)
	for i := 0; i < 10000; i++ {
		ch <- struct{}{}
		wg.Add(1)

		go func(i int) {
			defer func() {
				<-ch
				wg.Done()
			}()
			task(i)
		}(i)
	}
	wg.Wait()
}

func task(i int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	println(i)
}
