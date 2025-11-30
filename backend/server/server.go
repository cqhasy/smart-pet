package server

import "github.com/wailsapp/wails/v3/pkg/application"

type Server struct {
	Core *application.App
}

func NewServer(c *application.App) *Server {
	return &Server{Core: c}
}

func (s *Server) createMainWindow() {
	s.Core.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "main",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Width:           200,
		Height:          200,
		Frameless:       true, // 无边框
		AlwaysOnTop:     true,
		InitialPosition: application.WindowXY,
		X:               0,
		Y:               0,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             "/",
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
			WindowMaskDraggable:               true,
		},
	})
}

func (s *Server) InitServer(dialog *DialogServer) {
	// 创建主窗口
	s.createMainWindow()

	// 注册服务
	s.Core.RegisterService(application.NewService(dialog))
}

func (s *Server) Run() error {
	return s.Core.Run()
}
