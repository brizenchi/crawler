package parser

import (
	"fmt"
	"regexp"
	"crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) *engine.ParserResult {
	fmt.Println(string(contents))
	re := regexp.MustCompile(cityListRe)
	maches := re.FindAllSubmatch(contents, -1)

	result := &engine.ParserResult{}
	for _, m := range maches {
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}