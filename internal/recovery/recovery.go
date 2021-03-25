package recovery

import (
	"fmt"
	"github.com/themakers/plain/internal/dialog"
	"runtime/debug"
)

func Guard(thread string) {
	if rec := recover(); rec != nil {
		dialog.Error(fmt.Sprintf("%s %v", thread, rec), string(debug.Stack()))
		panic(rec)
	}
}
