package execute

import (
	"fmt"
	"time"

	"github.com/inoth/trigger/accumulator"
	"github.com/inoth/trigger/plugin"
	"github.com/inoth/trigger/plugin/execute"
)

const (
	name = "transcode"
)

type TranscodeExecute struct{}

func (e *TranscodeExecute) Execute(acc accumulator.Accumulator) error {
	fmt.Println("download remote vedio: ", acc.GetMatedata("url"))
	time.Sleep(time.Second)
	fmt.Println("transcode vedio")
	time.Sleep(time.Second)
	fmt.Println("upload vedio to cdn", acc.GetMatedata("url"))
	time.Sleep(time.Second)
	return nil
}

func init() {
	execute.Add(name, func() plugin.Execute {
		return &TranscodeExecute{}
	})
}
