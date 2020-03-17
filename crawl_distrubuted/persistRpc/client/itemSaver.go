package client

import (
	"WebPas/crawl_distrubuted/config"
	"WebPas/crawl_distrubuted/rpc_support"
	"WebPas/engine"
	"log"
)


func ItemSaver(host string)(chan engine.Item ,error){
	out := make(chan engine.Item)
	client, err := rpc_support.NewClient(host)
	if err !=nil {
		return nil,err
	}
	if err !=nil {
		return nil,err
	}
	go func() {
		itemCount := 1
		for {
			item := <- out
			log.Printf("Item Saver:got item #%d %v",itemCount,item)
			itemCount++
			//call rpc save item
			result := ""
			err = client.Call(config.ItemSaverService,item,&result)
			if err != nil || result != "ok"{
				log.Printf("result: %s",result)
			}
			if err != nil{
				log.Printf("Item save error : %v,%s",item,err)
			}
		}
	}()
	return out,nil
}