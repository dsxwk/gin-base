package event

import (
	"sync"
	"time"
)

// Event 事件
type Event struct {
	Name string
	Data interface{}
}

// Events 事件
type Events struct {
	handlers []func(event Event, timestamp time.Time)
	lock     sync.RWMutex
}

// NewEvents 创建新的事件
func NewEvents() *Events {
	return &Events{
		handlers: make([]func(event Event, timestamp time.Time), 0),
	}
}

// Publish 发布事件
func (e *Events) Publish(event Event) {
	e.lock.RLock()
	defer e.lock.RUnlock()

	timestamp := time.Now()

	for _, handler := range e.handlers {
		go handler(event, timestamp)
	}
}

// RegisterAllEvent 注册所有事件
func (e *Events) RegisterAllEvent(handler func(event Event, timestamp time.Time)) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.handlers = append(e.handlers, handler)
}
