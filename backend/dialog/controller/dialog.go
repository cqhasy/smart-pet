package controller

import (
	"changeme/backend/dialog/service"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"net/url"
	"time"
)

const (
	dialogWidth   = 260
	dialogHeight  = 160
	dialogTimeout = 3 * time.Second
)

type DialogService interface {
	SayHello() string
}

type DialogController struct {
	Core    *application.App
	dialog  *application.WebviewWindow
	Service DialogService
}

func NewDialogController(core *application.App, s *service.DialogService) *DialogController {
	return &DialogController{
		Core:    core,
		Service: s,
	}
}

func (c *DialogController) SayHello(screenX float64, screenY float64) string {
	line := c.Service.SayHello()
	if c.Core == nil {
		return line
	}

	if screenX < 0 {
		screenX = 0
	}
	if screenY < 0 {
		screenY = 0
	}

	if c.dialog != nil {
		c.dialog.Close()
		c.dialog = nil
	}
	dialog := c.Core.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:           "猫咪对话",
		Width:           dialogWidth,
		Height:          dialogHeight,
		InitialPosition: application.WindowXY,
		X:               int(screenX),
		Y:               int(screenY),
		Frameless:       true,
		AlwaysOnTop:     true,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             fmt.Sprintf("/dialog.html?message=%s", url.QueryEscape(line)),
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
		},
	})

	c.dialog = dialog

	go func(win *application.WebviewWindow) {
		time.Sleep(dialogTimeout)
		win.Close()
		if c.dialog == win {
			c.dialog = nil
		}
	}(dialog)

	return line
}
