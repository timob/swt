package swt

import "github.com/timob/javabind"

type CustomViewFormInterface interface {
	WidgetsCompositeInterface

	// public org.eclipse.swt.widgets.Control getContent()
	GetContent() *WidgetsControl

	// public org.eclipse.swt.widgets.Control getTopCenter()
	GetTopCenter() *WidgetsControl

	// public org.eclipse.swt.widgets.Control getTopLeft()
	GetTopLeft() *WidgetsControl

	// public org.eclipse.swt.widgets.Control getTopRight()
	GetTopRight() *WidgetsControl

	// public void setContent(org.eclipse.swt.widgets.Control)
	SetContent(a WidgetsControlInterface) 

	// public void setTopCenter(org.eclipse.swt.widgets.Control)
	SetTopCenter(a WidgetsControlInterface) 

	// public void setTopLeft(org.eclipse.swt.widgets.Control)
	SetTopLeft(a WidgetsControlInterface) 

	// public void setTopRight(org.eclipse.swt.widgets.Control)
	SetTopRight(a WidgetsControlInterface) 

	// public void setBorderVisible(boolean)
	SetBorderVisible(a bool) 

	// public void setTopCenterSeparate(boolean)
	SetTopCenterSeparate(a bool) 
}

type CustomViewForm struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.ViewForm(org.eclipse.swt.widgets.Composite, int)
func NewCustomViewForm(a WidgetsCompositeInterface, b int) (*CustomViewForm) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/ViewForm", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomViewForm{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *CustomViewForm) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
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

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *CustomViewForm) GetClientArea() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.Control getContent()
func (jbobject *CustomViewForm) GetContent() *WidgetsControl {
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

// public org.eclipse.swt.widgets.Control getTopCenter()
func (jbobject *CustomViewForm) GetTopCenter() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopCenter", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.widgets.Control getTopLeft()
func (jbobject *CustomViewForm) GetTopLeft() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopLeft", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.widgets.Control getTopRight()
func (jbobject *CustomViewForm) GetTopRight() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopRight", "org/eclipse/swt/widgets/Control")
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

// public void setContent(org.eclipse.swt.widgets.Control)
func (jbobject *CustomViewForm) SetContent(a WidgetsControlInterface)  {
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

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomViewForm) SetLayout(a WidgetsLayoutInterface)  {
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

// public void setTopCenter(org.eclipse.swt.widgets.Control)
func (jbobject *CustomViewForm) SetTopCenter(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopCenter", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTopLeft(org.eclipse.swt.widgets.Control)
func (jbobject *CustomViewForm) SetTopLeft(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopLeft", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTopRight(org.eclipse.swt.widgets.Control)
func (jbobject *CustomViewForm) SetTopRight(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopRight", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBorderVisible(boolean)
func (jbobject *CustomViewForm) SetBorderVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBorderVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopCenterSeparate(boolean)
func (jbobject *CustomViewForm) SetTopCenterSeparate(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopCenterSeparate", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomViewForm) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomViewForm) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomViewForm) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomViewForm) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomViewForm) HorizontalSpacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomViewForm) SetFieldHorizontalSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalSpacing", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomViewForm) VerticalSpacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomViewForm) SetFieldVerticalSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalSpacing", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomViewForm) BorderInsideRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ViewForm", "borderInsideRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomViewForm) SetFieldBorderInsideRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ViewForm", "borderInsideRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomViewForm) BorderMiddleRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ViewForm", "borderMiddleRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomViewForm) SetFieldBorderMiddleRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ViewForm", "borderMiddleRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomViewForm) BorderOutsideRGB() *GraphicsRGB {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ViewForm", "borderOutsideRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomViewForm) SetFieldBorderOutsideRGB(val GraphicsRGBInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ViewForm", "borderOutsideRGB", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


