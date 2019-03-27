// +build !ios

package table

/*
#cgo CFLAGS: -pipe -O2 -arch x86_64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -arch x86_64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../widgets -I. -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtMultimedia.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtDesigner.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtUiPlugin.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtWidgets.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtQuick.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtGui.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtQml.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtNetwork.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtDBus.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtXml.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/lib/QtCore.framework/Headers -I. -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/AGL.framework/Headers -I/Users/naponmeka/Qt/5.12.0/clang_64/mkspecs/macx-clang -F/Users/naponmeka/Qt/5.12.0/clang_64/lib
#cgo LDFLAGS: -stdlib=libc++ -headerpad_max_install_names -arch x86_64 -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12  -Wl,-rpath,/Users/naponmeka/Qt/5.12.0/clang_64/lib
#cgo LDFLAGS:  -F/Users/naponmeka/Qt/5.12.0/clang_64/lib -framework QtMultimedia -framework QtNetwork -framework QtCore -framework DiskArbitration -framework IOKit -framework QtGui -framework QtDesigner -framework QtWidgets -framework QtXml -framework QtQuick -framework QtQml -framework QtDBus -framework OpenGL -framework AGL
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
