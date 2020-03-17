package main

import (
	"WebPas/Scheduler"
	"WebPas/crawl_distrubuted/config"
	"WebPas/crawl_distrubuted/persistRpc/client"
	"WebPas/engine"
	"WebPas/zhenai/parser"
)

func main(){
	itemChan, err := client.ItemSaver(config.HostAndPort)
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


