package swt

import "github.com/timob/javabind"

type CustomControlEditorInterface interface {
	JavaLangObjectInterface

	// public void dispose()
	Dispose() 

	// public org.eclipse.swt.widgets.Control getEditor()
	GetEditor() *WidgetsControl

	// public void layout()
	Layout() 

	// public void setEditor(org.eclipse.swt.widgets.Control)
	SetEditor(a WidgetsControlInterface) 
}

type CustomControlEditor struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.ControlEditor(org.eclipse.swt.widgets.Composite)
func NewCustomControlEditor(a WidgetsCompositeInterface) (*CustomControlEditor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/ControlEditor", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomControlEditor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *CustomControlEditor) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.widgets.Control getEditor()
func (jbobject *CustomControlEditor) GetEditor() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEditor", "org/eclipse/swt/widgets/Control")
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

// public void layout()
func (jbobject *CustomControlEditor) Layout()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setEditor(org.eclipse.swt.widgets.Control)
func (jbobject *CustomControlEditor) SetEditor(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *CustomControlEditor) HorizontalAlignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomControlEditor) SetFieldHorizontalAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalAlignment", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomControlEditor) GrabHorizontal() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "grabHorizontal", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomControlEditor) SetFieldGrabHorizontal(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "grabHorizontal", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomControlEditor) MinimumWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "minimumWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomControlEditor) SetFieldMinimumWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "minimumWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomControlEditor) VerticalAlignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomControlEditor) SetFieldVerticalAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalAlignment", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomControlEditor) GrabVertical() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "grabVertical", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomControlEditor) SetFieldGrabVertical(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "grabVertical", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomControlEditor) MinimumHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "minimumHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomControlEditor) SetFieldMinimumHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "minimumHeight", val)
	if err != nil {
		panic(err)
	}

}


