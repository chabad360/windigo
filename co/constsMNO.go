/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type MDT uint32 // GetDpiForMonitor dpiType

const (
	MDT_EFFECTIVE_DPI MDT = 0
	MDT_ANGULAR_DPI   MDT = 1
	MDT_RAW_DPI       MDT = 2
	MDT_DEFAULT       MDT = MDT_EFFECTIVE_DPI
)

type MB uint32 // MessageBox uType

const (
	MB_ABORTRETRYIGNORE  MB = 0x00000002
	MB_CANCELTRYCONTINUE MB = 0x00000006
	MB_HELP              MB = 0x00004000
	MB_OK                MB = 0x00000000
	MB_OKCANCEL          MB = 0x00000001
	MB_RETRYCANCEL       MB = 0x00000005
	MB_YESNO             MB = 0x00000004
	MB_YESNOCANCEL       MB = 0x00000003

	MB_ICONEXCLAMATION MB = 0x00000030
	MB_ICONWARNING     MB = MB_ICONEXCLAMATION
	MB_ICONINFORMATION MB = 0x00000040
	MB_ICONASTERISK    MB = MB_ICONINFORMATION
	MB_ICONQUESTION    MB = 0x00000020
	MB_ICONSTOP        MB = MB_ICONERROR
	MB_ICONERROR       MB = 0x00000010
	MB_ICONHAND        MB = MB_ICONERROR

	MB_DEFBUTTON1 MB = 0x00000000
	MB_DEFBUTTON2 MB = 0x00000100
	MB_DEFBUTTON3 MB = 0x00000200
	MB_DEFBUTTON4 MB = 0x00000300

	MB_APPLMODAL   MB = 0x00000000
	MB_SYSTEMMODAL MB = 0x00001000
	MB_TASKMODAL   MB = 0x00002000

	MB_DEFAULT_DESKTOP_ONLY MB = 0x00020000
	MB_RIGHT                MB = 0x00080000
	MB_RTLREADING           MB = 0x00100000
	MB_SETFOREGROUND        MB = 0x00010000
	MB_TOPMOST              MB = 0x00040000
	MB_SERVICE_NOTIFICATION MB = 0x00200000
)

type MF uint32 // EnableMenuItem uEnable

const (
	MF_INSERT          MF = 0x00000000
	MF_CHANGE          MF = 0x00000080
	MF_APPEND          MF = 0x00000100
	MF_DELETE          MF = 0x00000200
	MF_REMOVE          MF = 0x00001000
	MF_BYCOMMAND       MF = 0x00000000
	MF_BYPOSITION      MF = 0x00000400
	MF_SEPARATOR       MF = 0x00000800
	MF_ENABLED         MF = 0x00000000
	MF_GRAYED          MF = 0x00000001
	MF_DISABLED        MF = 0x00000002
	MF_UNCHECKED       MF = 0x00000000
	MF_CHECKED         MF = 0x00000008
	MF_USECHECKBITMAPS MF = 0x00000200
	MF_STRING          MF = 0x00000000
	MF_BITMAP          MF = 0x00000004
	MF_OWNERDRAW       MF = 0x00000100
	MF_POPUP           MF = 0x00000010
	MF_MENUBARBREAK    MF = 0x00000020
	MF_MENUBREAK       MF = 0x00000040
	MF_UNHILITE        MF = 0x00000000
	MF_HILITE          MF = 0x00000080
	MF_DEFAULT         MF = 0x00001000
	MF_SYSMENU         MF = 0x00002000
	MF_HELP            MF = 0x00004000
	MF_RIGHTJUSTIFY    MF = 0x00004000
	MF_MOUSESELECT     MF = 0x00008000
)

type MFS uint32 // MENUITEMINFO fState

const (
	MFS_GRAYED    MFS = 0x00000003
	MFS_DISABLED  MFS = MFS_GRAYED
	MFS_CHECKED   MFS = MFS(MF_CHECKED)
	MFS_HILITE    MFS = MFS(MF_HILITE)
	MFS_ENABLED   MFS = MFS(MF_ENABLED)
	MFS_UNCHECKED MFS = MFS(MF_UNCHECKED)
	MFS_UNHILITE  MFS = MFS(MF_UNHILITE)
	MFS_DEFAULT   MFS = MFS(MF_DEFAULT)
)

type MFT uint32 // MENUITEMINFO fType

const (
	MFT_STRING       MFT = MFT(MF_STRING)
	MFT_BITMAP       MFT = MFT(MF_BITMAP)
	MFT_MENUBARBREAK MFT = MFT(MF_MENUBARBREAK)
	MFT_MENUBREAK    MFT = MFT(MF_MENUBREAK)
	MFT_OWNERDRAW    MFT = MFT(MF_OWNERDRAW)
	MFT_RADIOCHECK   MFT = 0x00000200
	MFT_SEPARATOR    MFT = MFT(MF_SEPARATOR)
	MFT_RIGHTORDER   MFT = 0x00002000
	MFT_RIGHTJUSTIFY MFT = MFT(MF_RIGHTJUSTIFY)
)

