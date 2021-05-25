package scheduler

import (
	"time"
)

type TScheduler interface {
	Start()
	Produce(task *TSchedulerTask)
}

type TSchedulerTask struct {
	Time time.Time
	Name string
	Text string
}

type sScheduler struct {
	queue        chan TSchedulerTask
	timeout      int
	isSleep      bool
	immedConsume func(task TSchedulerTask)
	delayConsume func(task []TSchedulerTask)
}

func NewScheduler(
	timeout int,
	immed func(task TSchedulerTask),
	delay func(task []TSchedulerTask),
) TScheduler {
	return &sScheduler{
		queue:        make(chan TSchedulerTask, 50),
		timeout:      timeout,
		isSleep:      false,
		immedConsume: immed,
		delayConsume: delay,
	}
}

func (sch *sScheduler) consume() {
	queueLen := len(sch.queue)
	if queueLen == 1 {
		task := <-sch.queue
		sch.immedConsume(task)
	} else if queueLen == 0 {
		// sch.isSleep = true
	} else {
		var task []TSchedulerTask
		for len(sch.queue) > 0 {
			task = append(task, <-sch.queue)
		}
		sch.delayConsume(task)
	}
}

func (sch *sScheduler) Start() {
	sch.isSleep = false
	timeTickerChan := time.NewTicker(time.Second * time.Duration(sch.timeout))
	for !sch.isSleep {
		go sch.consume()
		<-timeTickerChan.C
	}
}

func (sch sScheduler) Produce(task *TSchedulerTask) {
	sch.queue <- *task
}
