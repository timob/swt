package swt

import "github.com/timob/javabind"

type CustomPopupListInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public java.lang.String[] getItems()
	GetItems() []string

	// public int getMinimumWidth()
	GetMinimumWidth() int

	// public java.lang.String open(org.eclipse.swt.graphics.Rectangle)
	Open(a GraphicsRectangleInterface) string

	// public void select(java.lang.String)
	Select(a string) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setItems(java.lang.String[])
	SetItems(a []string) 

	// public void setMinimumWidth(int)
	SetMinimumWidth(a int) 
}

type CustomPopupList struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.PopupList(org.eclipse.swt.widgets.Shell)
func NewCustomPopupList(a WidgetsShellInterface) (*CustomPopupList) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/PopupList", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomPopupList{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.PopupList(org.eclipse.swt.widgets.Shell, int)
func NewCustomPopupList2(a WidgetsShellInterface, b int) (*CustomPopupList) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/PopupList", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomPopupList{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *CustomPopupList) GetFont() *GraphicsFont {
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

// public java.lang.String[] getItems()
func (jbobject *CustomPopupList) GetItems() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getMinimumWidth()
func (jbobject *CustomPopupList) GetMinimumWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimumWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String open(org.eclipse.swt.graphics.Rectangle)
func (jbobject *CustomPopupList) Open(a GraphicsRectangleInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", "java/lang/String", conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void select(java.lang.String)
func (jbobject *CustomPopupList) Select(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "select", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomPopupList) SetFont(a GraphicsFontInterface)  {
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

// public void setItems(java.lang.String[])
func (jbobject *CustomPopupList) SetItems(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItems", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMinimumWidth(int)
func (jbobject *CustomPopupList) SetMinimumWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimumWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


