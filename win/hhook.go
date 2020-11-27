/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hhook
type HHOOK HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-callnexthookex
func (hHook HHOOK) CallNextHookEx(nCode int32, wp WPARAM, lp LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.CallNextHookEx.Addr(), 4,
		uintptr(hHook), uintptr(nCode), uintptr(wp), uintptr(lp), 0, 0)
	return uintptr(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
func (hHook HHOOK) UnhookWindowsHookEx() {
	ret, _, lerr := syscall.Syscall(proc.UnhookWindowsHookEx.Addr(), 1,
		uintptr(hHook), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "UnhookWindowsHookEx"))
	}
}
