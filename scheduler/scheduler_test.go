package scheduler

import (
	"fmt"
	"testing"
	"time"
)

func immed(data TSchedulerTask) {
	fmt.Printf(time.Now().String(), data.name, data.text, data.time)
}

func delay(data []TSchedulerTask) {
	for _, v := range data {
		fmt.Printf(time.Now().String(), v.name, v.text, v.time)
	}
}

func TestScheduler(t *testing.T) {
	t.Run("Scheduler", func(t *testing.T) {
		sch := NewScheduler(500, immed, delay)
		sch.Start()

		timeTickerChan := time.NewTicker(time.Second * time.Duration(300))
		for !sch.isSleep {
			fmt.Println(time.Now().String())
			sch.Produce(&TSchedulerTask{name: "Test", text: time.Now().String(), time: time.Now()})
			<-timeTickerChan.C
		}
	})
}
