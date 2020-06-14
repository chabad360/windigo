/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type SC uint32 // WM_SYSCOMMAND

const (
	SC_SIZE         SC = 0xF000
	SC_MOVE         SC = 0xF010
	SC_MINIMIZE     SC = 0xF020
	SC_MAXIMIZE     SC = 0xF030
	SC_NEXTWINDOW   SC = 0xF040
	SC_PREVWINDOW   SC = 0xF050
	SC_CLOSE        SC = 0xF060
	SC_VSCROLL      SC = 0xF070
	SC_HSCROLL      SC = 0xF080
	SC_MOUSEMENU    SC = 0xF090
	SC_KEYMENU      SC = 0xF100
	SC_ARRANGE      SC = 0xF110
	SC_RESTORE      SC = 0xF120
	SC_TASKLIST     SC = 0xF130
	SC_SCREENSAVE   SC = 0xF140
	SC_HOTKEY       SC = 0xF150
	SC_DEFAULT      SC = 0xF160
	SC_MONITORPOWER SC = 0xF170
	SC_CONTEXTHELP  SC = 0xF180
	SC_SEPARATOR    SC = 0xF00F
)

type SEC uint32 // CreateFileMapping flProtect

const (
	SEC_NONE                   SEC = 0
	SEC_PARTITION_OWNER_HANDLE SEC = 0x00040000
	SEC_64K_PAGES              SEC = 0x00080000
	SEC_FILE                   SEC = 0x00800000
	SEC_IMAGE                  SEC = 0x01000000
	SEC_PROTECTED_IMAGE        SEC = 0x02000000
	SEC_RESERVE                SEC = 0x04000000
	SEC_COMMIT                 SEC = 0x08000000
	SEC_NOCACHE                SEC = 0x10000000
	SEC_WRITECOMBINE           SEC = 0x40000000
	SEC_LARGE_PAGES            SEC = 0x80000000
	SEC_IMAGE_NO_EXECUTE       SEC = SEC_IMAGE | SEC_NOCACHE
)

type SECTION uint32 // composes FILE_MAP

const (
	SECTION_QUERY                SECTION = 0x0001
	SECTION_MAP_WRITE            SECTION = 0x0002
	SECTION_MAP_READ             SECTION = 0x0004
	SECTION_MAP_EXECUTE          SECTION = 0x0008
	SECTION_EXTEND_SIZE          SECTION = 0x0010
	SECTION_MAP_EXECUTE_EXPLICIT SECTION = 0x0020
	SECTION_ALL_ACCESS           SECTION = SECTION(STANDARD_RIGHTS_REQUIRED) | SECTION_QUERY | SECTION_MAP_WRITE | SECTION_MAP_READ | SECTION_MAP_EXECUTE | SECTION_EXTEND_SIZE
)

type SECURITY uint32 // CreateFile dwFlagsAndAttributes

const (
	SECURITY_NONE             SECURITY = 0
	SECURITY_ANONYMOUS        SECURITY = 0 << 16
	SECURITY_IDENTIFICATION   SECURITY = 1 << 16
	SECURITY_IMPERSONATION    SECURITY = 2 << 16
	SECURITY_DELEGATION       SECURITY = 3 << 16
	SECURITY_CONTEXT_TRACKING SECURITY = 0x00040000
	SECURITY_EFFECTIVE_ONLY   SECURITY = 0x00080000
	SECURITY_SQOS_PRESENT     SECURITY = 0x00100000
	SECURITY_VALID_SQOS_FLAGS SECURITY = 0x001F0000
)

type SIZE int32 // WM_SIZE request

const (
	SIZE_RESTORED  SIZE = 0
	SIZE_MINIMIZED SIZE = 1
	SIZE_MAXIMIZED SIZE = 2
	SIZE_MAXSHOW   SIZE = 3
	SIZE_MAXHIDE   SIZE = 4
)

type SM int32 // GetSystemMetrics nIndex

