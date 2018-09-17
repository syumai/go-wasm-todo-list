package todo

type ToDoData struct {
	Title string
	Done  bool
	ID    int
}

type State struct {
	lastID int
	toDos  []ToDoData
}

var state = State{}

func appendToDo(title string) {
	state.lastID++
	state.toDos = append(state.toDos, ToDoData{
		Title: title,
		ID:    state.lastID,
	})
}

func updateToDo(id int, done bool) {
	var t *ToDoData
	for i := 0; i < len(state.toDos); i++ {
		p := &state.toDos[i]
		if p.ID == id {
			t = p
			break
		}
	}
	if t == nil {
		return
	}
	t.Done = done
}

func separateToDos(toDos []ToDoData) (doingToDos []ToDoData, doneToDos []ToDoData) {
	for _, toDo := range toDos {
		if toDo.Done {
			doneToDos = append(doneToDos, toDo)
		} else {
			doingToDos = append(doingToDos, toDo)
		}
	}
	return
}
