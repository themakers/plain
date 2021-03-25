package fonts

import (
	"fyne.io/fyne/v2"

	"github.com/themakers/plain/assets/fonts/monofur"
)

func FyneDefault(style fyne.TextStyle) fyne.Resource {
	return monofur.Fyne(style)
}