const (
	SM_CXSCREEN                    SM = 0
	SM_CYSCREEN                    SM = 1
	SM_CXVSCROLL                   SM = 2
	SM_CYHSCROLL                   SM = 3
	SM_CYCAPTION                   SM = 4
	SM_CXBORDER                    SM = 5
	SM_CYBORDER                    SM = 6
	SM_CXDLGFRAME                  SM = 7
	SM_CYDLGFRAME                  SM = 8
	SM_CYVTHUMB                    SM = 9
	SM_CXHTHUMB                    SM = 10
	SM_CXICON                      SM = 11
	SM_CYICON                      SM = 12
	SM_CXCURSOR                    SM = 13
	SM_CYCURSOR                    SM = 14
	SM_CYMENU                      SM = 15
	SM_CXFULLSCREEN                SM = 16
	SM_CYFULLSCREEN                SM = 17
	SM_CYKANJIWINDOW               SM = 18
	SM_MOUSEPRESENT                SM = 19
	SM_CYVSCROLL                   SM = 20
	SM_CXHSCROLL                   SM = 21
	SM_DEBUG                       SM = 22
	SM_SWAPBUTTON                  SM = 23
	SM_RESERVED1                   SM = 24
	SM_RESERVED2                   SM = 25
	SM_RESERVED3                   SM = 26
	SM_RESERVED4                   SM = 27
	SM_CXMIN                       SM = 28
	SM_CYMIN                       SM = 29
	SM_CXSIZE                      SM = 30
	SM_CYSIZE                      SM = 31
	SM_CXFRAME                     SM = 32
	SM_CYFRAME                     SM = 33
	SM_CXMINTRACK                  SM = 34
	SM_CYMINTRACK                  SM = 35
	SM_CXDOUBLECLK                 SM = 36
	SM_CYDOUBLECLK                 SM = 37
	SM_CXICONSPACING               SM = 38
	SM_CYICONSPACING               SM = 39
	SM_MENUDROPALIGNMENT           SM = 40
	SM_PENWINDOWS                  SM = 41
	SM_DBCSENABLED                 SM = 42
	SM_CMOUSEBUTTONS               SM = 43
	SM_CXFIXEDFRAME                SM = SM_CXDLGFRAME
	SM_CYFIXEDFRAME                SM = SM_CYDLGFRAME
	SM_CXSIZEFRAME                 SM = SM_CXFRAME
	SM_CYSIZEFRAME                 SM = SM_CYFRAME
	SM_SECURE                      SM = 44
	SM_CXEDGE                      SM = 45
	SM_CYEDGE                      SM = 46
	SM_CXMINSPACING                SM = 47
	SM_CYMINSPACING                SM = 48
	SM_CXSMICON                    SM = 49
	SM_CYSMICON                    SM = 50
	SM_CYSMCAPTION                 SM = 51
	SM_CXSMSIZE                    SM = 52
	SM_CYSMSIZE                    SM = 53
	SM_CXMENUSIZE                  SM = 54
	SM_CYMENUSIZE                  SM = 55
	SM_ARRANGE                     SM = 56
	SM_CXMINIMIZED                 SM = 57
	SM_CYMINIMIZED                 SM = 58
	SM_CXMAXTRACK                  SM = 59
	SM_CYMAXTRACK                  SM = 60
	SM_CXMAXIMIZED                 SM = 61
	SM_CYMAXIMIZED                 SM = 62
	SM_NETWORK                     SM = 63
	SM_CLEANBOOT                   SM = 67
	SM_CXDRAG                      SM = 68
	SM_CYDRAG                      SM = 69
	SM_SHOWSOUNDS                  SM = 70
	SM_CXMENUCHECK                 SM = 71
	SM_CYMENUCHECK                 SM = 72
	SM_SLOWMACHINE                 SM = 73
	SM_MIDEASTENABLED              SM = 74
	SM_MOUSEWHEELPRESENT           SM = 75
	SM_XVIRTUALSCREEN              SM = 76
	SM_YVIRTUALSCREEN              SM = 77
	SM_CXVIRTUALSCREEN             SM = 78
	SM_CYVIRTUALSCREEN             SM = 79
	SM_CMONITORS                   SM = 80
	SM_SAMEDISPLAYFORMAT           SM = 81
	SM_IMMENABLED                  SM = 82
	SM_CXFOCUSBORDER               SM = 83
	SM_CYFOCUSBORDER               SM = 84
	SM_TABLETPC                    SM = 86
	SM_MEDIACENTER                 SM = 87
	SM_STARTER                     SM = 88
	SM_SERVERR2                    SM = 89
	SM_MOUSEHORIZONTALWHEELPRESENT SM = 91
	SM_CXPADDEDBORDER              SM = 92
	SM_DIGITIZER                   SM = 94
	SM_MAXIMUMTOUCHES              SM = 95
	SM_CMETRICS                    SM = 97
	SM_REMOTESESSION               SM = 0x1000
	SM_SHUTTINGDOWN                SM = 0x2000
	SM_REMOTECONTROL               SM = 0x2001
	SM_CARETBLINKINGENABLED        SM = 0x2002
	SM_CONVERTIBLESLATEMODE        SM = 0x2003
	SM_SYSTEMDOCKED                SM = 0x2004
)

type SPI uint32 // SystemParametersInfo uiAction

