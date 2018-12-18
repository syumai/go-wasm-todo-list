package components

import (
	"strconv"

	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-wasm-todo-list/store"
	"github.com/syumai/go-wasm-todo-list/style"
)

var (
	listContainerStyle = style.Style(
		style.Prop{"width", "100%"},
		style.Prop{"display", "flex"},
		style.Prop{"flex-wrap", "wrap"},
	)

	listStyle = style.Style(
		style.Prop{"width", "50%"},
		style.Prop{"max-width", "250px"},
	)
)

func toDoItem(props h.Object) h.VNode {
	updateToDo := props.Get("updateToDo").(func(int, bool))

	toDo, ok := props.Get("toDo").(*store.ToDo)
	if !ok {
		return h.BlankElement
	}

	return h.H("li", nil,
		h.H("input", h.Object{
			"type":    "checkbox",
			"checked": toDo.Done,
			"onchange": h.NewEventCallback(0, func(e h.Value) {
				updateToDo(toDo.ID, !toDo.Done)
			}),
		}),
		h.Text(toDo.Title),
	)
}

func toDoList(props h.Object) h.VNode {
	v := props.Get("toDos")
	toDos, ok := v.([]*store.ToDo)
	if !ok {
		return h.BlankElement
	}

	elements := make(h.VNodes, len(toDos))
	for i, t := range toDos {
		elements[i] = toDoItem(h.Object{
			"toDo":       t,
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
				"onsubmit":     props.Get("appendToDo"),
			},
				h.H("input", h.Object{
					"type":        "text",
					"name":        "title",
					"placeholder": "Input title",
					"value":       props.String("title"),
					"oninput":     props.Get("setTitle"),
				}),
				h.H("button", nil, h.Text("Add")),
			),
		),
		h.H("div", h.Object{"style": listContainerStyle},
			h.H("div", h.Object{"style": listStyle},
				h.H("h3", nil, h.Text("Doing")),
				h.H("div", nil, h.Text("Count: "+strconv.Itoa(len(doingToDos)))),
				h.H(toDoList, h.Object{
					"updateToDo": props.Get("updateToDo"),
					"toDos":      doingToDos,
				}),
			),
			h.H("div", h.Object{"style": listStyle},
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
