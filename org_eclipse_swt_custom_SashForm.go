package swt

import "github.com/timob/javabind"

type CustomSashFormInterface interface {
	WidgetsCompositeInterface

	// public int getSashWidth()
	GetSashWidth() int

	// public org.eclipse.swt.widgets.Control getMaximizedControl()
	GetMaximizedControl() *WidgetsControl

	// public int[] getWeights()
	GetWeights() []int

	// public void setMaximizedControl(org.eclipse.swt.widgets.Control)
	SetMaximizedControl(a WidgetsControlInterface) 

	// public void setSashWidth(int)
	SetSashWidth(a int) 

	// public void setWeights(int[])
	SetWeights(a []int) 
}

type CustomSashForm struct {
	WidgetsComposite
}

// public org.eclipse.swt.custom.SashForm(org.eclipse.swt.widgets.Composite, int)
func NewCustomSashForm(a WidgetsCompositeInterface, b int) (*CustomSashForm) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/SashForm", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomSashForm{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getOrientation()
func (jbobject *CustomSashForm) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getSashWidth()
func (jbobject *CustomSashForm) GetSashWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSashWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getStyle()
func (jbobject *CustomSashForm) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Control getMaximizedControl()
func (jbobject *CustomSashForm) GetMaximizedControl() *WidgetsControl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaximizedControl", "org/eclipse/swt/widgets/Control")
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

// public int[] getWeights()
func (jbobject *CustomSashForm) GetWeights() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWeights", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public void setOrientation(int)
func (jbobject *CustomSashForm) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomSashForm) SetBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomSashForm) SetForeground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForeground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setLayout(org.eclipse.swt.widgets.Layout)
func (jbobject *CustomSashForm) SetLayout(a WidgetsLayoutInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLayout", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Layout"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMaximizedControl(org.eclipse.swt.widgets.Control)
func (jbobject *CustomSashForm) SetMaximizedControl(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaximizedControl", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSashWidth(int)
func (jbobject *CustomSashForm) SetSashWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSashWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setToolTipText(java.lang.String)
func (jbobject *CustomSashForm) SetToolTipText(a string)  {
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

// public void setWeights(int[])
func (jbobject *CustomSashForm) SetWeights(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWeights", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomSashForm) SASH_WIDTH() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "SASH_WIDTH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomSashForm) SetFieldSASH_WIDTH(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "SASH_WIDTH", val)
	if err != nil {
		panic(err)
	}

}


