package src

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

type CustomTreeModel_ITF interface {
	std_core.QAbstractItemModel_ITF
	CustomTreeModel_PTR() *CustomTreeModel
}

func (ptr *CustomTreeModel) CustomTreeModel_PTR() *CustomTreeModel {
	return ptr
}

func (ptr *CustomTreeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (ptr *CustomTreeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemModel_PTR().SetPointer(p)
	}
}

func PointerFromCustomTreeModel(ptr CustomTreeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CustomTreeModel_PTR().Pointer()
	}
	return nil
}

func NewCustomTreeModelFromPointer(ptr unsafe.Pointer) (n *CustomTreeModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CustomTreeModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CustomTreeModel:
			n = deduced

		case *std_core.QAbstractItemModel:
			n = &CustomTreeModel{QAbstractItemModel: *deduced}

		default:
			n = new(CustomTreeModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackCustomTreeModela8f92c_Constructor
func callbackCustomTreeModela8f92c_Constructor(ptr unsafe.Pointer) {
	this := NewCustomTreeModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectRemove(this.remove)
	this.ConnectAdd(this.add)
	this.ConnectEdit(this.edit)
	this.init()
}

//export callbackCustomTreeModela8f92c_Remove
func callbackCustomTreeModela8f92c_Remove(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomTreeModel) ConnectRemove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remove") {
			C.CustomTreeModela8f92c_ConnectRemove(ptr.Pointer())
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

func (ptr *CustomTreeModel) DisconnectRemove() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DisconnectRemove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *CustomTreeModel) Remove() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_Remove(ptr.Pointer())
	}
}

//export callbackCustomTreeModela8f92c_Add
func callbackCustomTreeModela8f92c_Add(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(*TreeItem))(NewTreeItemFromPointer(item))
	}

}

func (ptr *CustomTreeModel) ConnectAdd(f func(item *TreeItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomTreeModela8f92c_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(item *TreeItem) {
				signal.(func(*TreeItem))(item)
				f(item)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *CustomTreeModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *CustomTreeModel) Add(item TreeItem_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_Add(ptr.Pointer(), PointerFromTreeItem(item))
	}
}

//export callbackCustomTreeModela8f92c_Edit
func callbackCustomTreeModela8f92c_Edit(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(firstName), cGoUnpackString(lastName))
	}

}

func (ptr *CustomTreeModel) ConnectEdit(f func(firstName string, lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edit") {
			C.CustomTreeModela8f92c_ConnectEdit(ptr.Pointer())
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

func (ptr *CustomTreeModel) DisconnectEdit() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DisconnectEdit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edit")
	}
}

func (ptr *CustomTreeModel) Edit(firstName string, lastName string) {
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
		C.CustomTreeModela8f92c_Edit(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func CustomTreeModel_QRegisterMetaType() int {
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType()))
}

func (ptr *CustomTreeModel) QRegisterMetaType() int {
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType()))
}

func CustomTreeModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomTreeModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType2(typeNameC)))
}

func CustomTreeModel_QmlRegisterType() int {
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType()))
}

func (ptr *CustomTreeModel) QmlRegisterType() int {
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType()))
}

func CustomTreeModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomTreeModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomTreeModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTreeModela8f92c___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTreeModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.CustomTreeModela8f92c___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTreeModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTreeModela8f92c___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomTreeModela8f92c___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTreeModela8f92c___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTreeModel) __roleNames_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___roleNames_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.CustomTreeModela8f92c___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTreeModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTreeModela8f92c___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomTreeModel) __itemData_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___itemData_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.CustomTreeModela8f92c___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomTreeModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __match_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___match_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomTreeModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomTreeModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomTreeModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomTreeModela8f92c___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomTreeModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTreeModela8f92c___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTreeModel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomTreeModela8f92c___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomTreeModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTreeModela8f92c___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTreeModel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomTreeModela8f92c___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomTreeModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTreeModela8f92c___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTreeModel) __findChildren_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomTreeModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomTreeModela8f92c___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomTreeModel) __children_newList() unsafe.Pointer {
	return C.CustomTreeModela8f92c___children_newList(ptr.Pointer())
}

