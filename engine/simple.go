package engine

import (
	"crawler/fetcher"
	"fmt"
	"github.com/lexkong/log"
)

type Engine interface {
	Run(request ...Request)
}

type SimpleEngine struct {

}


func (e *SimpleEngine) Run(requests ...Request) {

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		content, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Error("", err)
		}
		parserResult := req.ParserFunc(content)
		requests = append(requests, parserResult.Request...)

		for _, item := range parserResult.Items {
			fmt.Println(item)
		}
	}
}