type MIIM uint32 // MENUITEMINFO fMask

const (
	MIIM_STATE      MIIM = 0x00000001
	MIIM_ID         MIIM = 0x00000002
	MIIM_SUBMENU    MIIM = 0x00000004
	MIIM_CHECKMARKS MIIM = 0x00000008
	MIIM_TYPE       MIIM = 0x00000010
	MIIM_DATA       MIIM = 0x00000020
	MIIM_STRING     MIIM = 0x00000040
	MIIM_BITMAP     MIIM = 0x00000080
	MIIM_FTYPE      MIIM = 0x00000100
)

type MIM uint32 // MENUINFO fMask

const (
	MIM_MAXHEIGHT       MIM = 0x00000001
	MIM_BACKGROUND      MIM = 0x00000002
	MIM_HELPID          MIM = 0x00000004
	MIM_MENUDATA        MIM = 0x00000008
	MIM_STYLE           MIM = 0x00000010
	MIM_APPLYTOSUBMENUS MIM = 0x80000000
)

type MK uint16 // WM_LBUTTONDOWN and similar

const (
	MK_LBUTTON  MK = 0x0001
	MK_RBUTTON  MK = 0x0002
	MK_SHIFT    MK = 0x0004
	MK_CONTROL  MK = 0x0008
	MK_MBUTTON  MK = 0x0010
	MK_XBUTTON1 MK = 0x0020
	MK_XBUTTON2 MK = 0x0040
)

type MNC uint32 // WM_MENUCHAR return

const (
	MNC_IGNORE  MNC = 0
	MNC_CLOSE   MNC = 1
	MNC_EXECUTE MNC = 2
	MNC_SELECT  MNC = 3
)

type MNS uint32 // MENUINFO dwStyle

const (
	MNS_NOCHECK     MNS = 0x80000000
	MNS_MODELESS    MNS = 0x40000000
	MNS_DRAGDROP    MNS = 0x20000000
	MNS_AUTODISMISS MNS = 0x10000000
	MNS_NOTIFYBYPOS MNS = 0x08000000
	MNS_CHECKORBMP  MNS = 0x04000000
)

type MOD uint16 // WM_HOTKEY

const (
	MOD_ALT                 MOD = 0x0001
	MOD_CONTROL             MOD = 0x0002
	MOD_SHIFT               MOD = 0x0004
	MOD_LEFT                MOD = 0x8000
	MOD_RIGHT               MOD = 0x4000
	MOD_ON_KEYUP            MOD = 0x0800
	MOD_IGNORE_ALL_MODIFIER MOD = 0x0400
)

type MONITOR uint32 // MonitorFromPoint dwFlags

const (
	MONITOR_DEFAULTTONULL    MONITOR = 0x00000000
	MONITOR_DEFAULTTOPRIMARY MONITOR = 0x00000001
	MONITOR_DEFAULTTONEAREST MONITOR = 0x00000002
)

type NM int32 // common control notification

const (
	nM_FIRST NM = 0

	NM_OUTOFMEMORY          NM = nM_FIRST - 1
	NM_CLICK                NM = nM_FIRST - 2
	NM_DBLCLK               NM = nM_FIRST - 3
	NM_RETURN               NM = nM_FIRST - 4
	NM_RCLICK               NM = nM_FIRST - 5
	NM_RDBLCLK              NM = nM_FIRST - 6
	NM_SETFOCUS             NM = nM_FIRST - 7
	NM_KILLFOCUS            NM = nM_FIRST - 8
	NM_CUSTOMDRAW           NM = nM_FIRST - 12
	NM_HOVER                NM = nM_FIRST - 13
	NM_NCHITTEST            NM = nM_FIRST - 14
	NM_KEYDOWN              NM = nM_FIRST - 15
	NM_RELEASEDCAPTURE      NM = nM_FIRST - 16
	NM_SETCURSOR            NM = nM_FIRST - 17
	NM_CHAR                 NM = nM_FIRST - 18
	NM_TOOLTIPSCREATED      NM = nM_FIRST - 19
	NM_LDOWN                NM = nM_FIRST - 20
	NM_RDOWN                NM = nM_FIRST - 21
	NM_THEMECHANGED         NM = nM_FIRST - 22
	NM_FONTCHANGED          NM = nM_FIRST - 23
	NM_CUSTOMTEXT           NM = nM_FIRST - 24
	NM_TVSTATEIMAGECHANGING NM = nM_FIRST - 24
)
