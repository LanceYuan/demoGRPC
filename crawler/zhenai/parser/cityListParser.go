package parser

import (
	"crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>(.*?)</a>`)

func CityList(contents []byte) engine.ParserResult {
	match := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, item := range match {
		result.Items = append(result.Items, string(item[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(item[1]),
			ParserFunc: PerCity,
		})
	}
	return result
}
