/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.12.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.12.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_CustomTreeModel5c58e0_t {
    QByteArrayData data[10];
    char stringdata0[89];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomTreeModel5c58e0_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomTreeModel5c58e0_t qt_meta_stringdata_CustomTreeModel5c58e0 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "CustomTreeModel5c58e0"
QT_MOC_LITERAL(1, 22, 6), // "remove"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 9), // "removeAll"
QT_MOC_LITERAL(4, 40, 3), // "add"
QT_MOC_LITERAL(5, 44, 15), // "TreeItem5c58e0*"
QT_MOC_LITERAL(6, 60, 4), // "item"
QT_MOC_LITERAL(7, 65, 4), // "edit"
QT_MOC_LITERAL(8, 70, 9), // "firstName"
QT_MOC_LITERAL(9, 80, 8) // "lastName"

    },
    "CustomTreeModel5c58e0\0remove\0\0removeAll\0"
    "add\0TreeItem5c58e0*\0item\0edit\0firstName\0"
    "lastName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomTreeModel5c58e0[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   34,    2, 0x06 /* Public */,
       3,    0,   35,    2, 0x06 /* Public */,
       4,    1,   36,    2, 0x06 /* Public */,
       7,    2,   39,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 5,    6,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    8,    9,

       0        // eod
};

void CustomTreeModel5c58e0::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CustomTreeModel5c58e0 *_t = static_cast<CustomTreeModel5c58e0 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->remove(); break;
        case 1: _t->removeAll(); break;
        case 2: _t->add((*reinterpret_cast< TreeItem5c58e0*(*)>(_a[1]))); break;
        case 3: _t->edit((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< TreeItem5c58e0* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomTreeModel5c58e0::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModel5c58e0::remove)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CustomTreeModel5c58e0::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModel5c58e0::removeAll)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (CustomTreeModel5c58e0::*)(TreeItem5c58e0 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModel5c58e0::add)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (CustomTreeModel5c58e0::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModel5c58e0::edit)) {
                *result = 3;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject CustomTreeModel5c58e0::staticMetaObject = { {
    &QAbstractItemModel::staticMetaObject,
    qt_meta_stringdata_CustomTreeModel5c58e0.data,
    qt_meta_data_CustomTreeModel5c58e0,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *CustomTreeModel5c58e0::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CustomTreeModel5c58e0::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CustomTreeModel5c58e0.stringdata0))
        return static_cast<void*>(this);
    return QAbstractItemModel::qt_metacast(_clname);
}

int CustomTreeModel5c58e0::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractItemModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void CustomTreeModel5c58e0::remove()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void CustomTreeModel5c58e0::removeAll()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void CustomTreeModel5c58e0::add(TreeItem5c58e0 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void CustomTreeModel5c58e0::edit(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_TreeItem5c58e0_t {
    QByteArrayData data[1];
    char stringdata0[15];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TreeItem5c58e0_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TreeItem5c58e0_t qt_meta_stringdata_TreeItem5c58e0 = {
    {
QT_MOC_LITERAL(0, 0, 14) // "TreeItem5c58e0"

    },
    "TreeItem5c58e0"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TreeItem5c58e0[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       0,    0, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

       0        // eod
};

void TreeItem5c58e0::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    Q_UNUSED(_o);
    Q_UNUSED(_id);
    Q_UNUSED(_c);
    Q_UNUSED(_a);
}

QT_INIT_METAOBJECT const QMetaObject TreeItem5c58e0::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_TreeItem5c58e0.data,
    qt_meta_data_TreeItem5c58e0,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TreeItem5c58e0::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TreeItem5c58e0::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TreeItem5c58e0.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TreeItem5c58e0::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
