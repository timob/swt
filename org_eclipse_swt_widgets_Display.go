package swt

import "github.com/timob/javabind"

type WidgetsDisplayInterface interface {
	GraphicsDeviceInterface

	// public void addFilter(int, org.eclipse.swt.widgets.Listener)
	AddFilter(a int, b WidgetsListenerInterface) 

	// public void addListener(int, org.eclipse.swt.widgets.Listener)
	AddListener(a int, b WidgetsListenerInterface) 

	// public void beep()
	Beep() 

	// public void close()
	Close() 

	// public org.eclipse.swt.widgets.Widget findWidget(long)
	FindWidget(a int64) *WidgetsWidget

	// public org.eclipse.swt.widgets.Widget findWidget(long, long)
	FindWidget2(a int64, b int64) *WidgetsWidget

	// public org.eclipse.swt.widgets.Widget findWidget(org.eclipse.swt.widgets.Widget, long)
	FindWidget3(a WidgetsWidgetInterface, b int64) *WidgetsWidget

	// public org.eclipse.swt.widgets.Shell getActiveShell()
	GetActiveShell() *WidgetsShell

	// public static org.eclipse.swt.widgets.Display getCurrent()
	GetCurrent() *WidgetsDisplay

	// public org.eclipse.swt.widgets.Control getCursorControl()
	GetCursorControl() *WidgetsControl

	// public org.eclipse.swt.graphics.Point getCursorLocation()
	GetCursorLocation() *GraphicsPoint

	// public org.eclipse.swt.graphics.Point[] getCursorSizes()
	GetCursorSizes() []*GraphicsPoint

	// public java.lang.Object getData(java.lang.String)
	GetData2(a string) *JavaLangObject

	// public java.lang.Object getData()
	GetData() *JavaLangObject

	// public static org.eclipse.swt.widgets.Display getDefault()
	GetDefault() *WidgetsDisplay

	// public org.eclipse.swt.widgets.Menu getMenuBar()
	GetMenuBar() *WidgetsMenu

	// public int getDismissalAlignment()
	GetDismissalAlignment() int

	// public int getDoubleClickTime()
	GetDoubleClickTime() int

	// public org.eclipse.swt.widgets.Control getFocusControl()
	GetFocusControl() *WidgetsControl

	// public boolean getHighContrast()
	GetHighContrast() bool

	// public int getIconDepth()
	GetIconDepth() int

	// public org.eclipse.swt.graphics.Point[] getIconSizes()
	GetIconSizes() []*GraphicsPoint

	// public org.eclipse.swt.widgets.Monitor[] getMonitors()
	GetMonitors() []*WidgetsMonitor

	// public org.eclipse.swt.widgets.Monitor getPrimaryMonitor()
	GetPrimaryMonitor() *WidgetsMonitor

	// public org.eclipse.swt.widgets.Shell[] getShells()
	GetShells() []*WidgetsShell

	// public org.eclipse.swt.widgets.Synchronizer getSynchronizer()
	GetSynchronizer() *WidgetsSynchronizer

	// public org.eclipse.swt.graphics.Cursor getSystemCursor(int)
	GetSystemCursor(a int) *GraphicsCursor

	// public org.eclipse.swt.graphics.Image getSystemImage(int)
	GetSystemImage(a int) *GraphicsImage

	// public org.eclipse.swt.widgets.Menu getSystemMenu()
	GetSystemMenu() *WidgetsMenu

	// public org.eclipse.swt.widgets.TaskBar getSystemTaskBar()
	GetSystemTaskBar() *WidgetsTaskBar

	// public org.eclipse.swt.widgets.Tray getSystemTray()
	GetSystemTray() *WidgetsTray

	// public boolean getTouchEnabled()
	GetTouchEnabled() bool

	// public org.eclipse.swt.graphics.Point map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, org.eclipse.swt.graphics.Point)
	Map3(a WidgetsControlInterface, b WidgetsControlInterface, c GraphicsPointInterface) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, int, int)
	Map(a WidgetsControlInterface, b WidgetsControlInterface, c int, d int) *GraphicsPoint

	// public org.eclipse.swt.graphics.Rectangle map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, org.eclipse.swt.graphics.Rectangle)
	Map4(a WidgetsControlInterface, b WidgetsControlInterface, c GraphicsRectangleInterface) *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, int, int, int, int)
	Map2(a WidgetsControlInterface, b WidgetsControlInterface, c int, d int, e int, f int) *GraphicsRectangle

	// public boolean post(org.eclipse.swt.widgets.Event)
	Post(a WidgetsEventInterface) bool

	// public boolean readAndDispatch()
	ReadAndDispatch() bool

	// public void removeFilter(int, org.eclipse.swt.widgets.Listener)
	RemoveFilter(a int, b WidgetsListenerInterface) 

	// public void removeListener(int, org.eclipse.swt.widgets.Listener)
	RemoveListener(a int, b WidgetsListenerInterface) 

	// public static java.lang.String getAppName()
	GetAppName() string

	// public static java.lang.String getAppVersion()
	GetAppVersion() string

	// public static void setAppName(java.lang.String)
	SetAppName(a string) 

	// public static void setAppVersion(java.lang.String)
	SetAppVersion(a string) 

	// public void setCursorLocation(int, int)
	SetCursorLocation(a int, b int) 

	// public void setCursorLocation(org.eclipse.swt.graphics.Point)
	SetCursorLocation2(a GraphicsPointInterface) 

	// public void setData(java.lang.String, java.lang.Object)
	SetData2(a string, b interface{}) 

	// public void setData(java.lang.Object)
	SetData(a interface{}) 

	// public void setSynchronizer(org.eclipse.swt.widgets.Synchronizer)
	SetSynchronizer(a WidgetsSynchronizerInterface) 

	// public boolean sleep()
	Sleep() bool

	// public void update()
	Update() 

	// public void wake()
	Wake() 
}

