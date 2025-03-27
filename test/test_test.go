package test

import (
	"context"
	"testing"
	"time"

	"github.com/inoth/trigger"
	"github.com/inoth/trigger/event"

	_ "github.com/inoth/trigger/plugin/after/all"
	_ "github.com/inoth/trigger/plugin/before/all"
	_ "github.com/inoth/trigger/plugin/execute/all"
)

func TestNewTrigger(t *testing.T) {
	tg := trigger.New()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 1)
		for i := range 5 {
			tg.SendEvent(event.NewEvent(
				event.SetMetadata("url", "http://www.baidu.com"),
				event.SetDelay(uint(5-i)),
				event.SetExecute("transcode"),
			))
		}
	}()

	if err := tg.Start(ctx); err != nil && err != context.DeadlineExceeded {
		t.Errorf("%v", err)
		return
	}

	t.Logf("ok")
}
