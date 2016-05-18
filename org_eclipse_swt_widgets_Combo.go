package swt

import "github.com/timob/javabind"

type WidgetsComboInterface interface {
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

	// public org.eclipse.swt.graphics.Point getCaretLocation()
	GetCaretLocation() *GraphicsPoint

	// public int getCaretPosition()
	GetCaretPosition() int

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

type WidgetsCombo struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.Combo(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsCombo(a WidgetsCompositeInterface, b int) (*WidgetsCombo) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Combo", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCombo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void add(java.lang.String)
func (jbobject *WidgetsCombo) Add(a string)  {
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
func (jbobject *WidgetsCombo) Add2(a string, b int)  {
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
func (jbobject *WidgetsCombo) AddModifyListener(a EventsModifyListenerInterface)  {
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
func (jbobject *WidgetsCombo) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsCombo) AddVerifyListener(a EventsVerifyListenerInterface)  {
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
func (jbobject *WidgetsCombo) ClearSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsCombo) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsCombo) Copy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cut()
func (jbobject *WidgetsCombo) Cut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void deselect(int)
func (jbobject *WidgetsCombo) Deselect(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void deselectAll()
func (jbobject *WidgetsCombo) DeselectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deselectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point getCaretLocation()
func (jbobject *WidgetsCombo) GetCaretLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretLocation", "org/eclipse/swt/graphics/Point")
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

// public int getCaretPosition()
func (jbobject *WidgetsCombo) GetCaretPosition() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretPosition", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getItem(int)
func (jbobject *WidgetsCombo) GetItem(a int) string {
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
func (jbobject *WidgetsCombo) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getItemHeight()
func (jbobject *WidgetsCombo) GetItemHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String[] getItems()
func (jbobject *WidgetsCombo) GetItems() []string {
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
func (jbobject *WidgetsCombo) GetListVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getListVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getOrientation()
func (jbobject *WidgetsCombo) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getSelection()
func (jbobject *WidgetsCombo) GetSelection() *GraphicsPoint {
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
func (jbobject *WidgetsCombo) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *WidgetsCombo) GetText() string {
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
func (jbobject *WidgetsCombo) GetTextHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTextLimit()
func (jbobject *WidgetsCombo) GetTextLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getVisibleItemCount()
func (jbobject *WidgetsCombo) GetVisibleItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisibleItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(java.lang.String)
func (jbobject *WidgetsCombo) IndexOf(a string) int {
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
func (jbobject *WidgetsCombo) IndexOf2(a string, b int) int {
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

// public void paste()
func (jbobject *WidgetsCombo) Paste()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paste", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void remove(int)
func (jbobject *WidgetsCombo) Remove(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void remove(int, int)
func (jbobject *WidgetsCombo) Remove2(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void remove(java.lang.String)
func (jbobject *WidgetsCombo) Remove3(a string)  {
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
func (jbobject *WidgetsCombo) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *WidgetsCombo) RemoveModifyListener(a EventsModifyListenerInterface)  {
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
func (jbobject *WidgetsCombo) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsCombo) RemoveVerifyListener(a EventsVerifyListenerInterface)  {
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
func (jbobject *WidgetsCombo) Select(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setItem(int, java.lang.String)
func (jbobject *WidgetsCombo) SetItem(a int, b string)  {
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
func (jbobject *WidgetsCombo) SetItems(a []string)  {
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

// public void setListVisible(boolean)
func (jbobject *WidgetsCombo) SetListVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setListVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setOrientation(int)
func (jbobject *WidgetsCombo) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCombo) SetSelection(a GraphicsPointInterface)  {
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
func (jbobject *WidgetsCombo) SetText(a string)  {
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
func (jbobject *WidgetsCombo) SetTextLimit(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextLimit", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setVisibleItemCount(int)
func (jbobject *WidgetsCombo) SetVisibleItemCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisibleItemCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsCombo) LIMIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/widgets/Combo", "LIMIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsCombo) SetFieldLIMIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/widgets/Combo", "LIMIT", val)
	if err != nil {
		panic(err)
	}

}


