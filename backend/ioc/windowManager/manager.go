package windowManager

import (
	"smartPet/backend/ioc/windowManager/event"
	"sync"
)

type WindowManager struct {
	events map[event.Event]struct{}
	mu     sync.Mutex
}

func NewWindowManager() *WindowManager {
	return &WindowManager{
		events: make(map[event.Event]struct{}),
	}
}

// 注册并启动事件
func (wm *WindowManager) RegisterEvent(e event.Event) {
	wm.mu.Lock()
	if _, ok := wm.events[e]; ok {
		wm.mu.Unlock()
		return
	}
	wm.events[e] = struct{}{}
	wm.mu.Unlock()
	e.Start()
}

// 停止并移除事件
func (wm *WindowManager) UnregisterEvent(e event.Event) {
	wm.mu.Lock()
	if _, ok := wm.events[e]; ok {
		delete(wm.events, e)
		wm.mu.Unlock()
		e.Stop()
		return
	}
	wm.mu.Unlock()
}

// 停止所有事件
func (wm *WindowManager) StopAllEvents() {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	for e := range wm.events {
		e.Stop()
		delete(wm.events, e)
	}
}
