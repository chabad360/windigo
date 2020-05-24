/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api"
)

// Native list view control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type ListView struct {
	controlNativeBase
}

// Optional; returns a ListView with a specific control ID.
func MakeListView(ctrlId api.ID) ListView {
	return ListView{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	lvc := api.LVCOLUMN{
		Mask:    api.LVCF_TEXT | api.LVCF_WIDTH,
		PszText: api.StrToUtf16Ptr(text),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(api.LVM_INSERTCOLUMN, 0xFFFF,
		api.LPARAM(unsafe.Pointer(&lvc)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return newListViewColumn(me, uint32(newIdx)) // return newly inserted column
}

func (me *ListView) AddColumns(texts []string, widths []uint32) *ListView {
	if len(texts) != len(widths) {
		panic("Mismatch lenghts of texts and widths.")
	}
	for i := range texts {
		me.AddColumn(texts[i], widths[i])
	}
	return me
}

func (me *ListView) AddItem(text string) *ListViewItem {
	lvi := api.LVITEM{
		Mask:    api.LVIF_TEXT,
		PszText: api.StrToUtf16Ptr(text),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(api.LVM_INSERTITEM, 0,
		api.LPARAM(unsafe.Pointer(&lvi)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return newListViewItem(me, uint32(newIdx)) // return newly inserted item
}

func (me *ListView) AddItems(texts []string) *ListView {
	for i := range texts {
		me.AddItem(texts[i])
	}
	return me
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles api.WS_EX, styles api.WS,
	lvExStyles api.LVS_EX, lvStyles api.LVS) *ListView {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles|api.WS_EX(lvExStyles),
		"SysListView32", "", styles|api.WS(lvStyles),
		x, y, width, height, parent)
	return me
}

// Calls CreateWindowEx(). List view control will have LVS_REPORT style.
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		api.WS_EX_CLIENTEDGE,
		api.WS_CHILD|api.WS_GROUP|api.WS_TABSTOP|api.WS_VISIBLE,
		api.LVS_EX_FULLROWSELECT,
		api.LVS_REPORT|api.LVS_SHOWSELALWAYS)
}

func (me *ListView) Column(index uint32) *ListViewColumn {
	numCols := me.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve column with index out of bounds.")
	}
	return newListViewColumn(me, index)
}

func (me *ListView) ColumnCount() uint32 {
	hHeader := api.HWND(me.sendLvmMessage(api.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(api.WM(api.HDM_GETITEMCOUNT), 0, 0)
	if int32(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(api.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

func (me *ListView) IsGroupViewEnabled() bool {
	return me.sendLvmMessage(api.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return newListViewItem(me, index)
}

func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(api.LVM_GETITEMCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(api.LVM_GETSELECTEDCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(api.WM_SETREDRAW, api.WPARAM(wp), 0)
	return me
}

func (me *ListView) SetView(view api.LV_VIEW) *ListView {
	if int32(me.sendLvmMessage(api.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

func (me *ListView) View() api.LV_VIEW {
	return api.LV_VIEW(me.sendLvmMessage(api.LVM_GETVIEW, 0, 0))
}

func (me *ListView) sendLvmMessage(msg api.LVM,
	wParam api.WPARAM, lParam api.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(api.WM(msg), wParam, lParam) // simple wrapper
}

//------------------------------------------------------------------------------

// A single column of a list view control.
type ListViewColumn struct {
	owner *ListView
	index uint32
}

func newListViewColumn(owner *ListView, index uint32) *ListViewColumn {
	return &ListViewColumn{
		owner: owner,
		index: index,
	}
}

func (me *ListViewColumn) Index() uint32 {
	return me.index
}

func (me *ListViewColumn) SetText(text string) *ListViewColumn {
	lvc := api.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     api.LVCF_TEXT,
		PszText:  api.StrToUtf16Ptr(text),
	}
	ret := me.owner.sendLvmMessage(api.LVM_SETCOLUMN,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return me
}

func (me *ListViewColumn) SetWidth(width uint32) *ListViewColumn {
	me.owner.sendLvmMessage(api.LVM_SETCOLUMNWIDTH,
		api.WPARAM(me.index), api.LPARAM(width))
	return me
}

func (me *ListViewColumn) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvc := api.LVCOLUMN{
		ISubItem:   int32(me.index),
		Mask:       api.LVCF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.sendLvmMessage(api.LVM_GETCOLUMN,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvc)))
	if ret < 0 {
		panic("LVM_GETCOLUMN failed to get text.")
	}
	return syscall.UTF16ToString(buf)
}

func (me *ListViewColumn) Width() uint32 {
	cx := me.owner.sendLvmMessage(api.LVM_GETCOLUMNWIDTH, api.WPARAM(me.index), 0)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return uint32(cx)
}

//------------------------------------------------------------------------------

// A single item row of a list view control.
type ListViewItem struct {
	owner *ListView
	index uint32
}

func newListViewItem(owner *ListView, index uint32) *ListViewItem {
	return &ListViewItem{
		owner: owner,
		index: index,
	}
}

func (me *ListViewItem) Delete() {
	if me.index >= me.owner.ItemCount() { // index out of bounds: ignore
		return
	}

	ret := me.owner.sendLvmMessage(api.LVM_DELETEITEM,
		api.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.\n", me.index))
	}
}

func (me *ListViewItem) Index() uint32 {
	return me.index
}

func (me *ListViewItem) IsCut() bool {
	sta := me.owner.sendLvmMessage(api.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(api.LVIS_CUT))
	return (api.LVIS(sta) & api.LVIS_CUT) != 0
}

func (me *ListViewItem) IsFocused() bool {
	sta := me.owner.sendLvmMessage(api.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(api.LVIS_FOCUSED))
	return (api.LVIS(sta) & api.LVIS_FOCUSED) != 0
}

func (me *ListViewItem) IsSelected() bool {
	sta := me.owner.sendLvmMessage(api.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(api.LVIS_SELECTED))
	return (api.LVIS(sta) & api.LVIS_SELECTED) != 0
}

func (me *ListViewItem) IsVisible() bool {
	return me.owner.sendLvmMessage(api.LVM_ISITEMVISIBLE,
		api.WPARAM(me.index), 0) != 0
}

func (me *ListViewItem) SetFocus() *ListViewItem {
	lvi := api.LVITEM{
		StateMask: api.LVIS_FOCUSED,
		State:     api.LVIS_FOCUSED,
	}
	ret := me.owner.sendLvmMessage(api.LVM_SETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_FOCUSED.")
	}
	return me
}

func (me *ListViewItem) SetSelected(selected bool) *ListViewItem {
	lvi := api.LVITEM{
		StateMask: api.LVIS_SELECTED,
	}
	if selected { // otherwise remains zero
		lvi.State = api.LVIS_SELECTED
	}
	ret := me.owner.sendLvmMessage(api.LVM_SETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_SELECTED.")
	}
	return me
}

func (me *ListViewItem) SetText(text string) *ListViewItem {
	me.SubItem(0).SetText(text)
	return me
}

func (me *ListViewItem) SubItem(index uint32) *ListViewSubItem {
	numCols := me.owner.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve sub item with index out of bounds.")
	}
	return newListViewSubItem(me, index)
}

func (me *ListViewItem) Text() string {
	return me.SubItem(0).Text()
}

func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(api.LVM_UPDATE, api.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}

//------------------------------------------------------------------------------

// A cell from a list view item row.
type ListViewSubItem struct {
	item  *ListViewItem
	index uint32
}

func newListViewSubItem(item *ListViewItem, index uint32) *ListViewSubItem {
	return &ListViewSubItem{
		item:  item,
		index: index,
	}
}

func (me *ListViewSubItem) Index() uint32 {
	return me.index
}

func (me *ListViewSubItem) SetText(text string) *ListViewSubItem {
	lvi := api.LVITEM{
		ISubItem: int32(me.index),
		PszText:  api.StrToUtf16Ptr(text),
	}
	ret := me.item.owner.sendLvmMessage(api.LVM_SETITEMTEXT,
		api.WPARAM(me.item.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

func (me *ListViewSubItem) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvi := api.LVITEM{
		ISubItem:   int32(me.index),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.item.owner.sendLvmMessage(api.LVM_GETITEMTEXT,
		api.WPARAM(me.item.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf)
}
