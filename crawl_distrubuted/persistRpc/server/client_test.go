package main

import (
	"WebPas/crawl_distrubuted/rpc_support"
	"WebPas/engine"
	"WebPas/engine/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	//start itemSaver server
	//start client
	//call
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
	go serveRpc(":1234","test1")

	time.Sleep(5*time.Second)
	client,err := rpc_support.NewClient(":1234")
	if err != nil {
		panic(err)
	}
	result := ""
	err = client.Call("ItemSaverService.Save",expected,&result)
	if err != nil || result != "ok"{
		t.Errorf("result: %s",result)
	}
}
