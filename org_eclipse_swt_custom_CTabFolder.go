package swt

import "github.com/timob/javabind"

type CustomCTabFolderInterface interface {
	WidgetsCompositeInterface

	// public void addCTabFolder2Listener(org.eclipse.swt.custom.CTabFolder2Listener)
	AddCTabFolder2Listener(a CustomCTabFolder2ListenerInterface) 

	// public void addCTabFolderListener(org.eclipse.swt.custom.CTabFolderListener)
	AddCTabFolderListener(a CustomCTabFolderListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public boolean getBorderVisible()
	GetBorderVisible() bool

	// public org.eclipse.swt.custom.CTabItem getItem(int)
	GetItem(a int) *CustomCTabItem

	// public org.eclipse.swt.custom.CTabItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *CustomCTabItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.custom.CTabItem[] getItems()
	GetItems() []*CustomCTabItem

	// public boolean getMinimized()
	GetMinimized() bool

	// public boolean getMinimizeVisible()
	GetMinimizeVisible() bool

	// public int getMinimumCharacters()
	GetMinimumCharacters() int

	// public boolean getMaximized()
	GetMaximized() bool

	// public boolean getMaximizeVisible()
	GetMaximizeVisible() bool

	// public boolean getMRUVisible()
	GetMRUVisible() bool

	// public org.eclipse.swt.custom.CTabFolderRenderer getRenderer()
	GetRenderer() *CustomCTabFolderRenderer

	// public org.eclipse.swt.custom.CTabItem getSelection()
	GetSelection() *CustomCTabItem

	// public org.eclipse.swt.graphics.Color getSelectionBackground()
	GetSelectionBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Color getSelectionForeground()
	GetSelectionForeground() *GraphicsColor

	// public int getSelectionIndex()
	GetSelectionIndex() int

	// public boolean getSimple()
	GetSimple() bool

	// public boolean getSingle()
	GetSingle() bool

	// public int getTabHeight()
	GetTabHeight() int

	// public int getTabPosition()
	GetTabPosition() int

	// public org.eclipse.swt.widgets.Control getTopRight()
	GetTopRight() *WidgetsControl

	// public int getTopRightAlignment()
	GetTopRightAlignment() int

	// public boolean getUnselectedCloseVisible()
	GetUnselectedCloseVisible() bool

	// public boolean getUnselectedImageVisible()
	GetUnselectedImageVisible() bool

	// public int indexOf(org.eclipse.swt.custom.CTabItem)
	IndexOf(a CustomCTabItemInterface) int

	// public void removeCTabFolder2Listener(org.eclipse.swt.custom.CTabFolder2Listener)
	RemoveCTabFolder2Listener(a CustomCTabFolder2ListenerInterface) 

	// public void removeCTabFolderListener(org.eclipse.swt.custom.CTabFolderListener)
	RemoveCTabFolderListener(a CustomCTabFolderListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setBackground(org.eclipse.swt.graphics.Color[], int[])
	SetBackground2(a []*GraphicsColor, b []int) 

	// public void setBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
	SetBackground3(a []*GraphicsColor, b []int, c bool) 

	// public void setBorderVisible(boolean)
	SetBorderVisible(a bool) 

	// public void setInsertMark(org.eclipse.swt.custom.CTabItem, boolean)
	SetInsertMark2(a CustomCTabItemInterface, b bool) 

	// public void setInsertMark(int, boolean)
	SetInsertMark(a int, b bool) 

	// public void setMaximizeVisible(boolean)
	SetMaximizeVisible(a bool) 

	// public void setMaximized(boolean)
	SetMaximized(a bool) 

	// public void setMinimizeVisible(boolean)
	SetMinimizeVisible(a bool) 

	// public void setMinimized(boolean)
	SetMinimized(a bool) 

	// public void setMinimumCharacters(int)
	SetMinimumCharacters(a int) 

	// public void setMRUVisible(boolean)
	SetMRUVisible(a bool) 

	// public void setRenderer(org.eclipse.swt.custom.CTabFolderRenderer)
	SetRenderer(a CustomCTabFolderRendererInterface) 

	// public void setSelection(org.eclipse.swt.custom.CTabItem)
	SetSelection2(a CustomCTabItemInterface) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelectionBackground(org.eclipse.swt.graphics.Color)
	SetSelectionBackground(a GraphicsColorInterface) 

	// public void setSelectionBackground(org.eclipse.swt.graphics.Color[], int[])
	SetSelectionBackground3(a []*GraphicsColor, b []int) 

	// public void setSelectionBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
	SetSelectionBackground4(a []*GraphicsColor, b []int, c bool) 

	// public void setSelectionBackground(org.eclipse.swt.graphics.Image)
	SetSelectionBackground2(a GraphicsImageInterface) 

	// public void setSelectionForeground(org.eclipse.swt.graphics.Color)
	SetSelectionForeground(a GraphicsColorInterface) 

	// public void setSimple(boolean)
	SetSimple(a bool) 

	// public void setSingle(boolean)
	SetSingle(a bool) 

	// public void setTabHeight(int)
	SetTabHeight(a int) 

	// public void setTabPosition(int)
	SetTabPosition(a int) 

	// public void setTopRight(org.eclipse.swt.widgets.Control)
	SetTopRight(a WidgetsControlInterface) 

	// public void setTopRight(org.eclipse.swt.widgets.Control, int)
	SetTopRight2(a WidgetsControlInterface, b int) 

	// public void setUnselectedCloseVisible(boolean)
	SetUnselectedCloseVisible(a bool) 

	// public void setUnselectedImageVisible(boolean)
	SetUnselectedImageVisible(a bool) 

	// public void showItem(org.eclipse.swt.custom.CTabItem)
	ShowItem(a CustomCTabItemInterface) 

	// public void showSelection()
	ShowSelection() 
}

type CustomCTabFolder struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.CTabFolder(org.eclipse.swt.widgets.Composite, int)
func NewCustomCTabFolder(a WidgetsCompositeInterface, b int) (*CustomCTabFolder) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabFolder", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCTabFolder{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addCTabFolder2Listener(org.eclipse.swt.custom.CTabFolder2Listener)
func (jbobject *CustomCTabFolder) AddCTabFolder2Listener(a CustomCTabFolder2ListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addCTabFolder2Listener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolder2Listener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addCTabFolderListener(org.eclipse.swt.custom.CTabFolderListener)
func (jbobject *CustomCTabFolder) AddCTabFolderListener(a CustomCTabFolderListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addCTabFolderListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomCTabFolder) AddSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *CustomCTabFolder) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
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

// public boolean getBorderVisible()
func (jbobject *CustomCTabFolder) GetBorderVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *CustomCTabFolder) GetClientArea() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientArea", "org/eclipse/swt/graphics/Rectangle")
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

// public org.eclipse.swt.custom.CTabItem getItem(int)
func (jbobject *CustomCTabFolder) GetItem(a int) *CustomCTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/CTabItem", a)
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
	unique_x := &CustomCTabItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.custom.CTabItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *CustomCTabFolder) GetItem2(a GraphicsPointInterface) *CustomCTabItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/CTabItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
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
	unique_x := &CustomCTabItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *CustomCTabFolder) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.custom.CTabItem[] getItems()
