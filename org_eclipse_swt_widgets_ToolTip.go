package swt

import "github.com/timob/javabind"

type WidgetsToolTipInterface interface {
	WidgetsWidgetInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public boolean getAutoHide()
	GetAutoHide() bool

	// public java.lang.String getMessage()
	GetMessage() string

	// public org.eclipse.swt.widgets.Shell getParent()
	GetParent() *WidgetsShell

	// public java.lang.String getText()
	GetText() string

	// public boolean getVisible()
	GetVisible() bool

	// public boolean isVisible()
	IsVisible() bool

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setAutoHide(boolean)
	SetAutoHide(a bool) 

	// public void setLocation(int, int)
	SetLocation(a int, b int) 

	// public void setLocation(org.eclipse.swt.graphics.Point)
	SetLocation2(a GraphicsPointInterface) 

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setVisible(boolean)
	SetVisible(a bool) 
}

type WidgetsToolTip struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.ToolTip(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsToolTip(a WidgetsShellInterface, b int) (*WidgetsToolTip) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ToolTip", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsToolTip{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsToolTip) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public boolean getAutoHide()
func (jbobject *WidgetsToolTip) GetAutoHide() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAutoHide", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getMessage()
func (jbobject *WidgetsToolTip) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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

// public org.eclipse.swt.widgets.Shell getParent()
func (jbobject *WidgetsToolTip) GetParent() *WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Shell")
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

// public java.lang.String getText()
func (jbobject *WidgetsToolTip) GetText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getText", "java/lang/String")
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
func (jbobject *WidgetsToolTip) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsToolTip) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsToolTip) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setAutoHide(boolean)
func (jbobject *WidgetsToolTip) SetAutoHide(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoHide", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLocation(int, int)
func (jbobject *WidgetsToolTip) SetLocation(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setLocation(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsToolTip) SetLocation2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMessage(java.lang.String)
func (jbobject *WidgetsToolTip) SetMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setText(java.lang.String)
func (jbobject *WidgetsToolTip) SetText(a string)  {
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
func (jbobject *WidgetsToolTip) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


