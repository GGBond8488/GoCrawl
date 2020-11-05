package main

import (
	"WebPas/Scheduler"
	"WebPas/engine"
	"WebPas/presist"
	"WebPas/zhenai/parser"
)

func main(){
	itemChan, err := presist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConEngine{

		Scheduler:   &Scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}


