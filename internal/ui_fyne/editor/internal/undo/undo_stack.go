package undo

import (
	"github.com/themakers/plain/lib/debouncer"
	"time"
)

type Stack struct {
	update func(text string, cursorRow, cursorColumn int)

	deb *debouncer.Debouncer

	stack []entry
	ptr   int
}

type entry struct {
	text                    string
	cursorRow, cursorColumn int
}

func New(update func(text string, cursorRow, cursorColumn int), initial string, cursorRow, cursorColumn int) *Stack {
	s := &Stack{
		update: update,
		deb:    debouncer.New(1*time.Second, 5*time.Second),
	}

	s.push(initial, cursorRow, cursorColumn)

	return s
}

func (us *Stack) push(text string, cursorRow, cursorColumn int) {
	if len(us.stack) == 0 {
		us.stack = append(us.stack, entry{
			text:         text,
			cursorRow:    cursorRow,
			cursorColumn: cursorColumn,
		})
	} else {
		us.stack = us.stack[:us.ptr+1]
		us.stack = append(us.stack, entry{
			text:         text,
			cursorRow:    cursorRow,
			cursorColumn: cursorColumn,
		})
		us.ptr++
	}
}

func (us *Stack) Change(text string, cursorRow, cursorColumn int) {
	if len(us.stack) > 0 && text == us.stack[us.ptr].text {
		return
	}

	us.deb.Trigger(func() {
		us.push(text, cursorRow, cursorColumn)
	})
}

func (us *Stack) Undo() {
	us.deb.Flush()
	if len(us.stack) > 0 && us.ptr > 0 {
		us.ptr--
		e := us.stack[us.ptr]
		us.update(e.text, e.cursorRow, e.cursorColumn)
	}
}

func (us *Stack) Redo() {
	us.deb.Flush()
	if len(us.stack) > 0 && us.ptr < len(us.stack)-1 {
		us.ptr++
		e := us.stack[us.ptr]
		us.update(e.text, e.cursorRow, e.cursorColumn)
	}
}
