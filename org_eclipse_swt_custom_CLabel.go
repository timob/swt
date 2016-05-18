package swt

import "github.com/timob/javabind"

type CustomCLabelInterface interface {
	WidgetsCanvasInterface

	// public int getAlignment()
	GetAlignment() int

	// public int getBottomMargin()
	GetBottomMargin() int

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public int getLeftMargin()
	GetLeftMargin() int

	// public int getRightMargin()
	GetRightMargin() int

	// public java.lang.String getText()
	GetText() string

	// public int getTopMargin()
	GetTopMargin() int

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setBackground(org.eclipse.swt.graphics.Color[], int[])
	SetBackground3(a []*GraphicsColor, b []int) 

	// public void setBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
	SetBackground4(a []*GraphicsColor, b []int, c bool) 

	// public void setBackground(org.eclipse.swt.graphics.Image)
	SetBackground2(a GraphicsImageInterface) 

	// public void setBottomMargin(int)
	SetBottomMargin(a int) 

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setLeftMargin(int)
	SetLeftMargin(a int) 

	// public void setMargins(int, int, int, int)
	SetMargins(a int, b int, c int, d int) 

	// public void setRightMargin(int)
	SetRightMargin(a int) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setTopMargin(int)
	SetTopMargin(a int) 
}

type CustomCLabel struct {
	WidgetsCanvas
}

// public org.eclipse.swt.custom.CLabel(org.eclipse.swt.widgets.Composite, int)
func NewCustomCLabel(a WidgetsCompositeInterface, b int) (*CustomCLabel) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CLabel", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCLabel{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *CustomCLabel) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public int getAlignment()
func (jbobject *CustomCLabel) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getBottomMargin()
func (jbobject *CustomCLabel) GetBottomMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBottomMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *CustomCLabel) GetImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "org/eclipse/swt/graphics/Image")
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
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

// public int getLeftMargin()
func (jbobject *CustomCLabel) GetLeftMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLeftMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getRightMargin()
func (jbobject *CustomCLabel) GetRightMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRightMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getStyle()
func (jbobject *CustomCLabel) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *CustomCLabel) GetText() string {
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

// public java.lang.String getToolTipText()
func (jbobject *CustomCLabel) GetToolTipText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToolTipText", "java/lang/String")
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

// public int getTopMargin()
func (jbobject *CustomCLabel) GetTopMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setAlignment(int)
func (jbobject *CustomCLabel) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomCLabel) SetBackground(a GraphicsColorInterface)  {
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

// public void setBackground(org.eclipse.swt.graphics.Color[], int[])
func (jbobject *CustomCLabel) SetBackground3(a []*GraphicsColor, b []int)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackground(org.eclipse.swt.graphics.Color[], int[], boolean)
func (jbobject *CustomCLabel) SetBackground4(a []*GraphicsColor, b []int, c bool)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Color")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackground(org.eclipse.swt.graphics.Image)
func (jbobject *CustomCLabel) SetBackground2(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBottomMargin(int)
func (jbobject *CustomCLabel) SetBottomMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBottomMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomCLabel) SetFont(a GraphicsFontInterface)  {
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

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *CustomCLabel) SetImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setLeftMargin(int)
func (jbobject *CustomCLabel) SetLeftMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLeftMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMargins(int, int, int, int)
func (jbobject *CustomCLabel) SetMargins(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMargins", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void setRightMargin(int)
func (jbobject *CustomCLabel) SetRightMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRightMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *CustomCLabel) SetText(a string)  {
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

// public void setToolTipText(java.lang.String)
func (jbobject *CustomCLabel) SetToolTipText(a string)  {
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

// public void setTopMargin(int)
func (jbobject *CustomCLabel) SetTopMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


