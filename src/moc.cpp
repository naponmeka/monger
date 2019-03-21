

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>


class CustomTreeModela8f92c: public QAbstractItemModel
{
Q_OBJECT
public:
	CustomTreeModela8f92c(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType();CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaTypes();callbackCustomTreeModela8f92c_Constructor(this);};
	void Signal_Remove() { callbackCustomTreeModela8f92c_Remove(this); };
	void Signal_Add(TreeItema8f92c* item) { callbackCustomTreeModela8f92c_Add(this, item); };
	void Signal_Edit(QString firstName, QString lastName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackCustomTreeModela8f92c_Edit(this, firstNamePacked, lastNamePacked); };
	 ~CustomTreeModela8f92c() { callbackCustomTreeModela8f92c_DestroyCustomTreeModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomTreeModela8f92c_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModela8f92c_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModela8f92c_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModela8f92c_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModela8f92c_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModela8f92c_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModela8f92c_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomTreeModela8f92c_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomTreeModela8f92c_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomTreeModela8f92c_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomTreeModela8f92c_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomTreeModela8f92c_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomTreeModela8f92c_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomTreeModela8f92c_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomTreeModela8f92c_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomTreeModela8f92c_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModela8f92c_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModela8f92c_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomTreeModela8f92c_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomTreeModela8f92c_ModelReset(this); };
	void resetInternalData() { callbackCustomTreeModela8f92c_ResetInternalData(this); };
	void revert() { callbackCustomTreeModela8f92c_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomTreeModela8f92c_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomTreeModela8f92c_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomTreeModela8f92c_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModela8f92c_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomTreeModela8f92c_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomTreeModela8f92c_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomTreeModela8f92c_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomTreeModela8f92c_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModela8f92c_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomTreeModela8f92c_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModela8f92c_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModela8f92c_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomTreeModela8f92c_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomTreeModela8f92c_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomTreeModela8f92c_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModela8f92c_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModela8f92c_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModela8f92c_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModela8f92c_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomTreeModela8f92c_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomTreeModela8f92c_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomTreeModela8f92c_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomTreeModela8f92c_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomTreeModela8f92c_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomTreeModela8f92c_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomTreeModela8f92c_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomTreeModela8f92c_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomTreeModela8f92c_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomTreeModela8f92c_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomTreeModela8f92c_CustomEvent(this, event); };
	void deleteLater() { callbackCustomTreeModela8f92c_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomTreeModela8f92c_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomTreeModela8f92c_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomTreeModela8f92c_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomTreeModela8f92c_TimerEvent(this, event); };
signals:
	void remove();
	void add(TreeItema8f92c* item);
	void edit(QString firstName, QString lastName);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomTreeModela8f92c*)


void CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaTypes() {
}

class TreeItema8f92c: public QObject
{
Q_OBJECT
public:
	TreeItema8f92c(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TreeItema8f92c_TreeItema8f92c_QRegisterMetaType();TreeItema8f92c_TreeItema8f92c_QRegisterMetaTypes();callbackTreeItema8f92c_Constructor(this);};
	 ~TreeItema8f92c() { callbackTreeItema8f92c_DestroyTreeItem(this); };
	bool event(QEvent * e) { return callbackTreeItema8f92c_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTreeItema8f92c_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTreeItema8f92c_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTreeItema8f92c_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTreeItema8f92c_CustomEvent(this, event); };
	void deleteLater() { callbackTreeItema8f92c_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTreeItema8f92c_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTreeItema8f92c_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTreeItema8f92c_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTreeItema8f92c_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(TreeItema8f92c*)


void TreeItema8f92c_TreeItema8f92c_QRegisterMetaTypes() {
}

void CustomTreeModela8f92c_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)()>(&CustomTreeModela8f92c::remove), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)()>(&CustomTreeModela8f92c::Signal_Remove));
}

void CustomTreeModela8f92c_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)()>(&CustomTreeModela8f92c::remove), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)()>(&CustomTreeModela8f92c::Signal_Remove));
}

void CustomTreeModela8f92c_Remove(void* ptr)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->remove();
}

void CustomTreeModela8f92c_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(TreeItema8f92c*)>(&CustomTreeModela8f92c::add), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(TreeItema8f92c*)>(&CustomTreeModela8f92c::Signal_Add));
}

void CustomTreeModela8f92c_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(TreeItema8f92c*)>(&CustomTreeModela8f92c::add), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(TreeItema8f92c*)>(&CustomTreeModela8f92c::Signal_Add));
}

void CustomTreeModela8f92c_Add(void* ptr, void* item)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->add(static_cast<TreeItema8f92c*>(item));
}

void CustomTreeModela8f92c_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(QString, QString)>(&CustomTreeModela8f92c::edit), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(QString, QString)>(&CustomTreeModela8f92c::Signal_Edit));
}

void CustomTreeModela8f92c_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(QString, QString)>(&CustomTreeModela8f92c::edit), static_cast<CustomTreeModela8f92c*>(ptr), static_cast<void (CustomTreeModela8f92c::*)(QString, QString)>(&CustomTreeModela8f92c::Signal_Edit));
}

void CustomTreeModela8f92c_Edit(void* ptr, struct Moc_PackedString firstName, struct Moc_PackedString lastName)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->edit(QString::fromUtf8(firstName.data, firstName.len), QString::fromUtf8(lastName.data, lastName.len));
}

int CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType()
{
	return qRegisterMetaType<CustomTreeModela8f92c*>();
}

int CustomTreeModela8f92c_CustomTreeModela8f92c_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomTreeModela8f92c*>(const_cast<const char*>(typeName));
}

int CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModela8f92c>();
#else
	return 0;
#endif
}

int CustomTreeModela8f92c_CustomTreeModela8f92c_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModela8f92c>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CustomTreeModela8f92c___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModela8f92c___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModela8f92c___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModela8f92c___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModela8f92c___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModela8f92c___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModela8f92c___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModela8f92c___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomTreeModela8f92c___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModela8f92c___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModela8f92c___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModela8f92c___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModela8f92c___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomTreeModela8f92c___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomTreeModela8f92c___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModela8f92c___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModela8f92c___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModela8f92c___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModela8f92c___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModela8f92c___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModela8f92c___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModela8f92c___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModela8f92c___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModela8f92c___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModela8f92c_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModela8f92c_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModela8f92c_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModela8f92c_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomTreeModela8f92c___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModela8f92c___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomTreeModela8f92c___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomTreeModela8f92c___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModela8f92c___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModela8f92c___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModela8f92c___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModela8f92c___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModela8f92c___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModela8f92c___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModela8f92c___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModela8f92c___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomTreeModela8f92c_NewCustomTreeModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModela8f92c(static_cast<QWindow*>(parent));
	} else {
		return new CustomTreeModela8f92c(static_cast<QObject*>(parent));
	}
}

void CustomTreeModela8f92c_DestroyCustomTreeModel(void* ptr)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->~CustomTreeModela8f92c();
}

void CustomTreeModela8f92c_DestroyCustomTreeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomTreeModela8f92c_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModela8f92c_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModela8f92c_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomTreeModela8f92c_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomTreeModela8f92c_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomTreeModela8f92c_SubmitDefault(void* ptr)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::submit();
}

void CustomTreeModela8f92c_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomTreeModela8f92c_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::resetInternalData();
}

void CustomTreeModela8f92c_RevertDefault(void* ptr)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::revert();
}

void CustomTreeModela8f92c_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomTreeModela8f92c_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomTreeModela8f92c_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModela8f92c_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomTreeModela8f92c_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomTreeModela8f92c_Index(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModela8f92c_IndexDefault(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModela8f92c_Parent(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModela8f92c_ParentDefault(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModela8f92c_SiblingDefault(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomTreeModela8f92c_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModela8f92c_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomTreeModela8f92c_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tbaf865 = static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tbaf865.prepend("WHITESPACE").constData()+10), tbaf865.size()-10 }; });
}

void* CustomTreeModela8f92c_Data(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModela8f92c_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModela8f92c_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomTreeModela8f92c_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::supportedDragActions();
}

long long CustomTreeModela8f92c_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::supportedDropActions();
}

long long CustomTreeModela8f92c_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::flags(*static_cast<QModelIndex*>(index));
}

char CustomTreeModela8f92c_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomTreeModela8f92c_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}



int CustomTreeModela8f92c_ColumnCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModela8f92c_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModela8f92c_RowCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModela8f92c_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomTreeModela8f92c_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
}

char CustomTreeModela8f92c_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomTreeModela8f92c_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomTreeModela8f92c_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModela8f92c_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
}

void CustomTreeModela8f92c_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::deleteLater();
}

void CustomTreeModela8f92c_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModela8f92c_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModela8f92c*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

int TreeItema8f92c_TreeItema8f92c_QRegisterMetaType()
{
	return qRegisterMetaType<TreeItema8f92c*>();
}

int TreeItema8f92c_TreeItema8f92c_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TreeItema8f92c*>(const_cast<const char*>(typeName));
}

int TreeItema8f92c_TreeItema8f92c_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItema8f92c>();
#else
	return 0;
#endif
}

int TreeItema8f92c_TreeItema8f92c_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItema8f92c>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TreeItema8f92c___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TreeItema8f92c___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TreeItema8f92c___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TreeItema8f92c___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItema8f92c___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItema8f92c___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItema8f92c___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItema8f92c___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItema8f92c___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItema8f92c___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItema8f92c___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItema8f92c___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItema8f92c___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItema8f92c___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItema8f92c___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TreeItema8f92c_NewTreeItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItema8f92c(static_cast<QWindow*>(parent));
	} else {
		return new TreeItema8f92c(static_cast<QObject*>(parent));
	}
}

void TreeItema8f92c_DestroyTreeItem(void* ptr)
{
	static_cast<TreeItema8f92c*>(ptr)->~TreeItema8f92c();
}

void TreeItema8f92c_DestroyTreeItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TreeItema8f92c_EventDefault(void* ptr, void* e)
{
	return static_cast<TreeItema8f92c*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TreeItema8f92c_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TreeItema8f92c*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TreeItema8f92c_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TreeItema8f92c_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItema8f92c_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TreeItema8f92c_DeleteLaterDefault(void* ptr)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::deleteLater();
}

void TreeItema8f92c_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItema8f92c_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TreeItema8f92c*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
