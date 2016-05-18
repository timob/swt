package swt

import "github.com/timob/javabind"

type WidgetsSpinnerInterface interface {
	WidgetsCompositeInterface

	// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
	AddModifyListener(a EventsModifyListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void copy()
	Copy() 

	// public void cut()
	Cut() 

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

	// public java.lang.String getText()
	GetText() string

	// public int getTextLimit()
	GetTextLimit() int

	// public int getDigits()
	GetDigits() int

	// public void paste()
	Paste() 

	// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
	RemoveModifyListener(a EventsModifyListenerInterface) 

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

	// public void setTextLimit(int)
	SetTextLimit(a int) 

	// public void setDigits(int)
	SetDigits(a int) 

	// public void setValues(int, int, int, int, int, int)
	SetValues(a int, b int, c int, d int, e int, f int) 
}

type WidgetsSpinner struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.Spinner(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsSpinner(a WidgetsCompositeInterface, b int) (*WidgetsSpinner) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Spinner", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsSpinner{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *WidgetsSpinner) AddModifyListener(a EventsModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsSpinner) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsSpinner) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *WidgetsSpinner) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
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

// public void copy()
func (jbobject *WidgetsSpinner) Copy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cut()
func (jbobject *WidgetsSpinner) Cut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getBorderWidth()
func (jbobject *WidgetsSpinner) GetBorderWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getIncrement()
func (jbobject *WidgetsSpinner) GetIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMaximum()
func (jbobject *WidgetsSpinner) GetMaximum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinimum()
func (jbobject *WidgetsSpinner) GetMinimum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getPageIncrement()
func (jbobject *WidgetsSpinner) GetPageIncrement() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPageIncrement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelection()
func (jbobject *WidgetsSpinner) GetSelection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *WidgetsSpinner) GetText() string {
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

// public int getTextLimit()
func (jbobject *WidgetsSpinner) GetTextLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getDigits()
func (jbobject *WidgetsSpinner) GetDigits() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDigits", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void paste()
func (jbobject *WidgetsSpinner) Paste()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paste", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *WidgetsSpinner) RemoveModifyListener(a EventsModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsSpinner) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsSpinner) SetIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMaximum(int)
func (jbobject *WidgetsSpinner) SetMaximum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimum(int)
func (jbobject *WidgetsSpinner) SetMinimum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setPageIncrement(int)
func (jbobject *WidgetsSpinner) SetPageIncrement(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPageIncrement", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsSpinner) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTextLimit(int)
func (jbobject *WidgetsSpinner) SetTextLimit(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextLimit", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setDigits(int)
func (jbobject *WidgetsSpinner) SetDigits(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDigits", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setValues(int, int, int, int, int, int)
func (jbobject *WidgetsSpinner) SetValues(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValues", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsSpinner) LIMIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/widgets/Spinner", "LIMIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsSpinner) SetFieldLIMIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/widgets/Spinner", "LIMIT", val)
	if err != nil {
		panic(err)
	}

}


