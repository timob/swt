package swt

import "github.com/timob/javabind"

type CustomStyledTextDropTargetEffectInterface interface {
	DndDropTargetEffectInterface
}

type CustomStyledTextDropTargetEffect struct {
	DndDropTargetEffect
}

// public org.eclipse.swt.custom.StyledTextDropTargetEffect(org.eclipse.swt.custom.StyledText)
func NewCustomStyledTextDropTargetEffect(a CustomStyledTextInterface) (*CustomStyledTextDropTargetEffect) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyledTextDropTargetEffect", conv_a.Value().Cast("org/eclipse/swt/custom/StyledText"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomStyledTextDropTargetEffect{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dragEnter(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *CustomStyledTextDropTargetEffect) DragEnter(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dragEnter", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dragLeave(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *CustomStyledTextDropTargetEffect) DragLeave(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dragLeave", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dragOver(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *CustomStyledTextDropTargetEffect) DragOver(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dragOver", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dropAccept(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *CustomStyledTextDropTargetEffect) DropAccept(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dropAccept", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


