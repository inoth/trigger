package output

import "github.com/inoth/trigger/plugin"

type Creator func() plugin.Output

var Outputs = make(map[string]Creator)

func Add(name string, creator Creator) {
	Outputs[name] = creator
}
