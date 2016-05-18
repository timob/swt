package swt

import "github.com/timob/javabind"

type WidgetsExpandItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public boolean getExpanded()
	GetExpanded() bool

	// public int getHeaderHeight()
	GetHeaderHeight() int

	// public int getHeight()
	GetHeight() int

	// public org.eclipse.swt.widgets.ExpandBar getParent()
	GetParent() *WidgetsExpandBar

	// public void setControl(org.eclipse.swt.widgets.Control)
	SetControl(a WidgetsControlInterface) 

	// public void setExpanded(boolean)
	SetExpanded(a bool) 

	// public void setHeight(int)
	SetHeight(a int) 
}

type WidgetsExpandItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.ExpandItem(org.eclipse.swt.widgets.ExpandBar, int)
func NewWidgetsExpandItem(a WidgetsExpandBarInterface, b int) (*WidgetsExpandItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ExpandItem", conv_a.Value().Cast("org/eclipse/swt/widgets/ExpandBar"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsExpandItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.ExpandItem(org.eclipse.swt.widgets.ExpandBar, int, int)
func NewWidgetsExpandItem2(a WidgetsExpandBarInterface, b int, c int) (*WidgetsExpandItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ExpandItem", conv_a.Value().Cast("org/eclipse/swt/widgets/ExpandBar"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsExpandItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Control getControl()
func (jbobject *WidgetsExpandItem) GetControl() *WidgetsControl {
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

// public boolean getExpanded()
func (jbobject *WidgetsExpandItem) GetExpanded() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpanded", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getHeaderHeight()
func (jbobject *WidgetsExpandItem) GetHeaderHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHeight()
func (jbobject *WidgetsExpandItem) GetHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.ExpandBar getParent()
func (jbobject *WidgetsExpandItem) GetParent() *WidgetsExpandBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/ExpandBar")
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
	unique_x := &WidgetsExpandBar{}
	unique_x.Callable = dst
	return unique_x
}

// public void setControl(org.eclipse.swt.widgets.Control)
func (jbobject *WidgetsExpandItem) SetControl(a WidgetsControlInterface)  {
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

// public void setExpanded(boolean)
func (jbobject *WidgetsExpandItem) SetExpanded(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpanded", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHeight(int)
func (jbobject *WidgetsExpandItem) SetHeight(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHeight", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsExpandItem) SetImage(a GraphicsImageInterface)  {
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
func (jbobject *WidgetsExpandItem) SetText(a string)  {
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


