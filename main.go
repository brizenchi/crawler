package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	// 启动engine
	e := &engine.ConcurrentEngine{}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}