package parser

import (
	"WebPas/engine"
	"regexp"
)
//<a href="http://www.zhenai.com/zhenghun/huanggang/2">下一页</a>
//<td width="180"><span class="grayL">性别：</span>女士</td>
const gender  = `<td width="180"><span class="grayL">性别：</span>([^<]*)+</td>`
const profileRegex  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)+</a>`
const nextPage  = `(href="http://www.zhenai.com/zhenghun/huanggang/[^"]+")`
var genderre = regexp.MustCompile(gender)
var re = regexp.MustCompile(profileRegex)
var city = regexp.MustCompile(nextPage)
func ParseCity(contents []byte)engine.ParseResult  {
	matchs := re.FindAllSubmatch(contents, -1)
	gendermatchs := genderre.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i,item := range matchs {
		name := string(item[2])
		gender := string(gendermatchs[i][1])
		url := string(item[1])
		result.Requests = append(result.Requests,
			engine.Request{
			Url:        url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes,name,gender,url)
			},
		})
	}

	citys := city.FindAllSubmatch(contents, -1)
	for _,cityItem := range citys{
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(cityItem[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
