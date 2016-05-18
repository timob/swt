package swt

import "github.com/timob/javabind"

type WidgetsFontDialogInterface interface {
	WidgetsDialogInterface

	// public boolean getEffectsVisible()
	GetEffectsVisible() bool

	// public org.eclipse.swt.graphics.FontData getFontData()
	GetFontData() *GraphicsFontData

	// public org.eclipse.swt.graphics.FontData[] getFontList()
	GetFontList() []*GraphicsFontData

	// public org.eclipse.swt.graphics.RGB getRGB()
	GetRGB() *GraphicsRGB

	// public org.eclipse.swt.graphics.FontData open()
	Open() *GraphicsFontData

	// public void setEffectsVisible(boolean)
	SetEffectsVisible(a bool) 

	// public void setFontData(org.eclipse.swt.graphics.FontData)
	SetFontData(a GraphicsFontDataInterface) 

	// public void setFontList(org.eclipse.swt.graphics.FontData[])
	SetFontList(a []*GraphicsFontData) 

	// public void setRGB(org.eclipse.swt.graphics.RGB)
	SetRGB(a GraphicsRGBInterface) 
}

type WidgetsFontDialog struct {
	WidgetsDialog
}

// public org.eclipse.swt.widgets.FontDialog(org.eclipse.swt.widgets.Shell)
func NewWidgetsFontDialog(a WidgetsShellInterface) (*WidgetsFontDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/FontDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsFontDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.FontDialog(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsFontDialog2(a WidgetsShellInterface, b int) (*WidgetsFontDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/FontDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsFontDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean getEffectsVisible()
func (jbobject *WidgetsFontDialog) GetEffectsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEffectsVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.FontData getFontData()
func (jbobject *WidgetsFontDialog) GetFontData() *GraphicsFontData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFontData", "org/eclipse/swt/graphics/FontData")
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
	unique_x := &GraphicsFontData{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.FontData[] getFontList()
func (jbobject *WidgetsFontDialog) GetFontList() []*GraphicsFontData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFontList", javabind.ObjectArrayType("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/FontData")
	dst := new([]*GraphicsFontData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.graphics.RGB getRGB()
func (jbobject *WidgetsFontDialog) GetRGB() *GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRGB", "org/eclipse/swt/graphics/RGB")
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

// public org.eclipse.swt.graphics.FontData open()
func (jbobject *WidgetsFontDialog) Open() *GraphicsFontData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", "org/eclipse/swt/graphics/FontData")
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
	unique_x := &GraphicsFontData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEffectsVisible(boolean)
func (jbobject *WidgetsFontDialog) SetEffectsVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEffectsVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFontData(org.eclipse.swt.graphics.FontData)
func (jbobject *WidgetsFontDialog) SetFontData(a GraphicsFontDataInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFontData", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFontList(org.eclipse.swt.graphics.FontData[])
func (jbobject *WidgetsFontDialog) SetFontList(a []*GraphicsFontData)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/FontData")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFontList", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRGB(org.eclipse.swt.graphics.RGB)
func (jbobject *WidgetsFontDialog) SetRGB(a GraphicsRGBInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRGB", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


