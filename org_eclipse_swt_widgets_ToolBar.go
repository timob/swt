package swt

import "github.com/timob/javabind"

type WidgetsToolBarInterface interface {
	WidgetsCompositeInterface

	// public org.eclipse.swt.widgets.ToolItem getItem(int)
	GetItem(a int) *WidgetsToolItem

	// public org.eclipse.swt.widgets.ToolItem getItem(org.eclipse.swt.graphics.Point)
	GetItem2(a GraphicsPointInterface) *WidgetsToolItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.ToolItem[] getItems()
	GetItems() []*WidgetsToolItem

	// public int getRowCount()
	GetRowCount() int

	// public int indexOf(org.eclipse.swt.widgets.ToolItem)
	IndexOf(a WidgetsToolItemInterface) int
}

type WidgetsToolBar struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.ToolBar(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsToolBar(a WidgetsCompositeInterface, b int) (*WidgetsToolBar) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ToolBar", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsToolBar{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsToolBar) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public org.eclipse.swt.widgets.ToolItem getItem(int)
func (jbobject *WidgetsToolBar) GetItem(a int) *WidgetsToolItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/ToolItem", a)
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
	unique_x := &WidgetsToolItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.ToolItem getItem(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsToolBar) GetItem2(a GraphicsPointInterface) *WidgetsToolItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/ToolItem", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
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
	unique_x := &WidgetsToolItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsToolBar) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.ToolItem[] getItems()
func (jbobject *WidgetsToolBar) GetItems() []*WidgetsToolItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/ToolItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/ToolItem")
	dst := new([]*WidgetsToolItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getRowCount()
func (jbobject *WidgetsToolBar) GetRowCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRowCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.ToolItem)
func (jbobject *WidgetsToolBar) IndexOf(a WidgetsToolItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/ToolItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void setToolTipText(java.lang.String)
func (jbobject *WidgetsToolBar) SetToolTipText(a string)  {
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


