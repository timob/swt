package swt

import "github.com/timob/javabind"

type WidgetsDateTimeInterface interface {
	WidgetsCompositeInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getDay()
	GetDay() int

	// public int getHours()
	GetHours() int

	// public int getMinutes()
	GetMinutes() int

	// public int getMonth()
	GetMonth() int

	// public int getSeconds()
	GetSeconds() int

	// public int getYear()
	GetYear() int

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setDate(int, int, int)
	SetDate(a int, b int, c int) 

	// public void setDay(int)
	SetDay(a int) 

	// public void setHours(int)
	SetHours(a int) 

	// public void setMinutes(int)
	SetMinutes(a int) 

	// public void setMonth(int)
	SetMonth(a int) 

	// public void setSeconds(int)
	SetSeconds(a int) 

	// public void setTime(int, int, int)
	SetTime(a int, b int, c int) 

	// public void setYear(int)
	SetYear(a int) 
}

type WidgetsDateTime struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.DateTime(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsDateTime(a WidgetsCompositeInterface, b int) (*WidgetsDateTime) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/DateTime", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDateTime{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsDateTime) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsDateTime) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public int getDay()
func (jbobject *WidgetsDateTime) GetDay() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDay", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHours()
func (jbobject *WidgetsDateTime) GetHours() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHours", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinutes()
func (jbobject *WidgetsDateTime) GetMinutes() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinutes", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMonth()
func (jbobject *WidgetsDateTime) GetMonth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSeconds()
func (jbobject *WidgetsDateTime) GetSeconds() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSeconds", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getYear()
func (jbobject *WidgetsDateTime) GetYear() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getYear", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isFocusControl()
func (jbobject *WidgetsDateTime) IsFocusControl() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isFocusControl", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsDateTime) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsDateTime) SetBackground(a GraphicsColorInterface)  {
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

// public void setEnabled(boolean)
func (jbobject *WidgetsDateTime) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsDateTime) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *WidgetsDateTime) SetForeground(a GraphicsColorInterface)  {
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

// public void setDate(int, int, int)
func (jbobject *WidgetsDateTime) SetDate(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDate", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setDay(int)
func (jbobject *WidgetsDateTime) SetDay(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDay", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHours(int)
func (jbobject *WidgetsDateTime) SetHours(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHours", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsDateTime) SetMenu(a WidgetsMenuInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMenu", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMinutes(int)
func (jbobject *WidgetsDateTime) SetMinutes(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinutes", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMonth(int)
func (jbobject *WidgetsDateTime) SetMonth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSeconds(int)
func (jbobject *WidgetsDateTime) SetSeconds(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSeconds", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTime(int, int, int)
func (jbobject *WidgetsDateTime) SetTime(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTime", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setYear(int)
func (jbobject *WidgetsDateTime) SetYear(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setYear", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


