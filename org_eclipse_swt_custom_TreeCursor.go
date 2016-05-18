package swt

import "github.com/timob/javabind"

type CustomTreeCursorInterface interface {
	WidgetsCanvasInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getColumn()
	GetColumn() int

	// public org.eclipse.swt.widgets.TreeItem getRow()
	GetRow() *WidgetsTreeItem

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setSelection(int, int)
	SetSelection(a int, b int) 

	// public void setSelection(org.eclipse.swt.widgets.TreeItem, int)
	SetSelection2(a WidgetsTreeItemInterface, b int) 
}

type CustomTreeCursor struct {
	WidgetsCanvas
}

// public org.eclipse.swt.custom.TreeCursor(org.eclipse.swt.widgets.Tree, int)
func NewCustomTreeCursor(a WidgetsTreeInterface, b int) (*CustomTreeCursor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TreeCursor", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTreeCursor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTreeCursor) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *CustomTreeCursor) GetBackground() *GraphicsColor {
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

// public int getColumn()
func (jbobject *CustomTreeCursor) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *CustomTreeCursor) GetForeground() *GraphicsColor {
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

// public org.eclipse.swt.widgets.TreeItem getRow()
func (jbobject *CustomTreeCursor) GetRow() *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRow", "org/eclipse/swt/widgets/TreeItem")
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
	unique_x := &WidgetsTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTreeCursor) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *CustomTreeCursor) SetBackground(a GraphicsColorInterface)  {
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
func (jbobject *CustomTreeCursor) SetForeground(a GraphicsColorInterface)  {
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
func (jbobject *CustomTreeCursor) SetSelection(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.widgets.TreeItem, int)
func (jbobject *CustomTreeCursor) SetSelection2(a WidgetsTreeItemInterface, b int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setVisible(boolean)
func (jbobject *CustomTreeCursor) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


