package swt

import "github.com/timob/javabind"

type WidgetsItemInterface interface {
	WidgetsWidgetInterface

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public java.lang.String getText()
	GetText() string

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setText(java.lang.String)
	SetText(a string) 
}

type WidgetsItem struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.Item(org.eclipse.swt.widgets.Widget, int)
func NewWidgetsItem(a WidgetsWidgetInterface, b int) (*WidgetsItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Item", conv_a.Value().Cast("org/eclipse/swt/widgets/Widget"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Item(org.eclipse.swt.widgets.Widget, int, int)
func NewWidgetsItem2(a WidgetsWidgetInterface, b int, c int) (*WidgetsItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Item", conv_a.Value().Cast("org/eclipse/swt/widgets/Widget"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsItem) GetImage() *GraphicsImage {
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
func (jbobject *WidgetsItem) GetText() string {
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

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsItem) SetImage(a GraphicsImageInterface)  {
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
func (jbobject *WidgetsItem) SetText(a string)  {
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


