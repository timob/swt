package swt

import "github.com/timob/javabind"

type WidgetsTableInterface interface {
	WidgetsCompositeInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void clear(int)
	Clear(a int) 

	// public void clear(int, int)
	Clear3(a int, b int) 

	// public void clear(int[])
	Clear2(a []int) 

	// public void clearAll()
	ClearAll() 

	// public void deselect(int)
	Deselect(a int) 

	// public void deselect(int, int)
	Deselect3(a int, b int) 

	// public void deselect(int[])
	Deselect2(a []int) 

	// public void deselectAll()
	DeselectAll() 

	// public org.eclipse.swt.widgets.TableColumn getColumn(int)
	GetColumn(a int) *WidgetsTableColumn

	// public int getColumnCount()
	GetColumnCount() int

	// public int[] getColumnOrder()
	GetColumnOrder() []int

	// public org.eclipse.swt.widgets.TableColumn[] getColumns()
	GetColumns() []*WidgetsTableColumn

	// public int getGridLineWidth()
	GetGridLineWidth() int

	// public int getHeaderHeight()
	GetHeaderHeight() int

	// public boolean getHeaderVisible()
	GetHeaderVisible() bool

	// public org.eclipse.swt.widgets.TableItem getItem(int)
	GetItem(a int) *WidgetsTableItem

	// public org.eclipse.swt.widgets.TableItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *WidgetsTableItem

	// public int getItemCount()
	GetItemCount() int

	// public int getItemHeight()
	GetItemHeight() int

	// public org.eclipse.swt.widgets.TableItem[] getItems()
	GetItems() []*WidgetsTableItem

	// public boolean getLinesVisible()
	GetLinesVisible() bool

	// public org.eclipse.swt.widgets.TableItem[] getSelection()
	GetSelection() []*WidgetsTableItem

	// public int getSelectionCount()
	GetSelectionCount() int

	// public int getSelectionIndex()
	GetSelectionIndex() int

	// public int[] getSelectionIndices()
	GetSelectionIndices() []int

	// public org.eclipse.swt.widgets.TableColumn getSortColumn()
	GetSortColumn() *WidgetsTableColumn

	// public int getSortDirection()
	GetSortDirection() int

	// public int getTopIndex()
	GetTopIndex() int

	// public int indexOf(org.eclipse.swt.widgets.TableColumn)
	IndexOf(a WidgetsTableColumnInterface) int

	// public int indexOf(org.eclipse.swt.widgets.TableItem)
	IndexOf2(a WidgetsTableItemInterface) int

	// public boolean isSelected(int)
	IsSelected(a int) bool

	// public void remove(int)
	Remove(a int) 

	// public void remove(int, int)
	Remove3(a int, b int) 

	// public void remove(int[])
	Remove2(a []int) 

	// public void removeAll()
	RemoveAll() 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void select(int)
	Select(a int) 

	// public void select(int, int)
	Select3(a int, b int) 

	// public void select(int[])
	Select2(a []int) 

	// public void selectAll()
	SelectAll() 

	// public void setColumnOrder(int[])
	SetColumnOrder(a []int) 

	// public void setHeaderVisible(boolean)
	SetHeaderVisible(a bool) 

	// public void setItemCount(int)
	SetItemCount(a int) 

	// public void setLinesVisible(boolean)
	SetLinesVisible(a bool) 

	// public void setSortColumn(org.eclipse.swt.widgets.TableColumn)
	SetSortColumn(a WidgetsTableColumnInterface) 

	// public void setSortDirection(int)
	SetSortDirection(a int) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelection(int, int)
	SetSelection3(a int, b int) 

	// public void setSelection(int[])
	SetSelection2(a []int) 

	// public void setSelection(org.eclipse.swt.widgets.TableItem)
	SetSelection4(a WidgetsTableItemInterface) 

	// public void setSelection(org.eclipse.swt.widgets.TableItem[])
	SetSelection5(a []*WidgetsTableItem) 

	// public void setTopIndex(int)
	SetTopIndex(a int) 

	// public void showColumn(org.eclipse.swt.widgets.TableColumn)
	ShowColumn(a WidgetsTableColumnInterface) 

	// public void showItem(org.eclipse.swt.widgets.TableItem)
	ShowItem(a WidgetsTableItemInterface) 

	// public void showSelection()
	ShowSelection() 
}

type WidgetsTable struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.Table(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsTable(a WidgetsCompositeInterface, b int) (*WidgetsTable) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Table", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTable) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void clear(int)
func (jbobject *WidgetsTable) Clear(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void clear(int, int)
func (jbobject *WidgetsTable) Clear3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void clear(int[])
func (jbobject *WidgetsTable) Clear2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void clearAll()
func (jbobject *WidgetsTable) ClearAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsTable) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public void deselect(int)
func (jbobject *WidgetsTable) Deselect(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int, int)
func (jbobject *WidgetsTable) Deselect3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int[])
func (jbobject *WidgetsTable) Deselect2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselectAll()
func (jbobject *WidgetsTable) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *WidgetsTable) GetClientArea() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.TableColumn getColumn(int)
func (jbobject *WidgetsTable) GetColumn(a int) *WidgetsTableColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", "org/eclipse/swt/widgets/TableColumn", a)
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
	unique_x := &WidgetsTableColumn{}
	unique_x.Callable = dst
	return unique_x
}

