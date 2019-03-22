

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


class CustomTreeModel92ca9a: public QAbstractItemModel
{
Q_OBJECT
public:
	CustomTreeModel92ca9a(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomTreeModel92ca9a_CustomTreeModel92ca9a_QRegisterMetaType();CustomTreeModel92ca9a_CustomTreeModel92ca9a_QRegisterMetaTypes();callbackCustomTreeModel92ca9a_Constructor(this);};
	void Signal_Remove() { callbackCustomTreeModel92ca9a_Remove(this); };
	void Signal_Add(TreeItem92ca9a* item) { callbackCustomTreeModel92ca9a_Add(this, item); };
	void Signal_Edit(QString firstName, QString lastName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackCustomTreeModel92ca9a_Edit(this, firstNamePacked, lastNamePacked); };
	 ~CustomTreeModel92ca9a() { callbackCustomTreeModel92ca9a_DestroyCustomTreeModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomTreeModel92ca9a_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModel92ca9a_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModel92ca9a_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModel92ca9a_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModel92ca9a_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModel92ca9a_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModel92ca9a_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomTreeModel92ca9a_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomTreeModel92ca9a_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomTreeModel92ca9a_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomTreeModel92ca9a_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomTreeModel92ca9a_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomTreeModel92ca9a_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomTreeModel92ca9a_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomTreeModel92ca9a_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomTreeModel92ca9a_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModel92ca9a_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModel92ca9a_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomTreeModel92ca9a_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomTreeModel92ca9a_ModelReset(this); };
	void resetInternalData() { callbackCustomTreeModel92ca9a_ResetInternalData(this); };
	void revert() { callbackCustomTreeModel92ca9a_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomTreeModel92ca9a_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomTreeModel92ca9a_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomTreeModel92ca9a_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModel92ca9a_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomTreeModel92ca9a_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomTreeModel92ca9a_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomTreeModel92ca9a_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomTreeModel92ca9a_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModel92ca9a_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomTreeModel92ca9a_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModel92ca9a_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModel92ca9a_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomTreeModel92ca9a_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomTreeModel92ca9a_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomTreeModel92ca9a_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModel92ca9a_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModel92ca9a_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModel92ca9a_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModel92ca9a_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomTreeModel92ca9a_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomTreeModel92ca9a_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomTreeModel92ca9a_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomTreeModel92ca9a_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomTreeModel92ca9a_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomTreeModel92ca9a_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomTreeModel92ca9a_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomTreeModel92ca9a_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomTreeModel92ca9a_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomTreeModel92ca9a_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomTreeModel92ca9a_CustomEvent(this, event); };
	void deleteLater() { callbackCustomTreeModel92ca9a_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomTreeModel92ca9a_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomTreeModel92ca9a_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomTreeModel92ca9a_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomTreeModel92ca9a_TimerEvent(this, event); };
signals:
	void remove();
	void add(TreeItem92ca9a* item);
	void edit(QString firstName, QString lastName);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomTreeModel92ca9a*)


void CustomTreeModel92ca9a_CustomTreeModel92ca9a_QRegisterMetaTypes() {
}

class TreeItem92ca9a: public QObject
{
Q_OBJECT
public:
	TreeItem92ca9a(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TreeItem92ca9a_TreeItem92ca9a_QRegisterMetaType();TreeItem92ca9a_TreeItem92ca9a_QRegisterMetaTypes();callbackTreeItem92ca9a_Constructor(this);};
	 ~TreeItem92ca9a() { callbackTreeItem92ca9a_DestroyTreeItem(this); };
	bool event(QEvent * e) { return callbackTreeItem92ca9a_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTreeItem92ca9a_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTreeItem92ca9a_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTreeItem92ca9a_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTreeItem92ca9a_CustomEvent(this, event); };
	void deleteLater() { callbackTreeItem92ca9a_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTreeItem92ca9a_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTreeItem92ca9a_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTreeItem92ca9a_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTreeItem92ca9a_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(TreeItem92ca9a*)


void TreeItem92ca9a_TreeItem92ca9a_QRegisterMetaTypes() {
}

void CustomTreeModel92ca9a_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)()>(&CustomTreeModel92ca9a::remove), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)()>(&CustomTreeModel92ca9a::Signal_Remove));
}

void CustomTreeModel92ca9a_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)()>(&CustomTreeModel92ca9a::remove), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)()>(&CustomTreeModel92ca9a::Signal_Remove));
}

void CustomTreeModel92ca9a_Remove(void* ptr)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->remove();
}

void CustomTreeModel92ca9a_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(TreeItem92ca9a*)>(&CustomTreeModel92ca9a::add), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(TreeItem92ca9a*)>(&CustomTreeModel92ca9a::Signal_Add));
}

void CustomTreeModel92ca9a_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(TreeItem92ca9a*)>(&CustomTreeModel92ca9a::add), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(TreeItem92ca9a*)>(&CustomTreeModel92ca9a::Signal_Add));
}

void CustomTreeModel92ca9a_Add(void* ptr, void* item)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->add(static_cast<TreeItem92ca9a*>(item));
}

void CustomTreeModel92ca9a_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(QString, QString)>(&CustomTreeModel92ca9a::edit), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(QString, QString)>(&CustomTreeModel92ca9a::Signal_Edit));
}

void CustomTreeModel92ca9a_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(QString, QString)>(&CustomTreeModel92ca9a::edit), static_cast<CustomTreeModel92ca9a*>(ptr), static_cast<void (CustomTreeModel92ca9a::*)(QString, QString)>(&CustomTreeModel92ca9a::Signal_Edit));
}

void CustomTreeModel92ca9a_Edit(void* ptr, struct Moc_PackedString firstName, struct Moc_PackedString lastName)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->edit(QString::fromUtf8(firstName.data, firstName.len), QString::fromUtf8(lastName.data, lastName.len));
}

