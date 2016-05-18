package swt

import "github.com/timob/javabind"

type WidgetsCoolBarInterface interface {
	WidgetsCompositeInterface

	// public org.eclipse.swt.widgets.CoolItem getItem(int)
	GetItem(a int) *WidgetsCoolItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.CoolItem[] getItems()
	GetItems() []*WidgetsCoolItem

	// public int indexOf(org.eclipse.swt.widgets.CoolItem)
	IndexOf(a WidgetsCoolItemInterface) int

	// public int[] getItemOrder()
	GetItemOrder() []int

	// public org.eclipse.swt.graphics.Point[] getItemSizes()
	GetItemSizes() []*GraphicsPoint

	// public boolean getLocked()
	GetLocked() bool

	// public int[] getWrapIndices()
	GetWrapIndices() []int

	// public void setLocked(boolean)
	SetLocked(a bool) 

	// public void setWrapIndices(int[])
	SetWrapIndices(a []int) 

	// public void setItemLayout(int[], int[], org.eclipse.swt.graphics.Point[])
	SetItemLayout(a []int, b []int, c []*GraphicsPoint) 
}

type WidgetsCoolBar struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.CoolBar(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsCoolBar(a WidgetsCompositeInterface, b int) (*WidgetsCoolBar) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/CoolBar", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCoolBar{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsCoolBar) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public org.eclipse.swt.widgets.CoolItem getItem(int)
func (jbobject *WidgetsCoolBar) GetItem(a int) *WidgetsCoolItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/CoolItem", a)
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
	unique_x := &WidgetsCoolItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsCoolBar) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.CoolItem[] getItems()
func (jbobject *WidgetsCoolBar) GetItems() []*WidgetsCoolItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/CoolItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/CoolItem")
	dst := new([]*WidgetsCoolItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int indexOf(org.eclipse.swt.widgets.CoolItem)
func (jbobject *WidgetsCoolBar) IndexOf(a WidgetsCoolItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/CoolItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int[] getItemOrder()
func (jbobject *WidgetsCoolBar) GetItemOrder() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemOrder", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.graphics.Point[] getItemSizes()
func (jbobject *WidgetsCoolBar) GetItemSizes() []*GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemSizes", javabind.ObjectArrayType("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/Point")
	dst := new([]*GraphicsPoint)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getLocked()
func (jbobject *WidgetsCoolBar) GetLocked() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocked", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int[] getWrapIndices()
func (jbobject *WidgetsCoolBar) GetWrapIndices() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWrapIndices", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public void setLocked(boolean)
func (jbobject *WidgetsCoolBar) SetLocked(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocked", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setWrapIndices(int[])
func (jbobject *WidgetsCoolBar) SetWrapIndices(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWrapIndices", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setCursor(org.eclipse.swt.graphics.Cursor)
func (jbobject *WidgetsCoolBar) SetCursor(a GraphicsCursorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCursor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Cursor"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setItemLayout(int[], int[], org.eclipse.swt.graphics.Point[])
func (jbobject *WidgetsCoolBar) SetItemLayout(a []int, b []int, c []*GraphicsPoint)  {
	conv_c := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Point")
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItemLayout", javabind.Void, a, b, conv_c.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void setOrientation(int)
func (jbobject *WidgetsCoolBar) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


