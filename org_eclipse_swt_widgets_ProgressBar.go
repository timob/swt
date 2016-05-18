package swt

import "github.com/timob/javabind"

type WidgetsProgressBarInterface interface {
	WidgetsControlInterface

	// public int getMaximum()
	GetMaximum() int

	// public int getMinimum()
	GetMinimum() int

	// public int getSelection()
	GetSelection() int

	// public int getState()
	GetState() int

	// public void setMaximum(int)
	SetMaximum(a int) 

	// public void setMinimum(int)
	SetMinimum(a int) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setState(int)
	SetState(a int) 
}

type WidgetsProgressBar struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.ProgressBar(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsProgressBar(a WidgetsCompositeInterface, b int) (*WidgetsProgressBar) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ProgressBar", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsProgressBar{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getMaximum()
func (jbobject *WidgetsProgressBar) GetMaximum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getMinimum()
func (jbobject *WidgetsProgressBar) GetMinimum() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinimum", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSelection()
func (jbobject *WidgetsProgressBar) GetSelection() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getState()
func (jbobject *WidgetsProgressBar) GetState() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setMaximum(int)
func (jbobject *WidgetsProgressBar) SetMaximum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMinimum(int)
func (jbobject *WidgetsProgressBar) SetMinimum(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinimum", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsProgressBar) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setState(int)
func (jbobject *WidgetsProgressBar) SetState(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


