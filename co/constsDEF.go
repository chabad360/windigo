/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

import (
	"syscall"
)

type DPI_AWARE_CTX int32 // SetProcessDpiAwarenessContext value

const (
	DPI_AWARE_CTX_UNAWARE           DPI_AWARE_CTX = -1
	DPI_AWARE_CTX_SYSTEM_AWARE      DPI_AWARE_CTX = -2
	DPI_AWARE_CTX_PER_MON_AWARE     DPI_AWARE_CTX = -3
	DPI_AWARE_CTX_PER_MON_AWARE_V2  DPI_AWARE_CTX = -4
	DPI_AWARE_CTX_UNAWARE_GDISCALED DPI_AWARE_CTX = -5
)

type EM WM // edit control message

const (
	EM_GETSEL              EM = 0x00B0
	EM_SETSEL              EM = 0x00B1
	EM_GETRECT             EM = 0x00B2
	EM_SETRECT             EM = 0x00B3
	EM_SETRECTNP           EM = 0x00B4
	EM_SCROLL              EM = 0x00B5
	EM_LINESCROLL          EM = 0x00B6
	EM_SCROLLCARET         EM = 0x00B7
	EM_GETMODIFY           EM = 0x00B8
	EM_SETMODIFY           EM = 0x00B9
	EM_GETLINECOUNT        EM = 0x00BA
	EM_LINEINDEX           EM = 0x00BB
	EM_SETHANDLE           EM = 0x00BC
	EM_GETHANDLE           EM = 0x00BD
	EM_GETTHUMB            EM = 0x00BE
	EM_LINELENGTH          EM = 0x00C1
	EM_REPLACESEL          EM = 0x00C2
	EM_GETLINE             EM = 0x00C4
	EM_LIMITTEXT           EM = 0x00C5
	EM_CANUNDO             EM = 0x00C6
	EM_UNDO                EM = 0x00C7
	EM_FMTLINES            EM = 0x00C8
	EM_LINEFROMCHAR        EM = 0x00C9
	EM_SETTABSTOPS         EM = 0x00CB
	EM_SETPASSWORDCHAR     EM = 0x00CC
	EM_EMPTYUNDOBUFFER     EM = 0x00CD
	EM_GETFIRSTVISIBLELINE EM = 0x00CE
	EM_SETREADONLY         EM = 0x00CF
	EM_SETWORDBREAKPROC    EM = 0x00D0
	EM_GETWORDBREAKPROC    EM = 0x00D1
	EM_GETPASSWORDCHAR     EM = 0x00D2
	EM_SETMARGINS          EM = 0x00D3
	EM_GETMARGINS          EM = 0x00D4
	EM_SETLIMITTEXT        EM = EM_LIMITTEXT
	EM_GETLIMITTEXT        EM = 0x00D5
	EM_POSFROMCHAR         EM = 0x00D6
	EM_CHARFROMPOS         EM = 0x00D7
	EM_SETIMESTATUS        EM = 0x00D8
	EM_GETIMESTATUS        EM = 0x00D9
)

type EMF uint32 // NMLVEMPTYMARKUP dwFlags

const (
	EMF_NULL     EMF = 0x00000000
	EMF_CENTERED EMF = 0x00000001
)

type ES WS // edit control style

const (
	ES_LEFT        ES = 0x0000
	ES_CENTER      ES = 0x0001
	ES_RIGHT       ES = 0x0002
	ES_MULTILINE   ES = 0x0004
	ES_UPPERCASE   ES = 0x0008
	ES_LOWERCASE   ES = 0x0010
	ES_PASSWORD    ES = 0x0020
	ES_AUTOVSCROLL ES = 0x0040
	ES_AUTOHSCROLL ES = 0x0080
	ES_NOHIDESEL   ES = 0x0100
	ES_OEMCONVERT  ES = 0x0400
	ES_READONLY    ES = 0x0800
	ES_WANTRETURN  ES = 0x1000
	ES_NUMBER      ES = 0x2000
)

type ERROR syscall.Errno // GetLastError result

const (
	ERROR_SUCCESS              ERROR = 0
	ERROR_INVALID_HANDLE       ERROR = 6
	ERROR_SHARING_VIOLATION    ERROR = 32
	ERROR_CLASS_ALREADY_EXISTS ERROR = 1410
)

type FAPPCOMMAND uint32 // WM_APPCOMMAND

const (
	FAPPCOMMAND_MOUSE FAPPCOMMAND = 0x8000
	FAPPCOMMAND_KEY   FAPPCOMMAND = 0
	FAPPCOMMAND_OEM   FAPPCOMMAND = 0x1000
)

type FILE_ATTRIBUTE uint32 // CreateFile dwFlagsAndAttributes

