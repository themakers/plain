package dialog

import (
	"fmt"

	"github.com/sqweek/dialog"
)

func Error(title, message string) {
	title = fmt.Sprintf("ERROR: %s", title)
	dialog.Message("%s\n\n%s", title, message).Title(title).Error()
}
