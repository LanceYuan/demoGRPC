package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	nameRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">(.*?)</a></th>`)
)

func PerCity(contents []byte) engine.ParserResult {
	match := nameRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, item := range match {
		result.Items = append(result.Items, string(item[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(item[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
