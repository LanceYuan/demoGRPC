package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChannel chan engine.Request
}

func (ss *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		ss.workChannel <- request
	}()
}

func (ss *SimpleScheduler) ConfigWorkChan(req chan engine.Request) {
	ss.workChannel = req
}
