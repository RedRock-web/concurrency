// @program: concurrency
// @author: edte
// @create: 2020-07-16 20:00
package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestA(t *testing.T) {
	c := cron.New()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("everything 3 seconds executing")
	})
	c.Start()
	defer c.Stop()

	select {
	case <-time.After(time.Second * 10):
		return
	}
}
