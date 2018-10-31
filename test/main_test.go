package main

import (
	"testing"
	"strconv"
	"time"
)

// mongo 测试案例

// 文件名以  _test.go 结尾
// 测试方法名以 Test开头,后面大写开头
// 测试方法参数为t *testing.T
func TestMngoId2oTime(t *testing.T) {
	mongoId := "5bcc58599d4ff4694cf21fa1"
	subString := mongoId[:8]
	num, err := strconv.ParseInt(subString, 16, 0)
	if err != nil {
		t.Fatal(err)
	}
	tim := time.Unix(num, 0)
	t.Logf("time:%s", tim)
}
