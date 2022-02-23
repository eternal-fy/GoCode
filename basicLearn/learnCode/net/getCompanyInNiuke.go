package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Schedules struct {
	Context, EncodeUrl, Name, Time, Url string
}
type PageInfoMsg struct {
	ElementCount,
	PageCount,
	//PageCurrent,
	PageSize int
}
type Company struct {
	CreateTime    int64       `json:"createTime"`
	Schedules     []Schedules `json:"schedules"`
	Name          string      `json:"name"`
	InterviewTime string      `json:"interviewTime"`
}
type MainData struct {
	CompanyList []Company   `json:"companyList"`
	PageInfo    PageInfoMsg `json:"pageInfo"`
}
type Resp struct {
	Code  int      `json:"code"`
	Msg   string   `json:"msg"`
	Datas MainData `json:"data"`
}

func main() {
	body, _ := http.Get("https://www.nowcoder.com/school/schedule/data?token=&query=&typeId=0&tab=0&propertyId=0&onlyFollow=false&page=1&_=1645516917838")
	bytes, _ := ioutil.ReadAll(body.Body)
	//println(string(bytes))
	var companys Resp
	json.Unmarshal(bytes, &companys)
	file, _ := os.OpenFile("d:/companys.txt", os.O_APPEND|os.O_CREATE, 0666)
	for _, m := range companys.Datas.CompanyList {
		schedules := m.Schedules
		url := ""
		for _, ans := range schedules {
			if ans.Name == "投递方式" || ans.Name == "投递链接" {
				url = ans.Url
			}

		}
		str := "公司名称：" + m.Name + " 投递链接：" + url + "\n"
		file.WriteString(str)

	}

}
