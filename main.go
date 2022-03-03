package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"wangyufengGoGoGo.github.io/Regional-division/model"
)

const (
	key      = "2d5b8b130f923b20c5ca64a366f59a9a"
	url      = "https://restapi.amap.com/v3/config/district"
	typeId   = 99
	user     = "root"
	password = "123456"
	server   = "localhost:3306"
	db       = "test"
)

func main() {
	client := resty.New()
	body := map[string]string{
		"key":         key,
		"subdistrict": "3",
	}
	resp, err := client.R().SetQueryParams(body).Get(url)
	if err != nil {
		panic(fmt.Errorf("%s:%S", "请求接口失败", err))
	}

	result := model.Result{}
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		panic(fmt.Errorf("%s:%S", "解析失败", err))
	}

	AssembleData(result.Districts[0].Districts, 0)

}

func AssembleData(districts []*model.Dict, parentId int) []*model.ResultDict {
	resultDicts := []*model.ResultDict{}
	for _, district := range districts {
		adcode, _ := strconv.Atoi(district.Adcode)
		dict := model.ResultDict{
			TypeId:       typeId,
			DictId:       adcode,
			DictCode:     district.Adcode,
			DictName:     district.Name,
			DictParentId: parentId,
			Modify:       0,
			Visible:      1,
		}
		if len(district.Districts) > 0 {
			resultDicts = append(resultDicts, AssembleData(district.Districts, adcode)...)
		}
		resultDicts = append(resultDicts, &dict)
		fmt.Println(dict)
	}
	return resultDicts
}

func saveData(resultDict []*model.ResultDict) {
	fmt.Println("%s:%s@tcp(%s)/%s", user, password, server, db)
	//TODO
}
