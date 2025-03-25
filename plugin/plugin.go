package plugin

import "github.com/inoth/trigger/accumulator"

type PluginWithID interface {
	ID() string
}

type Input interface {
	Init(accumulator.Accumulator) error
}

type Execute interface {
	Execute(accumulator.Accumulator) error
}

type Output interface {
	Output(accumulator.Accumulator) error
}
