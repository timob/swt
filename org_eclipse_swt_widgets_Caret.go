package swt

import "github.com/timob/javabind"

type WidgetsCaretInterface interface {
	WidgetsWidgetInterface

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public org.eclipse.swt.graphics.Point getLocation()
	GetLocation() *GraphicsPoint

	// public org.eclipse.swt.widgets.Canvas getParent()
	GetParent() *WidgetsCanvas

	// public org.eclipse.swt.graphics.Point getSize()
	GetSize() *GraphicsPoint

	// public boolean getVisible()
	GetVisible() bool

	// public boolean isVisible()
	IsVisible() bool

	// public void setBounds(int, int, int, int)
	SetBounds(a int, b int, c int, d int) 

	// public void setBounds(org.eclipse.swt.graphics.Rectangle)
	SetBounds2(a GraphicsRectangleInterface) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setLocation(int, int)
	SetLocation(a int, b int) 

	// public void setLocation(org.eclipse.swt.graphics.Point)
	SetLocation2(a GraphicsPointInterface) 

	// public void setSize(int, int)
	SetSize(a int, b int) 

	// public void setSize(org.eclipse.swt.graphics.Point)
	SetSize2(a GraphicsPointInterface) 

	// public void setVisible(boolean)
	SetVisible(a bool) 
}

type WidgetsCaret struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.Caret(org.eclipse.swt.widgets.Canvas, int)
func NewWidgetsCaret(a WidgetsCanvasInterface, b int) (*WidgetsCaret) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Caret", conv_a.Value().Cast("org/eclipse/swt/widgets/Canvas"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCaret{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsCaret) GetBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle")
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

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *WidgetsCaret) GetFont() *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFont", "org/eclipse/swt/graphics/Font")
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
	unique_x := &GraphicsFont{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsCaret) GetImage() *GraphicsImage {
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

// public org.eclipse.swt.graphics.Point getLocation()
func (jbobject *WidgetsCaret) GetLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocation", "org/eclipse/swt/graphics/Point")
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

// public org.eclipse.swt.widgets.Canvas getParent()
func (jbobject *WidgetsCaret) GetParent() *WidgetsCanvas {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Canvas")
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
	unique_x := &WidgetsCanvas{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Point getSize()
func (jbobject *WidgetsCaret) GetSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "org/eclipse/swt/graphics/Point")
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

// public boolean getVisible()
func (jbobject *WidgetsCaret) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsCaret) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setBounds(int, int, int, int)
func (jbobject *WidgetsCaret) SetBounds(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBounds", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void setBounds(org.eclipse.swt.graphics.Rectangle)
func (jbobject *WidgetsCaret) SetBounds2(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBounds", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsCaret) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *WidgetsCaret) SetImage(a GraphicsImageInterface)  {
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

// public void setLocation(int, int)
func (jbobject *WidgetsCaret) SetLocation(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setLocation(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCaret) SetLocation2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSize(int, int)
func (jbobject *WidgetsCaret) SetSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsCaret) SetSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setVisible(boolean)
func (jbobject *WidgetsCaret) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


