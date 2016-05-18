package swt

import "github.com/timob/javabind"

type WidgetsColorDialogInterface interface {
	WidgetsDialogInterface

	// public org.eclipse.swt.graphics.RGB getRGB()
	GetRGB() *GraphicsRGB

	// public org.eclipse.swt.graphics.RGB[] getRGBs()
	GetRGBs() []*GraphicsRGB

	// public org.eclipse.swt.graphics.RGB open()
	Open() *GraphicsRGB

	// public void setRGB(org.eclipse.swt.graphics.RGB)
	SetRGB(a GraphicsRGBInterface) 

	// public void setRGBs(org.eclipse.swt.graphics.RGB[])
	SetRGBs(a []*GraphicsRGB) 
}

type WidgetsColorDialog struct {
	WidgetsDialog
}

// public org.eclipse.swt.widgets.ColorDialog(org.eclipse.swt.widgets.Shell)
func NewWidgetsColorDialog(a WidgetsShellInterface) (*WidgetsColorDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ColorDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsColorDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.ColorDialog(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsColorDialog2(a WidgetsShellInterface, b int) (*WidgetsColorDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ColorDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsColorDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.RGB getRGB()
func (jbobject *WidgetsColorDialog) GetRGB() *GraphicsRGB {
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

// public org.eclipse.swt.graphics.RGB[] getRGBs()
func (jbobject *WidgetsColorDialog) GetRGBs() []*GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRGBs", javabind.ObjectArrayType("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/RGB")
	dst := new([]*GraphicsRGB)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.graphics.RGB open()
func (jbobject *WidgetsColorDialog) Open() *GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", "org/eclipse/swt/graphics/RGB")
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

// public void setRGB(org.eclipse.swt.graphics.RGB)
func (jbobject *WidgetsColorDialog) SetRGB(a GraphicsRGBInterface)  {
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

// public void setRGBs(org.eclipse.swt.graphics.RGB[])
func (jbobject *WidgetsColorDialog) SetRGBs(a []*GraphicsRGB)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/RGB")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRGBs", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


