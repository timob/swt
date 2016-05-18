package swt

import "github.com/timob/javabind"

type GraphicsDeviceDataInterface interface {
	JavaLangObjectInterface
}

type GraphicsDeviceData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.DeviceData()
func NewGraphicsDeviceData() (*GraphicsDeviceData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/DeviceData")
	if err != nil {
		panic(err)
	}
	x := &GraphicsDeviceData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *GraphicsDeviceData) Display_name() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "display_name", "java/lang/String")
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

func (jbobject *GraphicsDeviceData) SetFieldDisplay_name(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "display_name", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsDeviceData) Application_name() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "application_name", "java/lang/String")
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

func (jbobject *GraphicsDeviceData) SetFieldApplication_name(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "application_name", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsDeviceData) Application_class() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "application_class", "java/lang/String")
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

func (jbobject *GraphicsDeviceData) SetFieldApplication_class(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "application_class", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsDeviceData) Debug() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "debug", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsDeviceData) SetFieldDebug(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "debug", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsDeviceData) Tracking() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "tracking", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsDeviceData) SetFieldTracking(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "tracking", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsDeviceData) Objects() []*JavaLangObject {
	jret, err := jbobject.GetField(javabind.GetEnv(), "objects", javabind.ObjectArrayType("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/lang/Object")
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *GraphicsDeviceData) SetFieldObjects(val []*JavaLangObject) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "objects", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