func NewCustomTreeModel(parent std_core.QObject_ITF) *CustomTreeModel {
	tmpValue := NewCustomTreeModelFromPointer(C.CustomTreeModela8f92c_NewCustomTreeModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomTreeModela8f92c_DestroyCustomTreeModel
func callbackCustomTreeModela8f92c_DestroyCustomTreeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CustomTreeModel"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTreeModelFromPointer(ptr).DestroyCustomTreeModelDefault()
	}
}

func (ptr *CustomTreeModel) ConnectDestroyCustomTreeModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CustomTreeModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~CustomTreeModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CustomTreeModel", f)
		}
	}
}

func (ptr *CustomTreeModel) DisconnectDestroyCustomTreeModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CustomTreeModel")
	}
}

func (ptr *CustomTreeModel) DestroyCustomTreeModel() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DestroyCustomTreeModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomTreeModel) DestroyCustomTreeModelDefault() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DestroyCustomTreeModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTreeModela8f92c_DropMimeData
func callbackCustomTreeModela8f92c_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_InsertColumns
func callbackCustomTreeModela8f92c_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_InsertRows
func callbackCustomTreeModela8f92c_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_MoveColumns
func callbackCustomTreeModela8f92c_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTreeModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_MoveRows
func callbackCustomTreeModela8f92c_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomTreeModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_RemoveColumns
func callbackCustomTreeModela8f92c_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_RemoveRows
func callbackCustomTreeModela8f92c_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_SetData
func callbackCustomTreeModela8f92c_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTreeModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_SetHeaderData
func callbackCustomTreeModela8f92c_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomTreeModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_SetItemData
func callbackCustomTreeModela8f92c_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewCustomTreeModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *CustomTreeModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_Submit
func callbackCustomTreeModela8f92c_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *CustomTreeModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_ColumnsAboutToBeInserted
func callbackCustomTreeModela8f92c_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_ColumnsAboutToBeMoved
func callbackCustomTreeModela8f92c_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackCustomTreeModela8f92c_ColumnsAboutToBeRemoved
func callbackCustomTreeModela8f92c_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_ColumnsInserted
func callbackCustomTreeModela8f92c_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_ColumnsMoved
func callbackCustomTreeModela8f92c_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackCustomTreeModela8f92c_ColumnsRemoved
func callbackCustomTreeModela8f92c_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_DataChanged
func callbackCustomTreeModela8f92c_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackCustomTreeModela8f92c_FetchMore
func callbackCustomTreeModela8f92c_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewCustomTreeModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *CustomTreeModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackCustomTreeModela8f92c_HeaderDataChanged
func callbackCustomTreeModela8f92c_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_LayoutAboutToBeChanged
func callbackCustomTreeModela8f92c_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTreeModela8f92c_LayoutChanged
func callbackCustomTreeModela8f92c_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomTreeModela8f92c_ModelAboutToBeReset
func callbackCustomTreeModela8f92c_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomTreeModela8f92c_ModelReset
func callbackCustomTreeModela8f92c_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomTreeModela8f92c_ResetInternalData
func callbackCustomTreeModela8f92c_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTreeModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *CustomTreeModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackCustomTreeModela8f92c_Revert
func callbackCustomTreeModela8f92c_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTreeModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *CustomTreeModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_RevertDefault(ptr.Pointer())
	}
}

//export callbackCustomTreeModela8f92c_RowsAboutToBeInserted
func callbackCustomTreeModela8f92c_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackCustomTreeModela8f92c_RowsAboutToBeMoved
func callbackCustomTreeModela8f92c_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackCustomTreeModela8f92c_RowsAboutToBeRemoved
func callbackCustomTreeModela8f92c_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_RowsInserted
func callbackCustomTreeModela8f92c_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_RowsMoved
func callbackCustomTreeModela8f92c_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackCustomTreeModela8f92c_RowsRemoved
func callbackCustomTreeModela8f92c_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomTreeModela8f92c_Sort
func callbackCustomTreeModela8f92c_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewCustomTreeModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *CustomTreeModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackCustomTreeModela8f92c_RoleNames
func callbackCustomTreeModela8f92c_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewCustomTreeModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTreeModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.CustomTreeModela8f92c_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackCustomTreeModela8f92c_ItemData
func callbackCustomTreeModela8f92c_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__itemData_newList())
		for k, v := range NewCustomTreeModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTreeModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.CustomTreeModela8f92c_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackCustomTreeModela8f92c_MimeData
