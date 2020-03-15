package main

import (
	"WebPas/fetcher"
	"fmt"
	"regexp"
)



func main(){
		//tr:=&http.Transport{
		//	TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
		//}
		//client:=&http.Client{Transport:tr}
		//resp ,err:=client.Get("http://www.zhenai.com/zhenghun/wuhan")
		//if err!=nil{
		//	fmt.Println(err)
		//}
	fetch, err2 := fetcher.Fetch("http://www.zhenai.com/zhenghun/huanggang")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	const info = `<td width="180"><span class="grayL">性别：</span>([^<]*)+</td>`
	//const personalProfileRegexPink  = `<div data-v-8b1eac0c="" class="m-btn pink">([^<]*)+</div>`
	//<div class="m-btn purple" data-v-8b1eac0c="">离异</div>
	//const personalProfileRegexPurple  = `<div class="m-btn purple"[^>]*>([^<]*)+</div>`
	Purple := regexp.MustCompile(info)
	submatchPurple := Purple.FindAllSubmatch(fetch,-1)

	fmt.Printf("%s \n",submatchPurple)

}
