package main

import (
	"ChatApp/client"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	client *client.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) NewClient(name string, serverIP string) bool {
	a.client = client.NewClient(name)
	err := a.client.Connect(serverIP)

	return err == nil
}

func (a *App) DestroyClient() {
	a.client.Destroy()
}

func (a *App) Send(msg string) {
	err := a.client.Send(msg)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func (a *App) Receive() {
	for {
		data, err := a.client.Receive()
		if err != nil {
			break
		}
		runtime.EventsEmit(a.ctx, "receive-message", data)
	}
}
