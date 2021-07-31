package main

import (
	"fmt"

	"github.com/Wlademon/scheduler/scheduler"
	"github.com/Wlademon/scheduler/worker"
)

func main() {

	cp := new(worker.CommandPool)
	scheduler.GetEmptyPool(cp)
	cp.SetCommand("first", worker.Worker{})

	fmt.Println(cp.ExistProcessingFunc("first"))
}
