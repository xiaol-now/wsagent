package main

import (
	"testing"
)

func Test_T1(t *testing.T) {
	ch1 := make(chan int)
	NotifyClose(ch1)
	for {
		select {
		case e := <-ch1:
			t.Logf("%#v", e)
			break
		}
	}

}
func NotifyClose(receiver chan int) chan int {
	close(receiver)
	return receiver
}