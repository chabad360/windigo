/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// Returned by GetLastError(), also an HRESULT.
//
// We can't simply use syscall.Errno because it's an uintptr (8 bytes), thus a
// native struct with such a field type would be wrong.
type ERROR uint32

const (
	ERROR_SUCCESS                    ERROR = 0
	ERROR_INVALID_FUNCTION           ERROR = 1
	ERROR_FILE_NOT_FOUND             ERROR = 2
	ERROR_PATH_NOT_FOUND             ERROR = 3
	ERROR_TOO_MANY_OPEN_FILES        ERROR = 4
	ERROR_ACCESS_DENIED              ERROR = 5
	ERROR_INVALID_HANDLE             ERROR = 6
	ERROR_ARENA_TRASHED              ERROR = 7
	ERROR_NOT_ENOUGH_MEMORY          ERROR = 8
	ERROR_INVALID_BLOCK              ERROR = 9
	ERROR_BAD_ENVIRONMENT            ERROR = 10
	ERROR_BAD_FORMAT                 ERROR = 11
	ERROR_INVALID_ACCESS             ERROR = 12
	ERROR_INVALID_DATA               ERROR = 13
	ERROR_OUTOFMEMORY                ERROR = 14
	ERROR_INVALID_DRIVE              ERROR = 15
	ERROR_CURRENT_DIRECTORY          ERROR = 16
	ERROR_NOT_SAME_DEVICE            ERROR = 17
	ERROR_NO_MORE_FILES              ERROR = 18
	ERROR_WRITE_PROTECT              ERROR = 19
	ERROR_BAD_UNIT                   ERROR = 20
	ERROR_NOT_READY                  ERROR = 21
	ERROR_BAD_COMMAND                ERROR = 22
	ERROR_CRC                        ERROR = 23
	ERROR_BAD_LENGTH                 ERROR = 24
	ERROR_SEEK                       ERROR = 25
	ERROR_NOT_DOS_DISK               ERROR = 26
	ERROR_SECTOR_NOT_FOUND           ERROR = 27
	ERROR_OUT_OF_PAPER               ERROR = 28
	ERROR_WRITE_FAULT                ERROR = 29
	ERROR_READ_FAULT                 ERROR = 30
	ERROR_GEN_FAILURE                ERROR = 31
	ERROR_SHARING_VIOLATION          ERROR = 32
	ERROR_LOCK_VIOLATION             ERROR = 33
	ERROR_WRONG_DISK                 ERROR = 34
	ERROR_SHARING_BUFFER_EXCEEDED    ERROR = 36
	ERROR_HANDLE_EOF                 ERROR = 38
	ERROR_HANDLE_DISK_FULL           ERROR = 39
	ERROR_NOT_SUPPORTED              ERROR = 50
	ERROR_REM_NOT_LIST               ERROR = 51
	ERROR_DUP_NAME                   ERROR = 52
	ERROR_BAD_NETPATH                ERROR = 53
	ERROR_NETWORK_BUSY               ERROR = 54
	ERROR_DEV_NOT_EXIST              ERROR = 55
	ERROR_TOO_MANY_CMDS              ERROR = 56
	ERROR_ADAP_HDW_ERR               ERROR = 57
	ERROR_BAD_NET_RESP               ERROR = 58
	ERROR_UNEXP_NET_ERR              ERROR = 59
	ERROR_BAD_REM_ADAP               ERROR = 60
	ERROR_INVALID_PASSWORD           ERROR = 86
	ERROR_INVALID_PARAMETER          ERROR = 87
	ERROR_NET_WRITE_FAULT            ERROR = 88
	ERROR_NO_PROC_SLOTS              ERROR = 89
	ERROR_TOO_MANY_SEMAPHORES        ERROR = 100
	ERROR_EXCL_SEM_ALREADY_OWNED     ERROR = 101
	ERROR_SEM_IS_SET                 ERROR = 102
	ERROR_TOO_MANY_SEM_REQUESTS      ERROR = 103
	ERROR_INVALID_AT_INTERRUPT_TIME  ERROR = 104
	ERROR_SEM_OWNER_DIED             ERROR = 105
	ERROR_SEM_USER_LIMIT             ERROR = 106
	ERROR_DISK_CHANGE                ERROR = 107
	ERROR_DRIVE_LOCKED               ERROR = 108
	ERROR_BROKEN_PIPE                ERROR = 109
	ERROR_OPEN_FAILED                ERROR = 110
	ERROR_BUFFER_OVERFLOW            ERROR = 111
	ERROR_DISK_FULL                  ERROR = 112
	ERROR_NO_MORE_SEARCH_HANDLES     ERROR = 113
	ERROR_INVALID_TARGET_HANDLE      ERROR = 114
	ERROR_INVALID_CATEGORY           ERROR = 117
	ERROR_INVALID_VERIFY_SWITCH      ERROR = 118
	ERROR_BAD_DRIVER_LEVEL           ERROR = 119
	ERROR_CALL_NOT_IMPLEMENTED       ERROR = 120
	ERROR_SEM_TIMEOUT                ERROR = 121
	ERROR_INSUFFICIENT_BUFFER        ERROR = 122
	ERROR_INVALID_NAME               ERROR = 123
	ERROR_INVALID_LEVEL              ERROR = 124
	ERROR_NO_VOLUME_LABEL            ERROR = 125
	ERROR_MOD_NOT_FOUND              ERROR = 126
	ERROR_PROC_NOT_FOUND             ERROR = 127
	ERROR_WAIT_NO_CHILDREN           ERROR = 128
	ERROR_CHILD_NOT_COMPLETE         ERROR = 129
	ERROR_DIRECT_ACCESS_HANDLE       ERROR = 130
	ERROR_NEGATIVE_SEEK              ERROR = 131
	ERROR_SEEK_ON_DEVICE             ERROR = 132
	ERROR_IS_JOIN_TARGET             ERROR = 133
	ERROR_IS_JOINED                  ERROR = 134
	ERROR_IS_SUBSTED                 ERROR = 135
	ERROR_NOT_JOINED                 ERROR = 136
	ERROR_NOT_SUBSTED                ERROR = 137
	ERROR_JOIN_TO_JOIN               ERROR = 138
	ERROR_SUBST_TO_SUBST             ERROR = 139
	ERROR_JOIN_TO_SUBST              ERROR = 140
	ERROR_SUBST_TO_JOIN              ERROR = 141
	ERROR_BUSY_DRIVE                 ERROR = 142
	ERROR_SAME_DRIVE                 ERROR = 143
	ERROR_DIR_NOT_ROOT               ERROR = 144
	ERROR_DIR_NOT_EMPTY              ERROR = 145
	ERROR_IS_SUBST_PATH              ERROR = 146
	ERROR_IS_JOIN_PATH               ERROR = 147
	ERROR_PATH_BUSY                  ERROR = 148
	ERROR_IS_SUBST_TARGET            ERROR = 149
	ERROR_SYSTEM_TRACE               ERROR = 150
	ERROR_INVALID_EVENT_COUNT        ERROR = 151
	ERROR_TOO_MANY_MUXWAITERS        ERROR = 152
	ERROR_INVALID_LIST_FORMAT        ERROR = 153
	ERROR_LABEL_TOO_LONG             ERROR = 154
	ERROR_TOO_MANY_TCBS              ERROR = 155
	ERROR_SIGNAL_REFUSED             ERROR = 156
	ERROR_DISCARDED                  ERROR = 157
	ERROR_NOT_LOCKED                 ERROR = 158
	ERROR_BAD_THREADID_ADDR          ERROR = 159
	ERROR_BAD_ARGUMENTS              ERROR = 160
	ERROR_BAD_PATHNAME               ERROR = 161
	ERROR_SIGNAL_PENDING             ERROR = 162
	ERROR_MAX_THRDS_REACHED          ERROR = 164
	ERROR_LOCK_FAILED                ERROR = 167
	ERROR_BUSY                       ERROR = 170
	ERROR_DEVICE_SUPPORT_IN_PROGRESS ERROR = 171
	ERROR_CANCEL_VIOLATION           ERROR = 173
	ERROR_ATOMIC_LOCKS_NOT_SUPPORTED ERROR = 174
	ERROR_INVALID_SEGMENT_NUMBER     ERROR = 180
	ERROR_INVALID_ORDINAL            ERROR = 182
	ERROR_ALREADY_EXISTS             ERROR = 183
	ERROR_INVALID_FLAG_NUMBER        ERROR = 186
	ERROR_SEM_NOT_FOUND              ERROR = 187
	ERROR_INVALID_STARTING_CODESEG   ERROR = 188
	ERROR_INVALID_STACKSEG           ERROR = 189
	ERROR_INVALID_MODULETYPE         ERROR = 190
	ERROR_INVALID_EXE_SIGNATURE      ERROR = 191
	ERROR_EXE_MARKED_INVALID         ERROR = 192
	ERROR_BAD_EXE_FORMAT             ERROR = 193
	ERROR_ITERATED_DATA_EXCEEDS_64k  ERROR = 194
	ERROR_INVALID_MINALLOCSIZE       ERROR = 195
	ERROR_DYNLINK_FROM_INVALID_RING  ERROR = 196
	ERROR_IOPL_NOT_ENABLED           ERROR = 197
	ERROR_INVALID_SEGDPL             ERROR = 198
	ERROR_AUTODATASEG_EXCEEDS_64k    ERROR = 199
	ERROR_RING2SEG_MUST_BE_MOVABLE   ERROR = 200

	ERROR_MORE_DATA            ERROR = 234
	ERROR_NO_MORE_ITEMS        ERROR = 259
	ERROR_OLD_WIN_VERSION      ERROR = 1150
	ERROR_CLASS_ALREADY_EXISTS ERROR = 1410

	ERROR_S_OK    ERROR = 0
	ERROR_S_FALSE ERROR = 1

	ERROR_E_UNEXPECTED           ERROR = 0x8000FFFF
	ERROR_E_NOTIMPL              ERROR = 0x80004001
	ERROR_E_OUTOFMEMORY          ERROR = 0x8007000E
	ERROR_E_INVALIDARG           ERROR = 0x80070057
	ERROR_E_NOINTERFACE          ERROR = 0x80004002
	ERROR_E_POINTER              ERROR = 0x80004003
	ERROR_E_HANDLE               ERROR = 0x80070006
	ERROR_E_ABORT                ERROR = 0x80004004
	ERROR_E_FAIL                 ERROR = 0x80004005
	ERROR_E_ACCESSDENIED         ERROR = 0x80070005
	ERROR_E_PENDING              ERROR = 0x8000000A
	ERROR_E_BOUNDS               ERROR = 0x8000000B
	ERROR_E_CHANGED_STATE        ERROR = 0x8000000C
	ERROR_E_ILLEGAL_STATE_CHANGE ERROR = 0x8000000D
	ERROR_E_ILLEGAL_METHOD_CALL  ERROR = 0x8000000E

	ERROR_CO_E_NOTINITIALIZED     ERROR = 0x800401F0
	ERROR_CO_E_ALREADYINITIALIZED ERROR = 0x800401F1

	ERROR_RPC_E_SERVERFAULT  ERROR = 0x80010105
	ERROR_RPC_E_CHANGED_MODE ERROR = 0x80010106
)
