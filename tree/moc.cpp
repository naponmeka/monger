

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


class CustomTreeModelabdc68: public QAbstractItemModel
{
Q_OBJECT
public:
	CustomTreeModelabdc68(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomTreeModelabdc68_CustomTreeModelabdc68_QRegisterMetaType();CustomTreeModelabdc68_CustomTreeModelabdc68_QRegisterMetaTypes();callbackCustomTreeModelabdc68_Constructor(this);};
	void Signal_Remove() { callbackCustomTreeModelabdc68_Remove(this); };
	void Signal_RemoveAll() { callbackCustomTreeModelabdc68_RemoveAll(this); };
	void Signal_Add(TreeItemabdc68* item) { callbackCustomTreeModelabdc68_Add(this, item); };
	void Signal_Edit(QString firstName, QString lastName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackCustomTreeModelabdc68_Edit(this, firstNamePacked, lastNamePacked); };
	 ~CustomTreeModelabdc68() { callbackCustomTreeModelabdc68_DestroyCustomTreeModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomTreeModelabdc68_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModelabdc68_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModelabdc68_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModelabdc68_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomTreeModelabdc68_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomTreeModelabdc68_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomTreeModelabdc68_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomTreeModelabdc68_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomTreeModelabdc68_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomTreeModelabdc68_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomTreeModelabdc68_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomTreeModelabdc68_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomTreeModelabdc68_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomTreeModelabdc68_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomTreeModelabdc68_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomTreeModelabdc68_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModelabdc68_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomTreeModelabdc68_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomTreeModelabdc68_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomTreeModelabdc68_ModelReset(this); };
	void resetInternalData() { callbackCustomTreeModelabdc68_ResetInternalData(this); };
	void revert() { callbackCustomTreeModelabdc68_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomTreeModelabdc68_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomTreeModelabdc68_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomTreeModelabdc68_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomTreeModelabdc68_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomTreeModelabdc68_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomTreeModelabdc68_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomTreeModelabdc68_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomTreeModelabdc68_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelabdc68_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelabdc68_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelabdc68_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomTreeModelabdc68_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomTreeModelabdc68_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomTreeModelabdc68_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomTreeModelabdc68_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModelabdc68_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomTreeModelabdc68_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModelabdc68_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomTreeModelabdc68_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomTreeModelabdc68_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomTreeModelabdc68_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomTreeModelabdc68_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomTreeModelabdc68_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomTreeModelabdc68_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomTreeModelabdc68_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomTreeModelabdc68_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomTreeModelabdc68_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomTreeModelabdc68_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomTreeModelabdc68_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomTreeModelabdc68_CustomEvent(this, event); };
	void deleteLater() { callbackCustomTreeModelabdc68_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomTreeModelabdc68_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomTreeModelabdc68_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomTreeModelabdc68_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomTreeModelabdc68_TimerEvent(this, event); };
signals:
	void remove();
	void removeAll();
	void add(TreeItemabdc68* item);
	void edit(QString firstName, QString lastName);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomTreeModelabdc68*)


void CustomTreeModelabdc68_CustomTreeModelabdc68_QRegisterMetaTypes() {
}

class TreeItemabdc68: public QObject
{
Q_OBJECT
public:
	TreeItemabdc68(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TreeItemabdc68_TreeItemabdc68_QRegisterMetaType();TreeItemabdc68_TreeItemabdc68_QRegisterMetaTypes();callbackTreeItemabdc68_Constructor(this);};
	 ~TreeItemabdc68() { callbackTreeItemabdc68_DestroyTreeItem(this); };
	bool event(QEvent * e) { return callbackTreeItemabdc68_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTreeItemabdc68_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTreeItemabdc68_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTreeItemabdc68_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTreeItemabdc68_CustomEvent(this, event); };
	void deleteLater() { callbackTreeItemabdc68_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTreeItemabdc68_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTreeItemabdc68_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTreeItemabdc68_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTreeItemabdc68_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(TreeItemabdc68*)


void TreeItemabdc68_TreeItemabdc68_QRegisterMetaTypes() {
}

int TreeItemabdc68_TreeItemabdc68_QRegisterMetaType()
{
	return qRegisterMetaType<TreeItemabdc68*>();
}

int TreeItemabdc68_TreeItemabdc68_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TreeItemabdc68*>(const_cast<const char*>(typeName));
}

int TreeItemabdc68_TreeItemabdc68_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItemabdc68>();
#else
	return 0;
#endif
}

int TreeItemabdc68_TreeItemabdc68_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TreeItemabdc68>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TreeItemabdc68___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TreeItemabdc68___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TreeItemabdc68___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TreeItemabdc68___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemabdc68___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemabdc68___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemabdc68___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemabdc68___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemabdc68___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemabdc68___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemabdc68___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemabdc68___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TreeItemabdc68___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TreeItemabdc68___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TreeItemabdc68___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TreeItemabdc68_NewTreeItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TreeItemabdc68(static_cast<QWindow*>(parent));
	} else {
		return new TreeItemabdc68(static_cast<QObject*>(parent));
	}
}

void TreeItemabdc68_DestroyTreeItem(void* ptr)
{
	static_cast<TreeItemabdc68*>(ptr)->~TreeItemabdc68();
}

void TreeItemabdc68_DestroyTreeItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TreeItemabdc68_EventDefault(void* ptr, void* e)
{
	return static_cast<TreeItemabdc68*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TreeItemabdc68_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TreeItemabdc68*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TreeItemabdc68_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TreeItemabdc68_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItemabdc68_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TreeItemabdc68_DeleteLaterDefault(void* ptr)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::deleteLater();
}

void TreeItemabdc68_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TreeItemabdc68_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TreeItemabdc68*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void CustomTreeModelabdc68_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::remove), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::Signal_Remove));
}

