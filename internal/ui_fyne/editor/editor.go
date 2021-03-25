package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"

	"github.com/themakers/plain/internal/ui_fyne/editor/internal/undo"
	"github.com/themakers/plain/internal/ui_fyne/internal/hotkey_manager"
)

var _ fyne.Disableable = (*Editor)(nil)
var _ fyne.Draggable = (*Editor)(nil)
var _ fyne.Focusable = (*Editor)(nil)
var _ fyne.Tappable = (*Editor)(nil)
var _ fyne.Widget = (*Editor)(nil)
var _ desktop.Mouseable = (*Editor)(nil)
var _ desktop.Keyable = (*Editor)(nil)
var _ mobile.Keyboardable = (*Editor)(nil)

type Editor struct {
	*widget.Entry

	OnChanged func(string)

	ghm *hotkey_manager.Manager

	us *undo.Stack
}

func NewEditor(hm *hotkey_manager.Manager, text string, cursorRow, cursorColumn int) *Editor {
	ed := &Editor{
		Entry: widget.NewMultiLineEntry(),
		ghm:   hm,
	}
	ed.ExtendBaseWidget(ed)

	upd := func(text string, cursorRow, cursorColumn int) {
		ed.Entry.SetText(text)
		ed.Entry.CursorRow = cursorRow
		ed.Entry.CursorColumn = cursorColumn
		ed.Refresh()
	}

	ed.us = undo.New(upd, text, cursorRow, cursorColumn)

	upd(text, cursorRow, cursorColumn)

	ed.Entry.OnChanged = ed.onChanged

	return ed
}

func (e *Editor) Undo() {
	e.us.Undo()
}

func (e *Editor) Redo() {
	e.us.Redo()
}

func (e *Editor) onChanged(text string) {
	e.us.Change(text, e.CursorRow, e.CursorColumn)

	if e.OnChanged != nil {
		e.OnChanged(text)
	}
}

func (e *Editor) KeyDown(key *fyne.KeyEvent) {
	e.ghm.KeyDown(key.Name)
	e.Entry.KeyDown(key)
}

func (e *Editor) KeyUp(key *fyne.KeyEvent) {
	e.ghm.KeyUp(key.Name)
	e.Entry.KeyUp(key)
}
