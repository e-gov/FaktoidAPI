package spordifakt


// Initial design from http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html
var WorkQueue = make(chan WorkRequest, 100)
var WorkerQueue chan chan WorkRequest
var ResponseQueue = make(chan HResponse, 100)
var Workers []Worker

type WorkRequest struct {
	URL string
}

type Worker struct {
	ID 		int
	Work		chan WorkRequest
	WorkerQueue 	chan chan WorkRequest
	QuitChan	chan bool
	ResultChan	chan HResponse
}

func NewWorker(id int, workerQueue chan chan WorkRequest, rc chan HResponse) Worker{
	return Worker{
		ID:		id,
		Work:		make(chan WorkRequest),
		WorkerQueue: 	workerQueue,
		QuitChan:	make(chan bool),
		ResultChan: 	rc}
}

func (w *Worker) Start(){
	go func(){
		for{
			w.WorkerQueue <- w.Work

			select{
			case work := <- w.Work:
				r := load(work.URL)
				w.ResultChan <- *r
			case <- w.QuitChan:
				return
			}
		}
	}()
}

func (w *Worker) Stop(){
	go func(){
		w.QuitChan <- true
	}()
}

func StartDispatcher(n int){
	WorkerQueue = make(chan chan WorkRequest, n)

	for i:= 0; i < n; i++{
		worker := NewWorker(i + 1, WorkerQueue, ResponseQueue)
		worker.Start()
		Workers = append(Workers, worker)
	}

	go func(){
		for{
			select{
			case work := <-WorkQueue:
				go func(){
					worker := <- WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}

func StopDispatcher(){
	for _, w := range Workers{
		w.Stop()
	}
}