type WidgetsDisplay struct {
	GraphicsDevice
}

// public org.eclipse.swt.widgets.Display()
func NewWidgetsDisplay() (*WidgetsDisplay) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Display")
	if err != nil {
		panic(err)
	}
	x := &WidgetsDisplay{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Display(org.eclipse.swt.graphics.DeviceData)
func NewWidgetsDisplay2(a GraphicsDeviceDataInterface) (*WidgetsDisplay) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Display", conv_a.Value().Cast("org/eclipse/swt/graphics/DeviceData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDisplay{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addFilter(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsDisplay) AddFilter(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addFilter", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void addListener(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsDisplay) AddListener(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addListener", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void beep()
func (jbobject *WidgetsDisplay) Beep()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "beep", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void close()
func (jbobject *WidgetsDisplay) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.widgets.Widget findWidget(long)
func (jbobject *WidgetsDisplay) FindWidget(a int64) *WidgetsWidget {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "findWidget", "org/eclipse/swt/widgets/Widget", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Widget findWidget(long, long)
func (jbobject *WidgetsDisplay) FindWidget2(a int64, b int64) *WidgetsWidget {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "findWidget", "org/eclipse/swt/widgets/Widget", a, b)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Widget findWidget(org.eclipse.swt.widgets.Widget, long)
func (jbobject *WidgetsDisplay) FindWidget3(a WidgetsWidgetInterface, b int64) *WidgetsWidget {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "findWidget", "org/eclipse/swt/widgets/Widget", conv_a.Value().Cast("org/eclipse/swt/widgets/Widget"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Shell getActiveShell()
func (jbobject *WidgetsDisplay) GetActiveShell() *WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveShell", "org/eclipse/swt/widgets/Shell")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsDisplay) GetBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.widgets.Display getCurrent()
func (jbobject *WidgetsDisplay) GetCurrent() *WidgetsDisplay {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "getCurrent", "org/eclipse/swt/widgets/Display")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsDisplay{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Control getCursorControl()
func (jbobject *WidgetsDisplay) GetCursorControl() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCursorControl", "org/eclipse/swt/widgets/Control")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsControl{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point getCursorLocation()
func (jbobject *WidgetsDisplay) GetCursorLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCursorLocation", "org/eclipse/swt/graphics/Point")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point[] getCursorSizes()
func (jbobject *WidgetsDisplay) GetCursorSizes() []*GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCursorSizes", javabind.ObjectArrayType("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/Point")
	dst := new([]*GraphicsPoint)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Object getData(java.lang.String)
func (jbobject *WidgetsDisplay) GetData2(a string) *JavaLangObject {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getData", "java/lang/Object", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object getData()
func (jbobject *WidgetsDisplay) GetData() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getData", "java/lang/Object")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.widgets.Display getDefault()
func (jbobject *WidgetsDisplay) GetDefault() *WidgetsDisplay {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "getDefault", "org/eclipse/swt/widgets/Display")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsDisplay{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Menu getMenuBar()
func (jbobject *WidgetsDisplay) GetMenuBar() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMenuBar", "org/eclipse/swt/widgets/Menu")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsMenu{}
	unique_x.Callable = dst
	return unique_x
}

// public int getDismissalAlignment()
func (jbobject *WidgetsDisplay) GetDismissalAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDismissalAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getDoubleClickTime()
func (jbobject *WidgetsDisplay) GetDoubleClickTime() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDoubleClickTime", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Control getFocusControl()
func (jbobject *WidgetsDisplay) GetFocusControl() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFocusControl", "org/eclipse/swt/widgets/Control")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsControl{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getHighContrast()
func (jbobject *WidgetsDisplay) GetHighContrast() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHighContrast", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getDepth()
func (jbobject *WidgetsDisplay) GetDepth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDepth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getIconDepth()
func (jbobject *WidgetsDisplay) GetIconDepth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIconDepth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point[] getIconSizes()
func (jbobject *WidgetsDisplay) GetIconSizes() []*GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIconSizes", javabind.ObjectArrayType("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/Point")
	dst := new([]*GraphicsPoint)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.widgets.Monitor[] getMonitors()
func (jbobject *WidgetsDisplay) GetMonitors() []*WidgetsMonitor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitors", javabind.ObjectArrayType("org/eclipse/swt/widgets/Monitor"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Monitor")
	dst := new([]*WidgetsMonitor)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.widgets.Monitor getPrimaryMonitor()
func (jbobject *WidgetsDisplay) GetPrimaryMonitor() *WidgetsMonitor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrimaryMonitor", "org/eclipse/swt/widgets/Monitor")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsMonitor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Shell[] getShells()
func (jbobject *WidgetsDisplay) GetShells() []*WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getShells", javabind.ObjectArrayType("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Shell")
	dst := new([]*WidgetsShell)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.widgets.Synchronizer getSynchronizer()
func (jbobject *WidgetsDisplay) GetSynchronizer() *WidgetsSynchronizer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSynchronizer", "org/eclipse/swt/widgets/Synchronizer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsSynchronizer{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Color getSystemColor(int)
func (jbobject *WidgetsDisplay) GetSystemColor(a int) *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemColor", "org/eclipse/swt/graphics/Color", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Cursor getSystemCursor(int)
func (jbobject *WidgetsDisplay) GetSystemCursor(a int) *GraphicsCursor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemCursor", "org/eclipse/swt/graphics/Cursor", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsCursor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Image getSystemImage(int)
func (jbobject *WidgetsDisplay) GetSystemImage(a int) *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemImage", "org/eclipse/swt/graphics/Image", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Menu getSystemMenu()
func (jbobject *WidgetsDisplay) GetSystemMenu() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemMenu", "org/eclipse/swt/widgets/Menu")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsMenu{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Font getSystemFont()
func (jbobject *WidgetsDisplay) GetSystemFont() *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemFont", "org/eclipse/swt/graphics/Font")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsFont{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TaskBar getSystemTaskBar()
func (jbobject *WidgetsDisplay) GetSystemTaskBar() *WidgetsTaskBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemTaskBar", "org/eclipse/swt/widgets/TaskBar")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsTaskBar{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Tray getSystemTray()
func (jbobject *WidgetsDisplay) GetSystemTray() *WidgetsTray {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemTray", "org/eclipse/swt/widgets/Tray")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsTray{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getTouchEnabled()
func (jbobject *WidgetsDisplay) GetTouchEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTouchEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
func (jbobject *WidgetsDisplay) Internal_dispose_GC(a int64, b GraphicsGCDataInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "internal_dispose_GC", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public long internal_new_GC(org.eclipse.swt.graphics.GCData)
func (jbobject *WidgetsDisplay) Internal_new_GC(a GraphicsGCDataInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "internal_new_GC", javabind.Long, conv_a.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public org.eclipse.swt.graphics.Point map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsDisplay) Map3(a WidgetsControlInterface, b WidgetsControlInterface, c GraphicsPointInterface) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "map", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/Control"), conv_c.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, int, int)
func (jbobject *WidgetsDisplay) Map(a WidgetsControlInterface, b WidgetsControlInterface, c int, d int) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "map", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/Control"), c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, org.eclipse.swt.graphics.Rectangle)
func (jbobject *WidgetsDisplay) Map4(a WidgetsControlInterface, b WidgetsControlInterface, c GraphicsRectangleInterface) *GraphicsRectangle {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "map", "org/eclipse/swt/graphics/Rectangle", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/Control"), conv_c.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle map(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.Control, int, int, int, int)
func (jbobject *WidgetsDisplay) Map2(a WidgetsControlInterface, b WidgetsControlInterface, c int, d int, e int, f int) *GraphicsRectangle {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "map", "org/eclipse/swt/graphics/Rectangle", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/Control"), c, d, e, f)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean post(org.eclipse.swt.widgets.Event)
func (jbobject *WidgetsDisplay) Post(a WidgetsEventInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "post", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean readAndDispatch()
func (jbobject *WidgetsDisplay) ReadAndDispatch() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "readAndDispatch", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeFilter(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsDisplay) RemoveFilter(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeFilter", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void removeListener(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsDisplay) RemoveListener(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeListener", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public static java.lang.String getAppName()
func (jbobject *WidgetsDisplay) GetAppName() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "getAppName", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String getAppVersion()
func (jbobject *WidgetsDisplay) GetAppVersion() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "getAppVersion", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static void setAppName(java.lang.String)
func (jbobject *WidgetsDisplay) SetAppName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "setAppName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public static void setAppVersion(java.lang.String)
func (jbobject *WidgetsDisplay) SetAppVersion(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Display", "setAppVersion", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCursorLocation(int, int)
func (jbobject *WidgetsDisplay) SetCursorLocation(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCursorLocation", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setCursorLocation(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsDisplay) SetCursorLocation2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCursorLocation", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setData(java.lang.String, java.lang.Object)
func (jbobject *WidgetsDisplay) SetData2(a string, b interface{})  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setData", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void setData(java.lang.Object)
func (jbobject *WidgetsDisplay) SetData(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setData", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSynchronizer(org.eclipse.swt.widgets.Synchronizer)
func (jbobject *WidgetsDisplay) SetSynchronizer(a WidgetsSynchronizerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSynchronizer", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Synchronizer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean sleep()
func (jbobject *WidgetsDisplay) Sleep() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sleep", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void update()
func (jbobject *WidgetsDisplay) Update()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "update", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void wake()
func (jbobject *WidgetsDisplay) Wake()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "wake", javabind.Void)
	if err != nil {
		panic(err)
	}

}


