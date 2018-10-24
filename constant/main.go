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
	var test = 8
	_ = test

	file := filepath.Base("口碑平台配劵页面.xlsx")
	fmt.Println(file)

	pow := []int{1, 3, 4, 4}
	for i := range pow {
		fmt.Println(i)
	}

	fmt.Println(m)
	m["fe"] = Vertex{40.68433, -74.39967}
	delete(m, "Google")
	fmt.Println(m)

J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 6 {
				//break J //现在终止的是j 循环，而不是i的那个
				continue J
			}
			fmt.Println("i:", i)
		}
		fmt.Println("j:", j)
	}

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
