package theme

import (
    "image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"github.com/themakers/plain/assets/fonts"
)

var _ fyne.Theme = (*thm)(nil)

type thm struct {
	t fyne.Theme
}

func New(dark bool) fyne.Theme {
	t := &thm{}

	if dark {
		t.t = theme.DarkTheme()
	} else {
		t.t = theme.LightTheme()
	}

	return t
}

func (t *thm) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return t.t.Color(n, v)
}

func (t *thm) Font(style fyne.TextStyle) fyne.Resource {
	return fonts.FyneDefault(style)
}

func (t *thm) Icon(n fyne.ThemeIconName) fyne.Resource {
	return t.t.Icon(n)
}

func (t *thm) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 20
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInputBorder:
		return 2
	default:
		return t.t.Size(s)
	}
}
