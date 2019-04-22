

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


class CustomTreeModelc306a2: public QAbstractItemModel
{
Q_OBJECT
public:
	CustomTreeModelc306a2(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomTreeModelc306a2_CustomTreeModelc306a2_QRegisterMetaType();CustomTreeModelc306a2_CustomTreeModelc306a2_QRegisterMetaTypes();callbackCustomTreeModelc306a2_Constructor(this);};
	void Signal_Remove() { callbackCustomTreeModelc306a2_Remove(this); };
	void Signal_Add(TreeItemc306a2* item) { callbackCustomTreeModelc306a2_Add(this, item); };
	void Signal_Edit(QString firstName, QString lastName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackCustomTreeModelc306a2_Edit(this, firstNamePacked, lastNamePacked); };
	 ~CustomTreeModelc306a2() { callbackCustomTreeModelc306a2_DestroyCustomTreeModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomTreeModelc306a2_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModelc306a2_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModelc306a2_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModelc306a2_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModelc306a2_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModelc306a2_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModelc306a2_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomTreeModelc306a2_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomTreeModelc306a2_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomTreeModelc306a2_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomTreeModelc306a2_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomTreeModelc306a2_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomTreeModelc306a2_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomTreeModelc306a2_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomTreeModelc306a2_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomTreeModelc306a2_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModelc306a2_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModelc306a2_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomTreeModelc306a2_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomTreeModelc306a2_ModelReset(this); };
	void resetInternalData() { callbackCustomTreeModelc306a2_ResetInternalData(this); };
	void revert() { callbackCustomTreeModelc306a2_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomTreeModelc306a2_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomTreeModelc306a2_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomTreeModelc306a2_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelc306a2_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomTreeModelc306a2_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomTreeModelc306a2_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomTreeModelc306a2_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomTreeModelc306a2_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelc306a2_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelc306a2_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelc306a2_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelc306a2_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomTreeModelc306a2_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomTreeModelc306a2_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomTreeModelc306a2_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModelc306a2_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModelc306a2_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModelc306a2_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModelc306a2_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomTreeModelc306a2_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomTreeModelc306a2_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomTreeModelc306a2_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomTreeModelc306a2_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomTreeModelc306a2_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomTreeModelc306a2_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomTreeModelc306a2_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomTreeModelc306a2_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomTreeModelc306a2_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomTreeModelc306a2_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomTreeModelc306a2_CustomEvent(this, event); };
	void deleteLater() { callbackCustomTreeModelc306a2_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomTreeModelc306a2_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomTreeModelc306a2_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomTreeModelc306a2_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomTreeModelc306a2_TimerEvent(this, event); };
signals:
	void remove();
	void add(TreeItemc306a2* item);
	void edit(QString firstName, QString lastName);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomTreeModelc306a2*)


void CustomTreeModelc306a2_CustomTreeModelc306a2_QRegisterMetaTypes() {
}

class TreeItemc306a2: public QObject
{
Q_OBJECT
public:
	TreeItemc306a2(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TreeItemc306a2_TreeItemc306a2_QRegisterMetaType();TreeItemc306a2_TreeItemc306a2_QRegisterMetaTypes();callbackTreeItemc306a2_Constructor(this);};
	 ~TreeItemc306a2() { callbackTreeItemc306a2_DestroyTreeItem(this); };
	bool event(QEvent * e) { return callbackTreeItemc306a2_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTreeItemc306a2_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTreeItemc306a2_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTreeItemc306a2_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTreeItemc306a2_CustomEvent(this, event); };
	void deleteLater() { callbackTreeItemc306a2_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTreeItemc306a2_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTreeItemc306a2_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTreeItemc306a2_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTreeItemc306a2_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(TreeItemc306a2*)


void TreeItemc306a2_TreeItemc306a2_QRegisterMetaTypes() {
}

void CustomTreeModelc306a2_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)()>(&CustomTreeModelc306a2::remove), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)()>(&CustomTreeModelc306a2::Signal_Remove));
}

void CustomTreeModelc306a2_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)()>(&CustomTreeModelc306a2::remove), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)()>(&CustomTreeModelc306a2::Signal_Remove));
}

void CustomTreeModelc306a2_Remove(void* ptr)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->remove();
}

void CustomTreeModelc306a2_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(TreeItemc306a2*)>(&CustomTreeModelc306a2::add), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(TreeItemc306a2*)>(&CustomTreeModelc306a2::Signal_Add));
}

void CustomTreeModelc306a2_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(TreeItemc306a2*)>(&CustomTreeModelc306a2::add), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(TreeItemc306a2*)>(&CustomTreeModelc306a2::Signal_Add));
}

void CustomTreeModelc306a2_Add(void* ptr, void* item)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->add(static_cast<TreeItemc306a2*>(item));
}

void CustomTreeModelc306a2_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(QString, QString)>(&CustomTreeModelc306a2::edit), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(QString, QString)>(&CustomTreeModelc306a2::Signal_Edit));
}

void CustomTreeModelc306a2_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(QString, QString)>(&CustomTreeModelc306a2::edit), static_cast<CustomTreeModelc306a2*>(ptr), static_cast<void (CustomTreeModelc306a2::*)(QString, QString)>(&CustomTreeModelc306a2::Signal_Edit));
}

