package trigger

import (
	"context"

	"github.com/inoth/trigger/event"
)

type Trigger struct {
	option

	event chan event.Event
}

func New(opts ...Option) *Trigger {
	o := defaultOption()
	for _, opt := range opts {
		opt(&o)
	}
	return &Trigger{
		option: o,
		event:  make(chan event.Event, o.eventSize),
	}
}

func (t *Trigger) Stop(ctx context.Context) error {
	return nil
}

func (t *Trigger) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-t.event:
			if !e.CanExecute() {
				t.event <- e
				continue
			}
			go e.Execute(ctx)
		default:
		}
	}
}

func (t *Trigger) AddEvent(e *event.Event) {
	t.event <- *e
}
