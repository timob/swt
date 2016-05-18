package swt

import "github.com/timob/javabind"

type WidgetsScrollableInterface interface {
	WidgetsControlInterface

	// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
	ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle getClientArea()
	GetClientArea() *GraphicsRectangle

	// public org.eclipse.swt.widgets.ScrollBar getHorizontalBar()
	GetHorizontalBar() *WidgetsScrollBar

	// public int getScrollbarsMode()
	GetScrollbarsMode() int

	// public org.eclipse.swt.widgets.ScrollBar getVerticalBar()
	GetVerticalBar() *WidgetsScrollBar
}

type WidgetsScrollable struct {
	WidgetsControl
}

// public org.eclipse.swt.widgets.Scrollable(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsScrollable(a WidgetsCompositeInterface, b int) (*WidgetsScrollable) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Scrollable", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsScrollable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *WidgetsScrollable) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeTrim", "org/eclipse/swt/graphics/Rectangle", a, b, c, d)
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

// public int getBorderWidth()
func (jbobject *WidgetsScrollable) GetBorderWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *WidgetsScrollable) GetClientArea() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientArea", "org/eclipse/swt/graphics/Rectangle")
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

// public org.eclipse.swt.widgets.ScrollBar getHorizontalBar()
func (jbobject *WidgetsScrollable) GetHorizontalBar() *WidgetsScrollBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHorizontalBar", "org/eclipse/swt/widgets/ScrollBar")
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
	unique_x := &WidgetsScrollBar{}
	unique_x.Callable = dst
	return unique_x
}

// public int getScrollbarsMode()
func (jbobject *WidgetsScrollable) GetScrollbarsMode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScrollbarsMode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.ScrollBar getVerticalBar()
func (jbobject *WidgetsScrollable) GetVerticalBar() *WidgetsScrollBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVerticalBar", "org/eclipse/swt/widgets/ScrollBar")
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
	unique_x := &WidgetsScrollBar{}
	unique_x.Callable = dst
	return unique_x
}