func (jbobject *CustomCTabFolder) GetItems() []*CustomCTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/custom/CTabItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/CTabItem")
	dst := new([]*CustomCTabItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getMinimized()
func (jbobject *CustomCTabFolder) GetMinimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getMinimizeVisible()
func (jbobject *CustomCTabFolder) GetMinimizeVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimizeVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getMinimumCharacters()
func (jbobject *CustomCTabFolder) GetMinimumCharacters() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimumCharacters", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getMaximized()
func (jbobject *CustomCTabFolder) GetMaximized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getMaximizeVisible()
func (jbobject *CustomCTabFolder) GetMaximizeVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximizeVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getMRUVisible()
func (jbobject *CustomCTabFolder) GetMRUVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMRUVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.custom.CTabFolderRenderer getRenderer()
func (jbobject *CustomCTabFolder) GetRenderer() *CustomCTabFolderRenderer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRenderer", "org/eclipse/swt/custom/CTabFolderRenderer")
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
	unique_x := &CustomCTabFolderRenderer{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.custom.CTabItem getSelection()
func (jbobject *CustomCTabFolder) GetSelection() *CustomCTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", "org/eclipse/swt/custom/CTabItem")
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
	unique_x := &CustomCTabItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Color getSelectionBackground()
func (jbobject *CustomCTabFolder) GetSelectionBackground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionBackground", "org/eclipse/swt/graphics/Color")
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

// public org.eclipse.swt.graphics.Color getSelectionForeground()
func (jbobject *CustomCTabFolder) GetSelectionForeground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionForeground", "org/eclipse/swt/graphics/Color")
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

// public int getSelectionIndex()
func (jbobject *CustomCTabFolder) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getSimple()
func (jbobject *CustomCTabFolder) GetSimple() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSimple", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getSingle()
func (jbobject *CustomCTabFolder) GetSingle() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSingle", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getStyle()
func (jbobject *CustomCTabFolder) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTabHeight()
func (jbobject *CustomCTabFolder) GetTabHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTabPosition()
func (jbobject *CustomCTabFolder) GetTabPosition() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabPosition", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Control getTopRight()
func (jbobject *CustomCTabFolder) GetTopRight() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopRight", "org/eclipse/swt/widgets/Control")
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

// public int getTopRightAlignment()
func (jbobject *CustomCTabFolder) GetTopRightAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopRightAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getUnselectedCloseVisible()
func (jbobject *CustomCTabFolder) GetUnselectedCloseVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUnselectedCloseVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getUnselectedImageVisible()
func (jbobject *CustomCTabFolder) GetUnselectedImageVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUnselectedImageVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int indexOf(org.eclipse.swt.custom.CTabItem)
func (jbobject *CustomCTabFolder) IndexOf(a CustomCTabItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/custom/CTabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void removeCTabFolder2Listener(org.eclipse.swt.custom.CTabFolder2Listener)
func (jbobject *CustomCTabFolder) RemoveCTabFolder2Listener(a CustomCTabFolder2ListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeCTabFolder2Listener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolder2Listener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeCTabFolderListener(org.eclipse.swt.custom.CTabFolderListener)
func (jbobject *CustomCTabFolder) RemoveCTabFolderListener(a CustomCTabFolderListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeCTabFolderListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomCTabFolder) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void reskin(int)
func (jbobject *CustomCTabFolder) Reskin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reskin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCTabFolder) SetBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackground(org.eclipse.swt.graphics.Color[], int[])
func (jbobject *CustomCTabFolder) SetBackground2(a []*GraphicsColor, b []int)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
func (jbobject *CustomCTabFolder) SetBackground3(a []*GraphicsColor, b []int, c bool)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackgroundImage(org.eclipse.swt.graphics.Image)
func (jbobject *CustomCTabFolder) SetBackgroundImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackgroundImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBorderVisible(boolean)
func (jbobject *CustomCTabFolder) SetBorderVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBorderVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomCTabFolder) SetFont(a GraphicsFontInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFont", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Font"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCTabFolder) SetForeground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForeground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInsertMark(org.eclipse.swt.custom.CTabItem, boolean)
func (jbobject *CustomCTabFolder) SetInsertMark2(a CustomCTabItemInterface, b bool)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInsertMark", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInsertMark(int, boolean)
func (jbobject *CustomCTabFolder) SetInsertMark(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInsertMark", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setMaximizeVisible(boolean)
func (jbobject *CustomCTabFolder) SetMaximizeVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximizeVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomCTabFolder) SetLayout(a WidgetsLayoutInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLayout", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Layout"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMaximized(boolean)
func (jbobject *CustomCTabFolder) SetMaximized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimizeVisible(boolean)
func (jbobject *CustomCTabFolder) SetMinimizeVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimizeVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimized(boolean)
func (jbobject *CustomCTabFolder) SetMinimized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimumCharacters(int)
func (jbobject *CustomCTabFolder) SetMinimumCharacters(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimumCharacters", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMRUVisible(boolean)
func (jbobject *CustomCTabFolder) SetMRUVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMRUVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setRenderer(org.eclipse.swt.custom.CTabFolderRenderer)
func (jbobject *CustomCTabFolder) SetRenderer(a CustomCTabFolderRendererInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRenderer", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderRenderer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(org.eclipse.swt.custom.CTabItem)
func (jbobject *CustomCTabFolder) SetSelection2(a CustomCTabItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(int)
func (jbobject *CustomCTabFolder) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelectionBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCTabFolder) SetSelectionBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelectionBackground(org.eclipse.swt.graphics.Color[], int[])
func (jbobject *CustomCTabFolder) SetSelectionBackground3(a []*GraphicsColor, b []int)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelectionBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
func (jbobject *CustomCTabFolder) SetSelectionBackground4(a []*GraphicsColor, b []int, c bool)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelectionBackground(org.eclipse.swt.graphics.Image)
func (jbobject *CustomCTabFolder) SetSelectionBackground2(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelectionForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCTabFolder) SetSelectionForeground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionForeground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSimple(boolean)
func (jbobject *CustomCTabFolder) SetSimple(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSimple", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSingle(boolean)
func (jbobject *CustomCTabFolder) SetSingle(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSingle", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTabHeight(int)
func (jbobject *CustomCTabFolder) SetTabHeight(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabHeight", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTabPosition(int)
func (jbobject *CustomCTabFolder) SetTabPosition(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabPosition", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopRight(org.eclipse.swt.widgets.Control)
func (jbobject *CustomCTabFolder) SetTopRight(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopRight", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTopRight(org.eclipse.swt.widgets.Control, int)
func (jbobject *CustomCTabFolder) SetTopRight2(a WidgetsControlInterface, b int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopRight", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setUnselectedCloseVisible(boolean)
func (jbobject *CustomCTabFolder) SetUnselectedCloseVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUnselectedCloseVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setUnselectedImageVisible(boolean)
func (jbobject *CustomCTabFolder) SetUnselectedImageVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUnselectedImageVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showItem(org.eclipse.swt.custom.CTabItem)
func (jbobject *CustomCTabFolder) ShowItem(a CustomCTabItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showSelection()
func (jbobject *CustomCTabFolder) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolder) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolder) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolder) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolder) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolder) MIN_TAB_WIDTH() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "MIN_TAB_WIDTH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolder) SetFieldMIN_TAB_WIDTH(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "MIN_TAB_WIDTH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolder) BorderInsideRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolder", "borderInsideRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomCTabFolder) SetFieldBorderInsideRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolder", "borderInsideRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomCTabFolder) BorderMiddleRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolder", "borderMiddleRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomCTabFolder) SetFieldBorderMiddleRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolder", "borderMiddleRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomCTabFolder) BorderOutsideRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolder", "borderOutsideRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomCTabFolder) SetFieldBorderOutsideRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolder", "borderOutsideRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


