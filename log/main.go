package main

import (
	"log"
	"fmt"
	"os"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

func main() {
	//arr := []int{2, 3}
	//log.Print("Print array ", arr, "\n")
	//test_deferfatal()
	//test_deferpanic()
	//test_customize()
	//test_logrus()
	test_logrus1()

}

// log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口
// 由于是直接调用系统接口退出，defer函数不会被调用
func test_deferfatal() {
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

//log.Panic接口，该函数把日志内容刷到标准错误后调用 panic 函数,在Panic之后声明的defer是不会执行的。
func test_deferpanic() {
	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		fmt.Println("--second--")
	}()
}

// 参考 https://studygolang.com/articles/9184
func test_customize() {
	fileName := "Info_First.log"
	log.Print("open file")

	// os.Create(fileName) 创建文件
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	//prefix 我们在前面已经看到，就是在日志内容前面的东西。我们可以将其置为 "[Info]" 、 "[Warning]"等来帮助区分日志级别。
	// Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
	debugLog := log.New(logFile, "[Info]", log.Llongfile)
	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
	debugLog.SetFlags(log.LstdFlags | log.Llongfile)
	debugLog.SetPrefix("[Warn]")
	debugLog.Println("A Warn Message here ")
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

func test_logrus() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus", "animal2": "walrus",
	}).Info("A walrus appears")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//logrus.WithFields(logrus.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func test_logrus1() {
	log := logrus.New()

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }
	log.Level = logrus.WarnLevel
	log.Formatter = new(MyJSONFormatter)

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"animal": "rabbit",
		"size":   1,
	}).Errorf("%d rabbits are running", 2)

	log.Info("show the info log")

}

type MyJSONFormatter struct {
}

func (f *MyJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//Note this doesn't include Time, Level and Message which are available on
	//the Entry. Consult `godoc` on information about those fields or read the
	//source of the official loggers.
	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil

	//return append([]byte(entry.Message), '\n'), nil
}
