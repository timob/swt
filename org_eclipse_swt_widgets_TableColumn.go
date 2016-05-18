package swt

import "github.com/timob/javabind"

type WidgetsTableColumnInterface interface {
	WidgetsItemInterface

	// public void addControlListener(org.eclipse.swt.events.ControlListener)
	AddControlListener(a EventsControlListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getAlignment()
	GetAlignment() int

	// public boolean getMoveable()
	GetMoveable() bool

	// public org.eclipse.swt.widgets.Table getParent()
	GetParent() *WidgetsTable

	// public boolean getResizable()
	GetResizable() bool

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public int getWidth()
	GetWidth() int

	// public void pack()
	Pack() 

	// public void removeControlListener(org.eclipse.swt.events.ControlListener)
	RemoveControlListener(a EventsControlListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setResizable(boolean)
	SetResizable(a bool) 

	// public void setMoveable(boolean)
	SetMoveable(a bool) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 

	// public void setWidth(int)
	SetWidth(a int) 
}

type WidgetsTableColumn struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TableColumn(org.eclipse.swt.widgets.Table, int)
func NewWidgetsTableColumn(a WidgetsTableInterface, b int) (*WidgetsTableColumn) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TableColumn", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTableColumn{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TableColumn(org.eclipse.swt.widgets.Table, int, int)
func NewWidgetsTableColumn2(a WidgetsTableInterface, b int, c int) (*WidgetsTableColumn) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TableColumn", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTableColumn{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTableColumn) AddControlListener(a EventsControlListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addControlListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTableColumn) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public int getAlignment()
func (jbobject *WidgetsTableColumn) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getMoveable()
func (jbobject *WidgetsTableColumn) GetMoveable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMoveable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.Table getParent()
func (jbobject *WidgetsTableColumn) GetParent() *WidgetsTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Table")
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
	unique_x := &WidgetsTable{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getResizable()
func (jbobject *WidgetsTableColumn) GetResizable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResizable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getToolTipText()
func (jbobject *WidgetsTableColumn) GetToolTipText() string {
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
func (jbobject *WidgetsTableColumn) GetWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void pack()
func (jbobject *WidgetsTableColumn) Pack()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "pack", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTableColumn) RemoveControlListener(a EventsControlListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeControlListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTableColumn) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setAlignment(int)
func (jbobject *WidgetsTableColumn) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTableColumn) SetImage(a GraphicsImageInterface)  {
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

// public void setResizable(boolean)
func (jbobject *WidgetsTableColumn) SetResizable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResizable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMoveable(boolean)
func (jbobject *WidgetsTableColumn) SetMoveable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMoveable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsTableColumn) SetText(a string)  {
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
func (jbobject *WidgetsTableColumn) SetToolTipText(a string)  {
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
func (jbobject *WidgetsTableColumn) SetWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


