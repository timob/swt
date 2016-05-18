package swt

import "github.com/timob/javabind"

type WidgetsCanvasInterface interface {
	WidgetsCompositeInterface

	// public void drawBackground(org.eclipse.swt.graphics.GC, int, int, int, int)
	DrawBackground2(a GraphicsGCInterface, b int, c int, d int, e int) 

	// public org.eclipse.swt.widgets.Caret getCaret()
	GetCaret() *WidgetsCaret

	// public org.eclipse.swt.widgets.IME getIME()
	GetIME() *WidgetsIME

	// public void scroll(int, int, int, int, int, int, boolean)
	Scroll(a int, b int, c int, d int, e int, f int, g bool) 

	// public void setCaret(org.eclipse.swt.widgets.Caret)
	SetCaret(a WidgetsCaretInterface) 

	// public void setIME(org.eclipse.swt.widgets.IME)
	SetIME(a WidgetsIMEInterface) 
}

type WidgetsCanvas struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.Canvas(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsCanvas(a WidgetsCompositeInterface, b int) (*WidgetsCanvas) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Canvas", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsCanvas{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void drawBackground(org.eclipse.swt.graphics.GC, int, int, int, int)
func (jbobject *WidgetsCanvas) DrawBackground2(a GraphicsGCInterface, b int, c int, d int, e int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"), b, c, d, e)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.widgets.Caret getCaret()
func (jbobject *WidgetsCanvas) GetCaret() *WidgetsCaret {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaret", "org/eclipse/swt/widgets/Caret")
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
	unique_x := &WidgetsCaret{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.IME getIME()
func (jbobject *WidgetsCanvas) GetIME() *WidgetsIME {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIME", "org/eclipse/swt/widgets/IME")
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
	unique_x := &WidgetsIME{}
	unique_x.Callable = dst
	return unique_x
}

// public void scroll(int, int, int, int, int, int, boolean)
func (jbobject *WidgetsCanvas) Scroll(a int, b int, c int, d int, e int, f int, g bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "scroll", javabind.Void, a, b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}

}

// public void setCaret(org.eclipse.swt.widgets.Caret)
func (jbobject *WidgetsCanvas) SetCaret(a WidgetsCaretInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCaret", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Caret"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsCanvas) SetFont(a GraphicsFontInterface)  {
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

// public void setIME(org.eclipse.swt.widgets.IME)
func (jbobject *WidgetsCanvas) SetIME(a WidgetsIMEInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIME", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/IME"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


