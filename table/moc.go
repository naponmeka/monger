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
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
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

//export callbackCustomTableModele3d187_Constructor
func callbackCustomTableModele3d187_Constructor(ptr unsafe.Pointer) {
	this := NewCustomTableModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectRemove(this.remove)
	this.ConnectAdd(this.add)
	this.ConnectEdit(this.edit)
	this.ConnectRemoveAt(this.removeAt)
	this.ConnectEditAt(this.editAt)
	this.init()
}

//export callbackCustomTableModele3d187_Remove
func callbackCustomTableModele3d187_Remove(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *CustomTableModel) ConnectRemove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remove") {
			C.CustomTableModele3d187_ConnectRemove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CustomTableModel) DisconnectRemove() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectRemove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *CustomTableModel) Remove() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_Remove(ptr.Pointer())
	}
}

//export callbackCustomTableModele3d187_Add
func callbackCustomTableModele3d187_Add(ptr unsafe.Pointer, item C.uintptr_t) {
	var itemD TableItem
	if itemI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(item))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(item)))
		itemD = (*(*TableItem)(itemI))
	}
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		(*(*func(TableItem))(signal))(itemD)
	}

}

func (ptr *CustomTableModel) ConnectAdd(f func(item TableItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomTableModele3d187_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			f := func(item TableItem) {
				(*(*func(TableItem))(signal))(item)
				f(item)
			}
			qt.ConnectSignal(ptr.Pointer(), "add", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CustomTableModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *CustomTableModel) Add(item TableItem) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&item), unsafe.Pointer(&item))
		C.CustomTableModele3d187_Add(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&item))))
	}
}

//export callbackCustomTableModele3d187_Edit
func callbackCustomTableModele3d187_Edit(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(firstName), cGoUnpackString(lastName))
	}

}

func (ptr *CustomTableModel) ConnectEdit(f func(firstName string, lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edit") {
			C.CustomTableModele3d187_ConnectEdit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edit"); signal != nil {
			f := func(firstName string, lastName string) {
				(*(*func(string, string))(signal))(firstName, lastName)
				f(firstName, lastName)
			}
			qt.ConnectSignal(ptr.Pointer(), "edit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CustomTableModel) DisconnectEdit() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectEdit(ptr.Pointer())
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
		C.CustomTableModele3d187_Edit(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

//export callbackCustomTableModele3d187_RemoveAt
func callbackCustomTableModele3d187_RemoveAt(ptr unsafe.Pointer, i C.int) {
	if signal := qt.GetSignal(ptr, "removeAt"); signal != nil {
		(*(*func(int))(signal))(int(int32(i)))
	}

}

func (ptr *CustomTableModel) ConnectRemoveAt(f func(i int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeAt") {
			C.CustomTableModele3d187_ConnectRemoveAt(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "removeAt"); signal != nil {
			f := func(i int) {
				(*(*func(int))(signal))(i)
				f(i)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CustomTableModel) DisconnectRemoveAt() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectRemoveAt(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeAt")
	}
}

func (ptr *CustomTableModel) RemoveAt(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_RemoveAt(ptr.Pointer(), C.int(int32(i)))
	}
}

//export callbackCustomTableModele3d187_EditAt
func callbackCustomTableModele3d187_EditAt(ptr unsafe.Pointer, i C.int, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "editAt"); signal != nil {
		(*(*func(int, string, string))(signal))(int(int32(i)), cGoUnpackString(firstName), cGoUnpackString(lastName))
	}

}

func (ptr *CustomTableModel) ConnectEditAt(f func(i int, firstName string, lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "editAt") {
			C.CustomTableModele3d187_ConnectEditAt(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "editAt"); signal != nil {
			f := func(i int, firstName string, lastName string) {
				(*(*func(int, string, string))(signal))(i, firstName, lastName)
				f(i, firstName, lastName)
			}
			qt.ConnectSignal(ptr.Pointer(), "editAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CustomTableModel) DisconnectEditAt() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectEditAt(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "editAt")
	}
}

func (ptr *CustomTableModel) EditAt(i int, firstName string, lastName string) {
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
		C.CustomTableModele3d187_EditAt(ptr.Pointer(), C.int(int32(i)), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func CustomTableModel_QRegisterMetaType() int {
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QRegisterMetaType()))
}

func (ptr *CustomTableModel) QRegisterMetaType() int {
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QRegisterMetaType()))
}

func CustomTableModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomTableModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QRegisterMetaType2(typeNameC)))
}

func CustomTableModel_QmlRegisterType() int {
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QmlRegisterType()))
}

