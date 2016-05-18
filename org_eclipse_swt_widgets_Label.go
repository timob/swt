package swt

import "github.com/timob/javabind"

type WidgetsLabelInterface interface {
	WidgetsControlInterface

	// public int getAlignment()
	GetAlignment() int

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public java.lang.String getText()
	GetText() string

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setText(java.lang.String)
	SetText(a string) 
}

type WidgetsLabel struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.Label(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsLabel(a WidgetsCompositeInterface, b int) (*WidgetsLabel) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Label", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsLabel{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsLabel) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsLabel) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getBorderWidth()
func (jbobject *WidgetsLabel) GetBorderWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsLabel) GetImage() *GraphicsImage {
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

// public java.lang.String getText()
func (jbobject *WidgetsLabel) GetText() string {
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

// public void setAlignment(int)
func (jbobject *WidgetsLabel) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsLabel) SetImage(a GraphicsImageInterface)  {
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

// public void setText(java.lang.String)
func (jbobject *WidgetsLabel) SetText(a string)  {
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


