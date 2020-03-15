package parser

import (
	"WebPas/engine"
	"regexp"
)

const CityListRegex  = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRegex)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _,item := range matchs {
		result.Requests = append(result.Requests,
			engine.Request{
			Url:        string(item[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
