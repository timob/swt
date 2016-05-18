package swt

import "github.com/timob/javabind"

type CustomScrolledCompositeInterface interface {
	WidgetsCompositeInterface

	// public boolean getAlwaysShowScrollBars()
	GetAlwaysShowScrollBars() bool

	// public boolean getExpandHorizontal()
	GetExpandHorizontal() bool

	// public boolean getExpandVertical()
	GetExpandVertical() bool

	// public int getMinWidth()
	GetMinWidth() int

	// public int getMinHeight()
	GetMinHeight() int

	// public org.eclipse.swt.widgets.Control getContent()
	GetContent() *WidgetsControl

	// public boolean getShowFocusedControl()
	GetShowFocusedControl() bool

	// public org.eclipse.swt.graphics.Point getOrigin()
	GetOrigin() *GraphicsPoint

	// public void setOrigin(org.eclipse.swt.graphics.Point)
	SetOrigin2(a GraphicsPointInterface) 

	// public void setOrigin(int, int)
	SetOrigin(a int, b int) 

	// public void setAlwaysShowScrollBars(boolean)
	SetAlwaysShowScrollBars(a bool) 

	// public void setContent(org.eclipse.swt.widgets.Control)
	SetContent(a WidgetsControlInterface) 

	// public void setExpandHorizontal(boolean)
	SetExpandHorizontal(a bool) 

	// public void setExpandVertical(boolean)
	SetExpandVertical(a bool) 

	// public void setMinHeight(int)
	SetMinHeight(a int) 

	// public void setMinSize(org.eclipse.swt.graphics.Point)
	SetMinSize2(a GraphicsPointInterface) 

	// public void setMinSize(int, int)
	SetMinSize(a int, b int) 

	// public void setMinWidth(int)
	SetMinWidth(a int) 

	// public void setShowFocusedControl(boolean)
	SetShowFocusedControl(a bool) 

	// public void showControl(org.eclipse.swt.widgets.Control)
	ShowControl(a WidgetsControlInterface) 
}

type CustomScrolledComposite struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.ScrolledComposite(org.eclipse.swt.widgets.Composite, int)
func NewCustomScrolledComposite(a WidgetsCompositeInterface, b int) (*CustomScrolledComposite) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/ScrolledComposite", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomScrolledComposite{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean getAlwaysShowScrollBars()
func (jbobject *CustomScrolledComposite) GetAlwaysShowScrollBars() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlwaysShowScrollBars", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getExpandHorizontal()
func (jbobject *CustomScrolledComposite) GetExpandHorizontal() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpandHorizontal", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getExpandVertical()
func (jbobject *CustomScrolledComposite) GetExpandVertical() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpandVertical", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getMinWidth()
func (jbobject *CustomScrolledComposite) GetMinWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinHeight()
func (jbobject *CustomScrolledComposite) GetMinHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Control getContent()
func (jbobject *CustomScrolledComposite) GetContent() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContent", "org/eclipse/swt/widgets/Control")
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
	unique_x := &WidgetsControl{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getShowFocusedControl()
func (jbobject *CustomScrolledComposite) GetShowFocusedControl() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getShowFocusedControl", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Point getOrigin()
func (jbobject *CustomScrolledComposite) GetOrigin() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrigin", "org/eclipse/swt/graphics/Point")
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

// public void setOrigin(org.eclipse.swt.graphics.Point)
func (jbobject *CustomScrolledComposite) SetOrigin2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrigin", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setOrigin(int, int)
func (jbobject *CustomScrolledComposite) SetOrigin(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrigin", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setAlwaysShowScrollBars(boolean)
func (jbobject *CustomScrolledComposite) SetAlwaysShowScrollBars(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlwaysShowScrollBars", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setContent(org.eclipse.swt.widgets.Control)
func (jbobject *CustomScrolledComposite) SetContent(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setContent", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setExpandHorizontal(boolean)
func (jbobject *CustomScrolledComposite) SetExpandHorizontal(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpandHorizontal", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setExpandVertical(boolean)
func (jbobject *CustomScrolledComposite) SetExpandVertical(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpandVertical", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomScrolledComposite) SetLayout(a WidgetsLayoutInterface)  {
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

// public void setMinHeight(int)
func (jbobject *CustomScrolledComposite) SetMinHeight(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinHeight", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinSize(org.eclipse.swt.graphics.Point)
func (jbobject *CustomScrolledComposite) SetMinSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMinSize(int, int)
func (jbobject *CustomScrolledComposite) SetMinSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setMinWidth(int)
func (jbobject *CustomScrolledComposite) SetMinWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setShowFocusedControl(boolean)
func (jbobject *CustomScrolledComposite) SetShowFocusedControl(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setShowFocusedControl", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showControl(org.eclipse.swt.widgets.Control)
func (jbobject *CustomScrolledComposite) ShowControl(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showControl", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


