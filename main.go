package main

import (
	"fmt"

	"github.com/Wlademon/scheduler/scheduler"
)

func main() {
	cp := scheduler.GetEmptyCommandPool()
	pool := scheduler.GetEmptySchedulePool(&cp)
	c := cp.SetCommand("first", scheduler.Worker{
		ProcessingFunc: func(pool *scheduler.SchedulePool, args ...interface{}) (scheduler.ResultWork, error) {
			return scheduler.ResultWork{Value: 1}, nil
		},
	})
	pool.GetCommands()

	fmt.Println(cp.GetCommands(), c)
}
