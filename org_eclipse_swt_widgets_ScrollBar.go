package swt

import "github.com/timob/javabind"

type WidgetsScrollBarInterface interface {
	WidgetsWidgetInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public boolean getEnabled()
	GetEnabled() bool

	// public int getIncrement()
	GetIncrement() int

	// public int getMaximum()
	GetMaximum() int

	// public int getMinimum()
	GetMinimum() int

	// public int getPageIncrement()
	GetPageIncrement() int

	// public org.eclipse.swt.widgets.Scrollable getParent()
	GetParent() *WidgetsScrollable

	// public int getSelection()
	GetSelection() int

	// public org.eclipse.swt.graphics.Point getSize()
	GetSize() *GraphicsPoint

	// public int getThumb()
	GetThumb() int

	// public org.eclipse.swt.graphics.Rectangle getThumbBounds()
	GetThumbBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle getThumbTrackBounds()
	GetThumbTrackBounds() *GraphicsRectangle

	// public boolean getVisible()
	GetVisible() bool

	// public boolean isEnabled()
	IsEnabled() bool

	// public boolean isVisible()
	IsVisible() bool

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setEnabled(boolean)
	SetEnabled(a bool) 

	// public void setIncrement(int)
	SetIncrement(a int) 

	// public void setMaximum(int)
	SetMaximum(a int) 

	// public void setMinimum(int)
	SetMinimum(a int) 

	// public void setPageIncrement(int)
	SetPageIncrement(a int) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setThumb(int)
	SetThumb(a int) 

	// public void setValues(int, int, int, int, int, int)
	SetValues(a int, b int, c int, d int, e int, f int) 

	// public void setVisible(boolean)
	SetVisible(a bool) 
}

type WidgetsScrollBar struct {
	WidgetsWidget
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsScrollBar) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public boolean getEnabled()
func (jbobject *WidgetsScrollBar) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getIncrement()
func (jbobject *WidgetsScrollBar) GetIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMaximum()
func (jbobject *WidgetsScrollBar) GetMaximum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinimum()
func (jbobject *WidgetsScrollBar) GetMinimum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getPageIncrement()
func (jbobject *WidgetsScrollBar) GetPageIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPageIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Scrollable getParent()
func (jbobject *WidgetsScrollBar) GetParent() *WidgetsScrollable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Scrollable")
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
	unique_x := &WidgetsScrollable{}
	unique_x.Callable = dst
	return unique_x
}

// public int getSelection()
func (jbobject *WidgetsScrollBar) GetSelection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getSize()
func (jbobject *WidgetsScrollBar) GetSize() *GraphicsPoint {
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

// public int getThumb()
func (jbobject *WidgetsScrollBar) GetThumb() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThumb", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Rectangle getThumbBounds()
func (jbobject *WidgetsScrollBar) GetThumbBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThumbBounds", "org/eclipse/swt/graphics/Rectangle")
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

// public org.eclipse.swt.graphics.Rectangle getThumbTrackBounds()
func (jbobject *WidgetsScrollBar) GetThumbTrackBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThumbTrackBounds", "org/eclipse/swt/graphics/Rectangle")
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

// public boolean getVisible()
func (jbobject *WidgetsScrollBar) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isEnabled()
func (jbobject *WidgetsScrollBar) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsScrollBar) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsScrollBar) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setEnabled(boolean)
func (jbobject *WidgetsScrollBar) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setIncrement(int)
func (jbobject *WidgetsScrollBar) SetIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMaximum(int)
func (jbobject *WidgetsScrollBar) SetMaximum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimum(int)
func (jbobject *WidgetsScrollBar) SetMinimum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setPageIncrement(int)
func (jbobject *WidgetsScrollBar) SetPageIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPageIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsScrollBar) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setThumb(int)
func (jbobject *WidgetsScrollBar) SetThumb(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThumb", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setValues(int, int, int, int, int, int)
func (jbobject *WidgetsScrollBar) SetValues(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValues", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void setVisible(boolean)
func (jbobject *WidgetsScrollBar) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


