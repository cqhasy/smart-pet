package observer

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// FollowObserver 负责将跟随者移动到策略计算的位置
type FollowObserver struct {
	Target   *application.WebviewWindow
	Strategy PositionStrategy
}

func NewFollowObserver(target *application.WebviewWindow, strategy PositionStrategy) *FollowObserver {
	f := &FollowObserver{
		Strategy: strategy,
		Target:   target,
	}
	return f
}

// Update 根据被观察者的位置更新跟随窗口
// 如果目标窗口为空，则直接返回
func (o *FollowObserver) Update(subX, subY int) {
	if o.Target == nil || o.Strategy == nil {
		return
	}
	objX, objY := o.Strategy.Calc(subX, subY)
	if objX < 0 {
		objX = 0
	}
	if objY < 0 {
		objY = 0
	}
	o.Target.SetPosition(objX, objY)
}
