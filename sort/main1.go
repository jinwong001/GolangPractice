package main1

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

func (ps PersonSlice) Less(i, j int) bool {
	return ps[j].Age < ps[i].Age //重写Less()方法，从大到小
}

func main() {
	intList := []int{2, 4, 8, 5, 3}
	floatList := []float64{1.2, 3, 5.8, 9.2, 2, 1}
	stringList := []string{"a", "c", "b", "d", "y", "e"}
	//从小到大，升序
	sort.Ints(intList)
	// sort.Float64s(floatList)
	// sort.Strings(stringList)

	// 从大到小，降序
	//sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	// sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))
	// sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	pos := sort.SearchInts(intList, 0)

	ps := []Person{
		{"zhang san", 12},
		{"li si  ", 14},
		{"wang wu ", 13},
	}

	sort.Sort(PersonSlice(ps)) //按照Age 逆序
	fmt.Println(ps)

	fmt.Printf("%v\n%v\n%v\n", intList, floatList, stringList)
	fmt.Print(pos)
}
