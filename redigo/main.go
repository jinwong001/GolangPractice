package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

// 参考 href="https://blog.csdn.net/wangshubo1989/article/details/75050024/"
func main1() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	//HMSET runoobkey name "redis tutorial" description "redis basic commands for caching" likes 20 visitors 23000
	_, err = c.Do("HMSET", "runoobkey", "name", "redis tutorial", "description", "redis")
	keys, err := c.Do("KEYS", "runoobke")
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", keys)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}

func main() {
	//setJson()
	//setExpire()
	//setPush()
	pipe()
}

func setJson() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, err := json.Marshal(imap)
	if err != nil {
		fmt.Println("marshal imap")
		return
	}

	n, err := c.Do("SETNX", key, value) // 若给定的 key 已经存在，则 SETNX 不做任何动作。
	if err != nil {
		fmt.Println("setnx error:", err)
		return
	}

	if n == int64(1) {
		fmt.Println("success")
	} else {
		fmt.Println("setnx error:", n)
		return
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

func setExpire() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	key := "key1"

	n, err := c.Do("EXPIRE", key, 24*3600)
	if err != nil {
		fmt.Println("expire error", err)
	}
	if n == int64(1) {
		fmt.Println("success")
	} else {
		fmt.Println("set expire failed:", n)
	}
}

func setPush() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}

func pipe() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo") // 向输出管道写入命令
	c.Flush()            // 将输出管道缓冲清空，并写入服务器

	v1,err:=c.Receive()          // reply from SET    //Recevie按照FIFO顺序依次读取服务器的响应
	if err != nil {
		fmt.Println("c.Receive", err)
	} else {
		fmt.Println("value", v1)  //OK
	}

	v2, err := c.Receive() // reply from GET
	if err != nil {
		fmt.Println("c.Receive", err)
	} else {
		fmt.Println("value", string(v2.([]byte)))
	}
}
