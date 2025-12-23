package observer

import (
	"sync"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// FollowSub 作为被观察者，负责定期通知跟随观察者
type FollowSub struct {
	Tar       *application.WebviewWindow
	Observers map[string]*FollowObserver // 跟随观察者集合，按 id 管理，便于频繁增删
	mu        sync.Mutex
	ticker    *time.Ticker
	stopCh    chan struct{}
	Interval  time.Duration // 观察间隔，默认 16ms（约 60fps）
}

func NewFollowSub(tar *application.WebviewWindow) *FollowSub {
	return &FollowSub{
		Tar:       tar,
		Interval:  16 * time.Millisecond,
		Observers: make(map[string]*FollowObserver),
	}
}

// AddObserver 按 id 注册观察者；如果 id 已存在则覆盖
func (sub *FollowSub) AddObserver(id string, observer *FollowObserver) {
	sub.mu.Lock()
	defer sub.mu.Unlock()
	if sub.Observers == nil {
		sub.Observers = make(map[string]*FollowObserver)
	}
	sub.Observers[id] = observer
}

// RemoveObserver 按 id 移除观察者
func (sub *FollowSub) RemoveObserver(id string) {
	sub.mu.Lock()
	defer sub.mu.Unlock()
	delete(sub.Observers, id)
}

// Observe 启动轮询，定期通知观察者
func (sub *FollowSub) Observe() {
	sub.mu.Lock()
	if sub.stopCh != nil {
		sub.mu.Unlock()
		return
	}
	interval := sub.Interval
	if interval <= 0 {
		interval = 16 * time.Millisecond
	}
	sub.stopCh = make(chan struct{})
	sub.ticker = time.NewTicker(interval)
	sub.mu.Unlock()

	go func() {
		for {
			select {
			case <-sub.stopCh:
				return
			case <-sub.ticker.C:
				sub.Inform()
			}
		}
	}()
}

// Inform 通知所有观察者
func (sub *FollowSub) Inform() {
	sub.mu.Lock()
	tar := sub.Tar
	snapshot := make([]*FollowObserver, 0, len(sub.Observers))
	for _, ob := range sub.Observers {
		snapshot = append(snapshot, ob)
	}
	sub.mu.Unlock()

	if tar == nil {
		sub.Stop()
		return
	}

	x, y := tar.Position()
	for _, ob := range snapshot {
		ob.Update(x, y)
	}
}

// Stop 停止观察
func (sub *FollowSub) Stop() {
	sub.mu.Lock()
	if sub.stopCh != nil {
		close(sub.stopCh)
		sub.stopCh = nil
	}
	if sub.ticker != nil {
		sub.ticker.Stop()
		sub.ticker = nil
	}
	sub.mu.Unlock()
}
