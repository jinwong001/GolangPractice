package main2

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonWrapper struct{
	ps []Person
	by func(p,q *Person) bool
}

func (pw PersonWrapper) Len() int{
    return len(pw.ps)
}

func(pw PersonWrapper)Swap(i,j int){
	pw.ps[i],pw.ps[j]=pw.ps[j],pw.ps[i]
} 

func(pw PersonWrapper)Less(i,j int)bool{
	return pw.by(&pw.ps[i],&pw.ps[j])
}

func main() {
	ps := []Person{
		{"zhang san", 12},
		{"li si  ", 14},
		{"wang wu ", 13},
	}

	sort.Sort(PersonWrapper{ps,func(p,q *Person)bool{
         return p.Age>q.Age //按照Age 降序
	}}) 

	fmt.Println(ps)

	sort.Sort(PersonWrapper{ps,func(p,q *Person)bool{
		return p.Name>q.Name //按照Name升序
   }}) 

	fmt.Println(ps)
}
