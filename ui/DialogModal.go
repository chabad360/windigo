/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"github.com/rodrigocfd/windigo/co"
)

// Main application dialog.
type DialogModal struct {
	*_DialogBase
}

// Constructor.
func NewDialogModal(opts *DialogModalOpts) *DialogModal {
	if opts.DialogId == 0 {
		panic("Dialog resource ID must be specified.")
	}

	me := DialogModal{
		_DialogBase: _NewDialogBase(opts.DialogId),
	}

	me.defaultMessageHandling()
	return &me
}

// Creates the modal window and disables the parent.
// Will block until the window is closed.
func (me *DialogModal) Show(parent Parent) int {
	hInst := parent.Hwnd().GetInstance()
	ret := me._DialogBase.dialogBoxParam(hInst, parent)
	return int(ret) // value passed to EndDialog()
}

// Adds the messages which have a default processing.
func (me *DialogModal) defaultMessageHandling() {
	me.On().WmClose(func() {
		me.Hwnd().EndDialog(uintptr(co.MBID_CANCEL))
	})
}

//------------------------------------------------------------------------------

// Options for NewDialogModal().
type DialogModalOpts struct {
	// Dialog resource ID to be loaded.
	DialogId int
}
