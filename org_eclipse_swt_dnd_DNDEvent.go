package swt

import "github.com/timob/javabind"

type DndDNDEventInterface interface {
	WidgetsEventInterface
}

type DndDNDEvent struct {
	WidgetsEvent
}

func (jbobject *DndDNDEvent) DataType() *DndTransferData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "dataType", "org/eclipse/swt/dnd/TransferData")
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

func (jbobject *DndDNDEvent) SetFieldDataType(val DndTransferDataInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "dataType", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *DndDNDEvent) DataTypes() []*DndTransferData {
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

func (jbobject *DndDNDEvent) SetFieldDataTypes(val []*DndTransferData) {
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

func (jbobject *DndDNDEvent) Operations() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "operations", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDNDEvent) SetFieldOperations(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "operations", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDNDEvent) Feedback() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "feedback", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDNDEvent) SetFieldFeedback(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "feedback", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDNDEvent) Image() *GraphicsImage {
	jret, err := jbobject.GetField(javabind.GetEnv(), "image", "org/eclipse/swt/graphics/Image")
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
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *DndDNDEvent) SetFieldImage(val GraphicsImageInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "image", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *DndDNDEvent) OffsetX() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "offsetX", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDNDEvent) SetFieldOffsetX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "offsetX", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndDNDEvent) OffsetY() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "offsetY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndDNDEvent) SetFieldOffsetY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "offsetY", val)
	if err != nil {
		panic(err)
	}

}


