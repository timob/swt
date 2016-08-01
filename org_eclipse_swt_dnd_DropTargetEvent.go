package swt

import "github.com/timob/javabind"

type DndDropTargetEventInterface interface {
	EventsTypedEventInterface
}

type DndDropTargetEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.dnd.DropTargetEvent(org.eclipse.swt.dnd.DNDEvent)
func NewDndDropTargetEvent(a DndDNDEventInterface) (*DndDropTargetEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/dnd/DropTargetEvent", conv_a.Value().Cast("org/eclipse/swt/dnd/DNDEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &DndDropTargetEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *DndDropTargetEvent) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

func (jbobject *DndDropTargetEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDropTargetEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDropTargetEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDropTargetEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDropTargetEvent) Detail() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "detail", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDropTargetEvent) SetFieldDetail(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "detail", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDropTargetEvent) Operations() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "operations", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDropTargetEvent) SetFieldOperations(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "operations", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDropTargetEvent) Feedback() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "feedback", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDropTargetEvent) SetFieldFeedback(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "feedback", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDropTargetEvent) Item() *WidgetsWidget {
	jret, err := jbobject.GetField(javabind.GetEnv(), "item", "org/eclipse/swt/widgets/Widget")
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

func (jbobject *DndDropTargetEvent) SetFieldItem(val WidgetsWidgetInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "item", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *DndDropTargetEvent) CurrentDataType() *DndTransferData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "currentDataType", "org/eclipse/swt/dnd/TransferData")
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
	unique_x := &DndTransferData{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *DndDropTargetEvent) SetFieldCurrentDataType(val DndTransferDataInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "currentDataType", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *DndDropTargetEvent) DataTypes() []*DndTransferData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "dataTypes", javabind.ObjectArrayType("org/eclipse/swt/dnd/TransferData"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/dnd/TransferData")
	dst := new([]*DndTransferData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *DndDropTargetEvent) SetFieldDataTypes(val []*DndTransferData) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/dnd/TransferData")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "dataTypes", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


