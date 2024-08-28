package analysis

type State struct {
	// Map file name to content
	Documents map[string]string
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}
