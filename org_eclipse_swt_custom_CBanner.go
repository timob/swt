package swt

import "github.com/timob/javabind"

type CustomCBannerInterface interface {
	WidgetsCompositeInterface

	// public org.eclipse.swt.widgets.Control getBottom()
	GetBottom() *WidgetsControl

	// public org.eclipse.swt.widgets.Control getLeft()
	GetLeft() *WidgetsControl

	// public org.eclipse.swt.widgets.Control getRight()
	GetRight() *WidgetsControl

	// public org.eclipse.swt.graphics.Point getRightMinimumSize()
	GetRightMinimumSize() *GraphicsPoint

	// public int getRightWidth()
	GetRightWidth() int

	// public boolean getSimple()
	GetSimple() bool

	// public void setBottom(org.eclipse.swt.widgets.Control)
	SetBottom(a WidgetsControlInterface) 

	// public void setLeft(org.eclipse.swt.widgets.Control)
	SetLeft(a WidgetsControlInterface) 

	// public void setRight(org.eclipse.swt.widgets.Control)
	SetRight(a WidgetsControlInterface) 

	// public void setRightMinimumSize(org.eclipse.swt.graphics.Point)
	SetRightMinimumSize(a GraphicsPointInterface) 

	// public void setRightWidth(int)
	SetRightWidth(a int) 

	// public void setSimple(boolean)
	SetSimple(a bool) 
}

type CustomCBanner struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.CBanner(org.eclipse.swt.widgets.Composite, int)
func NewCustomCBanner(a WidgetsCompositeInterface, b int) (*CustomCBanner) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CBanner", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCBanner{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Control getBottom()
func (jbobject *CustomCBanner) GetBottom() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBottom", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *CustomCBanner) GetClientArea() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.Control getLeft()
func (jbobject *CustomCBanner) GetLeft() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLeft", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.widgets.Control getRight()
func (jbobject *CustomCBanner) GetRight() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRight", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.graphics.Point getRightMinimumSize()
func (jbobject *CustomCBanner) GetRightMinimumSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRightMinimumSize", "org/eclipse/swt/graphics/Point")
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

// public int getRightWidth()
func (jbobject *CustomCBanner) GetRightWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRightWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getSimple()
func (jbobject *CustomCBanner) GetSimple() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSimple", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setBottom(org.eclipse.swt.widgets.Control)
func (jbobject *CustomCBanner) SetBottom(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBottom", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomCBanner) SetLayout(a WidgetsLayoutInterface)  {
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

// public void setLeft(org.eclipse.swt.widgets.Control)
func (jbobject *CustomCBanner) SetLeft(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLeft", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRight(org.eclipse.swt.widgets.Control)
func (jbobject *CustomCBanner) SetRight(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRight", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRightMinimumSize(org.eclipse.swt.graphics.Point)
func (jbobject *CustomCBanner) SetRightMinimumSize(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRightMinimumSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRightWidth(int)
func (jbobject *CustomCBanner) SetRightWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRightWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSimple(boolean)
func (jbobject *CustomCBanner) SetSimple(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSimple", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


