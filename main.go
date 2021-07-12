package main

import (
	"fmt"
	"time"

	"github.com/Wlademon/scheduler/scheduler"
)

func main() {
	var test scheduler.Command = "123"
	pool := scheduler.GetEmptyPool()
	pool.AddRepeatCommand(test, []string{}, false, time.Hour)
	pool.AddScheduleCommand(test, []string{}, true, time.Hour*23)
	pool.Each(func(entity *scheduler.CommandEntity) bool {
		fmt.Println((*entity).Type())
		return true
	}, time.Now())

	fmt.Println(pool)
}
