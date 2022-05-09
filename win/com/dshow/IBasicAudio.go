package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/autom"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-ibasicaudio
type IBasicAudio interface {
	autom.IDispatch

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_balance
	GetBalance() int

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_volume
	GetVolume() int

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_balance
	PutBalance(balance int)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_volume
	PutVolume(volume int)
}

type _IBasicAudio struct{ autom.IDispatch }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IBasicAudio.Release().
//
// Example:
//
//		var gb dshow.IGraphBuilder // initialized somewhere
//
//		ba := dshow.NewIBasicAudio(
//			gb.QueryInterface(dshowco.IID_IBasicAudio),
//		)
//		defer ba.Release()
func NewIBasicAudio(base com.IUnknown) IBasicAudio {
	return &_IBasicAudio{IDispatch: autom.NewIDispatch(base)}
}

func (me *_IBasicAudio) GetBalance() int {
	var balance int32
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).GetBalance, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&balance)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(balance)
	} else {
		panic(hr)
	}
}

func (me *_IBasicAudio) GetVolume() int {
	var volume int32
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).GetVolume, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&volume)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(volume)
	} else {
		panic(hr)
	}
}

func (me *_IBasicAudio) PutBalance(balance int) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).PutBalance, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(balance)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IBasicAudio) PutVolume(volume int) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).PutVolume, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(volume)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
