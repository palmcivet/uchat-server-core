package controller

import "main/scheduler"

func Qq(sch scheduler.TScheduler) func(p []byte) {

	return func(p []byte) {
		sch.Produce(&scheduler.TSchedulerTask{})
	}
}
