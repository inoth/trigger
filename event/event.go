package event

import (
	"context"
	"log"
	"time"

	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin/execute"
	"github.com/inoth/trigger/plugin/input"
	"github.com/inoth/trigger/plugin/output"
)

type EventOption func(*Event)

type Event struct {
	metadata map[string]string
	// 初始化调用插件、获取原始数据
	input string
	// 执行插件
	execute string
	// 结束插件
	output string
	// 进入执行延迟（秒）
	delay uint

	executeAt time.Time
}

func NewEvent(opts ...EventOption) *Event {
	e := Event{
		metadata: make(map[string]string),
		input:    "default",
		execute:  "default",
		output:   "default",
		delay:    0,
	}
	for _, opt := range opts {
		opt(&e)
	}
	if e.delay > 0 {
		e.executeAt = time.Now().Add(time.Duration(e.delay) * time.Second)
	}
	return &e
}

func SetMetadata(key, value string) EventOption {
	return func(e *Event) {
		e.metadata[key] = value
	}
}

func SetInput(input string) EventOption {
	if input == "" {
		input = "default"
	}
	return func(e *Event) {
		e.input = input
	}
}

func SetExecute(execute string) EventOption {
	if execute == "" {
		execute = "default"
	}
	return func(e *Event) {
		e.execute = execute
	}
}

func SetOutput(output string) EventOption {
	if output == "" {
		output = "default"
	}
	return func(e *Event) {
		e.output = output
	}
}

func SetDelay(delay uint) EventOption {
	return func(e *Event) {
		e.delay = delay
	}
}

func (e *Event) CanExecute() bool {
	return e.delay == 0 || time.Now().After(e.executeAt)
}

func (e *Event) Execute(ctx context.Context) {
	inputCreator, ok := input.Inputs[e.input]
	if !ok {
		return
	}
	executeCreator, ok := execute.Executes[e.execute]
	if !ok {
		return
	}
	outputCreator, ok := output.Outputs[e.output]
	if !ok {
		return
	}

	acc := accumulator.NewAccumulator(e.metadata)

	inp := inputCreator()
	err := inp.Init(acc)
	if err != nil {
		log.Printf("Init event err %s", err)
		return
	}

	exc := executeCreator()
	err = exc.Execute(acc)
	if err != nil {
		log.Printf("Execute event err %s", err)
		return
	}

	outp := outputCreator()
	err = outp.Output(acc)
	if err != nil {
		log.Printf("Output event err %s", err)
		return
	}
}
