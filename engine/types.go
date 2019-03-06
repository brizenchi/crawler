package engine

type Request struct {
	Url string
	ParserFunc func([]byte) *ParserResult
}

type ParserResult struct {
	Request []Request
	Items	[]interface{}
}

func NilParser(b []byte) *ParserResult {
	return &ParserResult{}
}

type Scheduler struct {
	WorkerChan chan Request
}

func (s *Scheduler) Get() Request {
	return <- s.WorkerChan
}

func (s *Scheduler) Submit(request Request) {
	s.WorkerChan <- request
}