const (
	SPI_GETBEEP                     SPI = 0x0001
	SPI_SETBEEP                     SPI = 0x0002
	SPI_GETMOUSE                    SPI = 0x0003
	SPI_SETMOUSE                    SPI = 0x0004
	SPI_GETBORDER                   SPI = 0x0005
	SPI_SETBORDER                   SPI = 0x0006
	SPI_GETKEYBOARDSPEED            SPI = 0x000A
	SPI_SETKEYBOARDSPEED            SPI = 0x000B
	SPI_LANGDRIVER                  SPI = 0x000C
	SPI_ICONHORIZONTALSPACING       SPI = 0x000D
	SPI_GETSCREENSAVETIMEOUT        SPI = 0x000E
	SPI_SETSCREENSAVETIMEOUT        SPI = 0x000F
	SPI_GETSCREENSAVEACTIVE         SPI = 0x0010
	SPI_SETSCREENSAVEACTIVE         SPI = 0x0011
	SPI_GETGRIDGRANULARITY          SPI = 0x0012
	SPI_SETGRIDGRANULARITY          SPI = 0x0013
	SPI_SETDESKWALLPAPER            SPI = 0x0014
	SPI_SETDESKPATTERN              SPI = 0x0015
	SPI_GETKEYBOARDDELAY            SPI = 0x0016
	SPI_SETKEYBOARDDELAY            SPI = 0x0017
	SPI_ICONVERTICALSPACING         SPI = 0x0018
	SPI_GETICONTITLEWRAP            SPI = 0x0019
	SPI_SETICONTITLEWRAP            SPI = 0x001A
	SPI_GETMENUDROPALIGNMENT        SPI = 0x001B
	SPI_SETMENUDROPALIGNMENT        SPI = 0x001C
	SPI_SETDOUBLECLKWIDTH           SPI = 0x001D
	SPI_SETDOUBLECLKHEIGHT          SPI = 0x001E
	SPI_GETICONTITLELOGFONT         SPI = 0x001F
	SPI_SETDOUBLECLICKTIME          SPI = 0x0020
	SPI_SETMOUSEBUTTONSWAP          SPI = 0x0021
	SPI_SETICONTITLELOGFONT         SPI = 0x0022
	SPI_GETFASTTASKSWITCH           SPI = 0x0023
	SPI_SETFASTTASKSWITCH           SPI = 0x0024
	SPI_SETDRAGFULLWINDOWS          SPI = 0x0025
	SPI_GETDRAGFULLWINDOWS          SPI = 0x0026
	SPI_GETNONCLIENTMETRICS         SPI = 0x0029
	SPI_SETNONCLIENTMETRICS         SPI = 0x002A
	SPI_GETMINIMIZEDMETRICS         SPI = 0x002B
	SPI_SETMINIMIZEDMETRICS         SPI = 0x002C
	SPI_GETICONMETRICS              SPI = 0x002D
	SPI_SETICONMETRICS              SPI = 0x002E
	SPI_SETWORKAREA                 SPI = 0x002F
	SPI_GETWORKAREA                 SPI = 0x0030
	SPI_SETPENWINDOWS               SPI = 0x0031
	SPI_GETHIGHCONTRAST             SPI = 0x0042
	SPI_SETHIGHCONTRAST             SPI = 0x0043
	SPI_GETKEYBOARDPREF             SPI = 0x0044
	SPI_SETKEYBOARDPREF             SPI = 0x0045
	SPI_GETSCREENREADER             SPI = 0x0046
	SPI_SETSCREENREADER             SPI = 0x0047
	SPI_GETANIMATION                SPI = 0x0048
	SPI_SETANIMATION                SPI = 0x0049
	SPI_GETFONTSMOOTHING            SPI = 0x004A
	SPI_SETFONTSMOOTHING            SPI = 0x004B
	SPI_SETDRAGWIDTH                SPI = 0x004C
	SPI_SETDRAGHEIGHT               SPI = 0x004D
	SPI_SETHANDHELD                 SPI = 0x004E
	SPI_GETLOWPOWERTIMEOUT          SPI = 0x004F
	SPI_GETPOWEROFFTIMEOUT          SPI = 0x0050
	SPI_SETLOWPOWERTIMEOUT          SPI = 0x0051
	SPI_SETPOWEROFFTIMEOUT          SPI = 0x0052
	SPI_GETLOWPOWERACTIVE           SPI = 0x0053
	SPI_GETPOWEROFFACTIVE           SPI = 0x0054
	SPI_SETLOWPOWERACTIVE           SPI = 0x0055
	SPI_SETPOWEROFFACTIVE           SPI = 0x0056
	SPI_SETCURSORS                  SPI = 0x0057
	SPI_SETICONS                    SPI = 0x0058
	SPI_GETDEFAULTINPUTLANG         SPI = 0x0059
	SPI_SETDEFAULTINPUTLANG         SPI = 0x005A
	SPI_SETLANGTOGGLE               SPI = 0x005B
	SPI_GETWINDOWSEXTENSION         SPI = 0x005C
	SPI_SETMOUSETRAILS              SPI = 0x005D
	SPI_GETMOUSETRAILS              SPI = 0x005E
	SPI_SETSCREENSAVERRUNNING       SPI = 0x0061
	SPI_SCREENSAVERRUNNING          SPI = SPI_SETSCREENSAVERRUNNING
	SPI_GETFILTERKEYS               SPI = 0x0032
	SPI_SETFILTERKEYS               SPI = 0x0033
	SPI_GETTOGGLEKEYS               SPI = 0x0034
	SPI_SETTOGGLEKEYS               SPI = 0x0035
	SPI_GETMOUSEKEYS                SPI = 0x0036
	SPI_SETMOUSEKEYS                SPI = 0x0037
	SPI_GETSHOWSOUNDS               SPI = 0x0038
	SPI_SETSHOWSOUNDS               SPI = 0x0039
	SPI_GETSTICKYKEYS               SPI = 0x003A
	SPI_SETSTICKYKEYS               SPI = 0x003B
	SPI_GETACCESSTIMEOUT            SPI = 0x003C
	SPI_SETACCESSTIMEOUT            SPI = 0x003D
	SPI_GETSERIALKEYS               SPI = 0x003E
	SPI_SETSERIALKEYS               SPI = 0x003F
	SPI_GETSOUNDSENTRY              SPI = 0x0040
	SPI_SETSOUNDSENTRY              SPI = 0x0041
	SPI_GETSNAPTODEFBUTTON          SPI = 0x005F
	SPI_SETSNAPTODEFBUTTON          SPI = 0x0060
	SPI_GETMOUSEHOVERWIDTH          SPI = 0x0062
	SPI_SETMOUSEHOVERWIDTH          SPI = 0x0063
	SPI_GETMOUSEHOVERHEIGHT         SPI = 0x0064
	SPI_SETMOUSEHOVERHEIGHT         SPI = 0x0065
	SPI_GETMOUSEHOVERTIME           SPI = 0x0066
	SPI_SETMOUSEHOVERTIME           SPI = 0x0067
	SPI_GETWHEELSCROLLLINES         SPI = 0x0068
	SPI_SETWHEELSCROLLLINES         SPI = 0x0069
	SPI_GETMENUSHOWDELAY            SPI = 0x006A
	SPI_SETMENUSHOWDELAY            SPI = 0x006B
	SPI_GETWHEELSCROLLCHARS         SPI = 0x006C
	SPI_SETWHEELSCROLLCHARS         SPI = 0x006D
	SPI_GETSHOWIMEUI                SPI = 0x006E
	SPI_SETSHOWIMEUI                SPI = 0x006F
	SPI_GETMOUSESPEED               SPI = 0x0070
	SPI_SETMOUSESPEED               SPI = 0x0071
	SPI_GETSCREENSAVERRUNNING       SPI = 0x0072
	SPI_GETDESKWALLPAPER            SPI = 0x0073
	SPI_GETAUDIODESCRIPTION         SPI = 0x0074
	SPI_SETAUDIODESCRIPTION         SPI = 0x0075
	SPI_GETSCREENSAVESECURE         SPI = 0x0076
	SPI_SETSCREENSAVESECURE         SPI = 0x0077
	SPI_GETHUNGAPPTIMEOUT           SPI = 0x0078
	SPI_SETHUNGAPPTIMEOUT           SPI = 0x0079
	SPI_GETWAITTOKILLTIMEOUT        SPI = 0x007A
	SPI_SETWAITTOKILLTIMEOUT        SPI = 0x007B
	SPI_GETWAITTOKILLSERVICETIMEOUT SPI = 0x007C
	SPI_SETWAITTOKILLSERVICETIMEOUT SPI = 0x007D
	SPI_GETMOUSEDOCKTHRESHOLD       SPI = 0x007E
	SPI_SETMOUSEDOCKTHRESHOLD       SPI = 0x007F
	SPI_GETPENDOCKTHRESHOLD         SPI = 0x0080
	SPI_SETPENDOCKTHRESHOLD         SPI = 0x0081
	SPI_GETWINARRANGING             SPI = 0x0082
	SPI_SETWINARRANGING             SPI = 0x0083
	SPI_GETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0084
	SPI_SETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0085
	SPI_GETPENDRAGOUTTHRESHOLD      SPI = 0x0086
	SPI_SETPENDRAGOUTTHRESHOLD      SPI = 0x0087
	SPI_GETMOUSESIDEMOVETHRESHOLD   SPI = 0x0088
	SPI_SETMOUSESIDEMOVETHRESHOLD   SPI = 0x0089
	SPI_GETPENSIDEMOVETHRESHOLD     SPI = 0x008A
	SPI_SETPENSIDEMOVETHRESHOLD     SPI = 0x008B
	SPI_GETDRAGFROMMAXIMIZE         SPI = 0x008C
	SPI_SETDRAGFROMMAXIMIZE         SPI = 0x008D
	SPI_GETSNAPSIZING               SPI = 0x008E
	SPI_SETSNAPSIZING               SPI = 0x008F
	SPI_GETDOCKMOVING               SPI = 0x0090
	SPI_SETDOCKMOVING               SPI = 0x0091
)

