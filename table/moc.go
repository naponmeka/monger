package table

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type CustomTableModel_ITF interface {
	std_core.QAbstractTableModel_ITF
	CustomTableModel_PTR() *CustomTableModel
}

func (ptr *CustomTableModel) CustomTableModel_PTR() *CustomTableModel {
	return ptr
}

func (ptr *CustomTableModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func (ptr *CustomTableModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractTableModel_PTR().SetPointer(p)
	}
}

func PointerFromCustomTableModel(ptr CustomTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CustomTableModel_PTR().Pointer()
	}
	return nil
}

func NewCustomTableModelFromPointer(ptr unsafe.Pointer) (n *CustomTableModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CustomTableModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CustomTableModel:
			n = deduced

		case *std_core.QAbstractTableModel:
			n = &CustomTableModel{QAbstractTableModel: *deduced}

		default:
			n = new(CustomTableModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackCustomTableModelb2ccaa_Constructor
func callbackCustomTableModelb2ccaa_Constructor(ptr unsafe.Pointer) {
	this := NewCustomTableModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectRemove(this.remove)
	this.ConnectAdd(this.add)
	this.ConnectEdit(this.edit)
	this.init()
}

//export callbackCustomTableModelb2ccaa_Remove
func callbackCustomTableModelb2ccaa_Remove(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomTableModel) ConnectRemove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remove") {
			C.CustomTableModelb2ccaa_ConnectRemove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remove", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", f)
		}
	}
}

func (ptr *CustomTableModel) DisconnectRemove() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DisconnectRemove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *CustomTableModel) Remove() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_Remove(ptr.Pointer())
	}
}

//export callbackCustomTableModelb2ccaa_Add
func callbackCustomTableModelb2ccaa_Add(ptr unsafe.Pointer, item C.uintptr_t) {
	var itemD TableItem
	if itemI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(item))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(item)))
		itemD = itemI.(TableItem)
	}
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(TableItem))(itemD)
	}

}

func (ptr *CustomTableModel) ConnectAdd(f func(item TableItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomTableModelb2ccaa_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(item TableItem) {
				signal.(func(TableItem))(item)
				f(item)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *CustomTableModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *CustomTableModel) Add(item TableItem) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&item), item)
		C.CustomTableModelb2ccaa_Add(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&item))))
	}
}

//export callbackCustomTableModelb2ccaa_Edit
func callbackCustomTableModelb2ccaa_Edit(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(firstName), cGoUnpackString(lastName))
	}

}

func (ptr *CustomTableModel) ConnectEdit(f func(firstName string, lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edit") {
			C.CustomTableModelb2ccaa_ConnectEdit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "edit", func(firstName string, lastName string) {
				signal.(func(string, string))(firstName, lastName)
				f(firstName, lastName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edit", f)
		}
	}
}

func (ptr *CustomTableModel) DisconnectEdit() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DisconnectEdit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edit")
	}
}

func (ptr *CustomTableModel) Edit(firstName string, lastName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.CustomTableModelb2ccaa_Edit(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func CustomTableModel_QRegisterMetaType() int {
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QRegisterMetaType()))
}

func (ptr *CustomTableModel) QRegisterMetaType() int {
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QRegisterMetaType()))
}

func CustomTableModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomTableModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QRegisterMetaType2(typeNameC)))
}

func CustomTableModel_QmlRegisterType() int {
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QmlRegisterType()))
}

func (ptr *CustomTableModel) QmlRegisterType() int {
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QmlRegisterType()))
}

func CustomTableModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomTableModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CustomTableModelb2ccaa_CustomTableModelb2ccaa_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModelb2ccaa___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTableModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.CustomTableModelb2ccaa___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTableModelb2ccaa___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTableModelb2ccaa___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTableModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTableModelb2ccaa___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTableModel) __roleNames_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___roleNames_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.CustomTableModelb2ccaa___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModelb2ccaa___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTableModel) __itemData_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___itemData_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.CustomTableModelb2ccaa___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __match_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___match_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTableModelb2ccaa___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTableModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModelb2ccaa___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModelb2ccaa___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModelb2ccaa___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModelb2ccaa___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __children_newList() unsafe.Pointer {
	return C.CustomTableModelb2ccaa___children_newList(ptr.Pointer())
}

