package output

import (
	"fmt"

	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/output"
)

const (
	name = "default"
)

type DefaultOutput struct {
}

func (e *DefaultOutput) Output(acc accumulator.Accumulator) error {
	fmt.Printf("[%s][Output] matedata: %s\n", acc.GetMatedata("id"), acc.String())
	acc.SetBody([]byte("step:Output"))
	return nil
}

func init() {
	output.Add(name, func() plugin.Output {
		return &DefaultOutput{}
	})
}
