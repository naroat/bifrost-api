package system

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type EcApi struct{}

func (e *EcApi) Ele(c *gin.Context) {
	var content map[string]interface{}
	//http://api.web.ecapi.cn/taoke/getTbkActivityInfo?apkey=0177bba7-68dc-5176-51c4-2e0c888b9942&tbname=qq1401696973&pid=mm_96078848_2668800494_112660200064&activity_material_id=2668800494
	url := "http://api.web.ecapi.cn/taoke/doItemHighCommissionPromotionLink?apkey=0177bba7-68dc-5176-51c4-2e0c888b9942&itemid=529048526383&pid=mm_96078848_2668800494_112660200064&tbname=qq1401696973"
	//url := "http://api.web.ecapi.cn/taoke/doItemHighCommissionPromotionLink?apkey=登录会员中心查看&itemid=529048526383&pid=mm_0000_0000_000&tbname=xxx"
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
	err = json.Unmarshal(body, &content)
	checkErr(err)

	fmt.Println(content)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
