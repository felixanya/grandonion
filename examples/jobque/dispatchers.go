package jobque

type Dispatcher struct {
	WorkerPool 	chan chan Job
	maxWorkers 	int
	quit 		chan bool
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// start number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatcher()
}

func (d *Dispatcher) Stop() {
	for ch := range d.WorkerPool{
		close(ch)
	}
	close(d.WorkerPool)

	go func() {
		d.quit <- true
	}()
}

func (d *Dispatcher) dispatcher() {
	for {
		select {
		case job := <- JobQueue:
			// receive a job request
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <- d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		case <-d.quit:
			return
		}
	}
}
