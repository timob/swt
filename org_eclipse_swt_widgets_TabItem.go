package swt

import "github.com/timob/javabind"

type WidgetsTabItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public org.eclipse.swt.widgets.TabFolder getParent()
	GetParent() *WidgetsTabFolder

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public void setControl(org.eclipse.swt.widgets.Control)
	SetControl(a WidgetsControlInterface) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 
}

type WidgetsTabItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TabItem(org.eclipse.swt.widgets.TabFolder, int)
func NewWidgetsTabItem(a WidgetsTabFolderInterface, b int) (*WidgetsTabItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TabItem", conv_a.Value().Cast("org/eclipse/swt/widgets/TabFolder"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTabItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TabItem(org.eclipse.swt.widgets.TabFolder, int, int)
func NewWidgetsTabItem2(a WidgetsTabFolderInterface, b int, c int) (*WidgetsTabItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TabItem", conv_a.Value().Cast("org/eclipse/swt/widgets/TabFolder"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTabItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsTabItem) GetBounds() *GraphicsRectangle {
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
func (jbobject *WidgetsTabItem) GetControl() *WidgetsControl {
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

// public org.eclipse.swt.widgets.TabFolder getParent()
func (jbobject *WidgetsTabItem) GetParent() *WidgetsTabFolder {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/TabFolder")
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
	unique_x := &WidgetsTabFolder{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getToolTipText()
func (jbobject *WidgetsTabItem) GetToolTipText() string {
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

// public void setControl(org.eclipse.swt.widgets.Control)
func (jbobject *WidgetsTabItem) SetControl(a WidgetsControlInterface)  {
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

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTabItem) SetImage(a GraphicsImageInterface)  {
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
func (jbobject *WidgetsTabItem) SetText(a string)  {
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
func (jbobject *WidgetsTabItem) SetToolTipText(a string)  {
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


