package main

import (
	"syscall/js"

	"github.com/syumai/go-hyperscript/dom"
	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-wasm-todo-list/components"
	"github.com/syumai/go-wasm-todo-list/store"
)

var (
	renderer = dom.NewRenderer()
	body     = js.Global().Get("document").Get("body")
	state    = store.NewState()
)

var (
	appendToDo = h.NewEventCallback(h.EventCallbackFlgPreventDefault, func(h.Value) {
		state.AppendToDo(state.Title)
		state.SetTitle("")
	})

	setTitle = h.NewEventCallback(h.EventCallbackFlgPreventDefault, func(event h.Value) {
		state.SetTitle(event.Get("target").Get("value").String())
	})
)

func render() {
	app := h.H("div", nil,
		h.H("h2", nil, h.Text("ToDo List")),
		h.H(components.ToDo, h.Object{
			"toDos":      state.ToDos,
			"title":      state.Title,
			"setTitle":   setTitle,
			"appendToDo": appendToDo,
			"updateToDo": func(id int, done bool) { state.UpdateToDo(id, done) },
		}),
		h.H("a", h.Object{"href": "https://github.com/syumai/go-wasm-todo-list/"},
			h.Text("Show the code on GitHub"),
		),
	)
	renderer.Render(app, body)
}

func main() {
	state.Subscribe(render)
	render()
	select {}
}
