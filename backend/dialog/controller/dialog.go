package controller

import (
	"fmt"
	"net/url"
	"smartPet/backend/ioc/observer"
	"sync"
	"time"

	"smartPet/backend/dialog/service"
	"smartPet/backend/ioc/windowManager"
	"smartPet/backend/ioc/windowManager/event"

	"github.com/wailsapp/wails/v3/pkg/application"
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
	Core        *application.App
	Manager     *windowManager.WindowManager
	dialog      *application.WebviewWindow
	followEvent *event.FollowEvent
	mainWindow  *application.WebviewWindow
	ObserverSub *observer.FollowSub
	mu          sync.Mutex
	Service     DialogService
}

func NewDialogController(core *application.App, mainWin *application.WebviewWindow, s *service.DialogService) *DialogController {
	return &DialogController{
		Core:        core,
		Service:     s,
		mainWindow:  mainWin,
		Manager:     windowManager.NewWindowManager(),
		ObserverSub: observer.NewFollowSub(mainWin),
	}
}

func (c *DialogController) createDialogWindow(screenX float64, screenY float64, line string) *application.WebviewWindow {
	return c.Core.Window.NewWithOptions(application.WebviewWindowOptions{
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
	})
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

	c.mu.Lock()
	oldDialog := c.dialog
	oldEvent := c.followEvent
	c.dialog = nil
	c.followEvent = nil
	c.mu.Unlock()

	// 关闭旧对话框并停止旧的跟随
	if oldEvent != nil {
		c.Manager.UnregisterEvent(oldEvent)
	}
	if oldDialog != nil {
		oldDialog.Close()
	}

	c.mu.Lock()
	mainWin := c.mainWindow
	c.mu.Unlock()

	var offsetX, offsetY int
	if mainWin != nil {
		mainX, mainY := mainWin.Position()
		offsetX = int(screenX) - mainX
		offsetY = int(screenY) - mainY
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
	})

	c.mu.Lock()
	c.dialog = dialog
	if mainWin != nil {
		c.followEvent = event.NewFollowEvent(mainWin, dialog, offsetX, offsetY)
	}
	c.mu.Unlock()

	// 注册跟随事件
	if mainWin != nil && c.followEvent != nil {
		c.Manager.RegisterEvent(c.followEvent)
	}

	// 自动关闭
	go func(win *application.WebviewWindow) {
		time.Sleep(dialogTimeout)
		c.mu.Lock()
		if c.dialog == win {
			evt := c.followEvent
			c.dialog = nil
			c.followEvent = nil
			c.mu.Unlock()

			if evt != nil {
				c.Manager.UnregisterEvent(evt) // 停止跟随
			}
			win.Close()
		} else {
			c.mu.Unlock()
		}
	}(dialog)

	return line
}

//func (c *DialogController) SayHello(screenX float64, screenY float64) string {
//	line := c.Service.SayHello()
//	if c.Core == nil {
//		return line
//	}
//
//	if screenX < 0 {
//		screenX = 0
//	}
//	if screenY < 0 {
//		screenY = 0
//	}
//	// 第一次创建对话窗口，直接新建窗口并建立监听事件
//	if c.dialog == nil {
//		c.dialog = c.createDialogWindow(screenX, screenY, line)
//		c.ObserverSub.AddObserver(c.dialog.Name(), observer.NewFollowObserver(c.dialog))
//	}
//
//}
