package core

import (
	"syscall/js"

	"github.com/syumai/go-hyperscript/dom"
	h "github.com/syumai/go-hyperscript/hyperscript"
)

var rootComponent h.StatelessComponent

var (
	body     = js.Global().Get("document").Get("body")
	renderer = dom.NewRenderer()
)

func SetRootComponent(c h.StatelessComponent) {
	rootComponent = c
}

func Update() {
	if rootComponent == nil {
		println("RootComponent is not found")
		return
	}
	renderer.Render(h.H(rootComponent, nil), body)
}
