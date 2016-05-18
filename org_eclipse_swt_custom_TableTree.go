package swt

import "github.com/timob/javabind"

type CustomTableTreeInterface interface {
	WidgetsCompositeInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void addTreeListener(org.eclipse.swt.events.TreeListener)
	AddTreeListener(a EventsTreeListenerInterface) 

	// public void deselectAll()
	DeselectAll() 

	// public int getItemCount()
	GetItemCount() int

	// public int getItemHeight()
	GetItemHeight() int

	// public org.eclipse.swt.custom.TableTreeItem[] getItems()
	GetItems() []*CustomTableTreeItem

	// public org.eclipse.swt.custom.TableTreeItem[] getSelection()
	GetSelection() []*CustomTableTreeItem

	// public int getSelectionCount()
	GetSelectionCount() int

	// public org.eclipse.swt.widgets.Table getTable()
	GetTable() *WidgetsTable

	// public int indexOf(org.eclipse.swt.custom.TableTreeItem)
	IndexOf(a CustomTableTreeItemInterface) int

	// public org.eclipse.swt.custom.TableTreeItem getItem(int)
	GetItem(a int) *CustomTableTreeItem

	// public org.eclipse.swt.custom.TableTreeItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *CustomTableTreeItem

	// public void removeAll()
	RemoveAll() 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeTreeListener(org.eclipse.swt.events.TreeListener)
	RemoveTreeListener(a EventsTreeListenerInterface) 

	// public void selectAll()
	SelectAll() 

	// public void setSelection(org.eclipse.swt.custom.TableTreeItem[])
	SetSelection(a []*CustomTableTreeItem) 

	// public void showItem(org.eclipse.swt.custom.TableTreeItem)
	ShowItem(a CustomTableTreeItemInterface) 

	// public void showSelection()
	ShowSelection() 
}

type CustomTableTree struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.TableTree(org.eclipse.swt.widgets.Composite, int)
func NewCustomTableTree(a WidgetsCompositeInterface, b int) (*CustomTableTree) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTree", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTree{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTableTree) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void addTreeListener(org.eclipse.swt.events.TreeListener)
func (jbobject *CustomTableTree) AddTreeListener(a EventsTreeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addTreeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TreeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *CustomTableTree) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *CustomTableTree) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
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

// public void deselectAll()
func (jbobject *CustomTableTree) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *CustomTableTree) GetBackground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *CustomTableTree) GetClientArea() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientArea", "org/eclipse/swt/graphics/Rectangle")
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

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *CustomTableTree) GetForeground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *CustomTableTree) GetFont() *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFont", "org/eclipse/swt/graphics/Font")
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
	unique_x := &GraphicsFont{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *CustomTableTree) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *CustomTableTree) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.custom.TableTreeItem[] getItems()
func (jbobject *CustomTableTree) GetItems() []*CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/TableTreeItem")
	dst := new([]*CustomTableTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.custom.TableTreeItem[] getSelection()
func (jbobject *CustomTableTree) GetSelection() []*CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.ObjectArrayType("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/TableTreeItem")
	dst := new([]*CustomTableTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSelectionCount()
func (jbobject *CustomTableTree) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getStyle()
func (jbobject *CustomTableTree) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Table getTable()
func (jbobject *CustomTableTree) GetTable() *WidgetsTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTable", "org/eclipse/swt/widgets/Table")
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

// public int indexOf(org.eclipse.swt.custom.TableTreeItem)
func (jbobject *CustomTableTree) IndexOf(a CustomTableTreeItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public org.eclipse.swt.custom.TableTreeItem getItem(int)
func (jbobject *CustomTableTree) GetItem(a int) *CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/TableTreeItem", a)
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
	unique_x := &CustomTableTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.custom.TableTreeItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *CustomTableTree) GetItem2(a GraphicsPointInterface) *CustomTableTreeItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/TableTreeItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &CustomTableTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void removeAll()
func (jbobject *CustomTableTree) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomTableTree) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void removeTreeListener(org.eclipse.swt.events.TreeListener)
func (jbobject *CustomTableTree) RemoveTreeListener(a EventsTreeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeTreeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TreeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void selectAll()
func (jbobject *CustomTableTree) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomTableTree) SetBackground(a GraphicsColorInterface)  {
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
func (jbobject *CustomTableTree) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomTableTree) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *CustomTableTree) SetForeground(a GraphicsColorInterface)  {
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

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *CustomTableTree) SetMenu(a WidgetsMenuInterface)  {
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

// public void setSelection(org.eclipse.swt.custom.TableTreeItem[])
func (jbobject *CustomTableTree) SetSelection(a []*CustomTableTreeItem)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/TableTreeItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setToolTipText(java.lang.String)
func (jbobject *CustomTableTree) SetToolTipText(a string)  {
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

// public void showItem(org.eclipse.swt.custom.TableTreeItem)
func (jbobject *CustomTableTree) ShowItem(a CustomTableTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showSelection()
func (jbobject *CustomTableTree) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}


