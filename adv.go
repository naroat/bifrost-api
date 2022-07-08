package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var c map[string]interface{}
	//2668800494
	//url := "http://api.web.ecapi.cn/taoke/doItemHighCommissionPromotionLink?apkey=登录会员中心查看&itemid=529048526383&pid=mm_0000_0000_000&tbname=xxx"
	url := "http://127.0.0.1:8027/admin_api/passport/info"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	checkErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	//fmt.Printf("%v", string(body))
	//body = []byte(`{"code":-1,"data":{},"msg":"59.42.238.199 IP非法. 参见: http:\/\/doc.21ds.cn\/index?project_id=4685763267728060&doc_id=4736469516973094 RequestId：RFE403562228116389888"}`)
	err = json.Unmarshal(body, &c)
	checkErr(err)

	fmt.Println(c)
}

func mainBk() {
	//存储返回数据
	var c map[string]interface{}

	url := "http://api.web.ecapi.cn/taoke/doItemHighCommissionPromotionLink?apkey=登录会员中心查看&itemid=529048526383&pid=mm_0000_0000_000&tbname=xxx"
	//url := "https://www.baidu.com"
	resp, err := http.Get(url)
	checkErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	err = json.Unmarshal(body, &c)
	checkErr(err)
	fmt.Println(c)
	//fmt.Printf("%v", resp.Header.Get("Content-Type"))

	//{"code":-1,"data":{},"msg":"59.42.238.199 IP非法. 参见: http:\/\/doc.21ds.cn\/index?project_id=4685763267728060&doc_id=4736469516973094 RequestId：RFE403562228116389888"}
	//获取json

	//解析json

	return

	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	//fmt.Println(string(body))

	//Status:400 Bad Request
	//StatusCode:400
	//Proto:HTTP/1.1
	//ProtoMajor:1
	//ProtoMinor:1
	//Header:map[Connection:[keep-alive]
	//Content-Type:[application/json;charset=utf8]
	//Date:[Thu, 07 Jul 2022 08:08:59 GMT]
	//Server:[openresty]]
	//Body:0xc00009e040

	//fmt.Printf("%p\n", resp.Body)

	/*for k, v := range body {
		fmt.Println(k, v)
	}*/
	//showMap(resp)
	//fmt.Printf("%v\n", resp.Body)
}

func showMap(m interface{}) {
	collect := m.(map[string]string)
	for k, v := range collect {
		fmt.Println(k, v)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
