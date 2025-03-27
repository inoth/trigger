package event

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/internal/util"
	"github.com/inoth/trigger/plugin/after"
	"github.com/inoth/trigger/plugin/before"
	"github.com/inoth/trigger/plugin/execute"
)

type EventOption func(*Event)

type Event struct {
	id       string
	metadata map[string]string
	// 初始化调用插件、获取原始数据
	eventBefore string
	// 执行插件
	eventExecute string
	// 结束插件
	eventAfter string
	// 进入执行延迟（秒）
	delay uint
}

func NewEvent(opts ...EventOption) Event {
	e := Event{
		id:           util.UUID(),
		metadata:     make(map[string]string),
		eventBefore:  "default",
		eventExecute: "default",
		eventAfter:   "default",
		delay:        0,
	}
	for _, opt := range opts {
		opt(&e)
	}
	return e
}

func SetMetadata(key, value string) EventOption {
	return func(e *Event) {
		e.metadata[key] = value
	}
}

func SetBefore(eventBefore string) EventOption {
	if eventBefore == "" {
		eventBefore = "default"
	}
	return func(e *Event) {
		e.eventBefore = eventBefore
	}
}

func SetExecute(eventExecute string) EventOption {
	if eventExecute == "" {
		eventExecute = "default"
	}
	return func(e *Event) {
		e.eventExecute = eventExecute
	}
}

func SetAfter(eventAfter string) EventOption {
	if eventAfter == "" {
		eventAfter = "default"
	}
	return func(e *Event) {
		e.eventAfter = eventAfter
	}
}

func SetDelay(delay uint) EventOption {
	return func(e *Event) {
		e.delay = delay
	}
}

func (e *Event) before(acc accumulator.Accumulator) {
	beforeCreator, ok := before.Befores[e.eventBefore]
	if !ok {
		return
	}
	err := beforeCreator().Before(acc)
	if err != nil {
		log.Printf("before error: %v", err)
	}
}

func (e *Event) after(acc accumulator.Accumulator) {
	afterCreator, ok := after.Afters[e.eventAfter]
	if !ok {
		return
	}
	err := afterCreator().After(acc)
	if err != nil {
		log.Printf("after error: %v", err)
	}
}

func (e *Event) PluginID() string {
	return e.id
}

func (e *Event) Execute(ctx context.Context) {
	if e.delay > 0 {
		tk := time.NewTicker(time.Duration(e.delay) * time.Second)
		defer tk.Stop()
		<-tk.C
	}

	defer func(st time.Time) {
		log.Printf("[%s]执行完毕，耗时：%v\n", e.PluginID(), time.Since(st))
	}(time.Now())

	acc := accumulator.NewAccumulator(e.metadata)

	e.before(acc)

	executeCreator, ok := execute.Executes[e.eventExecute]
	if !ok {
		log.Printf("Execute event err %v\n", fmt.Errorf("execute %s not found", e.eventExecute))
		return
	}
	err := executeCreator().Execute(acc)
	if err != nil {
		log.Printf("Execute event err %v\n", err)
		return
	}

	e.after(acc)
}