func callbackCustomTreeModela8f92c_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewCustomTreeModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewCustomTreeModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *CustomTreeModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.CustomTreeModela8f92c_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__mimeData_indexes_newList())
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

//export callbackCustomTreeModela8f92c_Buddy
func callbackCustomTreeModela8f92c_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTreeModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTreeModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_Index
func callbackCustomTreeModela8f92c_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTreeModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *CustomTreeModel) Index(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_Parent
func callbackCustomTreeModela8f92c_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTreeModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTreeModel) Parent(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_Parent(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_Sibling
func callbackCustomTreeModela8f92c_Sibling(ptr unsafe.Pointer, row C.int, column C.int, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomTreeModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTreeModel) SiblingDefault(row int, column int, index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomTreeModela8f92c_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_Match
func callbackCustomTreeModela8f92c_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomTreeModelFromPointer(NewCustomTreeModelFromPointer(nil).__match_newList())
		for _, v := range NewCustomTreeModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomTreeModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomTreeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.CustomTreeModela8f92c_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackCustomTreeModela8f92c_Span
func callbackCustomTreeModela8f92c_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewCustomTreeModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTreeModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomTreeModela8f92c_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_MimeTypes
func callbackCustomTreeModela8f92c_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewCustomTreeModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *CustomTreeModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.CustomTreeModela8f92c_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackCustomTreeModela8f92c_Data
func callbackCustomTreeModela8f92c_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTreeModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *CustomTreeModel) Data(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTreeModela8f92c_Data(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomTreeModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTreeModela8f92c_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_HeaderData
func callbackCustomTreeModela8f92c_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomTreeModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *CustomTreeModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomTreeModela8f92c_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomTreeModela8f92c_SupportedDragActions
func callbackCustomTreeModela8f92c_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomTreeModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *CustomTreeModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTreeModela8f92c_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTreeModela8f92c_SupportedDropActions
func callbackCustomTreeModela8f92c_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomTreeModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *CustomTreeModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomTreeModela8f92c_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomTreeModela8f92c_Flags
func callbackCustomTreeModela8f92c_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewCustomTreeModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomTreeModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.CustomTreeModela8f92c_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackCustomTreeModela8f92c_CanDropMimeData
func callbackCustomTreeModela8f92c_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_CanFetchMore
func callbackCustomTreeModela8f92c_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_HasChildren
func callbackCustomTreeModela8f92c_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomTreeModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_ColumnCount
func callbackCustomTreeModela8f92c_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTreeModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTreeModel) ColumnCount(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_ColumnCount(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *CustomTreeModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTreeModela8f92c_RowCount
func callbackCustomTreeModela8f92c_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomTreeModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomTreeModel) RowCount(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_RowCount(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *CustomTreeModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomTreeModela8f92c_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomTreeModela8f92c_Event
func callbackCustomTreeModela8f92c_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomTreeModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_EventFilter
func callbackCustomTreeModela8f92c_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomTreeModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomTreeModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomTreeModela8f92c_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCustomTreeModela8f92c_ChildEvent
func callbackCustomTreeModela8f92c_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomTreeModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomTreeModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomTreeModela8f92c_ConnectNotify
func callbackCustomTreeModela8f92c_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTreeModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTreeModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTreeModela8f92c_CustomEvent
func callbackCustomTreeModela8f92c_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomTreeModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomTreeModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomTreeModela8f92c_DeleteLater
func callbackCustomTreeModela8f92c_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCustomTreeModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomTreeModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomTreeModela8f92c_Destroyed
func callbackCustomTreeModela8f92c_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCustomTreeModela8f92c_DisconnectNotify
func callbackCustomTreeModela8f92c_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomTreeModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomTreeModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomTreeModela8f92c_ObjectNameChanged
func callbackCustomTreeModela8f92c_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCustomTreeModela8f92c_TimerEvent
func callbackCustomTreeModela8f92c_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomTreeModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomTreeModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomTreeModela8f92c_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TreeItem_ITF interface {
	std_core.QObject_ITF
	TreeItem_PTR() *TreeItem
}

func (ptr *TreeItem) TreeItem_PTR() *TreeItem {
	return ptr
}

func (ptr *TreeItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *TreeItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTreeItem(ptr TreeItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TreeItem_PTR().Pointer()
	}
	return nil
}

func NewTreeItemFromPointer(ptr unsafe.Pointer) (n *TreeItem) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TreeItem)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TreeItem:
			n = deduced

		case *std_core.QObject:
			n = &TreeItem{QObject: *deduced}

		default:
			n = new(TreeItem)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTreeItema8f92c_Constructor
func callbackTreeItema8f92c_Constructor(ptr unsafe.Pointer) {
	this := NewTreeItemFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func TreeItem_QRegisterMetaType() int {
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QRegisterMetaType()))
}

func (ptr *TreeItem) QRegisterMetaType() int {
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QRegisterMetaType()))
}

func TreeItem_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QRegisterMetaType2(typeNameC)))
}

