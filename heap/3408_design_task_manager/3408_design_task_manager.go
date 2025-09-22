package designtaskmanager

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type Task struct {
	taskId   int
	priority int
}

type MaxHeap []Task

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}

	return h[i].taskId > h[j].taskId
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Task))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TaskManager struct {
	taskToPriority map[int][2]int
	priorHeap      *MaxHeap
}

func Constructor(tasks [][]int) TaskManager {
	priorHeap := &MaxHeap{}
	heap.Init(priorHeap)

	tm := TaskManager{
		priorHeap:      priorHeap,
		taskToPriority: map[int][2]int{},
	}

	for _, task := range tasks {
		tm.Add(task[0], task[1], task[2])
	}

	return tm
}

// It is guaranteed that taskId does not exist in the system.
func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.taskToPriority[taskId] = [2]int{userId, priority}
	task := Task{
		taskId:   taskId,
		priority: priority,
	}
	heap.Push(this.priorHeap, task)
}

// It is guaranteed that taskId exists in the system.
func (this *TaskManager) Edit(taskId int, newPriority int) {
	task := this.taskToPriority[taskId]
	task[1] = newPriority
	this.taskToPriority[taskId] = task

	newTask := Task{
		taskId:   taskId,
		priority: newPriority,
	}

	heap.Push(this.priorHeap, newTask)
}

// It is guaranteed that taskId exists in the system.
func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskToPriority, taskId)
}

func (this *TaskManager) ExecTop() int {
	for this.priorHeap.Len() > 0 {
		top := (*this.priorHeap)[0]
		if task, ok := this.taskToPriority[top.taskId]; ok && top.priority == task[1] {
			return task[0]
		}

		heap.Pop(this.priorHeap)
	}
	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
