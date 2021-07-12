package boomer

import "sort"

var (
	order = func(t1, t2 *Task) bool {
		return t1.Order < t2.Order
	}
)

// Task is like the "Locust object" in locust, the python version.
// When boomer receives a start message from master, it will spawn several goroutines to run Task.Fn.
// But users can keep some information in the python version, they can't do the same things in boomer.
// Because Task.Fn is a pure function.
type Task struct {
	// The Order is used to keep the tasks order.
	Order int
	// The weight is used to distribute goroutines over multiple tasks.
	Weight int
	// The Name is used to mark the task name.
	Name string
	// Fn is called by the goroutines allocated to this task, in a loop.
	Fn func(Context)
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Task) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(tasks []*Task) {
	os := &orderSorter{
		tasks: tasks,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(os)
}

// orderSorter joins a By function and a slice of Planets to be sorted.
type orderSorter struct {
	tasks []*Task
	by    func(p1, p2 *Task) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *orderSorter) Len() int {
	return len(s.tasks)
}

// Swap is part of sort.Interface.
func (s *orderSorter) Swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *orderSorter) Less(i, j int) bool {
	return s.by(s.tasks[i], s.tasks[j])
}

func Sort(tasks []*Task) {
	By(order).Sort(tasks)
}