void CustomTreeModelabdc68_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::remove), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::Signal_Remove));
}

void CustomTreeModelabdc68_Remove(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->remove();
}

void CustomTreeModelabdc68_ConnectRemoveAll(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::removeAll), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::Signal_RemoveAll));
}

void CustomTreeModelabdc68_DisconnectRemoveAll(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::removeAll), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)()>(&CustomTreeModelabdc68::Signal_RemoveAll));
}

void CustomTreeModelabdc68_RemoveAll(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->removeAll();
}

void CustomTreeModelabdc68_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(TreeItemabdc68*)>(&CustomTreeModelabdc68::add), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(TreeItemabdc68*)>(&CustomTreeModelabdc68::Signal_Add));
}

void CustomTreeModelabdc68_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(TreeItemabdc68*)>(&CustomTreeModelabdc68::add), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(TreeItemabdc68*)>(&CustomTreeModelabdc68::Signal_Add));
}

void CustomTreeModelabdc68_Add(void* ptr, void* item)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->add(static_cast<TreeItemabdc68*>(item));
}

void CustomTreeModelabdc68_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(QString, QString)>(&CustomTreeModelabdc68::edit), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(QString, QString)>(&CustomTreeModelabdc68::Signal_Edit));
}

void CustomTreeModelabdc68_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(QString, QString)>(&CustomTreeModelabdc68::edit), static_cast<CustomTreeModelabdc68*>(ptr), static_cast<void (CustomTreeModelabdc68::*)(QString, QString)>(&CustomTreeModelabdc68::Signal_Edit));
}

void CustomTreeModelabdc68_Edit(void* ptr, struct Moc_PackedString firstName, struct Moc_PackedString lastName)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->edit(QString::fromUtf8(firstName.data, firstName.len), QString::fromUtf8(lastName.data, lastName.len));
}

int CustomTreeModelabdc68_CustomTreeModelabdc68_QRegisterMetaType()
{
	return qRegisterMetaType<CustomTreeModelabdc68*>();
}

int CustomTreeModelabdc68_CustomTreeModelabdc68_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomTreeModelabdc68*>(const_cast<const char*>(typeName));
}

int CustomTreeModelabdc68_CustomTreeModelabdc68_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModelabdc68>();
#else
	return 0;
#endif
}

int CustomTreeModelabdc68_CustomTreeModelabdc68_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomTreeModelabdc68>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CustomTreeModelabdc68___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModelabdc68___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModelabdc68___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelabdc68___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelabdc68___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelabdc68___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelabdc68___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModelabdc68___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomTreeModelabdc68___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModelabdc68___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModelabdc68___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomTreeModelabdc68___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomTreeModelabdc68___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomTreeModelabdc68___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomTreeModelabdc68___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelabdc68___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomTreeModelabdc68___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomTreeModelabdc68___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelabdc68___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelabdc68___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelabdc68___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelabdc68___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomTreeModelabdc68___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomTreeModelabdc68___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomTreeModelabdc68_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelabdc68_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelabdc68_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelabdc68_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomTreeModelabdc68_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomTreeModelabdc68_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomTreeModelabdc68___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomTreeModelabdc68___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomTreeModelabdc68___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomTreeModelabdc68___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelabdc68___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelabdc68___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelabdc68___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelabdc68___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelabdc68___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomTreeModelabdc68___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomTreeModelabdc68___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomTreeModelabdc68___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomTreeModelabdc68_NewCustomTreeModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomTreeModelabdc68(static_cast<QWindow*>(parent));
	} else {
		return new CustomTreeModelabdc68(static_cast<QObject*>(parent));
	}
}

void CustomTreeModelabdc68_DestroyCustomTreeModel(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->~CustomTreeModelabdc68();
}

void CustomTreeModelabdc68_DestroyCustomTreeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomTreeModelabdc68_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModelabdc68_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomTreeModelabdc68_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomTreeModelabdc68_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomTreeModelabdc68_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomTreeModelabdc68_SubmitDefault(void* ptr)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::submit();
}

void CustomTreeModelabdc68_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomTreeModelabdc68_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::resetInternalData();
}

void CustomTreeModelabdc68_RevertDefault(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::revert();
}

void CustomTreeModelabdc68_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomTreeModelabdc68_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomTreeModelabdc68_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelabdc68_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomTreeModelabdc68_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomTreeModelabdc68_Index(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModelabdc68_IndexDefault(void* ptr, int row, int column, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(row);
	Q_UNUSED(column);
	Q_UNUSED(parent);

}

void* CustomTreeModelabdc68_Parent(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModelabdc68_ParentDefault(void* ptr, void* index)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);

}

void* CustomTreeModelabdc68_SiblingDefault(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomTreeModelabdc68_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomTreeModelabdc68_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomTreeModelabdc68_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tbaf865 = static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tbaf865.prepend("WHITESPACE").constData()+10), tbaf865.size()-10 }; });
}

void* CustomTreeModelabdc68_Data(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModelabdc68_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomTreeModelabdc68_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomTreeModelabdc68_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::supportedDragActions();
}

long long CustomTreeModelabdc68_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::supportedDropActions();
}

long long CustomTreeModelabdc68_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::flags(*static_cast<QModelIndex*>(index));
}

char CustomTreeModelabdc68_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomTreeModelabdc68_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}



int CustomTreeModelabdc68_ColumnCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelabdc68_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelabdc68_RowCount(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int CustomTreeModelabdc68_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomTreeModelabdc68_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
}

char CustomTreeModelabdc68_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomTreeModelabdc68_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomTreeModelabdc68_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModelabdc68_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
}

void CustomTreeModelabdc68_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::deleteLater();
}

void CustomTreeModelabdc68_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomTreeModelabdc68_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomTreeModelabdc68*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
