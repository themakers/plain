package hotkey_manager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"log"
)

var modifiers = map[fyne.KeyName]bool{
	desktop.KeyShiftLeft:    true,
	desktop.KeyShiftRight:   true,
	desktop.KeyControlLeft:  true,
	desktop.KeyControlRight: true,
	desktop.KeyAltLeft:      true,
	desktop.KeyAltRight:     true,
	desktop.KeySuperLeft:    true,
	desktop.KeySuperRight:   true,
	desktop.KeyMenu:         true,
}

type Manager struct {
	hotkeys []hotkey

	modifiers map[fyne.KeyName]bool
	keys      map[fyne.KeyName]bool
}

type hotkey struct {
	Keys     []fyne.KeyName
	Callback func()
}

func (hk *hotkey) match(keys ...fyne.KeyName) bool {
	if len(keys) == len(hk.Keys) {
		nmatch := 0
		for _, k1 := range keys {
			for _, k2 := range hk.Keys {
				if k1 == k2 {
					nmatch++
				}
			}
		}
		return nmatch == len(hk.Keys)
	} else {
		return false
	}
}

func NewHotkeyManager() *Manager {
	return &Manager{
		modifiers: map[fyne.KeyName]bool{},
		keys:      map[fyne.KeyName]bool{},
	}
}

func (hm *Manager) Register(cb func(), keys ...fyne.KeyName) {
	hm.hotkeys = append(hm.hotkeys, hotkey{
		Keys:     keys,
		Callback: cb,
	})
}

func (hm *Manager) trigger() {
	var keys []fyne.KeyName

	for k := range hm.modifiers {
		keys = append(keys, k)
	}

	for k := range hm.keys {
		keys = append(keys, k)
	}

	for _, h := range hm.hotkeys {
		if h.match(keys...) {
			h.Callback()
			break
		}
	}
}

func (hm *Manager) KeyDown(key fyne.KeyName) {

	if modifiers[key] {
        defer log.Println("[DOWN MOD]", key, hm.modifiers, hm.keys)
		hm.modifiers[key] = true
	} else if len(hm.modifiers) > 0 {
        defer log.Println("[DOWN KEY]", key, hm.modifiers, hm.keys)
		hm.keys[key] = true
		hm.trigger()
	}

}

func (hm *Manager) KeyUp(key fyne.KeyName) {

	if modifiers[key] {
        defer log.Println("[UP MOD]", key, hm.modifiers, hm.keys)
		delete(hm.modifiers, key)
	} else {
        defer log.Println("[UP KEY]", key, hm.modifiers, hm.keys)
		delete(hm.keys, key)
	}
}
