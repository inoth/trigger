package after

import (
	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/after"
)

const (
	name = "default"
)

type DefaultOutput struct {
}

func (e *DefaultOutput) After(acc accumulator.Accumulator) error {
	return nil
}

func init() {
	after.Add(name, func() plugin.After {
		return &DefaultOutput{}
	})
}
