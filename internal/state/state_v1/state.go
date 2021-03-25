package state_v1



type State struct {
	Version      string         `json:"Version" toml:"Version" yaml:"Version"`
	ActiveEditor string         `json:"ActiveEditor" toml:"ActiveEditor" yaml:"ActiveEditor"`
	Editors      []*EditorState `json:"Editors" toml:"Editors" yaml:"Editors"`
}

type EditorState struct {
	ID           string `json:"ID" toml:"ID" yaml:"ID"`
	Text         string `json:"Text" toml:"Text" yaml:"Text"`
	CursorRow    int    `json:"CursorRow" toml:"CursorRow" yaml:"CursorRow"`
	CursorColumn int    `json:"CursorColumn" toml:"CursorColumn" yaml:"CursorColumn"`
}

func (s *State) EditorByID(editor string) *EditorState {
	for _, e := range s.Editors {
		if e.ID == editor {
			return e
		}
	}
	return nil
}

func (s *State) EditorIndexByID(editor string) int {
	for i, e := range s.Editors {
		if e.ID == editor {
			return i
		}
	}
	return -1
}

func (s *State) ActiveEditorIndex() int {
	return s.EditorIndexByID(s.ActiveEditor)
}

func (s *State) Clone() *State {
	ns := *s
	for i, es := range s.Editors {
		nes := *es
		ns.Editors[i] = &nes
	}
	return &ns
}
