package swt

import "github.com/timob/javabind"

type WidgetsTreeInterface interface {
	WidgetsCompositeInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void addTreeListener(org.eclipse.swt.events.TreeListener)
	AddTreeListener(a EventsTreeListenerInterface) 

	// public void clear(int, boolean)
	Clear(a int, b bool) 

	// public void clearAll(boolean)
	ClearAll(a bool) 

	// public void deselect(org.eclipse.swt.widgets.TreeItem)
	Deselect(a WidgetsTreeItemInterface) 

	// public void deselectAll()
	DeselectAll() 

	// public org.eclipse.swt.widgets.TreeColumn getColumn(int)
	GetColumn(a int) *WidgetsTreeColumn

	// public int getColumnCount()
	GetColumnCount() int

	// public int[] getColumnOrder()
	GetColumnOrder() []int

	// public org.eclipse.swt.widgets.TreeColumn[] getColumns()
	GetColumns() []*WidgetsTreeColumn

	// public int getGridLineWidth()
	GetGridLineWidth() int

	// public int getHeaderHeight()
	GetHeaderHeight() int

	// public boolean getHeaderVisible()
	GetHeaderVisible() bool

	// public org.eclipse.swt.widgets.TreeItem getItem(int)
	GetItem(a int) *WidgetsTreeItem

	// public org.eclipse.swt.widgets.TreeItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *WidgetsTreeItem

	// public int getItemCount()
	GetItemCount() int

	// public int getItemHeight()
	GetItemHeight() int

	// public org.eclipse.swt.widgets.TreeItem[] getItems()
	GetItems() []*WidgetsTreeItem

	// public boolean getLinesVisible()
	GetLinesVisible() bool

	// public org.eclipse.swt.widgets.TreeItem getParentItem()
	GetParentItem() *WidgetsTreeItem

	// public org.eclipse.swt.widgets.TreeItem[] getSelection()
	GetSelection() []*WidgetsTreeItem

	// public int getSelectionCount()
	GetSelectionCount() int

	// public org.eclipse.swt.widgets.TreeColumn getSortColumn()
	GetSortColumn() *WidgetsTreeColumn

	// public int getSortDirection()
	GetSortDirection() int

	// public org.eclipse.swt.widgets.TreeItem getTopItem()
	GetTopItem() *WidgetsTreeItem

	// public int indexOf(org.eclipse.swt.widgets.TreeColumn)
	IndexOf(a WidgetsTreeColumnInterface) int

	// public int indexOf(org.eclipse.swt.widgets.TreeItem)
	IndexOf2(a WidgetsTreeItemInterface) int

	// public void removeAll()
	RemoveAll() 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeTreeListener(org.eclipse.swt.events.TreeListener)
	RemoveTreeListener(a EventsTreeListenerInterface) 

	// public void setInsertMark(org.eclipse.swt.widgets.TreeItem, boolean)
	SetInsertMark(a WidgetsTreeItemInterface, b bool) 

	// public void setItemCount(int)
	SetItemCount(a int) 

	// public void select(org.eclipse.swt.widgets.TreeItem)
	Select(a WidgetsTreeItemInterface) 

	// public void selectAll()
	SelectAll() 

	// public void setColumnOrder(int[])
	SetColumnOrder(a []int) 

	// public void setHeaderVisible(boolean)
	SetHeaderVisible(a bool) 

	// public void setLinesVisible(boolean)
	SetLinesVisible(a bool) 

	// public void setSelection(org.eclipse.swt.widgets.TreeItem)
	SetSelection(a WidgetsTreeItemInterface) 

	// public void setSelection(org.eclipse.swt.widgets.TreeItem[])
	SetSelection2(a []*WidgetsTreeItem) 

	// public void setSortColumn(org.eclipse.swt.widgets.TreeColumn)
	SetSortColumn(a WidgetsTreeColumnInterface) 

	// public void setSortDirection(int)
	SetSortDirection(a int) 

	// public void setTopItem(org.eclipse.swt.widgets.TreeItem)
	SetTopItem(a WidgetsTreeItemInterface) 

	// public void showColumn(org.eclipse.swt.widgets.TreeColumn)
	ShowColumn(a WidgetsTreeColumnInterface) 

	// public void showSelection()
	ShowSelection() 

	// public void showItem(org.eclipse.swt.widgets.TreeItem)
	ShowItem(a WidgetsTreeItemInterface) 
}

