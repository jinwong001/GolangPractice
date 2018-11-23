package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

func main() {
	//readPersons()
	readPersons1()
	//readXMLName()
}

type Result struct {
	Person []Person
}

type Person struct {
	Name      string
	Age       int
	Career    string
	Interests Interests
	Hobby     string `xml:"hobby,attr"`
	// tag 优先级高于，field，而且tag 也区分大小写，
}

type Interests struct {
	Interest []string
}

// 参考 href="http://blog.studygolang.com/2012/12/%E6%A0%87%E5%87%86%E5%BA%93-xml%E5%A4%84%E7%90%86%EF%BC%88%E4%B8%80%EF%BC%89/"
func readPersons() error {
	file := "xml/person.xml"
	//file:="/Users/wanny/Desktop/gowork/src/GolangPractice/xml/person.xml"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var persons Result
	err = xml.Unmarshal(content, &persons)
	if err != nil {
		return err
	}
	fmt.Printf("persons:%+v\n", persons)

	for _, p := range persons.Person {
		fmt.Printf("p:%+v\n", p)
	}
	return nil
}

type Result1 struct {
	XMLName xml.Name  `xml:"Persons"`
	Persons []Person1 `xml:"Person"`
}

type Person1 struct {
	Name      string
	Age       int
	Career    string
	Interests []string `xml:"Interests>Interest"`
}

func readPersons1() error {
	file := "xml/person.xml"
	//file:="/Users/wanny/Desktop/gowork/src/GolangPractice/xml/person.xml"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var persons Result1
	err = xml.Unmarshal(content, &persons)
	if err != nil {
		return err
	}
	fmt.Printf("persons:%+v\n", persons)

	for _, p := range persons.Persons {
		fmt.Printf("p:%+v\n", p)
	}
	return nil
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}

// 规则7： 如果某个XML元素的子元素的名字和 “a”或 “a>b>c”这种格式的tag的前缀匹配，
// Unmarshal会沿着XML结构向下寻找这样名字的元素，
// 然后将最里面的元素映射到struct的字段上。以”>”开始的tag和字段后面跟上”>”是等价的

// 规则11： 一个非指针的匿名struct字段会被这样处理：该字段的值是外部struct的一部分
type Result2 struct {
	XMLName xml.Name `xml:"Person"` // 一般建议根元素加上此字段,命令空间
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"` // 规则 7，可见字段名可以随意，
	Address                              // 规则11
}

func readXMLName() {

	v := Result2{Name: "none", Phone: "none"}
	data := `
    <Person>
        <FullName>Grace R. Emlin</FullName>
        <Company>Example Inc.</Company>
        <Email where="home">
            <Addr>gre@example.com</Addr>
        </Email>
        <Email where='work'>
            <Addr>gre@work.com</Addr>
        </Email>
        <Group>
            <Value>Friends</Value>
            <Value>Squash</Value>
        </Group>
        <City>Hanga Roa</City>
        <State>Easter Island</State>
    </Person>
`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}
