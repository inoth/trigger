package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/inoth/trigger"
	"github.com/inoth/trigger/event"

	_ "github.com/inoth/trigger/plugin/execute/default"
	_ "github.com/inoth/trigger/plugin/input/default"
	_ "github.com/inoth/trigger/plugin/output/default"
)

func TestNewTrigger(t *testing.T) {
	tg := trigger.New()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 2)
		for i := range 5 {
			go func() {
				tg.AddEvent(event.NewEvent(
					event.SetMetadata("id", fmt.Sprintf("%d", i)),
					event.SetDelay(uint(i)),
				))
			}()
		}
	}()

	if err := tg.Start(ctx); err != nil && err != context.DeadlineExceeded {
		t.Errorf("%v", err)
		return
	}

	t.Logf("ok")
}
