package swt

import "github.com/timob/javabind"

type WidgetsListInterface interface {
	WidgetsScrollableInterface

	// public void add(java.lang.String)
	Add(a string) 

	// public void add(java.lang.String, int)
	Add2(a string, b int) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void deselect(int)
	Deselect(a int) 

	// public void deselect(int, int)
	Deselect3(a int, b int) 

	// public void deselect(int[])
	Deselect2(a []int) 

	// public void deselectAll()
	DeselectAll() 

	// public int getFocusIndex()
	GetFocusIndex() int

	// public java.lang.String getItem(int)
	GetItem(a int) string

	// public int getItemCount()
	GetItemCount() int

	// public int getItemHeight()
	GetItemHeight() int

	// public java.lang.String[] getItems()
	GetItems() []string

	// public java.lang.String[] getSelection()
	GetSelection() []string

	// public int getSelectionCount()
	GetSelectionCount() int

	// public int getSelectionIndex()
	GetSelectionIndex() int

	// public int[] getSelectionIndices()
	GetSelectionIndices() []int

	// public int getTopIndex()
	GetTopIndex() int

	// public int indexOf(java.lang.String)
	IndexOf(a string) int

	// public int indexOf(java.lang.String, int)
	IndexOf2(a string, b int) int

	// public boolean isSelected(int)
	IsSelected(a int) bool

	// public void remove(int)
	Remove(a int) 

	// public void remove(int, int)
	Remove3(a int, b int) 

	// public void remove(java.lang.String)
	Remove4(a string) 

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

	// public void setItem(int, java.lang.String)
	SetItem(a int, b string) 

	// public void setItems(java.lang.String[])
	SetItems(a []string) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelection(int, int)
	SetSelection3(a int, b int) 

	// public void setSelection(int[])
	SetSelection2(a []int) 

	// public void setSelection(java.lang.String[])
	SetSelection4(a []string) 

	// public void setTopIndex(int)
	SetTopIndex(a int) 

	// public void showSelection()
	ShowSelection() 
}

type WidgetsList struct {
	WidgetsScrollable
}

// public org.eclipse.swt.widgets.List(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsList(a WidgetsCompositeInterface, b int) (*WidgetsList) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/List", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsList{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void add(java.lang.String)
func (jbobject *WidgetsList) Add(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void add(java.lang.String, int)
func (jbobject *WidgetsList) Add2(a string, b int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsList) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsList) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsList) Deselect(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int, int)
func (jbobject *WidgetsList) Deselect3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int[])
func (jbobject *WidgetsList) Deselect2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselectAll()
func (jbobject *WidgetsList) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getFocusIndex()
func (jbobject *WidgetsList) GetFocusIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFocusIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getItem(int)
func (jbobject *WidgetsList) GetItem(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "java/lang/String", a)
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

// public int getItemCount()
func (jbobject *WidgetsList) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *WidgetsList) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String[] getItems()
func (jbobject *WidgetsList) GetItems() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String[] getSelection()
func (jbobject *WidgetsList) GetSelection() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSelectionCount()
func (jbobject *WidgetsList) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelectionIndex()
func (jbobject *WidgetsList) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getSelectionIndices()
func (jbobject *WidgetsList) GetSelectionIndices() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndices", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int getTopIndex()
func (jbobject *WidgetsList) GetTopIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(java.lang.String)
func (jbobject *WidgetsList) IndexOf(a string) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int indexOf(java.lang.String, int)
func (jbobject *WidgetsList) IndexOf2(a string, b int) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public boolean isSelected(int)
func (jbobject *WidgetsList) IsSelected(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isSelected", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void remove(int)
func (jbobject *WidgetsList) Remove(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void remove(int, int)
func (jbobject *WidgetsList) Remove3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void remove(java.lang.String)
func (jbobject *WidgetsList) Remove4(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void remove(int[])
func (jbobject *WidgetsList) Remove2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void removeAll()
func (jbobject *WidgetsList) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsList) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsList) Select(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void select(int, int)
func (jbobject *WidgetsList) Select3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void select(int[])
func (jbobject *WidgetsList) Select2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void selectAll()
func (jbobject *WidgetsList) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setItem(int, java.lang.String)
func (jbobject *WidgetsList) SetItem(a int, b string)  {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItem", javabind.Void, a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setItems(java.lang.String[])
func (jbobject *WidgetsList) SetItems(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItems", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(int)
func (jbobject *WidgetsList) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int, int)
func (jbobject *WidgetsList) SetSelection3(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int[])
func (jbobject *WidgetsList) SetSelection2(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(java.lang.String[])
func (jbobject *WidgetsList) SetSelection4(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTopIndex(int)
func (jbobject *WidgetsList) SetTopIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showSelection()
func (jbobject *WidgetsList) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}


