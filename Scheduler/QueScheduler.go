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
	//初始化两个channel
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		//对应的两个队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			//获取队首的request和空闲worker
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0&&len(workerQ)>0{
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
				//如果接收到request，把request加入request队列
				case r := <-q.requestChan:
					requestQ = append(requestQ,r)
				//如果接收到worker的channel（说明这个worker空闲下来了），把这个加入队列
				case w:=<-q.workerChan:
					workerQ = append(workerQ,w)
					//如果可以成功将request送入wokerchannel,各自队列中的元素都少了一个
				case activeWorker<-activeRequest:
					workerQ = workerQ[1:]
					requestQ = requestQ[1:]
			}
		}

	}()
}