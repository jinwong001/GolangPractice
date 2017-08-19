package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net/url"
)

func main()  {
	//httpGet()
	httpPost()
	//httpPostForm()
	//httpClient()
}



func httpGet(){
	resp,err:=http.Get("http://localhost:8080/say")
	if err!=nil{
		fmt.Print("http Get fail",err)
		return
	}

	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Print("http read body fail",err)
	}
	fmt.Println(string(body));
}

func httpPost(){
	//Tips：使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。
	resp,err:=http.Post("http://localhost:8080/say","application/x-www-form-urlencoded",strings.NewReader("name=test"))
	if err!=nil{
		fmt.Print("http Get fail",err)
		return
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Print("http read body fail",err)
	}
	fmt.Println(string(body));
}

func httpPostForm(){
	resp,err:=http.PostForm("http://localhost:8080/say",url.Values{"name":{"test"},"parameter":{"p1","p2"}})
	if err!=nil{
		fmt.Print("http Get fail",err)
		return
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("http read body fail",err)
	}
	fmt.Println(string(body))
}




func httpClient(){
	client:=&http.Client{}
	req,err:=http.NewRequest("POST",":8080/say",strings.NewReader("name=test"))
	if err!=nil{
		fmt.Println("http Get fail",err)
	}
	//同上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递。
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Cookie","name=anny")
	resp,err:=client.Do(req)

	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("http read body fail",err)
	}
	fmt.Println(string(body))
}

