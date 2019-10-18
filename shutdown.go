package shutdown

import (
	"sync"
	"time"
)

const DefaultTimeoutSec = 30

type Sync struct {
	Signal     chan interface{}
	isShutdown bool
	timeout    int
	complete   chan interface{}
	lock       sync.Mutex
}

func New() *Sync {
	return NewTimeout(DefaultTimeoutSec)
}

func NewTimeout(timeoutSec int) *Sync {
	return &Sync{
		Signal:     make(chan interface{}, 0),
		isShutdown: false,
		timeout:    timeoutSec,
		complete:   make(chan interface{}, 0),
	}
}

func (s *Sync) IsShutdown() bool {
	return s.isShutdown
}

func (s *Sync) Start() {
	close(s.Signal)
}

func (s *Sync) Complete() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isShutdown {
		return
	}

	s.isShutdown = true
	close(s.complete)
}

func (s *Sync) WaitForTimeout() bool {
	select {
	case <-s.complete:
		return false
	case <-time.After(time.Duration(s.timeout) * time.Second):
		return true
	}
}
