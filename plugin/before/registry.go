package before

import "github.com/inoth/trigger/plugin"

type Creator func() plugin.Before

var Befores = make(map[string]Creator)

func Add(name string, creator Creator) {
	Befores[name] = creator
}
