package input

import "github.com/inoth/trigger/plugin"

type Creator func() plugin.Input

var Inputs = make(map[string]Creator)

func Add(name string, creator Creator) {
	Inputs[name] = creator
}
