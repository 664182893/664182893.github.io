package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("https://www.liwenzhou.com/")
	// if err != nil {
	// 	fmt.Printf("get failed, err:%v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("read from resp.Body failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Print(string(body))
	apiurl := "http://127.0.0.1:9090/get"
	data := url.Values{} //url Encode（比如传递参数会有些转义符，那么需要涉及到加密）
	data.Set("name", "xiaowangzi")
	u, err := url.ParseRequestURI(apiurl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed%v", err)
		return
	}
	fmt.Println(string(b))
}
