package execute

import "github.com/inoth/trigger/plugin"

type Creator func() plugin.Execute

var Executes = make(map[string]Creator)

func Add(name string, creator Creator) {
	Executes[name] = creator
}
