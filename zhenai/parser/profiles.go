package parser

import (
	"WebPas/engine"
	"WebPas/engine/model"
	"regexp"
)

/*
个人资料
离异32岁天蝎座(10.23-11.21)160cm47kg工作地:阿坝汶川月收入:3-5千医疗/护理大专
羌族籍贯:四川成都体型:苗条不吸烟不喝酒已购房已买车有孩子但不在身边是否想要孩子:不想要孩子何时结婚:三年内
兴趣爱好
喜欢的一道菜：清蒸鲈鱼 欣赏的一个名人：韩红 喜欢的一首歌：爱囚 喜欢的一本书：读者 喜欢做的事：跑步游泳
择偶条件
35岁以上 175cm以上 工作地:四川成都都江堰市 月薪:8千以上 大学本科 离异 体型:运动员型 不要喝酒 不要吸烟 有孩子且偶尔会一起住 是否想要孩子:不想要孩子
*/
const personalProfileRegexPink  = `<div class="m-btn pink"[^>]*>([^<]*)+</div>`

const personalProfileRegexPurple  = `<div class="m-btn purple"[^>]*>([^<]*)+</div>`

var idUrlRe = regexp.MustCompile(`https://album\.zhenai\.com/u/([\d]+)`)
func ParseProfile(contents []byte,name string,gender string,url string)engine.ParseResult  {
	Purple := regexp.MustCompile(personalProfileRegexPurple)
	submatchPurple := Purple.FindAllSubmatch(contents,-1)
	Pink := regexp.MustCompile(personalProfileRegexPink)
	submatchPink := Pink.FindAllSubmatch(contents,-1)
	var Profile model.PersonProfile
	if len(submatchPurple)==9&&len(submatchPink)==10{
		Profile = model.PersonProfile{
			Marriage:  string(submatchPurple[0][1]),
			Age:       string(submatchPurple[1][1]),
			Xinzuo:    string(submatchPurple[2][1]),
			Height:    string(submatchPurple[3][1]),
			Weight:    string(submatchPurple[4][1]),
			WorkWhere: string(submatchPurple[5][1]),
			Income:    string(submatchPurple[6][1]),
			Occupation: string(submatchPurple[7][1]),
			Education: string(submatchPurple[8][1]),
			Minzu:      string(submatchPink[0][1]),
			Jiguan:     string(submatchPink[1][1]),
			Figure:     string(submatchPink[2][1]),
			House:      string(submatchPink[5][1]),
			Car:        string(submatchPink[6][1]),
			Child:      string(submatchPink[7][1]),
		}
	}
	Profile.Name = name
	Profile.Gender = gender
	result := engine.ParseResult{
		Items:   []engine.Item{
			{
				Id:      ExtractString(url, idUrlRe),
				Url:     url,
				Type:    "zhenai",
				Payload: Profile,
			},
		},
	}
	return result
}

func ExtractString(url string,re *regexp.Regexp) string{
	match := re.FindSubmatch([]byte(url))
	if len(match)==2{
		return string(match[1])
	}else {
		return ""
	}
}

/*
离异
32岁
天蝎座(10.23-11.21)
160cm
47kg
工作地:阿坝汶川
月收入:3-5千
医疗/护理
大专
*/

