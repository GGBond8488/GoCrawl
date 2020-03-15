package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contens,err := ioutil.ReadFile("cityList_test.html")
	if err != nil{
		panic(err)
	}
	result := ParserCityList(contens)
	const resultSize  = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCitys = []string{
		"City 阿坝","City 阿克苏","City 阿拉善盟",
	}
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d"+"requests;but had %d",resultSize,len(result.Requests))
	}
	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
		t.Errorf("expected urls #%d:%s"+"but"+" was %s",i,url,result.Requests[i].Url)
		}
	}
	for i,city := range expectedCitys{
		if result.Items[i].(string) != city{
			t.Errorf("expected citys #%d:%s"+"but"+" was %s",i,city,result.Items[i].(string))
		}
	}
	if len(result.Items) != resultSize{
		t.Errorf("result should have %d"+"requests;but had %d",resultSize,len(result.Items))
	}
	//fmt.Printf("%s ",contens)
	//ParserCityList(contens)
}
