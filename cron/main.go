package main

import (
	"github.com/astaxie/beego/httplib"
	"log"
	"fmt"
	"time"
	"github.com/robfig/cron"
)

func main() {
	req := httplib.Get("https://getman.cn/echo").Debug(true)
	//.Debug(true)
	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)

	go func() {
		for {
			//fmt.Println(time.Now())
			<-time.After(3 * time.Second)
		}
	}()

	// cron 参考  https://www.cnblogs.com/zuxingyu/p/6023919.html
	c := cron.New()
	// * * * * * ？  秒 分 时  日  月  星期
	// * 匹配所有值
	// / 执行间隔
	// , 用于 多个时间 枚注
	// - 连续时间,与枚注类似
	// ？ 只用于 日期 及星期    可以使用 * 代替

	// @every 用法比较特殊，这是Go里面比较特色的用法。同样的还有
	// @yearly @annually @monthly @weekly @daily @midnight @hourly

	c.AddFunc("*/5 * * * * ?", func() {
		fmt.Print("cron1:")
		fmt.Println(time.Now())
	})

	c.AddFunc("@every 3s", func() {
		fmt.Print("cron2:")
		fmt.Println(time.Now())
	})
	c.Start()
	time.Sleep(1000 * time.Second)
}
