package main

import (
	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-wasm-todo-list/core"
	"github.com/syumai/go-wasm-todo-list/todo"
)

func main() {
	app := func(object h.Object) h.VNode {
		return h.H("div", nil,
			h.H("h2", nil, h.Text("ToDo List")),
			h.H(todo.ToDo, nil),
			h.H("a", h.Object{"href": "https://github.com/syumai/go-wasm-todo-list/"},
				h.Text("Show the code on GitHub"),
			),
		)
	}
	core.SetRootComponent(app)
	core.Update()
	select {}
}