type WidgetsTree struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.Tree(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsTree(a WidgetsCompositeInterface, b int) (*WidgetsTree) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Tree", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTree{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTree) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsTree) AddTreeListener(a EventsTreeListenerInterface)  {
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

// public void clear(int, boolean)
func (jbobject *WidgetsTree) Clear(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void clearAll(boolean)
func (jbobject *WidgetsTree) ClearAll(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearAll", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsTree) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public void deselect(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) Deselect(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deselectAll()
func (jbobject *WidgetsTree) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *WidgetsTree) GetClientArea() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.TreeColumn getColumn(int)
func (jbobject *WidgetsTree) GetColumn(a int) *WidgetsTreeColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", "org/eclipse/swt/widgets/TreeColumn", a)
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
	unique_x := &WidgetsTreeColumn{}
	unique_x.Callable = dst
	return unique_x
}

// public int getColumnCount()
func (jbobject *WidgetsTree) GetColumnCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumnCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getColumnOrder()
func (jbobject *WidgetsTree) GetColumnOrder() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumnOrder", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.widgets.TreeColumn[] getColumns()
func (jbobject *WidgetsTree) GetColumns() []*WidgetsTreeColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumns", javabind.ObjectArrayType("org/eclipse/swt/widgets/TreeColumn"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TreeColumn")
	dst := new([]*WidgetsTreeColumn)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getGridLineWidth()
func (jbobject *WidgetsTree) GetGridLineWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGridLineWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHeaderHeight()
func (jbobject *WidgetsTree) GetHeaderHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getHeaderVisible()
func (jbobject *WidgetsTree) GetHeaderVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.TreeItem getItem(int)
func (jbobject *WidgetsTree) GetItem(a int) *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TreeItem", a)
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

// public org.eclipse.swt.widgets.TreeItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsTree) GetItem2(a GraphicsPointInterface) *WidgetsTreeItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TreeItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
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
	unique_x := &WidgetsTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsTree) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *WidgetsTree) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TreeItem[] getItems()
func (jbobject *WidgetsTree) GetItems() []*WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TreeItem")
	dst := new([]*WidgetsTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getLinesVisible()
func (jbobject *WidgetsTree) GetLinesVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLinesVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.TreeItem getParentItem()
func (jbobject *WidgetsTree) GetParentItem() *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentItem", "org/eclipse/swt/widgets/TreeItem")
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

// public org.eclipse.swt.widgets.TreeItem[] getSelection()
func (jbobject *WidgetsTree) GetSelection() []*WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.ObjectArrayType("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TreeItem")
	dst := new([]*WidgetsTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSelectionCount()
func (jbobject *WidgetsTree) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TreeColumn getSortColumn()
func (jbobject *WidgetsTree) GetSortColumn() *WidgetsTreeColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSortColumn", "org/eclipse/swt/widgets/TreeColumn")
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
	unique_x := &WidgetsTreeColumn{}
	unique_x.Callable = dst
	return unique_x
}

// public int getSortDirection()
func (jbobject *WidgetsTree) GetSortDirection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSortDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TreeItem getTopItem()
func (jbobject *WidgetsTree) GetTopItem() *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopItem", "org/eclipse/swt/widgets/TreeItem")
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

// public int indexOf(org.eclipse.swt.widgets.TreeColumn)
func (jbobject *WidgetsTree) IndexOf(a WidgetsTreeColumnInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) IndexOf2(a WidgetsTreeItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void removeAll()
func (jbobject *WidgetsTree) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTree) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsTree) RemoveTreeListener(a EventsTreeListenerInterface)  {
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

// public void setInsertMark(org.eclipse.swt.widgets.TreeItem, boolean)
func (jbobject *WidgetsTree) SetInsertMark(a WidgetsTreeItemInterface, b bool)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInsertMark", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setItemCount(int)
func (jbobject *WidgetsTree) SetItemCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItemCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void select(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) Select(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void selectAll()
func (jbobject *WidgetsTree) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setColumnOrder(int[])
func (jbobject *WidgetsTree) SetColumnOrder(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setColumnOrder", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHeaderVisible(boolean)
func (jbobject *WidgetsTree) SetHeaderVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHeaderVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLinesVisible(boolean)
func (jbobject *WidgetsTree) SetLinesVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLinesVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) SetSelection(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(org.eclipse.swt.widgets.TreeItem[])
func (jbobject *WidgetsTree) SetSelection2(a []*WidgetsTreeItem)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/TreeItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSortColumn(org.eclipse.swt.widgets.TreeColumn)
func (jbobject *WidgetsTree) SetSortColumn(a WidgetsTreeColumnInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSortColumn", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSortDirection(int)
func (jbobject *WidgetsTree) SetSortDirection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSortDirection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopItem(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) SetTopItem(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showColumn(org.eclipse.swt.widgets.TreeColumn)
func (jbobject *WidgetsTree) ShowColumn(a WidgetsTreeColumnInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showColumn", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showSelection()
func (jbobject *WidgetsTree) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void showItem(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTree) ShowItem(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


