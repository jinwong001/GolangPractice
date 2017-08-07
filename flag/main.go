package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//var ip = flag.Int("ip", 2, "help message for flagname")
	//log.Printf("flag ip:%d",&ip)
	//
	////os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	////你可以使用标准的索引位置方式取得单个参数的值。
	////arg := os.Args[3]
	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	////fmt.Println(arg)
	//
	//wordPtr := flag.String("word", "foo", "a string")
	//
	//fmt.Println(&wordPtr)

	//os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
	argsWithProg := os.Args
	fmt.Println("argsWithProg", argsWithProg)
	if len(os.Args) > 1 {
		argsWithoutProg := os.Args[1:]
		args1 := argsWithProg[1]
		fmt.Println("argsWithoutProg", argsWithoutProg)
		fmt.Println("args1", args1)
	}

	//标志 word并带有一个简短的描述。这里的 flag.String 函数返回一个字符串指针（不是一个字符串值），在下面我们会看到是如何使用这个指针的。
	wordPtr := flag.String("word", "foo", "a string")
	//使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	// 运行时 go run flag/main.go  -word="hh" -numb=34 -fork=true

	fmt.Println("wordPtr", *wordPtr)
	fmt.Println("numbPtr", *numbPtr)
	fmt.Println("boolPtr", *boolPtr)
	fmt.Println("svar", svar)

}