type SS WS // static control style

const (
	SS_LEFT            SS = 0x00000000
	SS_CENTER          SS = 0x00000001
	SS_RIGHT           SS = 0x00000002
	SS_ICON            SS = 0x00000003
	SS_BLACKRECT       SS = 0x00000004
	SS_GRAYRECT        SS = 0x00000005
	SS_WHITERECT       SS = 0x00000006
	SS_BLACKFRAME      SS = 0x00000007
	SS_GRAYFRAME       SS = 0x00000008
	SS_WHITEFRAME      SS = 0x00000009
	SS_USERITEM        SS = 0x0000000A
	SS_SIMPLE          SS = 0x0000000B
	SS_LEFTNOWORDWRAP  SS = 0x0000000C
	SS_OWNERDRAW       SS = 0x0000000D
	SS_BITMAP          SS = 0x0000000E
	SS_ENHMETAFILE     SS = 0x0000000F
	SS_ETCHEDHORZ      SS = 0x00000010
	SS_ETCHEDVERT      SS = 0x00000011
	SS_ETCHEDFRAME     SS = 0x00000012
	SS_TYPEMASK        SS = 0x0000001F
	SS_REALSIZECONTROL SS = 0x00000040
	SS_NOPREFIX        SS = 0x00000080
	SS_NOTIFY          SS = 0x00000100
	SS_CENTERIMAGE     SS = 0x00000200
	SS_RIGHTJUST       SS = 0x00000400
	SS_REALSIZEIMAGE   SS = 0x00000800
	SS_SUNKEN          SS = 0x00001000
	SS_EDITCONTROL     SS = 0x00002000
	SS_ENDELLIPSIS     SS = 0x00004000
	SS_PATHELLIPSIS    SS = 0x00008000
	SS_WORDELLIPSIS    SS = 0x0000C000
	SS_ELLIPSISMASK    SS = 0x0000C000
)

