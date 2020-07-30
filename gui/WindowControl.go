/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Custom user control.
// Allows message and notification handling.
type WindowControl struct {
	windowBase
	setup windowControlSetup // Parameters that will be used to create the window.
}

// Retrieves the command ID for this control.
func (me *WindowControl) Id() int32 {
	return me.hwnd.GetDlgCtrlID()
}

// Exposes parameters that will be used to create the child window control.
func (me *WindowControl) Setup() *windowControlSetup {
	if me.windowBase.Hwnd() != 0 {
		panic("Cannot change setup after the control was created.")
	}
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the child window control.
func (me *WindowControl) Create(
	parent Window, ctrlId, x, y int32, width, height uint32) {

	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	me.defaultMessageHandling()

	x, y, width, height = globalDpi.multiply(x, y, width, height)

	me.windowBase.createWindow("WindowControl", me.setup.ExStyle,
		me.setup.ClassName, "", me.setup.Style, x, y, width, height, parent,
		win.HMENU(ctrlId), hInst)
}

func (me *WindowControl) defaultMessageHandling() {
	me.windowBase.OnMsg().WmNcPaint(func(p WmNcPaint) {
		me.Hwnd().DefWindowProc(co.WM_NCPAINT, p.WParam, p.LParam) // make system draw the scrollbar for us

		if (me.Hwnd().GetExStyle()&co.WS_EX_CLIENTEDGE) == 0 || // has no border
			!win.IsThemeActive() ||
			!win.IsAppThemed() {
			// No themed borders to be painted.
			return
		}

		rc := me.Hwnd().GetWindowRect() // window outmost coordinates, including margins
		me.Hwnd().ScreenToClientRc(rc)
		rc.Left += 2 // manual fix, because it comes up anchored at -2,-2
		rc.Top += 2
		rc.Right += 2
		rc.Bottom += 2

		hdc := me.Hwnd().GetWindowDC()
		defer me.Hwnd().ReleaseDC(hdc)

		hTheme := me.Hwnd().OpenThemeData("LISTVIEW") // borrow style from listview
		if hTheme != 0 {
			defer hTheme.CloseThemeData()

			// Clipping region; will draw only within this rectangle.
			// Draw only the borders to avoid flickering.
			rc2 := win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Left + 2, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE(0), rc, &rc2) // draw themed left border

			rc2 = win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Right, Bottom: rc.Top + 2}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE(0), rc, &rc2) // draw themed top border

			rc2 = win.RECT{Left: rc.Right - 2, Top: rc.Top, Right: rc.Right, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE(0), rc, &rc2) // draw themed right border

			rc2 = win.RECT{Left: rc.Left, Top: rc.Bottom - 2, Right: rc.Right, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE(0), rc, &rc2) // draw themed bottom border
		}
	})
}

//------------------------------------------------------------------------------

type windowControlSetup struct {
	wasInit bool // default to false

	classNameBuf     []uint16
	ClassName        string      // Optional, defaults to a hash generated by WNDCLASSEX parameters. Passed to RegisterClassEx.
	ClassStyle       co.CS       // Window class style, passed to RegisterClassEx.
	HCursor          win.HCURSOR // Window cursor, passed to RegisterClassEx.
	HBrushBackground win.HBRUSH  // Window background brush, passed to RegisterClassEx.

	Style   co.WS    // Window style, passed to CreateWindowEx.
	ExStyle co.WS_EX // Window extended style, passed to CreateWindowEx. For a border, use WS_EX_CLIENTEDGE
}

func (me *windowControlSetup) initOnce() {
	if !me.wasInit {
		me.wasInit = true

		me.ClassStyle = co.CS_DBLCLKS

		me.Style = co.WS_CHILD | co.WS_VISIBLE | co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS
		me.ExStyle = co.WS_EX(0)
	}
}

func (me *windowControlSetup) genWndClassEx(
	hInst win.HINSTANCE) *win.WNDCLASSEX {

	wcx := win.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 { // user specified a cursor
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}

	if me.HBrushBackground != 0 { // user specified a background brush
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = win.CreateSysColorBrush(co.COLOR_WINDOW)
	}

	if me.ClassName == "" { // user left class name blank
		me.ClassName = wcx.Hash() // generate hash after all other fields are set
	}
	me.classNameBuf = win.StrToSlice(me.ClassName)
	wcx.LpszClassName = uintptr(unsafe.Pointer(&me.classNameBuf[0]))

	return &wcx
}
