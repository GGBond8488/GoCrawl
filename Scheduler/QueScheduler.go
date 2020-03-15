package Scheduler

import "WebPas/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q QueueScheduler)WorkerChan()chan engine.Request  {
	return make(chan engine.Request)
}

func (q *QueueScheduler)Submit(r engine.Request)  {
	q.requestChan <- r
}

func (q *QueueScheduler)WorkerReady(w chan engine.Request)  {
	q.workerChan <- w
}


func (q *QueueScheduler)Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var acriveWorker chan engine.Request
			if len(requestQ)>0&&len(workerQ)>0{
				acriveWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ,r)
				case w:=<-q.workerChan:
					workerQ = append(workerQ,w)
			case acriveWorker<-activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}