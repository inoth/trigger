package accumulator

type Accumulator interface {
	GetMatedata(key string) string

	SetBody([]byte)
	Body() []byte
	String() string
}

type accumulator struct {
	matedata map[string]string

	body []byte
}

func NewAccumulator(matedata map[string]string) Accumulator {
	return &accumulator{
		matedata: matedata,
	}
}

func (a *accumulator) GetMatedata(key string) string {
	return a.matedata[key]
}

func (a *accumulator) SetBody(buf []byte) {
	a.body = buf
}

func (a *accumulator) Body() []byte {
	return a.body
}

func (a *accumulator) String() string {
	return string(a.body)
}
