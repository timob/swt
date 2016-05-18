package swt

import "github.com/timob/javabind"

type CustomTableCursorInterface interface {
	WidgetsCanvasInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public int getColumn()
	GetColumn() int

	// public org.eclipse.swt.widgets.TableItem getRow()
	GetRow() *WidgetsTableItem

	// public void setSelection(int, int)
	SetSelection(a int, b int) 

	// public void setSelection(org.eclipse.swt.widgets.TableItem, int)
	SetSelection2(a WidgetsTableItemInterface, b int) 
}

type CustomTableCursor struct {
	WidgetsCanvas
}

// public org.eclipse.swt.custom.TableCursor(org.eclipse.swt.widgets.Table, int)
func NewCustomTableCursor(a WidgetsTableInterface, b int) (*CustomTableCursor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableCursor", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableCursor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTableCursor) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setVisible(boolean)
func (jbobject *CustomTableCursor) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTableCursor) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public int getColumn()
func (jbobject *CustomTableCursor) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *CustomTableCursor) GetBackground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackground", "org/eclipse/swt/graphics/Color")
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
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *CustomTableCursor) GetForeground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getForeground", "org/eclipse/swt/graphics/Color")
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
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TableItem getRow()
func (jbobject *CustomTableCursor) GetRow() *WidgetsTableItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRow", "org/eclipse/swt/widgets/TableItem")
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
	unique_x := &WidgetsTableItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomTableCursor) SetBackground(a GraphicsColorInterface)  {
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

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomTableCursor) SetForeground(a GraphicsColorInterface)  {
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

// public void setSelection(int, int)
func (jbobject *CustomTableCursor) SetSelection(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.widgets.TableItem, int)
func (jbobject *CustomTableCursor) SetSelection2(a WidgetsTableItemInterface, b int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


