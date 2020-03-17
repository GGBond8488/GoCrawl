package presist

import (
	"WebPas/engine"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)


func ItemSaver()(chan engine.Item ,error){
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL(Host),
	)
	if err !=nil {
		return nil,err
	}
	go func() {
		itemCount := 1
		for {
			item := <- out
			log.Printf("Item Saver:got item #%d %v",itemCount,item)
			itemCount++
			err := Save(item,client,Index)
			if err != nil{
				log.Printf("Item save error : %v,%s",item,err)
			}
		}
	}()
	return out,nil
}

const Host  = "http://47.98.246.49:9200"
const Index  = "dating_profile"

func Save(item engine.Item,client *elastic.Client,index string)error{

	if item.Type == ""{
		return errors.New("type 不能为空")
	}
	indexService:= client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id !=""{
		indexService.Id(item.Id)
	}
	_,err :=indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
