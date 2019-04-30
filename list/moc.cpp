

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
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


class CustomListModelbd2412: public QAbstractListModel
{
Q_OBJECT
public:
	CustomListModelbd2412(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomListModelbd2412_CustomListModelbd2412_QRegisterMetaType();CustomListModelbd2412_CustomListModelbd2412_QRegisterMetaTypes();callbackCustomListModelbd2412_Constructor(this);};
	void Signal_Remove() { callbackCustomListModelbd2412_Remove(this); };
	void Signal_Add(quintptr item) { callbackCustomListModelbd2412_Add(this, item); };
	void Signal_Edit(QString Name, QString URI) { QByteArray t709a23 = Name.toUtf8(); Moc_PackedString NamePacked = { const_cast<char*>(t709a23.prepend("WHITESPACE").constData()+10), t709a23.size()-10 };QByteArray tcb371c = URI.toUtf8(); Moc_PackedString URIPacked = { const_cast<char*>(tcb371c.prepend("WHITESPACE").constData()+10), tcb371c.size()-10 };callbackCustomListModelbd2412_Edit(this, NamePacked, URIPacked); };
	void Signal_RemoveAt(QString iStr) { QByteArray tb2c797 = iStr.toUtf8(); Moc_PackedString iStrPacked = { const_cast<char*>(tb2c797.prepend("WHITESPACE").constData()+10), tb2c797.size()-10 };callbackCustomListModelbd2412_RemoveAt(this, iStrPacked); };
	 ~CustomListModelbd2412() { callbackCustomListModelbd2412_DestroyCustomListModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomListModelbd2412_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomListModelbd2412_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackCustomListModelbd2412_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomListModelbd2412_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModelbd2412_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModelbd2412_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModelbd2412_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModelbd2412_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModelbd2412_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModelbd2412_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomListModelbd2412_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomListModelbd2412_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomListModelbd2412_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomListModelbd2412_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomListModelbd2412_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomListModelbd2412_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomListModelbd2412_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomListModelbd2412_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomListModelbd2412_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModelbd2412_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModelbd2412_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomListModelbd2412_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomListModelbd2412_ModelReset(this); };
	void resetInternalData() { callbackCustomListModelbd2412_ResetInternalData(this); };
	void revert() { callbackCustomListModelbd2412_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomListModelbd2412_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomListModelbd2412_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomListModelbd2412_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModelbd2412_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomListModelbd2412_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomListModelbd2412_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomListModelbd2412_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomListModelbd2412_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModelbd2412_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModelbd2412_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomListModelbd2412_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomListModelbd2412_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomListModelbd2412_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomListModelbd2412_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomListModelbd2412_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomListModelbd2412_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomListModelbd2412_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomListModelbd2412_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomListModelbd2412_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomListModelbd2412_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomListModelbd2412_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomListModelbd2412_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomListModelbd2412_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomListModelbd2412_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomListModelbd2412_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomListModelbd2412_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomListModelbd2412_CustomEvent(this, event); };
	void deleteLater() { callbackCustomListModelbd2412_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomListModelbd2412_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomListModelbd2412_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomListModelbd2412_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomListModelbd2412_TimerEvent(this, event); };
signals:
	void remove();
	void add(quintptr item);
	void edit(QString Name, QString URI);
	void removeAt(QString iStr);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomListModelbd2412*)


void CustomListModelbd2412_CustomListModelbd2412_QRegisterMetaTypes() {
}

void CustomListModelbd2412_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)()>(&CustomListModelbd2412::remove), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)()>(&CustomListModelbd2412::Signal_Remove));
}

void CustomListModelbd2412_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)()>(&CustomListModelbd2412::remove), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)()>(&CustomListModelbd2412::Signal_Remove));
}

void CustomListModelbd2412_Remove(void* ptr)
{
	static_cast<CustomListModelbd2412*>(ptr)->remove();
}

void CustomListModelbd2412_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(quintptr)>(&CustomListModelbd2412::add), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(quintptr)>(&CustomListModelbd2412::Signal_Add));
}

void CustomListModelbd2412_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(quintptr)>(&CustomListModelbd2412::add), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(quintptr)>(&CustomListModelbd2412::Signal_Add));
}

void CustomListModelbd2412_Add(void* ptr, uintptr_t item)
{
	static_cast<CustomListModelbd2412*>(ptr)->add(item);
}

void CustomListModelbd2412_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString, QString)>(&CustomListModelbd2412::edit), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString, QString)>(&CustomListModelbd2412::Signal_Edit));
}

