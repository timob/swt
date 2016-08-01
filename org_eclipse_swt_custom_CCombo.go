package swt

import "github.com/timob/javabind"

type CustomCComboInterface interface {
	WidgetsCompositeInterface

	// public void add(java.lang.String)
	Add(a string) 

	// public void add(java.lang.String, int)
	Add2(a string, b int) 

	// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
	AddModifyListener(a EventsModifyListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
	AddVerifyListener(a EventsVerifyListenerInterface) 

	// public void clearSelection()
	ClearSelection() 

	// public void copy()
	Copy() 

	// public void cut()
	Cut() 

	// public void deselect(int)
	Deselect(a int) 

	// public void deselectAll()
	DeselectAll() 

	// public boolean getEditable()
	GetEditable() bool

	// public java.lang.String getItem(int)
	GetItem(a int) string

	// public int getItemCount()
	GetItemCount() int

	// public int getItemHeight()
	GetItemHeight() int

	// public java.lang.String[] getItems()
	GetItems() []string

	// public boolean getListVisible()
	GetListVisible() bool

	// public org.eclipse.swt.graphics.Point getSelection()
	GetSelection() *GraphicsPoint

	// public int getSelectionIndex()
	GetSelectionIndex() int

	// public java.lang.String getText()
	GetText() string

	// public int getTextHeight()
	GetTextHeight() int

	// public int getTextLimit()
	GetTextLimit() int

	// public int getVisibleItemCount()
	GetVisibleItemCount() int

	// public int indexOf(java.lang.String)
	IndexOf(a string) int

	// public int indexOf(java.lang.String, int)
	IndexOf2(a string, b int) int

	// public void paste()
	Paste() 

	// public void remove(int)
	Remove(a int) 

	// public void remove(int, int)
	Remove2(a int, b int) 

	// public void remove(java.lang.String)
	Remove3(a string) 

	// public void removeAll()
	RemoveAll() 

	// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
	RemoveModifyListener(a EventsModifyListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeVerifyListener(org.eclipse.swt.events.VerifyListener)
	RemoveVerifyListener(a EventsVerifyListenerInterface) 

	// public void select(int)
	Select(a int) 

	// public void setEditable(boolean)
	SetEditable(a bool) 

	// public void setItem(int, java.lang.String)
	SetItem(a int, b string) 

	// public void setItems(java.lang.String[])
	SetItems(a []string) 

	// public void setListVisible(boolean)
	SetListVisible(a bool) 

	// public void setSelection(org.eclipse.swt.graphics.Point)
	SetSelection(a GraphicsPointInterface) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setTextLimit(int)
	SetTextLimit(a int) 

	// public void setVisibleItemCount(int)
	SetVisibleItemCount(a int) 
}

type CustomCCombo struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.CCombo(org.eclipse.swt.widgets.Composite, int)
func NewCustomCCombo(a WidgetsCompositeInterface, b int) (*CustomCCombo) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CCombo", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCCombo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void add(java.lang.String)
func (jbobject *CustomCCombo) Add(a string)  {
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
func (jbobject *CustomCCombo) Add2(a string, b int)  {
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

// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *CustomCCombo) AddModifyListener(a EventsModifyListenerInterface)  {
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
func (jbobject *CustomCCombo) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
func (jbobject *CustomCCombo) AddVerifyListener(a EventsVerifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addVerifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void clearSelection()
func (jbobject *CustomCCombo) ClearSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *CustomCCombo) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public void copy()
func (jbobject *CustomCCombo) Copy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cut()
func (jbobject *CustomCCombo) Cut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int)
func (jbobject *CustomCCombo) Deselect(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselectAll()
func (jbobject *CustomCCombo) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.widgets.Control[] getChildren()
func (jbobject *CustomCCombo) GetChildren() []*WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChildren", javabind.ObjectArrayType("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Control")
	dst := new([]*WidgetsControl)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getEditable()
func (jbobject *CustomCCombo) GetEditable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEditable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getItem(int)
func (jbobject *CustomCCombo) GetItem(a int) string {
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
func (jbobject *CustomCCombo) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *CustomCCombo) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String[] getItems()
func (jbobject *CustomCCombo) GetItems() []string {
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

// public boolean getListVisible()
func (jbobject *CustomCCombo) GetListVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getListVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.Menu getMenu()
func (jbobject *CustomCCombo) GetMenu() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMenu", "org/eclipse/swt/widgets/Menu")
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
	unique_x := &WidgetsMenu{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point getSelection()
func (jbobject *CustomCCombo) GetSelection() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", "org/eclipse/swt/graphics/Point")
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

// public int getSelectionIndex()
func (jbobject *CustomCCombo) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Shell getShell()
func (jbobject *CustomCCombo) GetShell() *WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getShell", "org/eclipse/swt/widgets/Shell")
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
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public int getStyle()
func (jbobject *CustomCCombo) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *CustomCCombo) GetText() string {
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

// public int getTextHeight()
func (jbobject *CustomCCombo) GetTextHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTextLimit()
func (jbobject *CustomCCombo) GetTextLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getVisibleItemCount()
func (jbobject *CustomCCombo) GetVisibleItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisibleItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(java.lang.String)
func (jbobject *CustomCCombo) IndexOf(a string) int {
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
func (jbobject *CustomCCombo) IndexOf2(a string, b int) int {
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

// public boolean isFocusControl()
func (jbobject *CustomCCombo) IsFocusControl() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isFocusControl", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void paste()
func (jbobject *CustomCCombo) Paste()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paste", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void redraw()
func (jbobject *CustomCCombo) Redraw()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void redraw(int, int, int, int, boolean)
func (jbobject *CustomCCombo) Redraw2(a int, b int, c int, d int, e bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void remove(int)
func (jbobject *CustomCCombo) Remove(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void remove(int, int)
func (jbobject *CustomCCombo) Remove2(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void remove(java.lang.String)
func (jbobject *CustomCCombo) Remove3(a string)  {
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

// public void removeAll()
func (jbobject *CustomCCombo) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *CustomCCombo) RemoveModifyListener(a EventsModifyListenerInterface)  {
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
func (jbobject *CustomCCombo) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void removeVerifyListener(org.eclipse.swt.events.VerifyListener)
func (jbobject *CustomCCombo) RemoveVerifyListener(a EventsVerifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeVerifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void select(int)
func (jbobject *CustomCCombo) Select(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCCombo) SetBackground(a GraphicsColorInterface)  {
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

// public void setEditable(boolean)
func (jbobject *CustomCCombo) SetEditable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEnabled(boolean)
func (jbobject *CustomCCombo) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean setFocus()
func (jbobject *CustomCCombo) SetFocus() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setFocus", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomCCombo) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *CustomCCombo) SetForeground(a GraphicsColorInterface)  {
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

// public void setItem(int, java.lang.String)
func (jbobject *CustomCCombo) SetItem(a int, b string)  {
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
func (jbobject *CustomCCombo) SetItems(a []string)  {
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

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomCCombo) SetLayout(a WidgetsLayoutInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLayout", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Layout"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setListVisible(boolean)
func (jbobject *CustomCCombo) SetListVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setListVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *CustomCCombo) SetMenu(a WidgetsMenuInterface)  {
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

// public void setSelection(org.eclipse.swt.graphics.Point)
func (jbobject *CustomCCombo) SetSelection(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setText(java.lang.String)
func (jbobject *CustomCCombo) SetText(a string)  {
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

// public void setTextLimit(int)
func (jbobject *CustomCCombo) SetTextLimit(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextLimit", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setToolTipText(java.lang.String)
func (jbobject *CustomCCombo) SetToolTipText(a string)  {
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

// public void setVisible(boolean)
func (jbobject *CustomCCombo) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setVisibleItemCount(int)
func (jbobject *CustomCCombo) SetVisibleItemCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisibleItemCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean traverse(int)
func (jbobject *CustomCCombo) Traverse(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "traverse", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