func (ptr *TreeItem) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QRegisterMetaType2(typeNameC)))
}

func TreeItem_QmlRegisterType() int {
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QmlRegisterType()))
}

func (ptr *TreeItem) QmlRegisterType() int {
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QmlRegisterType()))
}

func TreeItem_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TreeItem) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TreeItema8f92c_TreeItema8f92c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TreeItem) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TreeItema8f92c___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TreeItem) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TreeItem) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TreeItema8f92c___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TreeItem) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TreeItema8f92c___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TreeItem) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TreeItem) __findChildren_newList2() unsafe.Pointer {
	return C.TreeItema8f92c___findChildren_newList2(ptr.Pointer())
}

func (ptr *TreeItem) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TreeItema8f92c___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TreeItem) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TreeItem) __findChildren_newList3() unsafe.Pointer {
	return C.TreeItema8f92c___findChildren_newList3(ptr.Pointer())
}

func (ptr *TreeItem) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TreeItema8f92c___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TreeItem) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TreeItem) __findChildren_newList() unsafe.Pointer {
	return C.TreeItema8f92c___findChildren_newList(ptr.Pointer())
}

func (ptr *TreeItem) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TreeItema8f92c___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TreeItem) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TreeItem) __children_newList() unsafe.Pointer {
	return C.TreeItema8f92c___children_newList(ptr.Pointer())
}

func NewTreeItem(parent std_core.QObject_ITF) *TreeItem {
	tmpValue := NewTreeItemFromPointer(C.TreeItema8f92c_NewTreeItem(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTreeItema8f92c_DestroyTreeItem
func callbackTreeItema8f92c_DestroyTreeItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TreeItem"); signal != nil {
		signal.(func())()
	} else {
		NewTreeItemFromPointer(ptr).DestroyTreeItemDefault()
	}
}

func (ptr *TreeItem) ConnectDestroyTreeItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TreeItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~TreeItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TreeItem", f)
		}
	}
}

func (ptr *TreeItem) DisconnectDestroyTreeItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TreeItem")
	}
}

func (ptr *TreeItem) DestroyTreeItem() {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_DestroyTreeItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TreeItem) DestroyTreeItemDefault() {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_DestroyTreeItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTreeItema8f92c_Event
func callbackTreeItema8f92c_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTreeItemFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TreeItem) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TreeItema8f92c_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTreeItema8f92c_EventFilter
func callbackTreeItema8f92c_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTreeItemFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TreeItem) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TreeItema8f92c_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTreeItema8f92c_ChildEvent
func callbackTreeItema8f92c_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTreeItemFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TreeItem) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTreeItema8f92c_ConnectNotify
func callbackTreeItema8f92c_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTreeItemFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TreeItem) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTreeItema8f92c_CustomEvent
func callbackTreeItema8f92c_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTreeItemFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TreeItem) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTreeItema8f92c_DeleteLater
func callbackTreeItema8f92c_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTreeItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TreeItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTreeItema8f92c_Destroyed
func callbackTreeItema8f92c_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTreeItema8f92c_DisconnectNotify
func callbackTreeItema8f92c_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTreeItemFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TreeItem) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTreeItema8f92c_ObjectNameChanged
func callbackTreeItema8f92c_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTreeItema8f92c_TimerEvent
func callbackTreeItema8f92c_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTreeItemFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TreeItem) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TreeItema8f92c_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
