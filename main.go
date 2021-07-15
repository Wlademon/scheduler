package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Wlademon/scheduler/scheduler"
)

func main() {
	var test scheduler.Command = "123"
	var test2 scheduler.Command = "123gf"
	pool := scheduler.GetEmptyPool()
	pool.AddRepeatCommand(test, []string{"test"}, false, time.Millisecond)
	pool.AddScheduleCommand(test2, "test", true, time.Hour*22)

	pool.Each(func(entity *scheduler.CommandEntity) bool {
		return false
	}, time.Now(), nil)
	pool.RemoveEntityByCommand(string(test2))
	fmt.Println(strings.Join(pool.GetCommands(), "\n"))

	fmt.Println(pool)
}
