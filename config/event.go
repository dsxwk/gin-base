package config

import "gin-base/common/extend/event"

var (
	EventBus = event.NewEvents() // 全局事件
)

func InitEvent() *event.Events {
	return EventBus
}
