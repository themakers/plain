package main

import (
	"context"
	"time"

	"github.com/themakers/plain/internal/recovery"
	"github.com/themakers/plain/internal/state/state_v1"
	"github.com/themakers/plain/internal/state_manager"
	"github.com/themakers/plain/internal/storage"
	"github.com/themakers/plain/internal/ui_gioui"
	"github.com/themakers/plain/lib/debouncer"
	"github.com/themakers/plain/lib/worker"
)

// TODO: Signals
// TODO: Hotkeys
// TODO: Help
// TODO: Tray icon
// TODO: Global Hotkey
// TODO: Quake mode
// TODO: Markdown
// TODO: Releases
// TODO: Install script for *nix
// TODO: Self-update for *nix
// TODO: Self-update for Win
// TODO: Check compatibility on Windows (OpenGL; Fall back to software renderer?)
// TODO: Update notifications (label)
// TODO: Make testable and write Tests?
// TODO: Watch for state file changes?
// TODO: Single instance mode + basic IPC comm

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer recovery.Guard("entry")

	var (
		stor = storage.New()
		wor  = worker.New()
		deb  = debouncer.New(2*time.Second, 15*time.Second)
		sm   = state_manager.New(stor.Load(), func(state *state_v1.State) {
			deb.Trigger(func() {
				stor.Save(state)
			})
		})
		saveState = func() { stor.Save(sm.State()) }
	)

	go func() {
		defer cancel()
		defer recovery.Guard("debouncer")
		wor.Work(ctx, saveState)
	}()

	defer wor.RunSync(saveState)

	ui_gioui.Run(ctx, sm)
}
