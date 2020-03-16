package presist

import (
	"WebPas/engine"
	"WebPas/engine/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expectedProfile := model.PersonProfile{

		Marriage:   "离异",
		Age:        "32岁",
		Xinzuo:     "天蝎座(10.23-11.21)",
		Height:     "160cm",
		Weight:     "47kg",
		WorkWhere:  "工作地:阿坝汶川",
		Income:     "月收入:3-5千",
		Occupation: "医疗/护理",
		Education:  "大专",
		Minzu:      "羌族",
		Jiguan:     "籍贯:四川成都",
		Figure:     "体型:苗条",
		House:      "已购房",
		Car:        "已买车",
		Child:      "有孩子但不在身边",
		Name:       "奕然",
		Gender:     "女士",
	}
	expected := engine.Item{
		Id:      "1381240850",
		Url:     "https://album.zhenai.com/u/1381240850",
		Type:    "zhenai",
		Payload: expectedProfile,
	}

	err := save(expected)
	if err != nil{
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://47.98.246.49:9200"))

	service ,err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil{
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(service.Source,&actual)
	if err != nil{
		panic(err)
	}
	fmt.Println(actual)
}