func NewCustomTableModel(parent std_core.QObject_ITF) *CustomTableModel {
	tmpValue := NewCustomTableModelFromPointer(C.CustomTableModelb2ccaa_NewCustomTableModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomTableModelb2ccaa_DestroyCustomTableModel
func callbackCustomTableModelb2ccaa_DestroyCustomTableModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CustomTableModel"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTableModelFromPointer(ptr).DestroyCustomTableModelDefault()
	}
}

func (ptr *CustomTableModel) ConnectDestroyCustomTableModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CustomTableModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~CustomTableModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CustomTableModel", f)
		}
	}
}

func (ptr *CustomTableModel) DisconnectDestroyCustomTableModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CustomTableModel")
	}
}

func (ptr *CustomTableModel) DestroyCustomTableModel() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DestroyCustomTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomTableModel) DestroyCustomTableModelDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DestroyCustomTableModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTableModelb2ccaa_DropMimeData
func callbackCustomTableModelb2ccaa_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_Index
func callbackCustomTableModelb2ccaa_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *CustomTableModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_Sibling
func callbackCustomTableModelb2ccaa_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *CustomTableModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_Flags
func callbackCustomTableModelb2ccaa_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.CustomTableModelb2ccaa_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackCustomTableModelb2ccaa_InsertColumns
func callbackCustomTableModelb2ccaa_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_InsertRows
func callbackCustomTableModelb2ccaa_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_MoveColumns
func callbackCustomTableModelb2ccaa_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTableModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_MoveRows
func callbackCustomTableModelb2ccaa_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTableModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_RemoveColumns
func callbackCustomTableModelb2ccaa_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_RemoveRows
func callbackCustomTableModelb2ccaa_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_SetData
func callbackCustomTableModelb2ccaa_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTableModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_SetHeaderData
func callbackCustomTableModelb2ccaa_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTableModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_SetItemData
func callbackCustomTableModelb2ccaa_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewCustomTableModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *CustomTableModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_Submit
func callbackCustomTableModelb2ccaa_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *CustomTableModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_ColumnsAboutToBeInserted
func callbackCustomTableModelb2ccaa_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_ColumnsAboutToBeMoved
func callbackCustomTableModelb2ccaa_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackCustomTableModelb2ccaa_ColumnsAboutToBeRemoved
func callbackCustomTableModelb2ccaa_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_ColumnsInserted
func callbackCustomTableModelb2ccaa_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_ColumnsMoved
func callbackCustomTableModelb2ccaa_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackCustomTableModelb2ccaa_ColumnsRemoved
func callbackCustomTableModelb2ccaa_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_DataChanged
func callbackCustomTableModelb2ccaa_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackCustomTableModelb2ccaa_FetchMore
func callbackCustomTableModelb2ccaa_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewCustomTableModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *CustomTableModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackCustomTableModelb2ccaa_HeaderDataChanged
func callbackCustomTableModelb2ccaa_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_LayoutAboutToBeChanged
func callbackCustomTableModelb2ccaa_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTableModelb2ccaa_LayoutChanged
func callbackCustomTableModelb2ccaa_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTableModelb2ccaa_ModelAboutToBeReset
func callbackCustomTableModelb2ccaa_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomTableModelb2ccaa_ModelReset
func callbackCustomTableModelb2ccaa_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomTableModelb2ccaa_ResetInternalData
func callbackCustomTableModelb2ccaa_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTableModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *CustomTableModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackCustomTableModelb2ccaa_Revert
func callbackCustomTableModelb2ccaa_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTableModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *CustomTableModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_RevertDefault(ptr.Pointer())
	}
}

