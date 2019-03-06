package engine

import (
	"crawler/fetcher"
	"fmt"
)

type ConcurrentEngine struct {

}


func (e *ConcurrentEngine) Run(requests ...Request) {
	out := make(chan []Request)
	scheduler := &Scheduler{
		WorkerChan: make(chan Request),
	}

	go func() {
		for _, v := range requests {
			scheduler.Submit(v)
		}
	}()

	for i := 0; i < 10; i++ {
		CreateWorker(scheduler.WorkerChan, out)
	}
	for _, req := range <-out {
		scheduler.Submit(req)
	}
	// 获取剩下的任务 并加入到新的队列中
	//reqs, err := Work(req)
	//requests = append(requests, <-out...)


}

func CreateWorker(in chan Request, out chan []Request) {
	//go func() {
		for {
			request := <- in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	//}()
}


func Work(req Request) ([]Request, error) {
	// 获取数据
	content, err := fetcher.Fetch(req.Url)
	if err != nil {
		return nil, err
	}
	// 解析数据
	parserResult := req.ParserFunc(content)
	// 处理数据
	for _, item := range parserResult.Items {
		fmt.Println(item)
	}
	return parserResult.Request, nil
}
//func GetContent(url string) (<-chan []byte, error) {
//	out := make(chan []byte)
//	go func() {
//		out, err := <-fetcher.Fetch(url)
//		if err != nil {
//			return nil, err
//		}
//	}()
//	return out, nil
//}