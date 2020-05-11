package api

import (
	"gowinui/api/proc"
	c "gowinui/consts"
	"syscall"
	"unsafe"
)

type HMONITOR HANDLE

func MonitorFromPoint(pt POINT, dwFlags c.MONITOR) HMONITOR {
	ret, _, _ := syscall.Syscall(proc.MonitorFromPoint.Addr(), 3,
		uintptr(pt.X), uintptr(pt.Y), uintptr(dwFlags))
	return HMONITOR(ret)
}

// Available in Windows 8.1.
func (hmon HMONITOR) GetDpiForMonitor(dpiType c.MDT) (uint32, uint32) {
	dpiX, dpiY := uint32(0), uint32(0)
	ret, _, _ := syscall.Syscall6(proc.GetDpiForMonitor.Addr(), 4,
		uintptr(hmon), uintptr(dpiType),
		uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)),
		0, 0)
	if ret != 0 {
		panic("GetDpiForMonitor failed.")
	}
	return dpiX, dpiY
}
