package main

import (
	"changeme/backend/dialog/controller"
	"changeme/backend/dialog/service"
	"changeme/backend/server"
	"embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

var appRef *application.App

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
}

// 后面改为通过wire依赖注入
func main() {
	app := application.New(application.Options{
		Name:        "pet",
		Description: "A demo of using raw HTML & CSS",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	se := server.NewServer(app)
	s := service.New()
	c := controller.NewDialogController(se.Core, s)
	dialog := server.NewDialogServer(c)

	se.InitServer(dialog)
	err := se.Run()

	if err != nil {
		log.Fatal(err)
	}
}
