package before

import (
	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/before"
)

const (
	name = "default"
)

type DefaultInput struct {
}

func (e *DefaultInput) Before(acc accumulator.Accumulator) error {
	return nil
}

func init() {
	before.Add(name, func() plugin.Before {
		return &DefaultInput{}
	})
}
