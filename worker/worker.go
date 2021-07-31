package worker

import "github.com/Wlademon/scheduler/scheduler"

type CommandPool struct {
	Commands map[scheduler.Command]Worker
}

func (p *CommandPool) SetCommand(nameCommand string, worker Worker) scheduler.Command {
	c := scheduler.Command(nameCommand)
	p.Commands[c] = worker

	return c
}

func (p CommandPool) ExistProcessingFunc(nameCommand string) bool {
	return p.Commands[scheduler.Command(nameCommand)].ProcessingFunc != nil
}

func (p CommandPool) ExistExecutionFunc(nameCommand string) bool {
	return p.Commands[scheduler.Command(nameCommand)].ExecutionFunc != nil
}

func (p CommandPool) GetProcessingFunc(command scheduler.Command) func(pool *scheduler.SchedulePool, args ...interface{}) (ResultWork, error) {
	return p.Commands[command].ProcessingFunc
}

func (p CommandPool) GetExecutionFunc(command scheduler.Command) func(entity *scheduler.CommandEntity) (ResultWork, error) {
	return p.Commands[command].ExecutionFunc
}

func (p *CommandPool) RemoveCommand(nameCommand string) {
	delete(p.Commands, scheduler.Command(nameCommand))
}

func (p CommandPool) GetCommands() []scheduler.Command {
	var buffer []scheduler.Command
	for command, _ := range p.Commands {
		buffer = append(buffer, command)
	}

	return buffer
}

type Worker struct {
	ProcessingFunc func(pool *scheduler.SchedulePool, args ...interface{}) (ResultWork, error)
	ExecutionFunc  func(entity *scheduler.CommandEntity) (ResultWork, error)
}

type ResultWork struct {
	Value interface{}
}
