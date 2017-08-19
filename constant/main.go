package main

import (
	"fmt"
	"path/filepath"
	_ "runtime"
)

const (
	A = iota
	B
	C
	D = iota
	E
)

func main() {

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	var test=8
	_ = test

	file := filepath.Base("口碑平台配劵页面.xlsx")
	fmt.Println(file)

	pow:=[]int{1,3,4,4}
	for i:=range pow{
		fmt.Println(i)
	}

	fmt.Println(m)
         m["fe"]=Vertex{40.68433, -74.39967}
	 delete(m,"Google")
	fmt.Println(m)




}

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

