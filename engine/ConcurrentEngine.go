package engine

type ConEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}
type Scheduler interface {
	ReadyNotified
	Submit(Request)
	WorkerChan()chan Request
	Run()
}

type ReadyNotified interface {
	WorkerReady( chan Request)
}

func (e *ConEngine) Run(seeds ...Request) {
	out := make(chan ParseResult,10)
	e.Scheduler.Run()
	for i := 0;i < e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _,item := range result.Items{
			go func() {
				e.ItemChan<- item
			}()
		}

		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult, ready ReadyNotified) {
	go func() {
		for{
			ready.WorkerReady(in)
			//TELL SCHEDULER IM READY
			request := <-in
			worker,err := Worker(request)
			if err != nil{
				continue
			}
			out<-worker
		}
	}()
}