func (ptr *CustomTableModel) QmlRegisterType() int {
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QmlRegisterType()))
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
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.CustomTableModele3d187_CustomTableModele3d187_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModele3d187___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTableModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___setItemData_roles_newList(ptr.Pointer())
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
		}(C.CustomTableModele3d187___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTableModele3d187___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTableModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTableModele3d187___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTableModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTableModele3d187___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTableModel) __roleNames_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___roleNames_newList(ptr.Pointer())
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
		}(C.CustomTableModele3d187___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModele3d187___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTableModel) __itemData_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___itemData_newList(ptr.Pointer())
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
		}(C.CustomTableModele3d187___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTableModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __match_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___match_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTableModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTableModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTableModele3d187_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTableModele3d187___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTableModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModele3d187___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomTableModele3d187___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModele3d187___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomTableModele3d187___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomTableModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModele3d187___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __findChildren_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomTableModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTableModele3d187___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTableModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTableModel) __children_newList() unsafe.Pointer {
	return C.CustomTableModele3d187___children_newList(ptr.Pointer())
}

func NewCustomTableModel(parent std_core.QObject_ITF) *CustomTableModel {
	tmpValue := NewCustomTableModelFromPointer(C.CustomTableModele3d187_NewCustomTableModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomTableModele3d187_DestroyCustomTableModel
func callbackCustomTableModele3d187_DestroyCustomTableModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CustomTableModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCustomTableModelFromPointer(ptr).DestroyCustomTableModelDefault()
	}
}

func (ptr *CustomTableModel) ConnectDestroyCustomTableModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CustomTableModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~CustomTableModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CustomTableModel", unsafe.Pointer(&f))
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
		C.CustomTableModele3d187_DestroyCustomTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomTableModel) DestroyCustomTableModelDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DestroyCustomTableModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTableModele3d187_DropMimeData
func callbackCustomTableModele3d187_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_Index
func callbackCustomTableModele3d187_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *CustomTableModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_Sibling
func callbackCustomTableModele3d187_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *CustomTableModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_Flags
func callbackCustomTableModele3d187_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.CustomTableModele3d187_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackCustomTableModele3d187_InsertColumns
func callbackCustomTableModele3d187_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_InsertRows
func callbackCustomTableModele3d187_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_MoveColumns
func callbackCustomTableModele3d187_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTableModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_MoveRows
func callbackCustomTableModele3d187_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTableModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_RemoveColumns
func callbackCustomTableModele3d187_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_RemoveRows
func callbackCustomTableModele3d187_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_SetData
func callbackCustomTableModele3d187_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTableModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_SetHeaderData
func callbackCustomTableModele3d187_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTableModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_SetItemData
func callbackCustomTableModele3d187_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
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
		return int8(C.CustomTableModele3d187_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_Submit
func callbackCustomTableModele3d187_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *CustomTableModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_ColumnsAboutToBeInserted
func callbackCustomTableModele3d187_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_ColumnsAboutToBeMoved
func callbackCustomTableModele3d187_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackCustomTableModele3d187_ColumnsAboutToBeRemoved
func callbackCustomTableModele3d187_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_ColumnsInserted
func callbackCustomTableModele3d187_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_ColumnsMoved
func callbackCustomTableModele3d187_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackCustomTableModele3d187_ColumnsRemoved
func callbackCustomTableModele3d187_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_DataChanged
func callbackCustomTableModele3d187_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackCustomTableModele3d187_FetchMore
func callbackCustomTableModele3d187_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewCustomTableModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *CustomTableModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackCustomTableModele3d187_HeaderDataChanged
func callbackCustomTableModele3d187_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_LayoutAboutToBeChanged
func callbackCustomTableModele3d187_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTableModele3d187_LayoutChanged
func callbackCustomTableModele3d187_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTableModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTableModele3d187_ModelAboutToBeReset
func callbackCustomTableModele3d187_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackCustomTableModele3d187_ModelReset
func callbackCustomTableModele3d187_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackCustomTableModele3d187_ResetInternalData
func callbackCustomTableModele3d187_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCustomTableModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *CustomTableModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackCustomTableModele3d187_Revert
func callbackCustomTableModele3d187_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCustomTableModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *CustomTableModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_RevertDefault(ptr.Pointer())
	}
}

