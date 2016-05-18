package swt

import "github.com/timob/javabind"

type CustomCTabItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public org.eclipse.swt.graphics.Image getDisabledImage()
	GetDisabledImage() *GraphicsImage

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.custom.CTabFolder getParent()
	GetParent() *CustomCTabFolder

	// public boolean getShowClose()
	GetShowClose() bool

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public boolean isShowing()
	IsShowing() bool

	// public void setControl(org.eclipse.swt.widgets.Control)
	SetControl(a WidgetsControlInterface) 

	// public void setDisabledImage(org.eclipse.swt.graphics.Image)
	SetDisabledImage(a GraphicsImageInterface) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setShowClose(boolean)
	SetShowClose(a bool) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 
}

type CustomCTabItem struct {
	WidgetsItem
}

// public org.eclipse.swt.custom.CTabItem(org.eclipse.swt.custom.CTabFolder, int)
func NewCustomCTabItem(a CustomCTabFolderInterface, b int) (*CustomCTabItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabItem", conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolder"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCTabItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.CTabItem(org.eclipse.swt.custom.CTabFolder, int, int)
func NewCustomCTabItem2(a CustomCTabFolderInterface, b int, c int) (*CustomCTabItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabItem", conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolder"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomCTabItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *CustomCTabItem) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *CustomCTabItem) GetBounds() *GraphicsRectangle {
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

// public org.eclipse.swt.widgets.Control getControl()
func (jbobject *CustomCTabItem) GetControl() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getControl", "org/eclipse/swt/widgets/Control")
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

// public org.eclipse.swt.graphics.Image getDisabledImage()
func (jbobject *CustomCTabItem) GetDisabledImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisabledImage", "org/eclipse/swt/graphics/Image")
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

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *CustomCTabItem) GetFont() *GraphicsFont {
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

// public org.eclipse.swt.custom.CTabFolder getParent()
func (jbobject *CustomCTabItem) GetParent() *CustomCTabFolder {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/custom/CTabFolder")
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
	unique_x := &CustomCTabFolder{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getShowClose()
func (jbobject *CustomCTabItem) GetShowClose() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getShowClose", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getToolTipText()
func (jbobject *CustomCTabItem) GetToolTipText() string {
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

// public boolean isShowing()
func (jbobject *CustomCTabItem) IsShowing() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isShowing", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setControl(org.eclipse.swt.widgets.Control)
func (jbobject *CustomCTabItem) SetControl(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setControl", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setDisabledImage(org.eclipse.swt.graphics.Image)
func (jbobject *CustomCTabItem) SetDisabledImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDisabledImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomCTabItem) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *CustomCTabItem) SetImage(a GraphicsImageInterface)  {
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

// public void setShowClose(boolean)
func (jbobject *CustomCTabItem) SetShowClose(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setShowClose", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *CustomCTabItem) SetText(a string)  {
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
func (jbobject *CustomCTabItem) SetToolTipText(a string)  {
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


