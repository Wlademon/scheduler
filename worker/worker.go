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

func (p *CommandPool) SetProcessingFunc(nameCommand string, f func(pool *scheduler.SchedulePool, args ...interface{}) (ResultWork, error)) {
	t := p.Commands[scheduler.Command(nameCommand)]
	t.ProcessingFunc = f
	p.Commands[scheduler.Command(nameCommand)] = t
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