// public int getColumnCount()
func (jbobject *WidgetsTable) GetColumnCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumnCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getColumnOrder()
func (jbobject *WidgetsTable) GetColumnOrder() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumnOrder", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.widgets.TableColumn[] getColumns()
func (jbobject *WidgetsTable) GetColumns() []*WidgetsTableColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumns", javabind.ObjectArrayType("org/eclipse/swt/widgets/TableColumn"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TableColumn")
	dst := new([]*WidgetsTableColumn)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getGridLineWidth()
func (jbobject *WidgetsTable) GetGridLineWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGridLineWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHeaderHeight()
func (jbobject *WidgetsTable) GetHeaderHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getHeaderVisible()
func (jbobject *WidgetsTable) GetHeaderVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.TableItem getItem(int)
func (jbobject *WidgetsTable) GetItem(a int) *WidgetsTableItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TableItem", a)
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

// public org.eclipse.swt.widgets.TableItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsTable) GetItem2(a GraphicsPointInterface) *WidgetsTableItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TableItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
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
	unique_x := &WidgetsTableItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsTable) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *WidgetsTable) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TableItem[] getItems()
func (jbobject *WidgetsTable) GetItems() []*WidgetsTableItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TableItem")
	dst := new([]*WidgetsTableItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getLinesVisible()
func (jbobject *WidgetsTable) GetLinesVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLinesVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.TableItem[] getSelection()
func (jbobject *WidgetsTable) GetSelection() []*WidgetsTableItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.ObjectArrayType("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TableItem")
	dst := new([]*WidgetsTableItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSelectionCount()
func (jbobject *WidgetsTable) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelectionIndex()
func (jbobject *WidgetsTable) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getSelectionIndices()
func (jbobject *WidgetsTable) GetSelectionIndices() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndices", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.widgets.TableColumn getSortColumn()
func (jbobject *WidgetsTable) GetSortColumn() *WidgetsTableColumn {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSortColumn", "org/eclipse/swt/widgets/TableColumn")
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
	unique_x := &WidgetsTableColumn{}
	unique_x.Callable = dst
	return unique_x
}

// public int getSortDirection()
func (jbobject *WidgetsTable) GetSortDirection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSortDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopIndex()
func (jbobject *WidgetsTable) GetTopIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.TableColumn)
func (jbobject *WidgetsTable) IndexOf(a WidgetsTableColumnInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TableColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.TableItem)
func (jbobject *WidgetsTable) IndexOf2(a WidgetsTableItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public boolean isSelected(int)
func (jbobject *WidgetsTable) IsSelected(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isSelected", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void remove(int)
func (jbobject *WidgetsTable) Remove(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void remove(int, int)
func (jbobject *WidgetsTable) Remove3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void remove(int[])
func (jbobject *WidgetsTable) Remove2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void removeAll()
func (jbobject *WidgetsTable) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTable) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void select(int)
func (jbobject *WidgetsTable) Select(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void select(int, int)
func (jbobject *WidgetsTable) Select3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void select(int[])
func (jbobject *WidgetsTable) Select2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void selectAll()
func (jbobject *WidgetsTable) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setColumnOrder(int[])
func (jbobject *WidgetsTable) SetColumnOrder(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setColumnOrder", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHeaderVisible(boolean)
func (jbobject *WidgetsTable) SetHeaderVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHeaderVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setItemCount(int)
func (jbobject *WidgetsTable) SetItemCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItemCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLinesVisible(boolean)
func (jbobject *WidgetsTable) SetLinesVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLinesVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setRedraw(boolean)
func (jbobject *WidgetsTable) SetRedraw(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRedraw", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSortColumn(org.eclipse.swt.widgets.TableColumn)
func (jbobject *WidgetsTable) SetSortColumn(a WidgetsTableColumnInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSortColumn", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSortDirection(int)
func (jbobject *WidgetsTable) SetSortDirection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSortDirection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsTable) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int, int)
func (jbobject *WidgetsTable) SetSelection3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int[])
func (jbobject *WidgetsTable) SetSelection2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.widgets.TableItem)
func (jbobject *WidgetsTable) SetSelection4(a WidgetsTableItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(org.eclipse.swt.widgets.TableItem[])
func (jbobject *WidgetsTable) SetSelection5(a []*WidgetsTableItem)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/TableItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTopIndex(int)
func (jbobject *WidgetsTable) SetTopIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showColumn(org.eclipse.swt.widgets.TableColumn)
func (jbobject *WidgetsTable) ShowColumn(a WidgetsTableColumnInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showColumn", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableColumn"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showItem(org.eclipse.swt.widgets.TableItem)
func (jbobject *WidgetsTable) ShowItem(a WidgetsTableItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showSelection()
func (jbobject *WidgetsTable) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}


