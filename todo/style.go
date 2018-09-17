package todo

type style map[string]string

func (s style) String() (str string) {
	for k, v := range s {
		str += k + ": " + v + ";"
	}
	return
}
