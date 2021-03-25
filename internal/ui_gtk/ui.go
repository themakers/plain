package ui_gtk

import (
	"context"
	"log"

	"github.com/gotk3/gotk3/gtk"

	"github.com/themakers/plain/internal/state_manager"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func Run(ctx context.Context, sm *state_manager.Manager) {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	perror(err)

	win.SetTitle("Plain")
    win.SetDefaultSize(800, 600)
    win.SetPosition(gtk.WIN_POS_CENTER)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	nb, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Unable to create notebook:", err)
	}

	nb.SetHExpand(true)
	nb.SetVExpand(true)

	entry, err := gtk.TextViewNew()
	perror(err)

	nbTab, err := gtk.LabelNew("Tab label")
	perror(err)

	nb.AppendPage(entry, nbTab)

	win.Add(nb)
	win.ShowAll()
	gtk.Main()
}
