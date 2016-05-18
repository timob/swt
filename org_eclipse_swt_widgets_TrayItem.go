package swt

import "github.com/timob/javabind"

type WidgetsTrayItemInterface interface {
	WidgetsItemInterface

	// public void addMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
	AddMenuDetectListener(a EventsMenuDetectListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public org.eclipse.swt.widgets.Tray getParent()
	GetParent() *WidgetsTray

	// public org.eclipse.swt.graphics.Image getHighlightImage()
	GetHighlightImage() *GraphicsImage

	// public org.eclipse.swt.widgets.ToolTip getToolTip()
	GetToolTip() *WidgetsToolTip

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public boolean getVisible()
	GetVisible() bool

	// public void removeMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
	RemoveMenuDetectListener(a EventsMenuDetectListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setHighlightImage(org.eclipse.swt.graphics.Image)
	SetHighlightImage(a GraphicsImageInterface) 

	// public void setToolTip(org.eclipse.swt.widgets.ToolTip)
	SetToolTip(a WidgetsToolTipInterface) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 

	// public void setVisible(boolean)
	SetVisible(a bool) 
}

type WidgetsTrayItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TrayItem(org.eclipse.swt.widgets.Tray, int)
func NewWidgetsTrayItem(a WidgetsTrayInterface, b int) (*WidgetsTrayItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TrayItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Tray"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTrayItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
func (jbobject *WidgetsTrayItem) AddMenuDetectListener(a EventsMenuDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMenuDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTrayItem) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public org.eclipse.swt.widgets.Tray getParent()
func (jbobject *WidgetsTrayItem) GetParent() *WidgetsTray {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Tray")
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

// public org.eclipse.swt.graphics.Image getHighlightImage()
func (jbobject *WidgetsTrayItem) GetHighlightImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHighlightImage", "org/eclipse/swt/graphics/Image")
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

// public org.eclipse.swt.widgets.ToolTip getToolTip()
func (jbobject *WidgetsTrayItem) GetToolTip() *WidgetsToolTip {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToolTip", "org/eclipse/swt/widgets/ToolTip")
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
	unique_x := &WidgetsToolTip{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getToolTipText()
func (jbobject *WidgetsTrayItem) GetToolTipText() string {
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

// public boolean getVisible()
func (jbobject *WidgetsTrayItem) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
func (jbobject *WidgetsTrayItem) RemoveMenuDetectListener(a EventsMenuDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMenuDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTrayItem) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setHighlightImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTrayItem) SetHighlightImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHighlightImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTrayItem) SetImage(a GraphicsImageInterface)  {
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

// public void setToolTip(org.eclipse.swt.widgets.ToolTip)
func (jbobject *WidgetsTrayItem) SetToolTip(a WidgetsToolTipInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setToolTip", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/ToolTip"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setToolTipText(java.lang.String)
func (jbobject *WidgetsTrayItem) SetToolTipText(a string)  {
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

// public void setVisible(boolean)
func (jbobject *WidgetsTrayItem) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