type STANDARD_RIGHTS uint32 // RegOpenKeyEx samDesired

const (
	STANDARD_RIGHTS_REQUIRED STANDARD_RIGHTS = 0x000F0000
	STANDARD_RIGHTS_READ     STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_WRITE    STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_EXECUTE  STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_ALL      STANDARD_RIGHTS = 0x001F0000
)

type SW int32 // ShowWindow nCmdShow

const (
	SW_HIDE            SW = 0
	SW_SHOWNORMAL      SW = 1
	SW_NORMAL          SW = 1
	SW_SHOWMINIMIZED   SW = 2
	SW_SHOWMAXIMIZED   SW = 3
	SW_MAXIMIZE        SW = 3
	SW_SHOWNOACTIVATE  SW = 4
	SW_SHOW            SW = 5
	SW_MINIMIZE        SW = 6
	SW_SHOWMINNOACTIVE SW = 7
	SW_SHOWNA          SW = 8
	SW_RESTORE         SW = 9
	SW_SHOWDEFAULT     SW = 10
	SW_FORCEMINIMIZE   SW = 11
	SW_MAX             SW = 11
)

type SWP uint32 // SetWindowPos, DeferWindowPos uFlags

const (
	SWP_NOSIZE         SWP = 0x0001
	SWP_NOMOVE         SWP = 0x0002
	SWP_NOZORDER       SWP = 0x0004
	SWP_NOREDRAW       SWP = 0x0008
	SWP_NOACTIVATE     SWP = 0x0010
	SWP_FRAMECHANGED   SWP = 0x0020
	SWP_SHOWWINDOW     SWP = 0x0040
	SWP_HIDEWINDOW     SWP = 0x0080
	SWP_NOCOPYBITS     SWP = 0x0100
	SWP_NOOWNERZORDER  SWP = 0x0200
	SWP_NOSENDCHANGING SWP = 0x0400
	SWP_DRAWFRAME      SWP = SWP_FRAMECHANGED
	SWP_NOREPOSITION   SWP = SWP_NOOWNERZORDER
	SWP_DEFERERASE     SWP = 0x2000
	SWP_ASYNCWINDOWPOS SWP = 0x4000
)

type SWP_HWND int32 // SetWindowPos hwndInsertAfter

