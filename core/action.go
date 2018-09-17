package core

import (
	"syscall/js"
)

// Action converts func to js.Callback which invokes Update internally.
func Action(i interface{}) js.Callback {
	switch f := i.(type) {
	case func():
		return js.NewCallback(func(v []js.Value) { f(); Update() })
	case func(...js.Value):
		return js.NewCallback(func(v []js.Value) { f(v...); Update() })
	default:
		return js.NewCallback(func(v []js.Value) { println("invalid func type") })
	}
}

// EventAction converts func with EventCallbackFlg to js.Callback which invokes Update internally.
func EventAction(flags js.EventCallbackFlag, i interface{}) js.Callback {
	switch f := i.(type) {
	case func(js.Value):
		return js.NewEventCallback(flags, func(e js.Value) { f(e); Update() })
	default:
		return js.NewCallback(func(v []js.Value) { println("invalid func type") })
	}
}
