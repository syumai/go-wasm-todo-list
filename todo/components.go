package todo

import (
	"strconv"
	"syscall/js"

	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-wasm-todo-list/core"
)

var (
	listContainerStyle = style{
		"width":     "100%",
		"display":   "flex",
		"flex-wrap": "wrap",
	}.String()

	listStyle = style{
		"width":     "50%",
		"max-width": "250px",
	}.String()
)

func toDoItem(props h.Object) h.VNode {
	toDo, ok := props.Get("toDo").(*ToDoData)
	if !ok {
		return h.BlankElement
	}

	return h.H("li", nil,
		h.H("input", h.Object{
			"type":    "checkbox",
			"checked": toDo.Done,
			"data-id": toDo.ID,
			"onchange": core.EventAction(0, func(e js.Value) {
				t := e.Get("target")
				id, err := strconv.Atoi(t.Get("dataset").Get("id").String())
				if err != nil {
					return
				}
				checked := t.Get("checked").Bool()
				updateToDo(id, checked)
			}),
		}),
		h.Text(toDo.Title),
	)
}

func toDoList(props h.Object) h.VNode {
	v := props.Get("toDos")
	toDos, ok := v.([]ToDoData)
	if !ok {
		return h.BlankElement
	}

	elements := make(h.VNodes, len(toDos))
	for i, t := range toDos {
		elements[i] = toDoItem(h.Object{"toDo": &t})
	}
	return h.H("ul", nil, elements...)
}

func ToDo(props h.Object) h.VNode {
	doingToDos, doneToDos := separateToDos(state.toDos)
	return h.H("div", nil,
		h.H("div", h.Object{"className": "input"},
			h.H("form", h.Object{
				"action":       "#",
				"autocomplete": "off",
				"onsubmit": core.EventAction(js.PreventDefault, func(e js.Value) {
					title := e.Get("target").Get("title").Get("value").String()
					if len(title) == 0 {
						return
					}
					appendToDo(title)
				}),
			},
				h.H("input", h.Object{
					"type":        "text",
					"name":        "title",
					"placeholder": "Input title",
				}),
				h.H("button", nil, h.Text("Add")),
			),
		),
		h.H("div", h.Object{"style": listContainerStyle},
			h.H("div", h.Object{"style": listStyle},
				h.H("h3", nil, h.Text("Doing")),
				h.H("div", nil, h.Text("Count: "+strconv.Itoa(len(doingToDos)))),
				h.H(toDoList, h.Object{"toDos": doingToDos}),
			),
			h.H("div", h.Object{"style": listStyle},
				h.H("h3", nil, h.Text("Done")),
				h.H("div", nil, h.Text("Count: "+strconv.Itoa(len(doneToDos)))),
				h.H(toDoList, h.Object{"toDos": doneToDos}),
			),
		),
	)
}
