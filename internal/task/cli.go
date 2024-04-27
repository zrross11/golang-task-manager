package task

// Task represents a single task
type Task struct {
	ID   int
	Name string
}

// TaskManager represents the task manager
type TaskManager struct {
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// AddTask adds a new task to the task manager
func (tm *TaskManager) AddTask(name string) {
}

// ListTasks lists all the tasks in the task manager
func (tm *TaskManager) ListTasks() {
}

// RemoveTask removes a task from the task manager
func (tm *TaskManager) RemoveTask(id int) {
}

// RunCLI runs the CLI task manager
func RunCLI() {
}
