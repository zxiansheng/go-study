package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/axgle/mahonia"
)

func main() {
	//	get()
	test()
}

func test() {
	response, _ := http.Get("https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=15850781443")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("type:", reflect.TypeOf(body))
	shows := mahonia.NewDecoder("GB18030")
	ret := shows.ConvertString(string(body))
	fmt.Println(ret)
	if response.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

func get() {
	response, _ := http.Get("https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=15850781443")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	shows := ConvertToString(string(body), "gbk", "utf-8")
	fmt.Println(shows)
	if response.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

// 编码转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
