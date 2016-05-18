package swt

import "github.com/timob/javabind"

type WidgetsCoolItemInterface interface {
	WidgetsItemInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public org.eclipse.swt.graphics.Point computeSize(int, int)
	ComputeSize(a int, b int) *GraphicsPoint

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public org.eclipse.swt.graphics.Point getMinimumSize()
	GetMinimumSize() *GraphicsPoint

	// public org.eclipse.swt.widgets.CoolBar getParent()
	GetParent() *WidgetsCoolBar

	// public org.eclipse.swt.graphics.Point getPreferredSize()
	GetPreferredSize() *GraphicsPoint

	// public org.eclipse.swt.graphics.Point getSize()
	GetSize() *GraphicsPoint

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setControl(org.eclipse.swt.widgets.Control)
	SetControl(a WidgetsControlInterface) 

	// public void setMinimumSize(int, int)
	SetMinimumSize(a int, b int) 

	// public void setMinimumSize(org.eclipse.swt.graphics.Point)
	SetMinimumSize2(a GraphicsPointInterface) 

	// public void setPreferredSize(int, int)
	SetPreferredSize(a int, b int) 

	// public void setPreferredSize(org.eclipse.swt.graphics.Point)
	SetPreferredSize2(a GraphicsPointInterface) 

	// public void setSize(int, int)
	SetSize(a int, b int) 

	// public void setSize(org.eclipse.swt.graphics.Point)
	SetSize2(a GraphicsPointInterface) 
}

type WidgetsCoolItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.CoolItem(org.eclipse.swt.widgets.CoolBar, int)
func NewWidgetsCoolItem(a WidgetsCoolBarInterface, b int) (*WidgetsCoolItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/CoolItem", conv_a.Value().Cast("org/eclipse/swt/widgets/CoolBar"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCoolItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.CoolItem(org.eclipse.swt.widgets.CoolBar, int, int)
func NewWidgetsCoolItem2(a WidgetsCoolBarInterface, b int, c int) (*WidgetsCoolItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/CoolItem", conv_a.Value().Cast("org/eclipse/swt/widgets/CoolBar"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCoolItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsCoolItem) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public org.eclipse.swt.graphics.Point computeSize(int, int)
func (jbobject *WidgetsCoolItem) ComputeSize(a int, b int) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeSize", "org/eclipse/swt/graphics/Point", a, b)
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

// public void dispose()
func (jbobject *WidgetsCoolItem) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsCoolItem) GetBounds() *GraphicsRectangle {
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
func (jbobject *WidgetsCoolItem) GetControl() *WidgetsControl {
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

// public org.eclipse.swt.graphics.Point getMinimumSize()
func (jbobject *WidgetsCoolItem) GetMinimumSize() *GraphicsPoint {
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

// public org.eclipse.swt.widgets.CoolBar getParent()
func (jbobject *WidgetsCoolItem) GetParent() *WidgetsCoolBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/CoolBar")
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
	unique_x := &WidgetsCoolBar{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point getPreferredSize()
func (jbobject *WidgetsCoolItem) GetPreferredSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreferredSize", "org/eclipse/swt/graphics/Point")
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

// public org.eclipse.swt.graphics.Point getSize()
func (jbobject *WidgetsCoolItem) GetSize() *GraphicsPoint {
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

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsCoolItem) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsCoolItem) SetControl(a WidgetsControlInterface)  {
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

// public void setMinimumSize(int, int)
func (jbobject *WidgetsCoolItem) SetMinimumSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimumSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setMinimumSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCoolItem) SetMinimumSize2(a GraphicsPointInterface)  {
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

// public void setPreferredSize(int, int)
func (jbobject *WidgetsCoolItem) SetPreferredSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreferredSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setPreferredSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCoolItem) SetPreferredSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreferredSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSize(int, int)
func (jbobject *WidgetsCoolItem) SetSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCoolItem) SetSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


