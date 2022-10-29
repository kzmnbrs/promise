package future

import "context"

type Func struct {
	Func func() error
}

func FromFunc(f func() error) Future {
	return &Func{Func: f}
}

func (f *Func) Await(_ context.Context, _ context.CancelFunc) error {
	return f.Func()
}
