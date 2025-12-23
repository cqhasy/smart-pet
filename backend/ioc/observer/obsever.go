package observer

/*
设计初衷：为实现子窗口跟随主窗口，并在子窗口频繁创建与销毁中保持逻辑清晰与可读，
逻辑可复用，这里决定抽象出窗口跟随逻辑，并通过观察者模式与策略模式进行优化。

具体思路：
创建一个主体,其下有被观察对象与待通知对象。
定义观察指标，如对于跟随，观察指标就是主窗口的位置。
定义观察间隔。
定义更新方法。
更新方法应基于通知对象的状态而有所变化。
暂不考虑多观察主体。
以下是具体实现。
*/

type Subject interface {
	Observe()
	Inform()
}

// Observer 用于通用观察者模式，这里未直接使用，但保留以拓展。
type Observer interface {
	Update()
}

// PositionStrategy 负责根据被观察者位置计算跟随者的位置
// 之后如果需要更复杂的行为（吸附、裁剪、避障），只需实现新的策略。
type PositionStrategy interface {
	Calc(subX, subY int) (objX, objY int)
}
