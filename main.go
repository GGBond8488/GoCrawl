package main

import (
	"WebPas/Scheduler"
	"WebPas/engine"
	"WebPas/presist"
	"WebPas/zhenai/parser"
)

func main(){

	e := engine.ConEngine{
		Scheduler:   &Scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: presist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}


