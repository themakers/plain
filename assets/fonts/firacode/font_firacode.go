package firacode

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed furacode-regular.ttf
var FiraCodeRegular []byte

//go:embed furacode-bold.ttf
var FiraCodeBold []byte

//go:embed furacode-light.ttf
var FiraCodeItalic []byte

func Fyne(style fyne.TextStyle) fyne.Resource {
	var (
		regular = func() fyne.Resource { return fyne.NewStaticResource("firacode-regular", FiraCodeRegular) }
		bold    = func() fyne.Resource { return fyne.NewStaticResource("firacode-bold", FiraCodeBold) }
		italic  = func() fyne.Resource { return fyne.NewStaticResource("firacode-italic", FiraCodeItalic) }
	)

	if style.Monospace {
		return regular()
	}
	if style.Bold {
		if style.Italic {
			return italic()
		}
		return bold()
	}
	if style.Italic {
		return italic()
	}
	return regular()
}
