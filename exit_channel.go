package mygomod

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	ExitCh   chan int
	NotifyCh chan string
	MsgCh    chan string
	Id       string
}

func NewTask(id string, exit chan int) *Task {
	return &Task{
		Id:       id,
		ExitCh:   exit,
		NotifyCh: make(chan string),
		MsgCh:    make(chan string, 10),
	}
}

func (t *Task) Do(val string) {
	t.MsgCh <- val
}

func (t *Task) Run() {
	fmt.Println("routine", t.Id, "start...")
loop:
	for {
		select {
		case msg := <-t.MsgCh:
			cnt, err := strconv.Atoi(msg)
			if err != nil {
				fmt.Println("err:", err)
			}
			fmt.Println("routine", t.Id, "cnt:", cnt)
			if cnt > 0 {
				time.Sleep(1 * time.Second)
				t.MsgCh <- strconv.Itoa(cnt - 1)
			}
		case <-t.ExitCh:
			break loop
		default:
			// fmt.Println("routine", t.Id, "sleep")
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("routine", t.Id, "stop...")
	t.NotifyCh <- t.Id
}

func (t *Task) Notify() chan string {
	return t.NotifyCh
}

var (
	taskMap map[string]*Task
	exitCh  = make(chan int)
)

func exitChannel() {
	taskMap = map[string]*Task{}
	for i := 0; i < 4; i++ {
		id := strconv.Itoa(i)
		task := NewTask(id, exitCh)
		taskMap[id] = task
		go task.Run()
	}

	taskMap["0"].Do("3")
	taskMap["1"].Do("4")
	time.Sleep(10 * time.Second)
	fmt.Println("exit now")
	// exitCh <- 1
	close(exitCh)
	for _, task := range taskMap {
		fmt.Println("routine", <-task.Notify(), "notified")
	}
}