void CustomListModelbd2412_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString, QString)>(&CustomListModelbd2412::edit), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString, QString)>(&CustomListModelbd2412::Signal_Edit));
}

void CustomListModelbd2412_Edit(void* ptr, struct Moc_PackedString Name, struct Moc_PackedString URI)
{
	static_cast<CustomListModelbd2412*>(ptr)->edit(QString::fromUtf8(Name.data, Name.len), QString::fromUtf8(URI.data, URI.len));
}

void CustomListModelbd2412_ConnectRemoveAt(void* ptr)
{
	QObject::connect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString)>(&CustomListModelbd2412::removeAt), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString)>(&CustomListModelbd2412::Signal_RemoveAt));
}

void CustomListModelbd2412_DisconnectRemoveAt(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString)>(&CustomListModelbd2412::removeAt), static_cast<CustomListModelbd2412*>(ptr), static_cast<void (CustomListModelbd2412::*)(QString)>(&CustomListModelbd2412::Signal_RemoveAt));
}

void CustomListModelbd2412_RemoveAt(void* ptr, struct Moc_PackedString iStr)
{
	static_cast<CustomListModelbd2412*>(ptr)->removeAt(QString::fromUtf8(iStr.data, iStr.len));
}

int CustomListModelbd2412_CustomListModelbd2412_QRegisterMetaType()
{
	return qRegisterMetaType<CustomListModelbd2412*>();
}

int CustomListModelbd2412_CustomListModelbd2412_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomListModelbd2412*>(const_cast<const char*>(typeName));
}

int CustomListModelbd2412_CustomListModelbd2412_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModelbd2412>();
#else
	return 0;
#endif
}

int CustomListModelbd2412_CustomListModelbd2412_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModelbd2412>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int CustomListModelbd2412_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModelbd2412_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModelbd2412_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModelbd2412___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModelbd2412___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModelbd2412___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModelbd2412___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModelbd2412___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModelbd2412___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModelbd2412___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModelbd2412___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomListModelbd2412___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomListModelbd2412___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModelbd2412___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModelbd2412___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModelbd2412___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModelbd2412___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomListModelbd2412___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomListModelbd2412___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModelbd2412___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModelbd2412___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModelbd2412___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModelbd2412___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModelbd2412___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModelbd2412___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModelbd2412___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModelbd2412___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModelbd2412___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModelbd2412_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModelbd2412_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModelbd2412_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModelbd2412___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModelbd2412___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomListModelbd2412___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomListModelbd2412___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModelbd2412___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModelbd2412___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModelbd2412___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModelbd2412___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModelbd2412___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModelbd2412___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModelbd2412___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModelbd2412___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomListModelbd2412_NewCustomListModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModelbd2412(static_cast<QWindow*>(parent));
	} else {
		return new CustomListModelbd2412(static_cast<QObject*>(parent));
	}
}

void CustomListModelbd2412_DestroyCustomListModel(void* ptr)
{
	static_cast<CustomListModelbd2412*>(ptr)->~CustomListModelbd2412();
}

void CustomListModelbd2412_DestroyCustomListModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomListModelbd2412_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* CustomListModelbd2412_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* CustomListModelbd2412_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long CustomListModelbd2412_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char CustomListModelbd2412_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModelbd2412_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModelbd2412_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomListModelbd2412_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomListModelbd2412_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomListModelbd2412_SubmitDefault(void* ptr)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::submit();
}

void CustomListModelbd2412_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomListModelbd2412_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::resetInternalData();
}

void CustomListModelbd2412_RevertDefault(void* ptr)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::revert();
}

void CustomListModelbd2412_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomListModelbd2412_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomListModelbd2412_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModelbd2412_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomListModelbd2412_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomListModelbd2412_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomListModelbd2412_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModelbd2412_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomListModelbd2412_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t9c6749 = static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t9c6749.prepend("WHITESPACE").constData()+10), t9c6749.size()-10 }; });
}

void* CustomListModelbd2412_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomListModelbd2412_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomListModelbd2412_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long CustomListModelbd2412_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::supportedDropActions();
}

char CustomListModelbd2412_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomListModelbd2412_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int CustomListModelbd2412_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int CustomListModelbd2412_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomListModelbd2412_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char CustomListModelbd2412_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomListModelbd2412_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomListModelbd2412_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModelbd2412_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void CustomListModelbd2412_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::deleteLater();
}

void CustomListModelbd2412_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModelbd2412_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModelbd2412*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
