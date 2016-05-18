package swt

import "github.com/timob/javabind"

type WidgetsIMEInterface interface {
	WidgetsWidgetInterface

	// public int getCaretOffset()
	GetCaretOffset() int

	// public int getCommitCount()
	GetCommitCount() int

	// public int getCompositionOffset()
	GetCompositionOffset() int

	// public int[] getRanges()
	GetRanges() []int

	// public org.eclipse.swt.graphics.TextStyle[] getStyles()
	GetStyles() []*GraphicsTextStyle

	// public java.lang.String getText()
	GetText() string

	// public boolean getWideCaret()
	GetWideCaret() bool

	// public void setCompositionOffset(int)
	SetCompositionOffset(a int) 
}

type WidgetsIME struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.IME(org.eclipse.swt.widgets.Canvas, int)
func NewWidgetsIME(a WidgetsCanvasInterface, b int) (*WidgetsIME) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/IME", conv_a.Value().Cast("org/eclipse/swt/widgets/Canvas"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsIME{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getCaretOffset()
func (jbobject *WidgetsIME) GetCaretOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCommitCount()
func (jbobject *WidgetsIME) GetCommitCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCommitCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCompositionOffset()
func (jbobject *WidgetsIME) GetCompositionOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCompositionOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getRanges()
func (jbobject *WidgetsIME) GetRanges() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRanges", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.graphics.TextStyle[] getStyles()
func (jbobject *WidgetsIME) GetStyles() []*GraphicsTextStyle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyles", javabind.ObjectArrayType("org/eclipse/swt/graphics/TextStyle"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/TextStyle")
	dst := new([]*GraphicsTextStyle)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String getText()
func (jbobject *WidgetsIME) GetText() string {
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

// public boolean getWideCaret()
func (jbobject *WidgetsIME) GetWideCaret() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWideCaret", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setCompositionOffset(int)
func (jbobject *WidgetsIME) SetCompositionOffset(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCompositionOffset", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


