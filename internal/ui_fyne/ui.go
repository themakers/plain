package ui_fyne

import (
	"context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	ftheme "fyne.io/fyne/v2/theme"

	"github.com/themakers/plain/internal/state/state_v1"
	"github.com/themakers/plain/internal/state_manager"
	"github.com/themakers/plain/internal/ui_fyne/editor"
	"github.com/themakers/plain/internal/ui_fyne/internal/hotkey_manager"
	"github.com/themakers/plain/internal/ui_fyne/internal/theme"
	"github.com/themakers/plain/internal/wisdom"
)

func Run(ctx context.Context, sm *state_manager.Manager) {
	app := app.New()
	win := app.NewWindow("plain")

	app.Settings().SetTheme(theme.New(true))

	go func() {
		<-ctx.Done()
		win.Close()
		// app.Quit()
	}()

	win.SetOnClosed(func() {
	})
	win.SetCloseIntercept(func() {
		win.Close()
	})

	hm := hotkey_manager.NewHotkeyManager()

	var (
		tabHost = container.NewAppTabs()
		editors = map[string]*editor.Editor{}
	)

	ec := len(sm.State().Editors)

	{ //> Hotkeys
		selectTab := func(i int) func() {
			return func() {
				tabHost.SelectTabIndex(i)
			}
		}

		selectTabRelative := func(o int) func() {
			return func() {
				n := tabHost.CurrentTabIndex()

				i := n + o
				if i < 0 {
					i = ec + i
				} else if i >= ec {
					i = 0 + (i - ec)
				}

				tabHost.SelectTabIndex(i)
			}
		}

		undo := func() {
			editors[tabHost.CurrentTab().Text].Undo()
		}

		redo := func() {
			editors[tabHost.CurrentTab().Text].Redo()
		}

		hm.Register(selectTabRelative(-1), desktop.KeyAltLeft, fyne.KeyLeft)
		hm.Register(selectTabRelative(-1), desktop.KeyAltRight, fyne.KeyLeft)

		hm.Register(selectTabRelative(1), desktop.KeyAltLeft, fyne.KeyRight)
		hm.Register(selectTabRelative(1), desktop.KeyAltRight, fyne.KeyRight)

		hm.Register(selectTab(0), desktop.KeyAltLeft, fyne.Key1)
		hm.Register(selectTab(0), desktop.KeyAltRight, fyne.Key1)
		hm.Register(selectTab(1), desktop.KeyAltLeft, fyne.Key2)
		hm.Register(selectTab(1), desktop.KeyAltRight, fyne.Key2)
		hm.Register(selectTab(2), desktop.KeyAltLeft, fyne.Key3)
		hm.Register(selectTab(2), desktop.KeyAltRight, fyne.Key3)
		hm.Register(selectTab(3), desktop.KeyAltLeft, fyne.Key4)
		hm.Register(selectTab(3), desktop.KeyAltRight, fyne.Key4)
		hm.Register(selectTab(4), desktop.KeyAltLeft, fyne.Key5)
		hm.Register(selectTab(4), desktop.KeyAltRight, fyne.Key5)
		hm.Register(selectTab(5), desktop.KeyAltLeft, fyne.Key6)
		hm.Register(selectTab(5), desktop.KeyAltRight, fyne.Key6)
		hm.Register(selectTab(6), desktop.KeyAltLeft, fyne.Key7)
		hm.Register(selectTab(6), desktop.KeyAltRight, fyne.Key7)
		hm.Register(selectTab(7), desktop.KeyAltLeft, fyne.Key8)
		hm.Register(selectTab(7), desktop.KeyAltRight, fyne.Key8)
		hm.Register(selectTab(8), desktop.KeyAltLeft, fyne.Key9)
		hm.Register(selectTab(8), desktop.KeyAltRight, fyne.Key9)
		hm.Register(selectTab(9), desktop.KeyAltLeft, fyne.Key0)
		hm.Register(selectTab(9), desktop.KeyAltRight, fyne.Key0)

		hm.Register(undo, desktop.KeyControlLeft, fyne.KeyZ)
		hm.Register(undo, desktop.KeyControlRight, fyne.KeyZ)

		hm.Register(redo, desktop.KeyControlLeft, desktop.KeyShiftLeft, fyne.KeyZ)
		hm.Register(redo, desktop.KeyControlRight, desktop.KeyShiftLeft, fyne.KeyZ)
		hm.Register(redo, desktop.KeyControlLeft, desktop.KeyShiftRight, fyne.KeyZ)
		hm.Register(redo, desktop.KeyControlRight, desktop.KeyShiftRight, fyne.KeyZ)
	}

	for en, es := range sm.State().Editors {
		(func(en int, es state_v1.EditorState) {
			ed := editor.NewEditor(hm, es.Text, es.CursorRow, es.CursorColumn)
			//ed.SetText(es.Text)

			editors[es.ID] = ed

			tab := container.NewTabItem(es.ID, ed)
			tabHost.Append(tab)

			ed.OnChanged = func(text string) {
				sm.SetText(es.ID, text)

				if text == "" {
					tab.Icon = nil
				} else {
					tab.Icon = ftheme.InfoIcon()
				}

				tabHost.Refresh()
			}

			ed.OnCursorChanged = func() {
				sm.SetCursor(es.ID, ed.CursorRow, ed.CursorColumn)
			}

			ed.SetPlaceHolder(wisdom.Touch())

			if es.Text == "" {
				tab.Icon = nil
			} else {
				tab.Icon = ftheme.InfoIcon()
			}

			tabHost.Refresh()
		})(en, *es)
	}

	//tabHost.SetTabLocation(container.TabLocationBottom)

	tabHost.OnChanged = func(tab *container.TabItem) {
		//editor := editors[tabID]
		//win.Canvas().Focus(editor)
		// editor.FocusGained()

		tabID := tab.Text
		sm.SetActiveEditor(tabID)

		es := sm.State().EditorByID(tabID)

		ed := tab.Content.(*editor.Editor)
		ed.CursorRow = es.CursorRow
		ed.CursorColumn = es.CursorColumn
		ed.Refresh()

		win.Canvas().Focus(ed)
	}

	//tabHost.SelectTabIndex(sm.State().ActiveEditorIndex())
	go func() { // FIXME: Hack
		time.Sleep(500 * time.Millisecond)
		tabHost.SelectTabIndex(sm.State().ActiveEditorIndex())
	}()

	win.SetContent(tabHost)

	win.SetMaster()
	win.Show()
	app.Run()
}
