package parser

import (
	"WebPas/engine/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contens,err := ioutil.ReadFile("profiles_test.html")
	if err != nil{
		panic(err)
	}
	result := ParseProfile(contens,"奕然","女士")

	profile := result.Items[0].(model.PersonProfile)
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
	if profile!=expectedProfile{
		t.Errorf("expected %v : but was %v",expectedProfile,profile)
	}
}