//export callbackCustomTableModelb2ccaa_RowsAboutToBeInserted
func callbackCustomTableModelb2ccaa_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackCustomTableModelb2ccaa_RowsAboutToBeMoved
func callbackCustomTableModelb2ccaa_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackCustomTableModelb2ccaa_RowsAboutToBeRemoved
func callbackCustomTableModelb2ccaa_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_RowsInserted
func callbackCustomTableModelb2ccaa_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_RowsMoved
func callbackCustomTableModelb2ccaa_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackCustomTableModelb2ccaa_RowsRemoved
func callbackCustomTableModelb2ccaa_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModelb2ccaa_Sort
func callbackCustomTableModelb2ccaa_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewCustomTableModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *CustomTableModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackCustomTableModelb2ccaa_RoleNames
func callbackCustomTableModelb2ccaa_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewCustomTableModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTableModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.CustomTableModelb2ccaa_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackCustomTableModelb2ccaa_ItemData
func callbackCustomTableModelb2ccaa_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__itemData_newList())
		for k, v := range NewCustomTableModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTableModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.CustomTableModelb2ccaa_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackCustomTableModelb2ccaa_MimeData
func callbackCustomTableModelb2ccaa_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewCustomTableModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewCustomTableModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *CustomTableModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.CustomTableModelb2ccaa_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_Buddy
func callbackCustomTableModelb2ccaa_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_Parent
func callbackCustomTableModelb2ccaa_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModelb2ccaa_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_Match
func callbackCustomTableModelb2ccaa_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__match_newList())
		for _, v := range NewCustomTableModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTableModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.CustomTableModelb2ccaa_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackCustomTableModelb2ccaa_Span
func callbackCustomTableModelb2ccaa_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewCustomTableModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomTableModelb2ccaa_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_MimeTypes
func callbackCustomTableModelb2ccaa_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewCustomTableModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *CustomTableModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.CustomTableModelb2ccaa_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackCustomTableModelb2ccaa_Data
func callbackCustomTableModelb2ccaa_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTableModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *CustomTableModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModelb2ccaa_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_HeaderData
func callbackCustomTableModelb2ccaa_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTableModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *CustomTableModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModelb2ccaa_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModelb2ccaa_SupportedDragActions
func callbackCustomTableModelb2ccaa_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *CustomTableModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTableModelb2ccaa_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTableModelb2ccaa_SupportedDropActions
func callbackCustomTableModelb2ccaa_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *CustomTableModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTableModelb2ccaa_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTableModelb2ccaa_CanDropMimeData
func callbackCustomTableModelb2ccaa_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_CanFetchMore
func callbackCustomTableModelb2ccaa_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_HasChildren
func callbackCustomTableModelb2ccaa_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_ColumnCount
func callbackCustomTableModelb2ccaa_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTableModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTableModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTableModelb2ccaa_RowCount
func callbackCustomTableModelb2ccaa_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTableModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTableModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModelb2ccaa_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTableModelb2ccaa_Event
func callbackCustomTableModelb2ccaa_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomTableModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_EventFilter
func callbackCustomTableModelb2ccaa_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomTableModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModelb2ccaa_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCustomTableModelb2ccaa_ChildEvent
func callbackCustomTableModelb2ccaa_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomTableModelb2ccaa_ConnectNotify
func callbackCustomTableModelb2ccaa_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTableModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTableModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTableModelb2ccaa_CustomEvent
func callbackCustomTableModelb2ccaa_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomTableModelb2ccaa_DeleteLater
func callbackCustomTableModelb2ccaa_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTableModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomTableModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTableModelb2ccaa_Destroyed
func callbackCustomTableModelb2ccaa_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCustomTableModelb2ccaa_DisconnectNotify
func callbackCustomTableModelb2ccaa_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTableModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTableModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTableModelb2ccaa_ObjectNameChanged
func callbackCustomTableModelb2ccaa_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCustomTableModelb2ccaa_TimerEvent
func callbackCustomTableModelb2ccaa_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModelb2ccaa_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}