void CustomTreeModelc306a2_Edit(void* ptr, struct Moc_PackedString firstName, struct Moc_PackedString lastName)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->edit(QString::fromUtf8(firstName.data, firstName.len), QString::fromUtf8(lastName.data, lastName.len));
}

int CustomTreeModelc306a2_CustomTreeModelc306a2_QRegisterMetaType()
{
	return qRegisterMetaType<CustomTreeModelc306a2*>();
}

int CustomTreeModelc306a2_CustomTreeModelc306a2_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomTreeModelc306a2*>(const_cast<const char*>(typeName));
}

int CustomTreeModelc306a2_CustomTreeModelc306a2_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModelc306a2>();
#else
	return 0;
#endif
}

int CustomTreeModelc306a2_CustomTreeModelc306a2_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModelc306a2>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CustomTreeModelc306a2___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModelc306a2___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModelc306a2___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelc306a2___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelc306a2___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelc306a2___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelc306a2___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModelc306a2___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomTreeModelc306a2___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModelc306a2___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModelc306a2___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModelc306a2___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModelc306a2___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomTreeModelc306a2___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomTreeModelc306a2___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelc306a2___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModelc306a2___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModelc306a2___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelc306a2___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelc306a2___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelc306a2___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelc306a2___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelc306a2___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelc306a2___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModelc306a2_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelc306a2_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelc306a2_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelc306a2_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelc306a2_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelc306a2_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomTreeModelc306a2___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelc306a2___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomTreeModelc306a2___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomTreeModelc306a2___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelc306a2___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelc306a2___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelc306a2___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelc306a2___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelc306a2___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelc306a2___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelc306a2___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelc306a2___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomTreeModelc306a2_NewCustomTreeModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelc306a2(static_cast<QWindow*>(parent));
	} else {
		return new CustomTreeModelc306a2(static_cast<QObject*>(parent));
	}
}

void CustomTreeModelc306a2_DestroyCustomTreeModel(void* ptr)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->~CustomTreeModelc306a2();
}

void CustomTreeModelc306a2_DestroyCustomTreeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomTreeModelc306a2_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModelc306a2_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModelc306a2_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomTreeModelc306a2_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomTreeModelc306a2_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomTreeModelc306a2_SubmitDefault(void* ptr)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::submit();
}

void CustomTreeModelc306a2_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomTreeModelc306a2_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::resetInternalData();
}

void CustomTreeModelc306a2_RevertDefault(void* ptr)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::revert();
}

void CustomTreeModelc306a2_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomTreeModelc306a2_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomTreeModelc306a2_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelc306a2_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomTreeModelc306a2_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomTreeModelc306a2_Index(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModelc306a2_IndexDefault(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModelc306a2_Parent(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModelc306a2_ParentDefault(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModelc306a2_SiblingDefault(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomTreeModelc306a2_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelc306a2_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomTreeModelc306a2_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tbaf865 = static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tbaf865.prepend("WHITESPACE").constData()+10), tbaf865.size()-10 }; });
}

void* CustomTreeModelc306a2_Data(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModelc306a2_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModelc306a2_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomTreeModelc306a2_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::supportedDragActions();
}

long long CustomTreeModelc306a2_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::supportedDropActions();
}

long long CustomTreeModelc306a2_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::flags(*static_cast<QModelIndex*>(index));
}

char CustomTreeModelc306a2_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomTreeModelc306a2_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}



int CustomTreeModelc306a2_ColumnCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelc306a2_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelc306a2_RowCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelc306a2_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomTreeModelc306a2_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
}

char CustomTreeModelc306a2_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomTreeModelc306a2_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomTreeModelc306a2_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModelc306a2_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
}

void CustomTreeModelc306a2_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::deleteLater();
}

void CustomTreeModelc306a2_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModelc306a2_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelc306a2*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

int TreeItemc306a2_TreeItemc306a2_QRegisterMetaType()
{
	return qRegisterMetaType<TreeItemc306a2*>();
}

int TreeItemc306a2_TreeItemc306a2_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TreeItemc306a2*>(const_cast<const char*>(typeName));
}

int TreeItemc306a2_TreeItemc306a2_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItemc306a2>();
#else
	return 0;
#endif
}

int TreeItemc306a2_TreeItemc306a2_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItemc306a2>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TreeItemc306a2___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TreeItemc306a2___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TreeItemc306a2___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TreeItemc306a2___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemc306a2___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemc306a2___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemc306a2___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemc306a2___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemc306a2___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemc306a2___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemc306a2___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemc306a2___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemc306a2___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemc306a2___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemc306a2___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TreeItemc306a2_NewTreeItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItemc306a2(static_cast<QWindow*>(parent));
	} else {
		return new TreeItemc306a2(static_cast<QObject*>(parent));
	}
}

void TreeItemc306a2_DestroyTreeItem(void* ptr)
{
	static_cast<TreeItemc306a2*>(ptr)->~TreeItemc306a2();
}

void TreeItemc306a2_DestroyTreeItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TreeItemc306a2_EventDefault(void* ptr, void* e)
{
	return static_cast<TreeItemc306a2*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TreeItemc306a2_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TreeItemc306a2*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TreeItemc306a2_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TreeItemc306a2_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItemc306a2_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TreeItemc306a2_DeleteLaterDefault(void* ptr)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::deleteLater();
}

void TreeItemc306a2_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItemc306a2_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemc306a2*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
