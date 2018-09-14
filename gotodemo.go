package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"

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
	shows := mahonia.NewDecoder("GB18030")
	ret := shows.ConvertString(string(body))
	fmt.Println("type:", reflect.TypeOf(ret))
	strsplit(ret)
	//	fmt.Println(ret)
	if response.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

func get() {
	//	http://api.k780.com:88/?app=phone.get&phone=13800138000&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json
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

// 解析手机号 运营商
func strsplit(vals string) {
	arr := strings.Split(vals, "=")
	jsonArr := arr[1]

	// other test
	ret, _ := json.Marshal(jsonArr)
	type Animal struct {
		CatName string
	}
	var animals []Animal
	err := json.Unmarshal(ret, &animals)

	fmt.Println(1111, ret, reflect.TypeOf(ret), 222, err, 333)

	jsonMobile := gjson.Get(jsonArr, "catName")

	fmt.Println(333, jsonMobile.String(), 444, reflect.TypeOf(jsonMobile.String()))
	fmt.Println(jsonArr)
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
