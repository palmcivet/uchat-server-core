package scheduler

import (
	"time"
)

type Scheduler interface {
	Start()
	Produce(task *SSchedulerTask)
}

type SSchedulerTask struct {
	Type int
	Time time.Time
	Name string
	Text string
}

type sScheduler struct {
	queue        chan SSchedulerTask
	timeout      int
	logger       func(task *SSchedulerTask)
	immedConsume func(task SSchedulerTask)
	delayConsume func(task []SSchedulerTask)
}

func NewScheduler(
	timeout int,
	immed func(task SSchedulerTask),
	delay func(task []SSchedulerTask),
	logger func(task *SSchedulerTask),
) Scheduler {
	return &sScheduler{
		queue:        make(chan SSchedulerTask, 50),
		timeout:      timeout,
		logger:       logger,
		immedConsume: immed,
		delayConsume: delay,
	}
}

func (sch *sScheduler) consume() {
	queueLen := len(sch.queue)
	if queueLen == 1 {
		task := <-sch.queue
		sch.immedConsume(task)
	} else if queueLen > 0 {
		var task []SSchedulerTask
		for len(sch.queue) > 0 {
			task = append(task, <-sch.queue)
		}
		sch.delayConsume(task)
	}
}

func (sch *sScheduler) Start() {
	timeTickerChan := time.NewTicker(time.Second * time.Duration(sch.timeout))
	for {
		go sch.consume()
		<-timeTickerChan.C
	}
}

func (sch sScheduler) Produce(task *SSchedulerTask) {
	sch.logger(task)
	sch.queue <- *task
}
