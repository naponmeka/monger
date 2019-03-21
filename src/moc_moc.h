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
struct qt_meta_stringdata_CustomTreeModela8f92c_t {
    QByteArrayData data[9];
    char stringdata0[79];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomTreeModela8f92c_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomTreeModela8f92c_t qt_meta_stringdata_CustomTreeModela8f92c = {
    {
QT_MOC_LITERAL(0, 0, 21), // "CustomTreeModela8f92c"
QT_MOC_LITERAL(1, 22, 6), // "remove"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 3), // "add"
QT_MOC_LITERAL(4, 34, 15), // "TreeItema8f92c*"
QT_MOC_LITERAL(5, 50, 4), // "item"
QT_MOC_LITERAL(6, 55, 4), // "edit"
QT_MOC_LITERAL(7, 60, 9), // "firstName"
QT_MOC_LITERAL(8, 70, 8) // "lastName"

    },
    "CustomTreeModela8f92c\0remove\0\0add\0"
    "TreeItema8f92c*\0item\0edit\0firstName\0"
    "lastName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomTreeModela8f92c[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   29,    2, 0x06 /* Public */,
       3,    1,   30,    2, 0x06 /* Public */,
       6,    2,   33,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    7,    8,

       0        // eod
};

void CustomTreeModela8f92c::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CustomTreeModela8f92c *_t = static_cast<CustomTreeModela8f92c *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->remove(); break;
        case 1: _t->add((*reinterpret_cast< TreeItema8f92c*(*)>(_a[1]))); break;
        case 2: _t->edit((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< TreeItema8f92c* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomTreeModela8f92c::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModela8f92c::remove)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CustomTreeModela8f92c::*)(TreeItema8f92c * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModela8f92c::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (CustomTreeModela8f92c::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomTreeModela8f92c::edit)) {
                *result = 2;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject CustomTreeModela8f92c::staticMetaObject = { {
    &QAbstractItemModel::staticMetaObject,
    qt_meta_stringdata_CustomTreeModela8f92c.data,
    qt_meta_data_CustomTreeModela8f92c,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *CustomTreeModela8f92c::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CustomTreeModela8f92c::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CustomTreeModela8f92c.stringdata0))
        return static_cast<void*>(this);
    return QAbstractItemModel::qt_metacast(_clname);
}

int CustomTreeModela8f92c::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractItemModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    }
    return _id;
}

// SIGNAL 0
void CustomTreeModela8f92c::remove()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void CustomTreeModela8f92c::add(TreeItema8f92c * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void CustomTreeModela8f92c::edit(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_TreeItema8f92c_t {
    QByteArrayData data[1];
    char stringdata0[15];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TreeItema8f92c_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TreeItema8f92c_t qt_meta_stringdata_TreeItema8f92c = {
    {
QT_MOC_LITERAL(0, 0, 14) // "TreeItema8f92c"

    },
    "TreeItema8f92c"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TreeItema8f92c[] = {

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

void TreeItema8f92c::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    Q_UNUSED(_o);
    Q_UNUSED(_id);
    Q_UNUSED(_c);
    Q_UNUSED(_a);
}

QT_INIT_METAOBJECT const QMetaObject TreeItema8f92c::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_TreeItema8f92c.data,
    qt_meta_data_TreeItema8f92c,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TreeItema8f92c::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TreeItema8f92c::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TreeItema8f92c.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TreeItema8f92c::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
