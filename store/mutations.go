package store

func (s *State) SetTitle(title string) {
	defer s.onUpdated()

	s.Title = title
}

func (s *State) AppendToDo(title string) {
	defer s.onUpdated()

	s.LastID++
	s.ToDos = append(s.ToDos, &ToDo{
		Title: title,
		ID:    s.LastID,
	})
}

func (s *State) UpdateToDo(id int, done bool) {
	defer s.onUpdated()

	var t *ToDo
	for i := 0; i < len(s.ToDos); i++ {
		p := s.ToDos[i]
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
