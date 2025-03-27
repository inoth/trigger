package plugin

import "github.com/inoth/trigger/accumulator"

type Before interface {
	Before(accumulator.Accumulator) error
}

type Execute interface {
	Execute(accumulator.Accumulator) error
}

type After interface {
	After(accumulator.Accumulator) error
}
