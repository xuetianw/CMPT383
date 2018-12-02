package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO: initialize struct; start nWorkers workers as goroutines
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
}

func (queue WorkQueue) Shutdown() {
	// TODO: close .Jobs and remove all remaining jobs from the channel.
}
