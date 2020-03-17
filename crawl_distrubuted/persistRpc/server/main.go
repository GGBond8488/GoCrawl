package main

import (
	"WebPas/crawl_distrubuted/config"
	"WebPas/crawl_distrubuted/persistRpc"
	"WebPas/crawl_distrubuted/rpc_support"
	"github.com/olivere/elastic/v7"
)

func main() {
	err := serveRpc(config.HostAndPort,config.Index)
	if err != nil{
		panic(err)
	}
}

func serveRpc(host string,index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL(config.Host))
	if err != nil{
		return err
	}
	return rpc_support.ServiceRpc(host,&persistRpc.ItemSaverService{
		Client: client,
		Index:  index,
	})
}