const (
	FILE_ATTRIBUTE_READONLY              FILE_ATTRIBUTE = 0x00000001
	FILE_ATTRIBUTE_HIDDEN                FILE_ATTRIBUTE = 0x00000002
	FILE_ATTRIBUTE_SYSTEM                FILE_ATTRIBUTE = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY             FILE_ATTRIBUTE = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE               FILE_ATTRIBUTE = 0x00000020
	FILE_ATTRIBUTE_DEVICE                FILE_ATTRIBUTE = 0x00000040
	FILE_ATTRIBUTE_NORMAL                FILE_ATTRIBUTE = 0x00000080
	FILE_ATTRIBUTE_TEMPORARY             FILE_ATTRIBUTE = 0x00000100
	FILE_ATTRIBUTE_SPARSE_FILE           FILE_ATTRIBUTE = 0x00000200
	FILE_ATTRIBUTE_REPARSE_POINT         FILE_ATTRIBUTE = 0x00000400
	FILE_ATTRIBUTE_COMPRESSED            FILE_ATTRIBUTE = 0x00000800
	FILE_ATTRIBUTE_OFFLINE               FILE_ATTRIBUTE = 0x00001000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   FILE_ATTRIBUTE = 0x00002000
	FILE_ATTRIBUTE_ENCRYPTED             FILE_ATTRIBUTE = 0x00004000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      FILE_ATTRIBUTE = 0x00008000
	FILE_ATTRIBUTE_VIRTUAL               FILE_ATTRIBUTE = 0x00010000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         FILE_ATTRIBUTE = 0x00020000
	FILE_ATTRIBUTE_EA                    FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_PINNED                FILE_ATTRIBUTE = 0x00080000
	FILE_ATTRIBUTE_UNPINNED              FILE_ATTRIBUTE = 0x00100000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS FILE_ATTRIBUTE = 0x00400000
)

type FILE_DISPO uint32 // CreateFile dwCreationDisposition

const (
	FILE_DISPO_CREATE_ALWAYS     FILE_DISPO = 2
	FILE_DISPO_CREATE_NEW        FILE_DISPO = 1
	FILE_DISPO_OPEN_ALWAYS       FILE_DISPO = 4
	FILE_DISPO_OPEN_EXISTING     FILE_DISPO = 3
	FILE_DISPO_TRUNCATE_EXISTING FILE_DISPO = 5
)

type FILE_FLAG uint32 // CreateFile dwFlagsAndAttributes

const (
	FILE_FLAG_NONE                  FILE_FLAG = 0
	FILE_FLAG_WRITE_THROUGH         FILE_FLAG = 0x80000000
	FILE_FLAG_OVERLAPPED            FILE_FLAG = 0x40000000
	FILE_FLAG_NO_BUFFERING          FILE_FLAG = 0x20000000
	FILE_FLAG_RANDOM_ACCESS         FILE_FLAG = 0x10000000
	FILE_FLAG_SEQUENTIAL_SCAN       FILE_FLAG = 0x08000000
	FILE_FLAG_DELETE_ON_CLOSE       FILE_FLAG = 0x04000000
	FILE_FLAG_BACKUP_SEMANTICS      FILE_FLAG = 0x02000000
	FILE_FLAG_POSIX_SEMANTICS       FILE_FLAG = 0x01000000
	FILE_FLAG_SESSION_AWARE         FILE_FLAG = 0x00800000
	FILE_FLAG_OPEN_REPARSE_POINT    FILE_FLAG = 0x00200000
	FILE_FLAG_OPEN_NO_RECALL        FILE_FLAG = 0x00100000
	FILE_FLAG_FIRST_PIPE_INSTANCE   FILE_FLAG = 0x00080000
	FILE_FLAG_OPEN_REQUIRING_OPLOCK FILE_FLAG = 0x00040000
)

type FILE_MAP uint32 // MapViewOfFile dwDesiredAccess

const (
	FILE_MAP_WRITE           FILE_MAP = FILE_MAP(SECTION_MAP_WRITE)
	FILE_MAP_READ            FILE_MAP = FILE_MAP(SECTION_MAP_READ)
	FILE_MAP_ALL_ACCESS      FILE_MAP = FILE_MAP(SECTION_ALL_ACCESS)
	FILE_MAP_EXECUTE         FILE_MAP = FILE_MAP(SECTION_MAP_EXECUTE_EXPLICIT)
	FILE_MAP_COPY            FILE_MAP = 0x00000001
	FILE_MAP_RESERVE         FILE_MAP = 0x80000000
	FILE_MAP_TARGETS_INVALID FILE_MAP = 0x40000000
	FILE_MAP_LARGE_PAGES     FILE_MAP = 0x20000000
)

type FILE_SETPTR uint32 // SetFilePointer dwMoveMethod

const (
	FILE_SETPTR_BEGIN   FILE_SETPTR = 0
	FILE_SETPTR_CURRENT FILE_SETPTR = 1
	FILE_SETPTR_END     FILE_SETPTR = 2
)

type FILE_SHARE uint32 // CreateFile dwShareMode

const (
	FILE_SHARE_NONE   FILE_SHARE = 0
	FILE_SHARE_READ   FILE_SHARE = 0x00000001
	FILE_SHARE_WRITE  FILE_SHARE = 0x00000002
	FILE_SHARE_DELETE FILE_SHARE = 0x00000004
)

type FW uint32 // LOGFONT lfWeight

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT FW = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD   FW = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD  FW = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK      FW = FW_HEAVY
)
