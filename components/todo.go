package components

import (
	"strconv"

	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-wasm-todo-list/store"
	"github.com/syumai/go-wasm-todo-list/style"
)

func toDoItem(props h.Object) h.VNode {
	updateToDo := props.Get("updateToDo").(func(int, bool))
	toDo := props.Get("toDo").(*store.ToDo)
	return h.H("li", nil,
		h.H("input", h.Object{
			"type":    "checkbox",
			"checked": toDo.Done,
			"onchange": h.FuncOf(func(this h.Value, args []h.Value) interface{} {
				updateToDo(toDo.ID, !toDo.Done)
				return nil
			}),
		}),
		h.Text(toDo.Title),
	)
}

func toDoList(props h.Object) h.VNode {
	toDos := props.Get("toDos").([]*store.ToDo)
	elements := make(h.VNodes, len(toDos))
	for i, toDo := range toDos {
		elements[i] = toDoItem(h.Object{
			"toDo":       toDo,
			"updateToDo": props.Get("updateToDo"),
		})
	}
	return h.H("ul", nil, elements...)
}

func ToDo(props h.Object) h.VNode {
	toDos := props.Get("toDos").([]*store.ToDo)
	doingToDos, doneToDos := store.SeparateToDos(toDos)
	return h.H("div", nil,
		h.H("div", h.Object{"className": "input"},
			h.H("form", h.Object{
				"autocomplete": "off",
				"onsubmit":     props.Func("appendToDo"),
			},
				h.H("input", h.Object{
					"type":        "text",
					"name":        "title",
					"placeholder": "Input title",
					"value":       props.String("title"),
					"oninput":     props.Func("setTitle"),
				}),
				h.H("button", nil, h.Text("Add")),
			),
		),
		h.H("div", h.Object{"style": style.ListContainerStyle},
			h.H("div", h.Object{"style": style.ListStyle},
				h.H("h3", nil, h.Text("Doing")),
				h.H("div", nil, h.Text("Count: "+strconv.Itoa(len(doingToDos)))),
				h.H(toDoList, h.Object{
					"updateToDo": props.Get("updateToDo"),
					"toDos":      doingToDos,
				}),
			),
			h.H("div", h.Object{"style": style.ListStyle},
				h.H("h3", nil, h.Text("Done")),
				h.H("div", nil, h.Text("Count: "+strconv.Itoa(len(doneToDos)))),
				h.H(toDoList, h.Object{
					"updateToDo": props.Get("updateToDo"),
					"toDos":      doneToDos,
				}),
			),
		),
	)
}
