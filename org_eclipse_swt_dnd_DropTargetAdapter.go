package swt

import "github.com/timob/javabind"

type DndDropTargetAdapterInterface interface {
	JavaLangObjectInterface

	// public void dragEnter(org.eclipse.swt.dnd.DropTargetEvent)
	DragEnter(a DndDropTargetEventInterface) 

	// public void dragLeave(org.eclipse.swt.dnd.DropTargetEvent)
	DragLeave(a DndDropTargetEventInterface) 

	// public void dragOperationChanged(org.eclipse.swt.dnd.DropTargetEvent)
	DragOperationChanged(a DndDropTargetEventInterface) 

	// public void dragOver(org.eclipse.swt.dnd.DropTargetEvent)
	DragOver(a DndDropTargetEventInterface) 

	// public void drop(org.eclipse.swt.dnd.DropTargetEvent)
	Drop(a DndDropTargetEventInterface) 

	// public void dropAccept(org.eclipse.swt.dnd.DropTargetEvent)
	DropAccept(a DndDropTargetEventInterface) 
}

type DndDropTargetAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.dnd.DropTargetAdapter()
func NewDndDropTargetAdapter() (*DndDropTargetAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/dnd/DropTargetAdapter")
	if err != nil {
		panic(err)
	}
	x := &DndDropTargetAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dragEnter(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *DndDropTargetAdapter) DragEnter(a DndDropTargetEventInterface)  {
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
func (jbobject *DndDropTargetAdapter) DragLeave(a DndDropTargetEventInterface)  {
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

// public void dragOperationChanged(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *DndDropTargetAdapter) DragOperationChanged(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dragOperationChanged", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dragOver(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *DndDropTargetAdapter) DragOver(a DndDropTargetEventInterface)  {
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

// public void drop(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *DndDropTargetAdapter) Drop(a DndDropTargetEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drop", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/dnd/DropTargetEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dropAccept(org.eclipse.swt.dnd.DropTargetEvent)
func (jbobject *DndDropTargetAdapter) DropAccept(a DndDropTargetEventInterface)  {
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


