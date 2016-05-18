package swt

import "github.com/timob/javabind"

type WidgetsSliderInterface interface {
	WidgetsControlInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getIncrement()
	GetIncrement() int

	// public int getMaximum()
	GetMaximum() int

	// public int getMinimum()
	GetMinimum() int

	// public int getPageIncrement()
	GetPageIncrement() int

	// public int getSelection()
	GetSelection() int

	// public int getThumb()
	GetThumb() int

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

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
}

type WidgetsSlider struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.Slider(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsSlider(a WidgetsCompositeInterface, b int) (*WidgetsSlider) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Slider", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsSlider{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsSlider) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsSlider) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeSize", "org/eclipse/swt/graphics/Point", a, b, c)
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

// public int getIncrement()
func (jbobject *WidgetsSlider) GetIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMaximum()
func (jbobject *WidgetsSlider) GetMaximum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinimum()
func (jbobject *WidgetsSlider) GetMinimum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getPageIncrement()
func (jbobject *WidgetsSlider) GetPageIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPageIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelection()
func (jbobject *WidgetsSlider) GetSelection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getThumb()
func (jbobject *WidgetsSlider) GetThumb() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThumb", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsSlider) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setIncrement(int)
func (jbobject *WidgetsSlider) SetIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMaximum(int)
func (jbobject *WidgetsSlider) SetMaximum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimum(int)
func (jbobject *WidgetsSlider) SetMinimum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setPageIncrement(int)
func (jbobject *WidgetsSlider) SetPageIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPageIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsSlider) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setThumb(int)
func (jbobject *WidgetsSlider) SetThumb(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThumb", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setValues(int, int, int, int, int, int)
func (jbobject *WidgetsSlider) SetValues(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValues", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}


