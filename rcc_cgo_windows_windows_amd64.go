package main

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -Wall -W -Wextra -fexceptions -mthreads -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../naponmeka -I. -I/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/include -I/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/include/QtGui -I/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/include/QtCore -Irelease -I/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -L/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/lib -lQt5Gui -lQt5Core  -lmingw32 -lqtmain -lshell32 -L/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/lib -lQt5Core -lopengl32
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
