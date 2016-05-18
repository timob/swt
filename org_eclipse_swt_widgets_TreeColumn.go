package swt

import "github.com/timob/javabind"

type WidgetsTreeColumnInterface interface {
	WidgetsItemInterface

	// public void addControlListener(org.eclipse.swt.events.ControlListener)
	AddControlListener(a EventsControlListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getAlignment()
	GetAlignment() int

	// public boolean getMoveable()
	GetMoveable() bool

	// public org.eclipse.swt.widgets.Tree getParent()
	GetParent() *WidgetsTree

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

	// public void setMoveable(boolean)
	SetMoveable(a bool) 

	// public void setResizable(boolean)
	SetResizable(a bool) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 

	// public void setWidth(int)
	SetWidth(a int) 
}

type WidgetsTreeColumn struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TreeColumn(org.eclipse.swt.widgets.Tree, int)
func NewWidgetsTreeColumn(a WidgetsTreeInterface, b int) (*WidgetsTreeColumn) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeColumn", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeColumn{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TreeColumn(org.eclipse.swt.widgets.Tree, int, int)
func NewWidgetsTreeColumn2(a WidgetsTreeInterface, b int, c int) (*WidgetsTreeColumn) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeColumn", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeColumn{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTreeColumn) AddControlListener(a EventsControlListenerInterface)  {
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
func (jbobject *WidgetsTreeColumn) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsTreeColumn) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getMoveable()
func (jbobject *WidgetsTreeColumn) GetMoveable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMoveable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.Tree getParent()
func (jbobject *WidgetsTreeColumn) GetParent() *WidgetsTree {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Tree")
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
	unique_x := &WidgetsTree{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getResizable()
func (jbobject *WidgetsTreeColumn) GetResizable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResizable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getToolTipText()
func (jbobject *WidgetsTreeColumn) GetToolTipText() string {
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
func (jbobject *WidgetsTreeColumn) GetWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void pack()
func (jbobject *WidgetsTreeColumn) Pack()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "pack", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTreeColumn) RemoveControlListener(a EventsControlListenerInterface)  {
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
func (jbobject *WidgetsTreeColumn) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsTreeColumn) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTreeColumn) SetImage(a GraphicsImageInterface)  {
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

// public void setMoveable(boolean)
func (jbobject *WidgetsTreeColumn) SetMoveable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMoveable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setResizable(boolean)
func (jbobject *WidgetsTreeColumn) SetResizable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResizable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsTreeColumn) SetText(a string)  {
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
func (jbobject *WidgetsTreeColumn) SetToolTipText(a string)  {
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
func (jbobject *WidgetsTreeColumn) SetWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


