package state_manager

import (
    "github.com/themakers/plain/internal/state/state_v1"
    "log"
    "sync"
)

type Manager struct {
	state *state_v1.State

	lock sync.Mutex

	save func()
}

func New(init *state_v1.State, save func(state *state_v1.State)) *Manager {
	sm := &Manager{
		state: init,
	}

	sm.save = func() {
		save(sm.State())
	}

	return sm
}

func (sm *Manager) txn(fn func(state *state_v1.State)) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	defer sm.save()

	fn(sm.state)
}

func (sm *Manager) SetActiveEditor(editor string) {
    log.Println("EVENT: SetActiveEditor =>", editor)
	sm.txn(func(state *state_v1.State) {
		state.ActiveEditor = editor
	})
}

func (sm *Manager) SetText(editor, text string) {
    log.Println("EVENT: SetActiveEditor =>", editor, text)
	sm.txn(func(state *state_v1.State) {
		editor := state.EditorByID(editor)
		editor.Text = text
	})
}

func (sm *Manager) SetCursor(editor string, row, col int) {
    log.Println("EVENT: SetCursor =>", editor, row, col)
	sm.txn(func(state *state_v1.State) {
		editor := state.EditorByID(editor)

		editor.CursorRow = row
		editor.CursorColumn = col
	})
}

func (sm *Manager) State() *state_v1.State { return sm.state.Clone() }
