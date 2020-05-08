package ui

import (
	"fmt"
	c "gowinui/consts"
)

// Custom hash for WM_NOTIFY messages.
type nfyHash struct {
	IdFrom c.ID
	Code   int32
}

// Keeps all user message handlers.
type windowOn struct {
	msgs        map[c.WM]func(p wmBase) uintptr
	cmds        map[c.ID]func(p WmCommand)
	nfys        map[nfyHash]func(p wmBase) uintptr
	loopStarted bool
}

func makeWindowOn() windowOn {
	msgs := make(map[c.WM]func(p wmBase) uintptr)
	cmds := make(map[c.ID]func(p WmCommand))
	nfys := make(map[nfyHash]func(p wmBase) uintptr)

	return windowOn{
		msgs:        msgs,
		cmds:        cmds,
		nfys:        nfys,
		loopStarted: false,
	}
}

func (me *windowOn) addMsg(msg c.WM, userFunc func(p wmBase) uintptr) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after application loop started.", msg))
	}
	me.msgs[msg] = userFunc
}

func (me *windowOn) addNfy(idFrom c.ID, code int32,
	userFunc func(p wmBase) uintptr) {

	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add motify message %d/%d after application loop started.",
			idFrom, code))
	}
	me.nfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowOn) processMessage(p wmBase) (uintptr, bool) {
	switch p.Msg {
	case c.WM_COMMAND:
		paramCmd := makeWmCommand(p)
		if userFunc, hasCmd := me.cmds[paramCmd.ControlId]; hasCmd {
			userFunc(paramCmd)
			return 0, true
		}
	case c.WM_NOTIFY:
		paramNfy := makeWmNotify(p)
		hash := nfyHash{
			IdFrom: c.ID(paramNfy.NmHdr.IdFrom),
			Code:   int32(paramNfy.NmHdr.Code),
		}
		if userFunc, hasNfy := me.nfys[hash]; hasNfy {
			return userFunc(p), true
		}
	default:
		if userFunc, hasMsg := me.msgs[p.Msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}
