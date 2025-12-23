package event

type EventType int

type Event interface {
	GetType() EventType
	Start()
	Stop()
}

const (
	Follow EventType = 1 // 跟随
)
