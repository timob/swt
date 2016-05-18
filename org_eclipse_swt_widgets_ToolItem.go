package swt

import "github.com/timob/javabind"

type WidgetsToolItemInterface interface {
	WidgetsItemInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public org.eclipse.swt.graphics.Image getDisabledImage()
	GetDisabledImage() *GraphicsImage

	// public boolean getEnabled()
	GetEnabled() bool

	// public org.eclipse.swt.graphics.Image getHotImage()
	GetHotImage() *GraphicsImage

	// public org.eclipse.swt.widgets.ToolBar getParent()
	GetParent() *WidgetsToolBar

	// public boolean getSelection()
	GetSelection() bool

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public int getWidth()
	GetWidth() int

	// public boolean isEnabled()
	IsEnabled() bool

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setControl(org.eclipse.swt.widgets.Control)
	SetControl(a WidgetsControlInterface) 

	// public void setDisabledImage(org.eclipse.swt.graphics.Image)
	SetDisabledImage(a GraphicsImageInterface) 

	// public void setEnabled(boolean)
	SetEnabled(a bool) 

	// public void setHotImage(org.eclipse.swt.graphics.Image)
	SetHotImage(a GraphicsImageInterface) 

	// public void setSelection(boolean)
	SetSelection(a bool) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 

	// public void setWidth(int)
	SetWidth(a int) 
}

type WidgetsToolItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.ToolItem(org.eclipse.swt.widgets.ToolBar, int)
func NewWidgetsToolItem(a WidgetsToolBarInterface, b int) (*WidgetsToolItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ToolItem", conv_a.Value().Cast("org/eclipse/swt/widgets/ToolBar"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsToolItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.ToolItem(org.eclipse.swt.widgets.ToolBar, int, int)
func NewWidgetsToolItem2(a WidgetsToolBarInterface, b int, c int) (*WidgetsToolItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ToolItem", conv_a.Value().Cast("org/eclipse/swt/widgets/ToolBar"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsToolItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsToolItem) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void dispose()
func (jbobject *WidgetsToolItem) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsToolItem) GetBounds() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.Control getControl()
func (jbobject *WidgetsToolItem) GetControl() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getControl", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.graphics.Image getDisabledImage()
func (jbobject *WidgetsToolItem) GetDisabledImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisabledImage", "org/eclipse/swt/graphics/Image")
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

// public boolean getEnabled()
func (jbobject *WidgetsToolItem) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Image getHotImage()
func (jbobject *WidgetsToolItem) GetHotImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHotImage", "org/eclipse/swt/graphics/Image")
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

// public org.eclipse.swt.widgets.ToolBar getParent()
func (jbobject *WidgetsToolItem) GetParent() *WidgetsToolBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/ToolBar")
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

// public boolean getSelection()
func (jbobject *WidgetsToolItem) GetSelection() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getToolTipText()
func (jbobject *WidgetsToolItem) GetToolTipText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToolTipText", "java/lang/String")
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

// public int getWidth()
func (jbobject *WidgetsToolItem) GetWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isEnabled()
func (jbobject *WidgetsToolItem) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsToolItem) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setControl(org.eclipse.swt.widgets.Control)
func (jbobject *WidgetsToolItem) SetControl(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setControl", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setDisabledImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsToolItem) SetDisabledImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDisabledImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEnabled(boolean)
func (jbobject *WidgetsToolItem) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHotImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsToolItem) SetHotImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHotImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsToolItem) SetImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(boolean)
func (jbobject *WidgetsToolItem) SetSelection(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsToolItem) SetText(a string)  {
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

// public void setToolTipText(java.lang.String)
func (jbobject *WidgetsToolItem) SetToolTipText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setToolTipText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setWidth(int)
func (jbobject *WidgetsToolItem) SetWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


