package tests

import (
	"gin-base/common/global"
	"gin-base/config"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

// 测试事件发布和订阅
func TestEvent_PublishAndSubscribe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	var (
		eventName string
		eventData interface{}
		eventTime time.Time
	)

	// 订阅事件
	global.Event.Subscribe(func(event config.Event, timestamp time.Time) {
		defer wg.Done()
		eventName = event.Name
		eventData = event.Data
		eventTime = timestamp
	})

	// 发布事件
	global.Event.Publish(config.Event{
		Name: "test_event",
		Data: "test_data",
	})

	// 等待处理完成
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// 验证事件
		assert.Equal(t, "test_event", eventName)
		assert.Equal(t, "test_data", eventData)
		assert.WithinDuration(t, time.Now(), eventTime, 2*time.Second)
	case <-time.After(2 * time.Second):
		t.Error("event handler did not complete in time")
	}
}

// 取消订阅
func TestEvent_Unsubscribe(t *testing.T) {
	var (
		triggered bool
		mu        sync.Mutex
	)

	// 使用独立事件实例, 防止和其他测试用例冲突
	events := config.NewEvents()

	// 订阅事件
	id := events.Subscribe(func(event config.Event, timestamp time.Time) {
		mu.Lock()
		defer mu.Unlock()
		triggered = true
	})

	// 取消事件
	events.Unsubscribe(id)

	// 发布事件
	events.Publish(config.Event{
		Name: "unsub_test",
		Data: "should not be handled",
	})

	// 等待一段时间确认 handler 未被触发
	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	assert.False(t, triggered, "handler should not be triggered after unsubscribe")
}
