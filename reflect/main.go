package main

import (
	"reflect"
	"fmt"
)

// WsMessage websocket 消息类型
type WsMessage struct {
	Type      string `json:"type"`                // 消息类型
	ID        string `json:"id,omitempty"`        // 消息ID
	Data      string `json:"data,omitempty"`      // 消息 json字符串
	Timestamp int64  `json:"timestamp,omitempty"` // 时间戳, 后台生成
}

func main() {
	var wm *WsMessage
	t:=reflect.TypeOf(wm).String()
	fmt.Print(t)
}
