package store

type State struct {
	LastID int
	Title  string
	ToDos  []*ToDo
	hooks  []func()
}

func NewState() *State {
	return &State{}
}

func (s *State) Subscribe(f func()) {
	s.hooks = append(s.hooks, f)
}

func (s *State) onUpdated() {
	for _, hook := range s.hooks {
		hook()
	}
}
