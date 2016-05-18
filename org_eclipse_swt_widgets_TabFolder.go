package swt

import "github.com/timob/javabind"

type WidgetsTabFolderInterface interface {
	WidgetsCompositeInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public org.eclipse.swt.widgets.TabItem getItem(int)
	GetItem(a int) *WidgetsTabItem

	// public org.eclipse.swt.widgets.TabItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *WidgetsTabItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.TabItem[] getItems()
	GetItems() []*WidgetsTabItem

	// public org.eclipse.swt.widgets.TabItem[] getSelection()
	GetSelection() []*WidgetsTabItem

	// public int getSelectionIndex()
	GetSelectionIndex() int

	// public int indexOf(org.eclipse.swt.widgets.TabItem)
	IndexOf(a WidgetsTabItemInterface) int

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelection(org.eclipse.swt.widgets.TabItem)
	SetSelection2(a WidgetsTabItemInterface) 

	// public void setSelection(org.eclipse.swt.widgets.TabItem[])
	SetSelection3(a []*WidgetsTabItem) 
}

type WidgetsTabFolder struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.TabFolder(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsTabFolder(a WidgetsCompositeInterface, b int) (*WidgetsTabFolder) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TabFolder", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTabFolder{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTabFolder) AddSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *WidgetsTabFolder) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsTabFolder) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.TabItem getItem(int)
func (jbobject *WidgetsTabFolder) GetItem(a int) *WidgetsTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TabItem", a)
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
	unique_x := &WidgetsTabItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TabItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsTabFolder) GetItem2(a GraphicsPointInterface) *WidgetsTabItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TabItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
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
	unique_x := &WidgetsTabItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsTabFolder) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TabItem[] getItems()
func (jbobject *WidgetsTabFolder) GetItems() []*WidgetsTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TabItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TabItem")
	dst := new([]*WidgetsTabItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.widgets.TabItem[] getSelection()
func (jbobject *WidgetsTabFolder) GetSelection() []*WidgetsTabItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.ObjectArrayType("org/eclipse/swt/widgets/TabItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TabItem")
	dst := new([]*WidgetsTabItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSelectionIndex()
func (jbobject *WidgetsTabFolder) GetSelectionIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.TabItem)
func (jbobject *WidgetsTabFolder) IndexOf(a WidgetsTabItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsTabFolder) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void setSelection(int)
func (jbobject *WidgetsTabFolder) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.widgets.TabItem)
func (jbobject *WidgetsTabFolder) SetSelection2(a WidgetsTabItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(org.eclipse.swt.widgets.TabItem[])
func (jbobject *WidgetsTabFolder) SetSelection3(a []*WidgetsTabItem)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/TabItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TabItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


