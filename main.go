package main

import (
	"fmt"
	"time"

	config "github.com/devfabric/timingwheel/config"
	timingwheel "github.com/devfabric/timingwheel/wheel"
)

func main() {
	tmConfig, err := config.LoadTmWheelConfig("./")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tw := timingwheel.NewTimingWheel(time.Millisecond, tmConfig.Millisecond)
	tw.Start()
	defer tw.Stop()

	go func() {
		// exitC := make(chan time.Time, 1)
		tw.AfterFunc(time.Second, func() {
			fmt.Println("The timer fires")
			// exitC <- time.Now().UTC()
		})

	}()
	//<-exitC
	time.Sleep(2 * time.Second)
}
