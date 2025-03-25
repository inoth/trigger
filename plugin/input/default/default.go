package input

import (
	"fmt"

	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/input"
)

const (
	name = "default"
)

type DefaultInput struct {
}

func (e *DefaultInput) Init(acc accumulator.Accumulator) error {
	fmt.Printf("[Init] matedata: %s\n", acc.String())
	acc.SetBody([]byte("step:Init"))
	return nil
}

func init() {
	input.Add(name, func() plugin.Input {
		return &DefaultInput{}
	})
}
