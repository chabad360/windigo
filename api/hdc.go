package api

import (
	"gowinui/api/proc"
	"syscall"
	"unsafe"
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

func (hdc HDC) GetDeviceCaps(index int32) int32 {
	ret, _, _ := syscall.Syscall(proc.GetDeviceCaps.Addr(), 2,
		uintptr(hdc), uintptr(index), 0)
	return int32(ret)
}
