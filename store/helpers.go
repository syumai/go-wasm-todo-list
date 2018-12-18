package store

func SeparateToDos(toDos []*ToDo) (doingToDos, doneToDos []*ToDo) {
	for _, toDo := range toDos {
		if toDo.Done {
			doneToDos = append(doneToDos, toDo)
		} else {
			doingToDos = append(doingToDos, toDo)
		}
	}
	return
}
