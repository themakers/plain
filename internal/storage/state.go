package storage

import (
	"github.com/themakers/plain/internal/state/state_v1"
	"gopkg.in/yaml.v2"
)

func newInitialState() *state_v1.State {
	eids := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	st := &state_v1.State{
		Version:      "1",
		ActiveEditor: eids[0],
	}

	for _, eid := range eids {
		st.Editors = append(st.Editors, &state_v1.EditorState{ID: eid})
	}

	return st
}

func mustMarshal(state *state_v1.State) []byte {
	if data, err := yaml.Marshal(state); err != nil {
		panic(err)
	} else {
		return data
	}
}

func mustUnmarshal(data []byte) *state_v1.State {
	var state state_v1.State
	if err := yaml.Unmarshal(data, &state); err != nil {
		panic(err)
	} else {
		return &state
	}
}
