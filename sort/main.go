package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (ps PersonSlice) Len() int {
	return len(ps)
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// func (ps PersonSlice) Less(i, j int) bool {
// 	return ps[j].Age < ps[i].Age //重写Less()方法，从大到小
// }

type ByName struct{
	PersonSlice
}

func( bn ByName)Less(i,j int) bool{
	return bn.PersonSlice[i].Name<bn.PersonSlice[j].Name
}

type ByAge struct{
	PersonSlice
}

func( ba ByAge)Less(i,j int) bool{
	return ba.PersonSlice[i].Age<ba.PersonSlice[j].Age
}




func main() {
	ps := []Person{
		{"zhang san", 12},
		{"li si  ", 14},
		{"wang wu ", 13},
	}

	sort.Sort(ByAge{ps}) //按照Age 升序
	fmt.Println(ps)

	
	sort.Sort(ByName{ps}) //按照Name 升序
	fmt.Println(ps)
}