int CustomTreeModel92ca9a_CustomTreeModel92ca9a_QRegisterMetaType()
{
	return qRegisterMetaType<CustomTreeModel92ca9a*>();
}

int CustomTreeModel92ca9a_CustomTreeModel92ca9a_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomTreeModel92ca9a*>(const_cast<const char*>(typeName));
}

int CustomTreeModel92ca9a_CustomTreeModel92ca9a_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModel92ca9a>();
#else
	return 0;
#endif
}

int CustomTreeModel92ca9a_CustomTreeModel92ca9a_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModel92ca9a>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CustomTreeModel92ca9a___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModel92ca9a___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModel92ca9a___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModel92ca9a___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModel92ca9a___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModel92ca9a___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModel92ca9a___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModel92ca9a___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomTreeModel92ca9a___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModel92ca9a___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModel92ca9a___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModel92ca9a___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModel92ca9a___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomTreeModel92ca9a___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomTreeModel92ca9a___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModel92ca9a___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModel92ca9a___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModel92ca9a___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModel92ca9a___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModel92ca9a___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModel92ca9a___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModel92ca9a___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModel92ca9a___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModel92ca9a___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModel92ca9a_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModel92ca9a_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModel92ca9a_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModel92ca9a_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModel92ca9a_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModel92ca9a_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomTreeModel92ca9a___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModel92ca9a___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomTreeModel92ca9a___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomTreeModel92ca9a___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModel92ca9a___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModel92ca9a___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModel92ca9a___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModel92ca9a___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModel92ca9a___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModel92ca9a___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModel92ca9a___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModel92ca9a___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomTreeModel92ca9a_NewCustomTreeModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModel92ca9a(static_cast<QWindow*>(parent));
	} else {
		return new CustomTreeModel92ca9a(static_cast<QObject*>(parent));
	}
}

void CustomTreeModel92ca9a_DestroyCustomTreeModel(void* ptr)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->~CustomTreeModel92ca9a();
}

void CustomTreeModel92ca9a_DestroyCustomTreeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomTreeModel92ca9a_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModel92ca9a_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModel92ca9a_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomTreeModel92ca9a_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomTreeModel92ca9a_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomTreeModel92ca9a_SubmitDefault(void* ptr)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::submit();
}

void CustomTreeModel92ca9a_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomTreeModel92ca9a_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::resetInternalData();
}

void CustomTreeModel92ca9a_RevertDefault(void* ptr)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::revert();
}

void CustomTreeModel92ca9a_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomTreeModel92ca9a_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomTreeModel92ca9a_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModel92ca9a_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomTreeModel92ca9a_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomTreeModel92ca9a_Index(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModel92ca9a_IndexDefault(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModel92ca9a_Parent(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModel92ca9a_ParentDefault(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModel92ca9a_SiblingDefault(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomTreeModel92ca9a_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModel92ca9a_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomTreeModel92ca9a_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tbaf865 = static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tbaf865.prepend("WHITESPACE").constData()+10), tbaf865.size()-10 }; });
}

void* CustomTreeModel92ca9a_Data(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModel92ca9a_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModel92ca9a_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomTreeModel92ca9a_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::supportedDragActions();
}

long long CustomTreeModel92ca9a_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::supportedDropActions();
}

long long CustomTreeModel92ca9a_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::flags(*static_cast<QModelIndex*>(index));
}

char CustomTreeModel92ca9a_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomTreeModel92ca9a_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}



int CustomTreeModel92ca9a_ColumnCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModel92ca9a_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModel92ca9a_RowCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModel92ca9a_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomTreeModel92ca9a_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
}

char CustomTreeModel92ca9a_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomTreeModel92ca9a_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomTreeModel92ca9a_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModel92ca9a_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
}

void CustomTreeModel92ca9a_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::deleteLater();
}

void CustomTreeModel92ca9a_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModel92ca9a_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModel92ca9a*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

int TreeItem92ca9a_TreeItem92ca9a_QRegisterMetaType()
{
	return qRegisterMetaType<TreeItem92ca9a*>();
}

int TreeItem92ca9a_TreeItem92ca9a_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TreeItem92ca9a*>(const_cast<const char*>(typeName));
}

int TreeItem92ca9a_TreeItem92ca9a_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItem92ca9a>();
#else
	return 0;
#endif
}

int TreeItem92ca9a_TreeItem92ca9a_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItem92ca9a>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TreeItem92ca9a___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TreeItem92ca9a___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TreeItem92ca9a___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TreeItem92ca9a___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItem92ca9a___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItem92ca9a___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItem92ca9a___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItem92ca9a___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItem92ca9a___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItem92ca9a___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItem92ca9a___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItem92ca9a___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItem92ca9a___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItem92ca9a___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItem92ca9a___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TreeItem92ca9a_NewTreeItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItem92ca9a(static_cast<QWindow*>(parent));
	} else {
		return new TreeItem92ca9a(static_cast<QObject*>(parent));
	}
}

void TreeItem92ca9a_DestroyTreeItem(void* ptr)
{
	static_cast<TreeItem92ca9a*>(ptr)->~TreeItem92ca9a();
}

void TreeItem92ca9a_DestroyTreeItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TreeItem92ca9a_EventDefault(void* ptr, void* e)
{
	return static_cast<TreeItem92ca9a*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TreeItem92ca9a_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TreeItem92ca9a*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TreeItem92ca9a_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TreeItem92ca9a_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItem92ca9a_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TreeItem92ca9a_DeleteLaterDefault(void* ptr)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::deleteLater();
}

void TreeItem92ca9a_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItem92ca9a_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TreeItem92ca9a*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
