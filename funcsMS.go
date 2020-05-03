package winffi

import (
	"syscall"
	"winffi/consts"
	"winffi/procs"
)

func (hwnd HWND) MessageBox(message, caption string, flags consts.MB) int32 {
	ret, _, _ := syscall.Syscall6(procs.MessageBox.Addr(), 4,
		uintptr(0), toUtf16ToUintptr(message), toUtf16ToUintptr(caption),
		uintptr(flags), 0, 0)
	return int32(ret)
}