//export callbackCustomTableModele3d187_RowsAboutToBeInserted
func callbackCustomTableModele3d187_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackCustomTableModele3d187_RowsAboutToBeMoved
func callbackCustomTableModele3d187_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackCustomTableModele3d187_RowsAboutToBeRemoved
func callbackCustomTableModele3d187_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_RowsInserted
func callbackCustomTableModele3d187_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_RowsMoved
func callbackCustomTableModele3d187_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackCustomTableModele3d187_RowsRemoved
func callbackCustomTableModele3d187_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTableModele3d187_Sort
func callbackCustomTableModele3d187_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewCustomTableModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *CustomTableModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackCustomTableModele3d187_RoleNames
func callbackCustomTableModele3d187_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
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
		}(C.CustomTableModele3d187_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackCustomTableModele3d187_ItemData
func callbackCustomTableModele3d187_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
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
		}(C.CustomTableModele3d187_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackCustomTableModele3d187_MimeData
func callbackCustomTableModele3d187_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
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
		tmpValue := std_core.NewQMimeDataFromPointer(C.CustomTableModele3d187_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
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

//export callbackCustomTableModele3d187_Buddy
func callbackCustomTableModele3d187_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_Parent
func callbackCustomTableModele3d187_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTableModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTableModele3d187_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_Match
func callbackCustomTableModele3d187_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTableModelFromPointer(NewCustomTableModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
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
		}(C.CustomTableModele3d187_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackCustomTableModele3d187_Span
func callbackCustomTableModele3d187_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewCustomTableModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTableModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomTableModele3d187_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_MimeTypes
func callbackCustomTableModele3d187_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewCustomTableModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *CustomTableModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.CustomTableModele3d187_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackCustomTableModele3d187_Data
func callbackCustomTableModele3d187_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTableModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *CustomTableModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModele3d187_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_HeaderData
func callbackCustomTableModele3d187_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTableModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *CustomTableModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTableModele3d187_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTableModele3d187_SupportedDragActions
func callbackCustomTableModele3d187_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *CustomTableModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTableModele3d187_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTableModele3d187_SupportedDropActions
func callbackCustomTableModele3d187_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewCustomTableModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *CustomTableModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTableModele3d187_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTableModele3d187_CanDropMimeData
func callbackCustomTableModele3d187_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_CanFetchMore
func callbackCustomTableModele3d187_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_HasChildren
func callbackCustomTableModele3d187_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTableModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_ColumnCount
func callbackCustomTableModele3d187_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTableModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTableModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTableModele3d187_RowCount
func callbackCustomTableModele3d187_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTableModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTableModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTableModele3d187_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTableModele3d187_Event
func callbackCustomTableModele3d187_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomTableModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_EventFilter
func callbackCustomTableModele3d187_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTableModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomTableModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTableModele3d187_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCustomTableModele3d187_ChildEvent
func callbackCustomTableModele3d187_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomTableModele3d187_ConnectNotify
func callbackCustomTableModele3d187_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTableModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTableModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTableModele3d187_CustomEvent
func callbackCustomTableModele3d187_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomTableModele3d187_DeleteLater
func callbackCustomTableModele3d187_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCustomTableModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomTableModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTableModele3d187_Destroyed
func callbackCustomTableModele3d187_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackCustomTableModele3d187_DisconnectNotify
func callbackCustomTableModele3d187_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTableModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTableModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTableModele3d187_ObjectNameChanged
func callbackCustomTableModele3d187_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackCustomTableModele3d187_TimerEvent
func callbackCustomTableModele3d187_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomTableModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomTableModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTableModele3d187_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
