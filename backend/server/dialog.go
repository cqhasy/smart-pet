package server

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"smartPet/backend/dialog/controller"
	"smartPet/backend/dialog/service"
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

func RegisterDialogServer(s *Server) {
	se := service.New()
	c := controller.NewDialogController(s.Core, s.MainWindow, se)
	dialog := NewDialogServer(c)
	s.Core.RegisterService(application.NewService(dialog))
}

func (d *DialogServer) SayHello(screenX float64, screenY float64) string {
	return d.Service.SayHello(screenX, screenY)
}
