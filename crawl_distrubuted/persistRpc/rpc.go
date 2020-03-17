package persistRpc

import (
	"WebPas/engine"
	"WebPas/presist"
	"github.com/olivere/elastic/v7"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (saver ItemSaverService)Save(item engine.Item,result *string)error {
	err := presist.Save(item,saver.Client,saver.Index)
	log.Printf("Item %v saved",item)
	if err == nil {
		*result = "ok"
	}else {
		log.Printf("Error saving %v,%s",item,err)
	}
	return err
}