package taskmanager

func Handler(taskQueue *TaskQueue) error {
	for {
		task, err := taskQueue.Dequeue()
		// error handling or exit when task queue is empty
		if err != nil {
			if err == ErrQueueEmpty {
				return nil
			}
			return err
		}

		taskAction := task.GetAction()
		if taskAction != nil {
			taskAction()
		}
	}
}