const (
	SWP_HWND_BOTTOM    SWP_HWND = 1
	SWP_HWND_NOTOPMOST SWP_HWND = -2
	SWP_HWND_TOP       SWP_HWND = 0
	SWP_HWND_TOPMOST   SWP_HWND = -1
)

type VER uint32 // VerifyVersionInfo dwTypeMask

const (
	VER_BUILDNUMBER      VER = 0x0000004
	VER_MAJORVERSION     VER = 0x0000002
	VER_MINORVERSION     VER = 0x0000001
	VER_PLATFORMID       VER = 0x0000008
	VER_PRODUCT_TYPE     VER = 0x0000080
	VER_SERVICEPACKMAJOR VER = 0x0000020
	VER_SERVICEPACKMINOR VER = 0x0000010
	VER_SUITENAME        VER = 0x0000040
)

type VER_COND uint8 // VerifyVersionInfo dwlConditionMask

const (
	VER_COND_EQUAL         VER_COND = 1
	VER_COND_GREATER       VER_COND = 2
	VER_COND_GREATER_EQUAL VER_COND = 3
	VER_COND_LESS          VER_COND = 4
	VER_COND_LESS_EQUAL    VER_COND = 5

	VER_COND_AND VER_COND = 6
	VER_COND_OR  VER_COND = 7
)

type VK uint16 // virtual key code

