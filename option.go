package trigger

type Option func(opt *option)

type option struct {
	eventSize uint
}

func defaultOption() option {
	return option{
		eventSize: 1000,
	}
}

func SetEventSize(size uint) Option {
	return func(opt *option) {
		opt.eventSize = size
	}
}
