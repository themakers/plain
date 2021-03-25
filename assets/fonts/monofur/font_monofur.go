package monofur

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed monofur-regular.ttf
var MonofurRegular []byte

//go:embed monofur-bold.ttf
var MonofurBold []byte

//go:embed monofur-italic.ttf
var MonofurItalic []byte

func Fyne(style fyne.TextStyle) fyne.Resource {
	var (
		regular = func() fyne.Resource { return fyne.NewStaticResource("monofur-regular", MonofurRegular) }
		bold    = func() fyne.Resource { return fyne.NewStaticResource("monofur-bold", MonofurBold) }
		italic  = func() fyne.Resource { return fyne.NewStaticResource("monofur-italic", MonofurItalic) }
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
