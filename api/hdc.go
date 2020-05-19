/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"syscall"
	"unsafe"
	"wingows/api/proc"
	c "wingows/consts"
)

type HDC HANDLE

func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) []HMONITOR {
	hMons := []HMONITOR{}
	syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lp LPARAM) uintptr {
				hMons = append(hMons, hMon)
				return uintptr(1)
			}), 0, 0, 0)
	return hMons
}

func (hdc HDC) GetDeviceCaps(index c.GDC) int32 {
	ret, _, _ := syscall.Syscall(proc.GetDeviceCaps.Addr(), 2,
		uintptr(hdc), uintptr(index), 0)
	return int32(ret)
}

func (hdc HDC) GetTextExtentPoint32(lpString string, cLength uint32) *SIZE {
	sz := &SIZE{}
	ret, _, _ := syscall.Syscall6(proc.GetTextExtentPoint32.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpString))),
		uintptr(cLength), uintptr(unsafe.Pointer(sz)), 0, 0)
	if ret == 0 {
		panic("GetTextExtentPoint32 failed.")
	}
	return sz
}

func (hdc HDC) GetTextFace() string {
	buf := [32]uint16{} // LF_FACESIZE
	ret, _, _ := syscall.Syscall(proc.GetTextFace.Addr(), 3,
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic("GetTextFace failed.")
	}
	return syscall.UTF16ToString(buf[:])
}

func (hdc HDC) LineTo(x, y int32) {
	ret, _, _ := syscall.Syscall(proc.LineTo.Addr(), 3,
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic("LineTo failed.")
	}
}

func (hdc HDC) SelectObjectBitmap(b HBITMAP) HBITMAP {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0xFFFFFFFF { // HGDI_ERROR
		panic("SelectObject failed to HBITMAP.")
	}
	return HBITMAP(ret)
}

func (hdc HDC) SelectObjectBrush(b HBRUSH) HBRUSH {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0xFFFFFFFF { // HGDI_ERROR
		panic("SelectObject failed to HBRUSH.")
	}
	return HBRUSH(ret)
}

func (hdc HDC) SelectObjectFont(f HFONT) HFONT {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(f), 0)
	if ret == 0xFFFFFFFF { // HGDI_ERROR
		panic("SelectObject failed to HFONT.")
	}
	return HFONT(ret)
}

func (hdc HDC) SelectObjectPen(p HPEN) HPEN {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(p), 0)
	if ret == 0xFFFFFFFF { // HGDI_ERROR
		panic("SelectObject failed to HPEN.")
	}
	return HPEN(ret)
}

func (hdc HDC) SelectObjectRgn(r HRGN) HRGN {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(r), 0)
	if ret == 0xFFFFFFFF { // HGDI_ERROR
		panic("SelectObject failed to HRGN.")
	}
	return HRGN(ret)
}

func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, _ := syscall.Syscall(proc.SetBkColor.Addr(), 2,
		uintptr(hdc), uintptr(color), 0)
	if ret == 0xFFFFFFFF { // CLR_INVALID
		panic("SetBkColor failed.")
	}
	return COLORREF(ret)
}

func (hdc HDC) SetBkMode(mode c.BKMODE) c.BKMODE {
	ret, _, _ := syscall.Syscall(proc.SetBkMode.Addr(), 2,
		uintptr(hdc), uintptr(mode), 0)
	if ret == 0 {
		panic("SetBkMode failed.")
	}
	return c.BKMODE(ret)
}
