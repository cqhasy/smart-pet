package server

import (
	"changeme/backend/dialog/controller"
)

type DialogController interface {
	SayHello(screenX float64, screenY float64) string
}

// DialogServer 负责为业务层添加中间件逻辑和兜底等措施
type DialogServer struct {
	Service DialogController
}

func NewDialogServer(service *controller.DialogController) *DialogServer {
	return &DialogServer{
		Service: service,
	}
}

func (d *DialogServer) SayHello(screenX float64, screenY float64) string {
	return d.Service.SayHello(screenX, screenY)
}
