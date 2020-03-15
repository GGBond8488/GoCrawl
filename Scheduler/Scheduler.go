package Scheduler

import "WebPas/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (ssche *SimpleScheduler)WorkerChan()chan engine.Request {
	return ssche.workerChan
}

func (ssche *SimpleScheduler)Submit(r engine.Request)  {
	go func() {ssche.workerChan <- r}()
}
func (ssche *SimpleScheduler)Run()  {
	ssche.workerChan = make(chan engine.Request)
}

func (ssche *SimpleScheduler)WorkerReady(chan engine.Request)  {

}

