package show2b

import "github.com/gdamore/tcell/v2"

type KeyBindings map[string][]tcell.Key

var keyBindings = KeyBindings{
	"switch_focus": {tcell.KeyTab},
}

func NewKeyBinding() KeyBindings {
	return keyBindings
}

func (kb KeyBindings) SearchKey(k tcell.Key) string {
	for name, bind := range kb {
		for _, b := range bind {
			if b == k {
				return name
			}
		}
	}

	return ""
}
