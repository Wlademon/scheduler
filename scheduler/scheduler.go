package scheduler

import (
	"time"

	"github.com/satori/go.uuid"
)

type PoolCommand struct {
	eCommand []CommandEntity
}

func GetEmptyPool() PoolCommand {
	return PoolCommand{}
}

func (p *PoolCommand) Each(f func(entity *CommandEntity) bool, timeNow time.Time, af func()) {
	var buffer []CommandEntity
	for _, entity := range p.eCommand {
		res := false
		if entity.SendNow(timeNow) {
			res = f(&entity)
		}
		if res {
			entity.Sent(timeNow)
		}
		if !entity.IsOnce() || !res {
			buffer = append(buffer, entity)
		}
	}
	if af != nil {
		af()
	}

	p.eCommand = buffer
}

func (p *PoolCommand) AddCommandEntity(entity CommandEntity) {
	entity.GetId()
	p.eCommand = append(p.eCommand, entity)
}

func (p *PoolCommand) AddRepeatCommand(command Command, args interface{}, once bool, timer time.Duration) {
	var temp = new(RepeatCommand)
	temp.ExCommand = SimpleCommand{
		CCommand: command,
		Args:     args,
	}
	temp.LastSend = time.Unix(0, 0)
	temp.Once = once
	temp.Timer = timer
	p.AddCommandEntity(temp)
}

func (p *PoolCommand) AddScheduleCommand(command Command, args interface{}, once bool, hmc time.Duration) {
	var temp = new(ScheduleCommand)
	temp.ExCommand = SimpleCommand{
		CCommand: command,
		Args:     args,
	}
	temp.LastSend = time.Unix(0, 0)
	temp.Once = once
	temp.Hmc = hmc
	p.AddCommandEntity(temp)
}

type TypeSchedule string

const (
	Repeat   TypeSchedule = "timer"
	Schedule TypeSchedule = "schedule"
)

type CommandEntity interface {
	GetId() string
	Command() ExecCommand
	SendNow(timeNow time.Time) bool
	Type() TypeSchedule
	IsOnce() bool
	Sent(timeNow time.Time)
}

type RepeatCommand struct {
	id        string
	ExCommand ExecCommand
	LastSend  time.Time
	Once      bool
	Timer     time.Duration
}

func (r RepeatCommand) Command() ExecCommand {
	return r.ExCommand
}

func (r RepeatCommand) SendNow(timeNow time.Time) bool {
	return r.LastSend.Add(r.Timer).Unix() <= timeNow.Unix()
}

func (r RepeatCommand) Type() TypeSchedule {
	return Repeat
}

func (r RepeatCommand) IsOnce() bool {
	return r.Once
}

func (r *RepeatCommand) Sent(timeNow time.Time) {
	r.LastSend = timeNow
}

func (r *RepeatCommand) GetId() string {
	if r.id == "" {
		r.id = uuid.NewV4().String()
	}
	return r.id
}

type ScheduleCommand struct {
	id        string
	ExCommand ExecCommand
	LastSend  time.Time
	Once      bool
	Hmc       time.Duration
}

func (s ScheduleCommand) Command() ExecCommand {
	return s.ExCommand
}

func (s ScheduleCommand) Type() TypeSchedule {
	return Schedule
}

func (s ScheduleCommand) IsOnce() bool {
	return s.Once
}

func (s *ScheduleCommand) Sent(timeNow time.Time) {
	s.LastSend = timeNow
}

func (s *ScheduleCommand) GetId() string {
	if s.id == "" {
		s.id = uuid.NewV4().String()
	}
	return s.id
}

func (s ScheduleCommand) SendNow(timeNow time.Time) bool {
	location := timeNow.Location()
	cYear, cMonth, cDay := timeNow.Date()
	lYear, lMonth, lDay := s.LastSend.Date()

	sNowDate := time.Date(cYear, cMonth, cDay, 0, 0, 0, 0, location)
	isBigStartDate := timeNow.Sub(sNowDate) > s.Hmc
	lastNotEqualNow := lYear != cYear || lMonth != cMonth || lDay != cDay

	return isBigStartDate && lastNotEqualNow
}

type Command string

type ExecCommand interface {
	GetCommand() Command
	GetArgs() interface{}
}

type SimpleCommand struct {
	CCommand Command
	Args     interface{}
}

func (s SimpleCommand) GetCommand() Command {
	return s.CCommand
}

func (s SimpleCommand) GetArgs() interface{} {
	return s.Args
}
