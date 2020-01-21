package mockery

import "errors"

type Finder interface {
	Find(num int) error
}

type LinearFinder struct {
	data []int
}

func (l *LinearFinder) Find(num int) error {
	for ele := range l.data {
		if ele == num {
			return nil
		}
	}
	return errors.New("not found")
}
