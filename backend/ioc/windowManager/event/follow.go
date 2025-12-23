package event

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type FollowEvent struct {
	Sub              atomic.Pointer[application.WebviewWindow]
	Obj              atomic.Pointer[application.WebviewWindow]
	offsetX, offsetY int
	stopCh           chan struct{}
	running          bool
	mu               sync.Mutex
}

func NewFollowEvent(sub, obj *application.WebviewWindow, offsetX, offsetY int) *FollowEvent {
	e := &FollowEvent{
		offsetX: offsetX,
		offsetY: offsetY,
	}
	e.Sub.Store(sub)
	e.Obj.Store(obj)
	return e
}

func (e *FollowEvent) GetType() EventType {
	return Follow
}

func (e *FollowEvent) Start() {
	e.mu.Lock()
	// 检测是不是已经启动或者窗口是否为null
	if e.running {
		e.mu.Unlock()
		return
	}
	// 可以启动就初始化启动条件
	e.stopCh = make(chan struct{})
	e.running = true
	e.mu.Unlock()

	// 启动轮询观察者，定期检查主窗口位置并更新跟随窗口
	go func() {
		ticker := time.NewTicker(16 * time.Millisecond) // 约 60fps
		defer ticker.Stop()

		for {
			select {
			case <-e.stopCh:
				e.mu.Lock()
				e.running = false
				e.mu.Unlock()
				return
			case <-ticker.C:
				e.updatePosition()
			}
		}
	}()
}

func (e *FollowEvent) Stop() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.running && e.stopCh != nil {
		close(e.stopCh)
		e.stopCh = nil
		e.running = false
	}
}

// updatePosition 根据当前偏移立即更新跟随窗口位置
func (e *FollowEvent) updatePosition() {
	// 快速退出：已停止
	e.mu.Lock()
	if !e.running {
		e.mu.Unlock()
		return
	}
	offsetX, offsetY := e.offsetX, e.offsetY
	e.mu.Unlock()

	sub := e.Sub.Load()
	obj := e.Obj.Load()
	if sub == nil || obj == nil {
		// 任一窗口已销毁，自动停止
		e.Stop()
		return
	}

	subX, subY := sub.Position()
	newObjX := subX + offsetX
	newObjY := subY + offsetY
	if newObjX < 0 {
		newObjX = 0
	}
	if newObjY < 0 {
		newObjY = 0
	}
	obj.SetPosition(newObjX, newObjY)
}
