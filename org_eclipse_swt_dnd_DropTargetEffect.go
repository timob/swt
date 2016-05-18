package swt

import "github.com/timob/javabind"

type DndDropTargetEffectInterface interface {
	DndDropTargetAdapterInterface

	// public org.eclipse.swt.widgets.Control getControl()
	GetControl() *WidgetsControl

	// public org.eclipse.swt.widgets.Widget getItem(int, int)
	GetItem(a int, b int) *WidgetsWidget
}

type DndDropTargetEffect struct {
	DndDropTargetAdapter
}

// public org.eclipse.swt.dnd.DropTargetEffect(org.eclipse.swt.widgets.Control)
func NewDndDropTargetEffect(a WidgetsControlInterface) (*DndDropTargetEffect) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/dnd/DropTargetEffect", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &DndDropTargetEffect{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Control getControl()
func (jbobject *DndDropTargetEffect) GetControl() *WidgetsControl {
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

// public org.eclipse.swt.widgets.Widget getItem(int, int)
func (jbobject *DndDropTargetEffect) GetItem(a int, b int) *WidgetsWidget {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/Widget", a, b)
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
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}


