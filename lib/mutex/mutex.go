package mutex

import "sync"

type M func() func()

func New() M {
	var m sync.Mutex

	return func() func() {
		m.Lock()
		return m.Unlock
	}
}