const (
	VK_LBUTTON             VK = 0x01
	VK_RBUTTON             VK = 0x02
	VK_CANCEL              VK = 0x03
	VK_MBUTTON             VK = 0x04
	VK_XBUTTON1            VK = 0x05
	VK_XBUTTON2            VK = 0x06
	VK_BACK                VK = 0x08
	VK_TAB                 VK = 0x09
	VK_CLEAR               VK = 0x0C
	VK_RETURN              VK = 0x0D
	VK_SHIFT               VK = 0x10
	VK_CONTROL             VK = 0x11
	VK_MENU                VK = 0x12
	VK_PAUSE               VK = 0x13
	VK_CAPITAL             VK = 0x14
	VK_KANA                VK = 0x15
	VK_HANGEUL             VK = 0x15
	VK_HANGUL              VK = 0x15
	VK_JUNJA               VK = 0x17
	VK_FINAL               VK = 0x18
	VK_HANJA               VK = 0x19
	VK_KANJI               VK = 0x19
	VK_ESCAPE              VK = 0x1B
	VK_CONVERT             VK = 0x1C
	VK_NONCONVERT          VK = 0x1D
	VK_ACCEPT              VK = 0x1E
	VK_MODECHANGE          VK = 0x1F
	VK_SPACE               VK = 0x20
	VK_PRIOR               VK = 0x21
	VK_NEXT                VK = 0x22
	VK_END                 VK = 0x23
	VK_HOME                VK = 0x24
	VK_LEFT                VK = 0x25
	VK_UP                  VK = 0x26
	VK_RIGHT               VK = 0x27
	VK_DOWN                VK = 0x28
	VK_SELECT              VK = 0x29
	VK_PRINT               VK = 0x2A
	VK_EXECUTE             VK = 0x2B
	VK_SNAPSHOT            VK = 0x2C
	VK_INSERT              VK = 0x2D
	VK_DELETE              VK = 0x2E
	VK_HELP                VK = 0x2F
	VK_LWIN                VK = 0x5B
	VK_RWIN                VK = 0x5C
	VK_APPS                VK = 0x5D
	VK_SLEEP               VK = 0x5F
	VK_NUMPAD0             VK = 0x60
	VK_NUMPAD1             VK = 0x61
	VK_NUMPAD2             VK = 0x62
	VK_NUMPAD3             VK = 0x63
	VK_NUMPAD4             VK = 0x64
	VK_NUMPAD5             VK = 0x65
	VK_NUMPAD6             VK = 0x66
	VK_NUMPAD7             VK = 0x67
	VK_NUMPAD8             VK = 0x68
	VK_NUMPAD9             VK = 0x69
	VK_MULTIPLY            VK = 0x6A
	VK_ADD                 VK = 0x6B
	VK_SEPARATOR           VK = 0x6C
	VK_SUBTRACT            VK = 0x6D
	VK_DECIMAL             VK = 0x6E
	VK_DIVIDE              VK = 0x6F
	VK_F1                  VK = 0x70
	VK_F2                  VK = 0x71
	VK_F3                  VK = 0x72
	VK_F4                  VK = 0x73
	VK_F5                  VK = 0x74
	VK_F6                  VK = 0x75
	VK_F7                  VK = 0x76
	VK_F8                  VK = 0x77
	VK_F9                  VK = 0x78
	VK_F10                 VK = 0x79
	VK_F11                 VK = 0x7A
	VK_F12                 VK = 0x7B
	VK_F13                 VK = 0x7C
	VK_F14                 VK = 0x7D
	VK_F15                 VK = 0x7E
	VK_F16                 VK = 0x7F
	VK_F17                 VK = 0x80
	VK_F18                 VK = 0x81
	VK_F19                 VK = 0x82
	VK_F20                 VK = 0x83
	VK_F21                 VK = 0x84
	VK_F22                 VK = 0x85
	VK_F23                 VK = 0x86
	VK_F24                 VK = 0x87
	VK_NUMLOCK             VK = 0x90
	VK_SCROLL              VK = 0x91
	VK_OEM_NEC_EQUAL       VK = 0x92
	VK_OEM_FJ_JISHO        VK = 0x92
	VK_OEM_FJ_MASSHOU      VK = 0x93
	VK_OEM_FJ_TOUROKU      VK = 0x94
	VK_OEM_FJ_LOYA         VK = 0x95
	VK_OEM_FJ_ROYA         VK = 0x96
	VK_LSHIFT              VK = 0xA0
	VK_RSHIFT              VK = 0xA1
	VK_LCONTROL            VK = 0xA2
	VK_RCONTROL            VK = 0xA3
	VK_LMENU               VK = 0xA4
	VK_RMENU               VK = 0xA5
	VK_BROWSER_BACK        VK = 0xA6
	VK_BROWSER_FORWARD     VK = 0xA7
	VK_BROWSER_REFRESH     VK = 0xA8
	VK_BROWSER_STOP        VK = 0xA9
	VK_BROWSER_SEARCH      VK = 0xAA
	VK_BROWSER_FAVORITES   VK = 0xAB
	VK_BROWSER_HOME        VK = 0xAC
	VK_VOLUME_MUTE         VK = 0xAD
	VK_VOLUME_DOWN         VK = 0xAE
	VK_VOLUME_UP           VK = 0xAF
	VK_MEDIA_NEXT_TRACK    VK = 0xB0
	VK_MEDIA_PREV_TRACK    VK = 0xB1
	VK_MEDIA_STOP          VK = 0xB2
	VK_MEDIA_PLAY_PAUSE    VK = 0xB3
	VK_LAUNCH_MAIL         VK = 0xB4
	VK_LAUNCH_MEDIA_SELECT VK = 0xB5
	VK_LAUNCH_APP1         VK = 0xB6
	VK_LAUNCH_APP2         VK = 0xB7
	VK_OEM_1               VK = 0xBA
	VK_OEM_PLUS            VK = 0xBB
	VK_OEM_COMMA           VK = 0xBC
	VK_OEM_MINUS           VK = 0xBD
	VK_OEM_PERIOD          VK = 0xBE
	VK_OEM_2               VK = 0xBF
	VK_OEM_3               VK = 0xC0
	VK_OEM_4               VK = 0xDB
	VK_OEM_5               VK = 0xDC
	VK_OEM_6               VK = 0xDD
	VK_OEM_7               VK = 0xDE
	VK_OEM_8               VK = 0xDF
	VK_OEM_AX              VK = 0xE1
	VK_OEM_102             VK = 0xE2
	VK_ICO_HELP            VK = 0xE3
	VK_ICO_00              VK = 0xE4
	VK_PROCESSKEY          VK = 0xE5
	VK_ICO_CLEAR           VK = 0xE6
	VK_PACKET              VK = 0xE7
	VK_OEM_RESET           VK = 0xE9
	VK_OEM_JUMP            VK = 0xEA
	VK_OEM_PA1             VK = 0xEB
	VK_OEM_PA2             VK = 0xEC
	VK_OEM_PA3             VK = 0xED
	VK_OEM_WSCTRL          VK = 0xEE
	VK_OEM_CUSEL           VK = 0xEF
	VK_OEM_ATTN            VK = 0xF0
	VK_OEM_FINISH          VK = 0xF1
	VK_OEM_COPY            VK = 0xF2
	VK_OEM_AUTO            VK = 0xF3
	VK_OEM_ENLW            VK = 0xF4
	VK_OEM_BACKTAB         VK = 0xF5
	VK_ATTN                VK = 0xF6
	VK_CRSEL               VK = 0xF7
	VK_EXSEL               VK = 0xF8
	VK_EREOF               VK = 0xF9
	VK_PLAY                VK = 0xFA
	VK_ZOOM                VK = 0xFB
	VK_NONAME              VK = 0xFC
	VK_PA1                 VK = 0xFD
	VK_OEM_CLEAR           VK = 0xFE
)

type VS_PART int32 // https://docs.microsoft.com/en-us/windows/win32/controls/parts-and-states

