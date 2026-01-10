package main

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	tasks []Task
}

func NewScheduler() Scheduler {
	return Scheduler{
		tasks: []Task{},
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)

	s.tasks = normalizeUp(s.tasks, len(s.tasks)-1)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	for i, task := range s.tasks {
		if task.Identifier == taskID {
			s.tasks[i].Priority = newPriority
			s.normalize(i)

			return
		}
	}
}

func (s *Scheduler) GetTask() Task {
	result := s.tasks[0]
	s.tasks[0] = s.tasks[len(s.tasks)-1]
	s.tasks = s.tasks[:len(s.tasks)-1]
	s.normalize(0)

	return result
}

func (s *Scheduler) normalize(i int) {
	var isChanged bool

	if s.tasks, isChanged = normalizeDown(s.tasks, i); !isChanged {
		s.tasks = normalizeUp(s.tasks, i)
	}
}

func normalizeDown(tasks []Task, i int) ([]Task, bool) {
	isChanged := false

	for {
		leftChildIndex := 2*i + 1
		rightChildIndex := leftChildIndex + 1
		largestChildIndex := i

		if leftChildIndex < len(tasks) && tasks[leftChildIndex].Priority > tasks[largestChildIndex].Priority {
			largestChildIndex = leftChildIndex
		}

		if rightChildIndex < len(tasks) && tasks[rightChildIndex].Priority > tasks[largestChildIndex].Priority {
			largestChildIndex = rightChildIndex
		}

		if largestChildIndex == i {
			break
		}

		tasks[i], tasks[largestChildIndex] = tasks[largestChildIndex], tasks[i]
		i = largestChildIndex
		isChanged = true
	}

	return tasks, isChanged
}

func normalizeUp(tasks []Task, i int) []Task {
	parentIndex := (i - 1) / 2

	for i > 0 && tasks[parentIndex].Priority < tasks[i].Priority {
		tasks[i], tasks[parentIndex] = tasks[parentIndex], tasks[i]

		i = parentIndex
		parentIndex = (i - 1) / 2
	}

	return tasks
}
