package execute

import (
	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/execute"
)

const (
	name = "default"
)

type DefaultExecute struct {
}

func (e *DefaultExecute) Execute(acc accumulator.Accumulator) error {
	return nil
}

func init() {
	execute.Add(name, func() plugin.Execute {
		return &DefaultExecute{}
	})
}