const ( // list view styles
	VS_PART_LVP_LISTITEM         VS_PART = 1
	VS_PART_LVP_LISTGROUP        VS_PART = 2
	VS_PART_LVP_LISTDETAIL       VS_PART = 3
	VS_PART_LVP_LISTSORTEDDETAIL VS_PART = 4
	VS_PART_LVP_EMPTYTEXT        VS_PART = 5
	VS_PART_LVP_GROUPHEADER      VS_PART = 6
	VS_PART_LVP_GROUPHEADERLINE  VS_PART = 7
	VS_PART_LVP_EXPANDBUTTON     VS_PART = 8
	VS_PART_LVP_COLLAPSEBUTTON   VS_PART = 9
	VS_PART_LVP_COLUMNDETAIL     VS_PART = 10
)

type VS_STATE int32 // https://docs.microsoft.com/en-us/windows/win32/controls/parts-and-states

const ( // list view states
	VS_STATE_LISS_NORMAL           VS_STATE = 1
	VS_STATE_LISS_HOT              VS_STATE = 2
	VS_STATE_LISS_SELECTED         VS_STATE = 3
	VS_STATE_LISS_DISABLED         VS_STATE = 4
	VS_STATE_LISS_SELECTEDNOTFOCUS VS_STATE = 5
	VS_STATE_LISS_HOTSELECTED      VS_STATE = 6

	VS_STATE_LVGH_OPEN                       VS_STATE = 1
	VS_STATE_LVGH_OPENHOT                    VS_STATE = 2
	VS_STATE_LVGH_OPENSELECTED               VS_STATE = 3
	VS_STATE_LVGH_OPENSELECTEDHOT            VS_STATE = 4
	VS_STATE_LVGH_OPENSELECTEDNOTFOCUSED     VS_STATE = 5
	VS_STATE_LVGH_OPENSELECTEDNOTFOCUSEDHOT  VS_STATE = 6
	VS_STATE_LVGH_OPENMIXEDSELECTION         VS_STATE = 7
	VS_STATE_LVGH_OPENMIXEDSELECTIONHOT      VS_STATE = 8
	VS_STATE_LVGH_CLOSE                      VS_STATE = 9
	VS_STATE_LVGH_CLOSEHOT                   VS_STATE = 10
	VS_STATE_LVGH_CLOSESELECTED              VS_STATE = 11
	VS_STATE_LVGH_CLOSESELECTEDHOT           VS_STATE = 12
	VS_STATE_LVGH_CLOSESELECTEDNOTFOCUSED    VS_STATE = 13
	VS_STATE_LVGH_CLOSESELECTEDNOTFOCUSEDHOT VS_STATE = 14
	VS_STATE_LVGH_CLOSEMIXEDSELECTION        VS_STATE = 15
	VS_STATE_LVGH_CLOSEMIXEDSELECTIONHOT     VS_STATE = 16

	VS_STATE_LVGHL_OPEN                       VS_STATE = 1
	VS_STATE_LVGHL_OPENHOT                    VS_STATE = 2
	VS_STATE_LVGHL_OPENSELECTED               VS_STATE = 3
	VS_STATE_LVGHL_OPENSELECTEDHOT            VS_STATE = 4
	VS_STATE_LVGHL_OPENSELECTEDNOTFOCUSED     VS_STATE = 5
	VS_STATE_LVGHL_OPENSELECTEDNOTFOCUSEDHOT  VS_STATE = 6
	VS_STATE_LVGHL_OPENMIXEDSELECTION         VS_STATE = 7
	VS_STATE_LVGHL_OPENMIXEDSELECTIONHOT      VS_STATE = 8
	VS_STATE_LVGHL_CLOSE                      VS_STATE = 9
	VS_STATE_LVGHL_CLOSEHOT                   VS_STATE = 10
	VS_STATE_LVGHL_CLOSESELECTED              VS_STATE = 11
	VS_STATE_LVGHL_CLOSESELECTEDHOT           VS_STATE = 12
	VS_STATE_LVGHL_CLOSESELECTEDNOTFOCUSED    VS_STATE = 13
	VS_STATE_LVGHL_CLOSESELECTEDNOTFOCUSEDHOT VS_STATE = 14
	VS_STATE_LVGHL_CLOSEMIXEDSELECTION        VS_STATE = 15
	VS_STATE_LVGHL_CLOSEMIXEDSELECTIONHOT     VS_STATE = 16

	VS_STATE_LVEB_NORMAL VS_STATE = 1
	VS_STATE_LVEB_HOVER  VS_STATE = 2
	VS_STATE_LVEB_PUSHED VS_STATE = 3

	VS_STATE_LVCB_NORMAL VS_STATE = 1
	VS_STATE_LVCB_HOVER  VS_STATE = 2
	VS_STATE_LVCB_PUSHED VS_STATE = 3
)
