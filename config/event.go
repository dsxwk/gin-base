package config

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

var (
	EventBus = NewEvents() // 全局事件
)

func InitEvent() *Events {
	return EventBus
}

// Event 事件
type Event struct {
	Name string
	Data interface{}
}

// Events 事件
type Events struct {
	lock     sync.RWMutex
	handlers map[string]func(event Event, timestamp time.Time)
}

// NewEvents 创建新的事件
func NewEvents() *Events {
	return &Events{
		handlers: make(map[string]func(event Event, timestamp time.Time)),
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

// Subscribe 订阅事件
func (e *Events) Subscribe(handler func(event Event, timestamp time.Time)) string {
	e.lock.Lock()
	defer e.lock.Unlock()

	eventId := uuid.NewString()
	e.handlers[eventId] = handler

	return eventId
}

// Unsubscribe 取消订阅
func (e *Events) Unsubscribe(eventId string) {
	e.lock.Lock()
	defer e.lock.Unlock()

	delete(e.handlers, eventId)
}
