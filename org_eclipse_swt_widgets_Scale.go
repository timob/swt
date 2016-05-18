package swt

import "github.com/timob/javabind"

type WidgetsScaleInterface interface {
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
}

type WidgetsScale struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.Scale(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsScale(a WidgetsCompositeInterface, b int) (*WidgetsScale) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Scale", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsScale{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsScale) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsScale) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsScale) GetIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMaximum()
func (jbobject *WidgetsScale) GetMaximum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinimum()
func (jbobject *WidgetsScale) GetMinimum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getPageIncrement()
func (jbobject *WidgetsScale) GetPageIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPageIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelection()
func (jbobject *WidgetsScale) GetSelection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsScale) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsScale) SetIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMaximum(int)
func (jbobject *WidgetsScale) SetMaximum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimum(int)
func (jbobject *WidgetsScale) SetMinimum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setPageIncrement(int)
func (jbobject *WidgetsScale) SetPageIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPageIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsScale) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


