package server

import (
	"smartPet/backend/config"
	"smartPet/backend/util"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Server struct {
	Core       *application.App
	MainWindow *application.WebviewWindow
}

func NewServer(c *application.App) *Server {
	return &Server{Core: c}
}

func (s *Server) createMainWindow(conf *config.MainWindowConf) *application.WebviewWindow {
	maskImageBytes := util.TurnImgToTransparent(conf.Mask)

	main := s.Core.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "main",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Width:           conf.Width,
		Height:          conf.Height,
		DisableResize:   true, // 禁止调整窗口大小
		Frameless:       true, // 无边框
		AlwaysOnTop:     true,
		InitialPosition: application.WindowXY,
		X:               conf.X,
		Y:               conf.Y,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             "/",
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
			WindowMask:                        maskImageBytes,
			WindowMaskDraggable:               true, // 关闭遮罩拖动，只让猫本身拖动
		},
	})

	return main
}

func (s *Server) InitServer(conf *config.MainWindowConf) {
	// 创建主窗口
	s.MainWindow = s.createMainWindow(conf)

	// 注册服务
	RegisterDialogServer(s)
}

func (s *Server) Run() error {
	return s.Core.Run()
}
