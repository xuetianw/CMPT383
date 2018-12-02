package blockchain

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO

	q.NumWorkers = nWorkers
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, maxJobs)
	q.StopRequests = make(chan int, nWorkers)
	for i := uint(0); i < q.NumWorkers; i++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	// Run tasks from the queue, unless we have been asked to stop.
	for running {
		// TODO: run tasks from Jobs
		// TODO: stop when a signal is sent on StopRequests

		select {
		case _ = <-queue.StopRequests:
			running = false
			continue
		default:
		}

		select {
		case _ = <-queue.StopRequests:
			running = false
			continue
		case work := <-queue.Jobs:
			queue.Results <- work.Run()
		}
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO
	for i := uint(0); i < queue.NumWorkers; i++ {
		queue.StopRequests <- 0
	}
}
