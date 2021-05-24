package scheduler

import "time"

type TSchedulerTask struct {
	time time.Time
	name string
	text string
}

type tScheduler struct {
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
) *tScheduler {
	scheduler := new(tScheduler)
	scheduler.timeout = timeout
	scheduler.immedConsume = immed
	scheduler.delayConsume = delay
	return scheduler
}

func (sch *tScheduler) consume() {
	queueLen := len(sch.queue)
	if queueLen == 1 {
		task := <-sch.queue
		sch.immedConsume(task)
	} else if queueLen == 0 {
		sch.isSleep = true
	} else {
		var task []TSchedulerTask
		for len(sch.queue) > 0 {
			task = append(task, <-sch.queue)
		}
		sch.delayConsume(task)
	}
}

func (sch *tScheduler) Start() {
	sch.isSleep = false
	timeTickerChan := time.NewTicker(time.Second * time.Duration(sch.timeout))
	for !sch.isSleep {
		sch.consume()
		<-timeTickerChan.C
	}
}

func (sch tScheduler) Produce(task *TSchedulerTask) {
	sch.queue <- *task
}
