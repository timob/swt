package swt

import "github.com/timob/javabind"

type WidgetsShellInterface interface {
	WidgetsDecorationsInterface

	// public static org.eclipse.swt.widgets.Shell gtk_new(org.eclipse.swt.widgets.Display, long)
	Gtk_new(a WidgetsDisplayInterface, b int64) *WidgetsShell

	// public static org.eclipse.swt.widgets.Shell internal_new(org.eclipse.swt.widgets.Display, long)
	Internal_new(a WidgetsDisplayInterface, b int64) *WidgetsShell

	// public void addShellListener(org.eclipse.swt.events.ShellListener)
	AddShellListener(a EventsShellListenerInterface) 

	// public void close()
	Close() 

	// public org.eclipse.swt.widgets.ToolBar getToolBar()
	GetToolBar() *WidgetsToolBar

	// public int getAlpha()
	GetAlpha() int

	// public boolean getFullScreen()
	GetFullScreen() bool

	// public org.eclipse.swt.graphics.Point getMinimumSize()
	GetMinimumSize() *GraphicsPoint

	// public boolean getModified()
	GetModified() bool

	// public int getImeInputMode()
	GetImeInputMode() int

	// public org.eclipse.swt.widgets.Shell[] getShells()
	GetShells() []*WidgetsShell

	// public void open()
	Open() 

	// public void removeShellListener(org.eclipse.swt.events.ShellListener)
	RemoveShellListener(a EventsShellListenerInterface) 

	// public void setActive()
	SetActive() 

	// public void setAlpha(int)
	SetAlpha(a int) 

	// public void setFullScreen(boolean)
	SetFullScreen(a bool) 

	// public void setImeInputMode(int)
	SetImeInputMode(a int) 

	// public void setMinimumSize(int, int)
	SetMinimumSize(a int, b int) 

	// public void setMinimumSize(org.eclipse.swt.graphics.Point)
	SetMinimumSize2(a GraphicsPointInterface) 

	// public void setModified(boolean)
	SetModified(a bool) 

	// public void forceActive()
	ForceActive() 
}

type WidgetsShell struct {
	WidgetsDecorations
}

// public org.eclipse.swt.widgets.Shell()
func NewWidgetsShell() (*WidgetsShell) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell")
	if err != nil {
		panic(err)
	}
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell(int)
func NewWidgetsShell2(a int) (*WidgetsShell) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell", a)
	if err != nil {
		panic(err)
	}
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell(org.eclipse.swt.widgets.Display)
func NewWidgetsShell3(a WidgetsDisplayInterface) (*WidgetsShell) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell(org.eclipse.swt.widgets.Display, int)
func NewWidgetsShell5(a WidgetsDisplayInterface, b int) (*WidgetsShell) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell(org.eclipse.swt.widgets.Shell)
func NewWidgetsShell4(a WidgetsShellInterface) (*WidgetsShell) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsShell6(a WidgetsShellInterface, b int) (*WidgetsShell) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsShell{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.eclipse.swt.widgets.Shell gtk_new(org.eclipse.swt.widgets.Display, long)
func (jbobject *WidgetsShell) Gtk_new(a WidgetsDisplayInterface, b int64) *WidgetsShell {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Shell", "gtk_new", "org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"), b)
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
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.widgets.Shell internal_new(org.eclipse.swt.widgets.Display, long)
func (jbobject *WidgetsShell) Internal_new(a WidgetsDisplayInterface, b int64) *WidgetsShell {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/widgets/Shell", "internal_new", "org/eclipse/swt/widgets/Shell", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"), b)
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
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public void addShellListener(org.eclipse.swt.events.ShellListener)
func (jbobject *WidgetsShell) AddShellListener(a EventsShellListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addShellListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void close()
func (jbobject *WidgetsShell) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *WidgetsShell) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeTrim", "org/eclipse/swt/graphics/Rectangle", a, b, c, d)
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

// public org.eclipse.swt.widgets.ToolBar getToolBar()
func (jbobject *WidgetsShell) GetToolBar() *WidgetsToolBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToolBar", "org/eclipse/swt/widgets/ToolBar")
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
	unique_x := &WidgetsToolBar{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isEnabled()
func (jbobject *WidgetsShell) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsShell) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getAlpha()
func (jbobject *WidgetsShell) GetAlpha() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlpha", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getFullScreen()
func (jbobject *WidgetsShell) GetFullScreen() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFullScreen", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Point getLocation()
func (jbobject *WidgetsShell) GetLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocation", "org/eclipse/swt/graphics/Point")
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

// public boolean getMaximized()
func (jbobject *WidgetsShell) GetMaximized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Point getMinimumSize()
func (jbobject *WidgetsShell) GetMinimumSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimumSize", "org/eclipse/swt/graphics/Point")
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

// public boolean getModified()
func (jbobject *WidgetsShell) GetModified() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModified", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Point getSize()
func (jbobject *WidgetsShell) GetSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "org/eclipse/swt/graphics/Point")
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

// public boolean getVisible()
func (jbobject *WidgetsShell) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Region getRegion()
func (jbobject *WidgetsShell) GetRegion() *GraphicsRegion {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRegion", "org/eclipse/swt/graphics/Region")
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
	unique_x := &GraphicsRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public int getImeInputMode()
func (jbobject *WidgetsShell) GetImeInputMode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImeInputMode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Shell[] getShells()
func (jbobject *WidgetsShell) GetShells() []*WidgetsShell {
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

// public void open()
func (jbobject *WidgetsShell) Open()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "open", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean print(org.eclipse.swt.graphics.GC)
func (jbobject *WidgetsShell) Print(a GraphicsGCInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void removeShellListener(org.eclipse.swt.events.ShellListener)
func (jbobject *WidgetsShell) RemoveShellListener(a EventsShellListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeShellListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setActive()
func (jbobject *WidgetsShell) SetActive()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setActive", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setAlpha(int)
func (jbobject *WidgetsShell) SetAlpha(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlpha", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEnabled(boolean)
func (jbobject *WidgetsShell) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFullScreen(boolean)
func (jbobject *WidgetsShell) SetFullScreen(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFullScreen", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImeInputMode(int)
func (jbobject *WidgetsShell) SetImeInputMode(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImeInputMode", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMaximized(boolean)
func (jbobject *WidgetsShell) SetMaximized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMenuBar(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsShell) SetMenuBar(a WidgetsMenuInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMenuBar", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMinimized(boolean)
func (jbobject *WidgetsShell) SetMinimized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimumSize(int, int)
func (jbobject *WidgetsShell) SetMinimumSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimumSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setMinimumSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsShell) SetMinimumSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimumSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setModified(boolean)
func (jbobject *WidgetsShell) SetModified(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModified", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setRegion(org.eclipse.swt.graphics.Region)
func (jbobject *WidgetsShell) SetRegion(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegion", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setText(java.lang.String)
func (jbobject *WidgetsShell) SetText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setVisible(boolean)
func (jbobject *WidgetsShell) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void dispose()
func (jbobject *WidgetsShell) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void forceActive()
func (jbobject *WidgetsShell) ForceActive()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "forceActive", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsShell) GetBounds() *GraphicsRectangle {
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


