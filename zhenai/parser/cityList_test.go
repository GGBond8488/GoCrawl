package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents,err := ioutil.ReadFile("cityList_test.html")
	if err != nil{
		panic(err)
	}
	result := ParserCityList(contents)
	const resultSize  = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//var expectedCitys = []string{
	//	"阿坝","阿克苏","阿拉善盟",
	//}
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d"+"requests;but had %d",resultSize,len(result.Requests))
	}
	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
		t.Errorf("expected urls #%d:%s"+"but"+" was %s",i,url,result.Requests[i].Url)
		}
	}
	//for i,city := range expectedCitys{
	//	if result.Items[i].Payload.(string) != city{
	//		t.Errorf("expected citys #%d:%s"+"but"+" was %s",i,city,result.Items[i].Payload.(string))
	//	}
	//}
	//if len(result.Items) != resultSize{
	//	t.Errorf("result should have %d"+"requests;but had %d",resultSize,len(result.Items))
	//}
	//fmt.Printf("%s ",contens)
	//ParserCityList(contens)
}
