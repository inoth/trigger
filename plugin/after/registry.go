package after

import "github.com/inoth/trigger/plugin"

type Creator func() plugin.After

var Afters = make(map[string]Creator)

func Add(name string, creator Creator) {
	Afters[name] = creator
}
