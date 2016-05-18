package swt

import "github.com/timob/javabind"

type WidgetsButtonInterface interface {
	WidgetsControlInterface

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getAlignment()
	GetAlignment() int

	// public boolean getGrayed()
	GetGrayed() bool

	// public org.eclipse.swt.graphics.Image getImage()
	GetImage() *GraphicsImage

	// public boolean getSelection()
	GetSelection() bool

	// public java.lang.String getText()
	GetText() string

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setGrayed(boolean)
	SetGrayed(a bool) 

	// public void setImage(org.eclipse.swt.graphics.Image)
	SetImage(a GraphicsImageInterface) 

	// public void setSelection(boolean)
	SetSelection(a bool) 

	// public void setText(java.lang.String)
	SetText(a string) 
}

type WidgetsButton struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.Button(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsButton(a WidgetsCompositeInterface, b int) (*WidgetsButton) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Button", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsButton{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsButton) AddSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsButton) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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
func (jbobject *WidgetsButton) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getGrayed()
func (jbobject *WidgetsButton) GetGrayed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGrayed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsButton) GetImage() *GraphicsImage {
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

// public boolean getSelection()
func (jbobject *WidgetsButton) GetSelection() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getText()
func (jbobject *WidgetsButton) GetText() string {
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

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsButton) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setAlignment(int)
func (jbobject *WidgetsButton) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setGrayed(boolean)
func (jbobject *WidgetsButton) SetGrayed(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGrayed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsButton) SetImage(a GraphicsImageInterface)  {
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

// public void setSelection(boolean)
func (jbobject *WidgetsButton) SetSelection(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsButton) SetText(a string)  {
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


