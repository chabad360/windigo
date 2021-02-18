/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"

	proc "github.com/rodrigocfd/windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}
