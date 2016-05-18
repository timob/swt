package swt

import "github.com/timob/javabind"

type GraphicsFontDataInterface interface {
	JavaLangObjectInterface

	// public int getHeight()
	GetHeight() int

	// public java.lang.String getLocale()
	GetLocale() string

	// public java.lang.String getName()
	GetName() string

	// public int getStyle()
	GetStyle() int

	// public void setHeight(int)
	SetHeight(a int) 

	// public void setLocale(java.lang.String)
	SetLocale(a string) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setStyle(int)
	SetStyle(a int) 
}

type GraphicsFontData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.FontData()
func NewGraphicsFontData() (*GraphicsFontData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/FontData")
	if err != nil {
		panic(err)
	}
	x := &GraphicsFontData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.FontData(java.lang.String)
func NewGraphicsFontData2(a string) (*GraphicsFontData) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/FontData", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsFontData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.FontData(java.lang.String, int, int)
func NewGraphicsFontData3(a string, b int, c int) (*GraphicsFontData) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/FontData", conv_a.Value().Cast("java/lang/String"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsFontData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsFontData) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int getHeight()
func (jbobject *GraphicsFontData) GetHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getLocale()
func (jbobject *GraphicsFontData) GetLocale() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocale", "java/lang/String")
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

// public java.lang.String getName()
func (jbobject *GraphicsFontData) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public int getStyle()
func (jbobject *GraphicsFontData) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int hashCode()
func (jbobject *GraphicsFontData) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setHeight(int)
func (jbobject *GraphicsFontData) SetHeight(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHeight", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLocale(java.lang.String)
func (jbobject *GraphicsFontData) SetLocale(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocale", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setName(java.lang.String)
func (jbobject *GraphicsFontData) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStyle(int)
func (jbobject *GraphicsFontData) SetStyle(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyle", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toString()
func (jbobject *GraphicsFontData) ToString() string {
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

func (jbobject *GraphicsFontData) Name() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "name", "java/lang/String")
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

func (jbobject *GraphicsFontData) SetFieldName(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "name", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsFontData) Height() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsFontData) SetFieldHeight(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsFontData) Style() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "style", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsFontData) SetFieldStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "style", val)
	if err != nil {
		panic(err)
	}

}


