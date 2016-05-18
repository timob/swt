package swt

import "github.com/timob/javabind"

type LayoutFormDataInterface interface {
	JavaLangObjectInterface
}

type LayoutFormData struct {
	JavaLangObject
}

// public org.eclipse.swt.layout.FormData()
func NewLayoutFormData() (*LayoutFormData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormData")
	if err != nil {
		panic(err)
	}
	x := &LayoutFormData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormData(int, int)
func NewLayoutFormData2(a int, b int) (*LayoutFormData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormData", a, b)
	if err != nil {
		panic(err)
	}
	x := &LayoutFormData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutFormData) ToString() string {
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

func (jbobject *LayoutFormData) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormData) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormData) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormData) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormData) Left() *LayoutFormAttachment {
	jret, err := jbobject.GetField(javabind.GetEnv(), "left", "org/eclipse/swt/layout/FormAttachment")
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
	unique_x := &LayoutFormAttachment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *LayoutFormData) SetFieldLeft(val LayoutFormAttachmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "left", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *LayoutFormData) Right() *LayoutFormAttachment {
	jret, err := jbobject.GetField(javabind.GetEnv(), "right", "org/eclipse/swt/layout/FormAttachment")
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
	unique_x := &LayoutFormAttachment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *LayoutFormData) SetFieldRight(val LayoutFormAttachmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "right", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *LayoutFormData) Top() *LayoutFormAttachment {
	jret, err := jbobject.GetField(javabind.GetEnv(), "top", "org/eclipse/swt/layout/FormAttachment")
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
	unique_x := &LayoutFormAttachment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *LayoutFormData) SetFieldTop(val LayoutFormAttachmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "top", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *LayoutFormData) Bottom() *LayoutFormAttachment {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bottom", "org/eclipse/swt/layout/FormAttachment")
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
	unique_x := &LayoutFormAttachment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *LayoutFormData) SetFieldBottom(val LayoutFormAttachmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "bottom", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


