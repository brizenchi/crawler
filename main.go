package main

import "crawler/rpc"

func main() {
	// 启动engine
	//e := &engine.ConcurrentEngine{}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//})
	rpc.Test()
}