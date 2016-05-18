package swt

import "github.com/timob/javabind"

type DndTransferDataInterface interface {
	JavaLangObjectInterface
}

type DndTransferData struct {
	JavaLangObject
}

// public org.eclipse.swt.dnd.TransferData()
func NewDndTransferData() (*DndTransferData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/dnd/TransferData")
	if err != nil {
		panic(err)
	}
	x := &DndTransferData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *DndTransferData) Type() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *DndTransferData) SetFieldType(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndTransferData) Length() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndTransferData) SetFieldLength(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "length", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndTransferData) Format() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "format", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndTransferData) SetFieldFormat(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "format", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndTransferData) PValue() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "pValue", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *DndTransferData) SetFieldPValue(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "pValue", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *DndTransferData) Result() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "result", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *DndTransferData) SetFieldResult(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "result", val)
	if err != nil {
		panic(err)
	}

}


