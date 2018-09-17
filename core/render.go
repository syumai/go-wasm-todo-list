package core

import (
	"syscall/js"

	h "github.com/syumai/go-hyperscript"
)

var rootComponent h.StatelessComponent

var body = js.Global().Get("document").Get("body")

func SetRootComponent(c h.StatelessComponent) {
	rootComponent = c
}

func Update() {
	if rootComponent == nil {
		println("RootComponent is not found")
		return
	}
	body.Set("innerHTML", "")
	h.Render(h.H(rootComponent, nil), body)
}
