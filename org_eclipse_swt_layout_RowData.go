package swt

import "github.com/timob/javabind"

type LayoutRowDataInterface interface {
	JavaLangObjectInterface
}

type LayoutRowData struct {
	JavaLangObject
}

// public org.eclipse.swt.layout.RowData()
func NewLayoutRowData() (*LayoutRowData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/RowData")
	if err != nil {
		panic(err)
	}
	x := &LayoutRowData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.RowData(int, int)
func NewLayoutRowData2(a int, b int) (*LayoutRowData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/RowData", a, b)
	if err != nil {
		panic(err)
	}
	x := &LayoutRowData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.RowData(org.eclipse.swt.graphics.Point)
func NewLayoutRowData3(a GraphicsPointInterface) (*LayoutRowData) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/RowData", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &LayoutRowData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutRowData) ToString() string {
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

func (jbobject *LayoutRowData) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowData) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowData) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowData) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowData) Exclude() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "exclude", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowData) SetFieldExclude(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "exclude", val)
	if err != nil {
		panic(err)
	}

}


