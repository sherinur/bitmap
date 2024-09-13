package taskmanager

type Queue interface {
	Enqueue(task Task)
	Dequeue() (Task, error)
}

type TaskQueue struct {
	tasks []Task
	size  int
}

func (q *TaskQueue) Enqueue(task Task) {
	q.tasks = append(q.tasks, task)
	q.size++
}

func (q *TaskQueue) Dequeue() (Task, error) {
	if len(q.tasks) == 0 {
		return Task{}, ErrQueueEmpty
	}
	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	q.size--
	return task, nil
}

func (taskQueue *TaskQueue) GetTasks() []Task {
	return taskQueue.tasks
}

func (taskQueue *TaskQueue) GetSize() int {
	return taskQueue.size
}
