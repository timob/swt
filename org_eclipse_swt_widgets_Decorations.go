package swt

import "github.com/timob/javabind"

type WidgetsDecorationsInterface interface {
	WidgetsCanvasInterface

	// public org.eclipse.swt.widgets.Button getDefaultButton()
	GetDefaultButton() *WidgetsButton

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public org.eclipse.swt.graphics.Image[] getImages()
	GetImages() []*GraphicsImage

	// public boolean getMaximized()
	GetMaximized() bool

	// public org.eclipse.swt.widgets.Menu getMenuBar()
	GetMenuBar() *WidgetsMenu

	// public boolean getMinimized()
	GetMinimized() bool

	// public java.lang.String getText()
	GetText() string

	// public void setDefaultButton(org.eclipse.swt.widgets.Button)
	SetDefaultButton(a WidgetsButtonInterface) 

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setImages(org.eclipse.swt.graphics.Image[])
	SetImages(a []*GraphicsImage) 

	// public void setMaximized(boolean)
	SetMaximized(a bool) 

	// public void setMenuBar(org.eclipse.swt.widgets.Menu)
	SetMenuBar(a WidgetsMenuInterface) 

	// public void setMinimized(boolean)
	SetMinimized(a bool) 

	// public void setText(java.lang.String)
	SetText(a string) 
}

type WidgetsDecorations struct {
	WidgetsCanvas
}

// public org.eclipse.swt.widgets.Decorations(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsDecorations(a WidgetsCompositeInterface, b int) (*WidgetsDecorations) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Decorations", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDecorations{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Button getDefaultButton()
func (jbobject *WidgetsDecorations) GetDefaultButton() *WidgetsButton {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDefaultButton", "org/eclipse/swt/widgets/Button")
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
	unique_x := &WidgetsButton{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsDecorations) GetImage() *GraphicsImage {
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

// public org.eclipse.swt.graphics.Image[] getImages()
func (jbobject *WidgetsDecorations) GetImages() []*GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImages", javabind.ObjectArrayType("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/Image")
	dst := new([]*GraphicsImage)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getMaximized()
func (jbobject *WidgetsDecorations) GetMaximized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.Menu getMenuBar()
func (jbobject *WidgetsDecorations) GetMenuBar() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMenuBar", "org/eclipse/swt/widgets/Menu")
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
	unique_x := &WidgetsMenu{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getMinimized()
func (jbobject *WidgetsDecorations) GetMinimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getText()
func (jbobject *WidgetsDecorations) GetText() string {
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

// public boolean isReparentable()
func (jbobject *WidgetsDecorations) IsReparentable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isReparentable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setDefaultButton(org.eclipse.swt.widgets.Button)
func (jbobject *WidgetsDecorations) SetDefaultButton(a WidgetsButtonInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultButton", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Button"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsDecorations) SetImage(a GraphicsImageInterface)  {
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

// public void setImages(org.eclipse.swt.graphics.Image[])
func (jbobject *WidgetsDecorations) SetImages(a []*GraphicsImage)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Image")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImages", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMaximized(boolean)
func (jbobject *WidgetsDecorations) SetMaximized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMenuBar(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsDecorations) SetMenuBar(a WidgetsMenuInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMenuBar", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMinimized(boolean)
func (jbobject *WidgetsDecorations) SetMinimized(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimized", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsDecorations) SetText(a string)  {
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


