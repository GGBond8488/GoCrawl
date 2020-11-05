package engine

import (
	"WebPas/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r := requests[0]
		requests = requests[1:]
		parseResult,err := Worker(r)
		if err!=nil{
			continue
		}
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items {
			log.Printf("Got Item %v \n", item)
		}
	}
}

func Worker(r Request) (ParseResult,error) {
	log.Printf("Fetching url:%v",r.Url)
	body , err := fetcher.Fetch(r.Url)
	if err!=nil{
		log.Println("Fetcher err:",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}
