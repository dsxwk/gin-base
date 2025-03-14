package config

import "gin-base/common/extend/event"

// InitEvent 初始化事件
func InitEvent() *event.Events {
	return event.NewEvents()
}
