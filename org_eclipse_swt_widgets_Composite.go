package swt

import "github.com/timob/javabind"

type WidgetsCompositeInterface interface {
	WidgetsScrollableInterface

	// public void changed(org.eclipse.swt.widgets.Control[])
	Changed(a []*WidgetsControl) 

	// public void drawBackground(org.eclipse.swt.graphics.GC, int, int, int, int, int, int)
	DrawBackground(a GraphicsGCInterface, b int, c int, d int, e int, f int, g int) 

	// public int getBackgroundMode()
	GetBackgroundMode() int

	// public org.eclipse.swt.widgets.Control[] getChildren()
	GetChildren() []*WidgetsControl

	// public org.eclipse.swt.widgets.Layout getLayout()
	GetLayout() *WidgetsLayout

	// public boolean getLayoutDeferred()
	GetLayoutDeferred() bool

	// public org.eclipse.swt.widgets.Control[] getTabList()
	GetTabList() []*WidgetsControl

	// public boolean isLayoutDeferred()
	IsLayoutDeferred() bool

	// public void layout()
	Layout() 

	// public void layout(boolean)
	Layout2(a bool) 

	// public void layout(boolean, boolean)
	Layout3(a bool, b bool) 

	// public void layout(org.eclipse.swt.widgets.Control[])
	Layout4(a []*WidgetsControl) 

	// public void layout(org.eclipse.swt.widgets.Control[], int)
	Layout5(a []*WidgetsControl, b int) 

	// public void setBackgroundMode(int)
	SetBackgroundMode(a int) 

	// public void setLayout(org.eclipse.swt.widgets.Layout)
	SetLayout(a WidgetsLayoutInterface) 

	// public void setLayoutDeferred(boolean)
	SetLayoutDeferred(a bool) 

	// public void setTabList(org.eclipse.swt.widgets.Control[])
	SetTabList(a []*WidgetsControl) 
}

type WidgetsComposite struct {
	WidgetsScrollable
}

// public org.eclipse.swt.widgets.Composite(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsComposite(a WidgetsCompositeInterface, b int) (*WidgetsComposite) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Composite", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsComposite{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void changed(org.eclipse.swt.widgets.Control[])
func (jbobject *WidgetsComposite) Changed(a []*WidgetsControl)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Control")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "changed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsComposite) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public void drawBackground(org.eclipse.swt.graphics.GC, int, int, int, int, int, int)
func (jbobject *WidgetsComposite) DrawBackground(a GraphicsGCInterface, b int, c int, d int, e int, f int, g int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"), b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int getBackgroundMode()
func (jbobject *WidgetsComposite) GetBackgroundMode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackgroundMode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Control[] getChildren()
func (jbobject *WidgetsComposite) GetChildren() []*WidgetsControl {
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

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *WidgetsComposite) GetClientArea() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.Layout getLayout()
func (jbobject *WidgetsComposite) GetLayout() *WidgetsLayout {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLayout", "org/eclipse/swt/widgets/Layout")
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
	unique_x := &WidgetsLayout{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getLayoutDeferred()
func (jbobject *WidgetsComposite) GetLayoutDeferred() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLayoutDeferred", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.Control[] getTabList()
func (jbobject *WidgetsComposite) GetTabList() []*WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabList", javabind.ObjectArrayType("org/eclipse/swt/widgets/Control"))
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

// public boolean isLayoutDeferred()
func (jbobject *WidgetsComposite) IsLayoutDeferred() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLayoutDeferred", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void layout()
func (jbobject *WidgetsComposite) Layout()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void layout(boolean)
func (jbobject *WidgetsComposite) Layout2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void layout(boolean, boolean)
func (jbobject *WidgetsComposite) Layout3(a bool, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void layout(org.eclipse.swt.widgets.Control[])
func (jbobject *WidgetsComposite) Layout4(a []*WidgetsControl)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Control")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void layout(org.eclipse.swt.widgets.Control[], int)
func (jbobject *WidgetsComposite) Layout5(a []*WidgetsControl, b int)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Control")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackgroundMode(int)
func (jbobject *WidgetsComposite) SetBackgroundMode(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackgroundMode", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean setFocus()
func (jbobject *WidgetsComposite) SetFocus() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setFocus", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *WidgetsComposite) SetLayout(a WidgetsLayoutInterface)  {
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

// public void setLayoutDeferred(boolean)
func (jbobject *WidgetsComposite) SetLayoutDeferred(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLayoutDeferred", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTabList(org.eclipse.swt.widgets.Control[])
func (jbobject *WidgetsComposite) SetTabList(a []*WidgetsControl)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Control")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabList", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *WidgetsComposite) EmbeddedHandle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "embeddedHandle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *WidgetsComposite) SetFieldEmbeddedHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "embeddedHandle", val)
	if err != nil {
		panic(err)
	}

}


