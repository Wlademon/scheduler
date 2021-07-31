package scheduler

type CommandPool struct {
	Commands map[Command]Worker
}

func GetEmptyCommandPool() CommandPool {
	return CommandPool{
		Commands: make(map[Command]Worker),
	}
}

func (p *CommandPool) SetCommand(nameCommand string, worker Worker) Command {
	c := Command(nameCommand)
	p.Commands[c] = worker

	return c
}

func (p CommandPool) ExistProcessingFunc(nameCommand string) bool {
	return p.Commands[Command(nameCommand)].ProcessingFunc != nil
}

func (p CommandPool) ExistExecutionFunc(nameCommand string) bool {
	return p.Commands[Command(nameCommand)].ExecutionFunc != nil
}

func (p CommandPool) GetProcessingFunc(command Command) func(pool *SchedulePool, args ...interface{}) (ResultWork, error) {
	return p.Commands[command].ProcessingFunc
}

func (p CommandPool) GetExecutionFunc(command Command) func(entity *CommandEntity) (ResultWork, error) {
	return p.Commands[command].ExecutionFunc
}

func (p *CommandPool) RemoveCommand(nameCommand string) {
	delete(p.Commands, Command(nameCommand))
}

func (p CommandPool) GetCommands() []Command {
	var buffer []Command
	for command, _ := range p.Commands {
		buffer = append(buffer, command)
	}

	return buffer
}

type Worker struct {
	ProcessingFunc func(pool *SchedulePool, args ...interface{}) (ResultWork, error)
	ExecutionFunc  func(entity *CommandEntity) (ResultWork, error)
}

type ResultWork struct {
	Value interface{}
}
