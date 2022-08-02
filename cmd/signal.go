package cmd

import (
	"context"
	"log"
	"os"
)

func SignalObserve(cancel context.CancelFunc) {
	ch := make(chan os.Signal)
	// signal.Notify(ch, syscall.SIGHUP, syscall.SIGURG)
	go func() {
		select {
		case s := <-ch:
			log.Println("接收了信号: ", s.String())
			cancel()
		}
	}()
}
