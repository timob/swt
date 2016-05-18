package swt

import "github.com/timob/javabind"

type GraphicsImageLoaderInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.graphics.ImageData[] load(java.lang.String)
	Load(a string) []*GraphicsImageData

	// public void save(java.lang.String, int)
	Save(a string, b int) 

	// public void addImageLoaderListener(org.eclipse.swt.graphics.ImageLoaderListener)
	AddImageLoaderListener(a GraphicsImageLoaderListenerInterface) 

	// public void removeImageLoaderListener(org.eclipse.swt.graphics.ImageLoaderListener)
	RemoveImageLoaderListener(a GraphicsImageLoaderListenerInterface) 

	// public boolean hasListeners()
	HasListeners() bool

	// public void notifyListeners(org.eclipse.swt.graphics.ImageLoaderEvent)
	NotifyListeners(a GraphicsImageLoaderEventInterface) 
}

type GraphicsImageLoader struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.ImageLoader()
func NewGraphicsImageLoader() (*GraphicsImageLoader) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageLoader")
	if err != nil {
		panic(err)
	}
	x := &GraphicsImageLoader{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.ImageData[] load(java.lang.String)
func (jbobject *GraphicsImageLoader) Load(a string) []*GraphicsImageData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "load", javabind.ObjectArrayType("org/eclipse/swt/graphics/ImageData"), conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/ImageData")
	dst := new([]*GraphicsImageData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void save(java.lang.String, int)
func (jbobject *GraphicsImageLoader) Save(a string, b int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "save", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addImageLoaderListener(org.eclipse.swt.graphics.ImageLoaderListener)
func (jbobject *GraphicsImageLoader) AddImageLoaderListener(a GraphicsImageLoaderListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addImageLoaderListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/ImageLoaderListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeImageLoaderListener(org.eclipse.swt.graphics.ImageLoaderListener)
func (jbobject *GraphicsImageLoader) RemoveImageLoaderListener(a GraphicsImageLoaderListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeImageLoaderListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/ImageLoaderListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean hasListeners()
func (jbobject *GraphicsImageLoader) HasListeners() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasListeners", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void notifyListeners(org.eclipse.swt.graphics.ImageLoaderEvent)
func (jbobject *GraphicsImageLoader) NotifyListeners(a GraphicsImageLoaderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "notifyListeners", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/ImageLoaderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *GraphicsImageLoader) Data() []*GraphicsImageData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "data", javabind.ObjectArrayType("org/eclipse/swt/graphics/ImageData"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/ImageData")
	dst := new([]*GraphicsImageData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *GraphicsImageLoader) SetFieldData(val []*GraphicsImageData) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/ImageData")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "data", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsImageLoader) LogicalScreenWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "logicalScreenWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoader) SetFieldLogicalScreenWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "logicalScreenWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageLoader) LogicalScreenHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "logicalScreenHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoader) SetFieldLogicalScreenHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "logicalScreenHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageLoader) BackgroundPixel() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "backgroundPixel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoader) SetFieldBackgroundPixel(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "backgroundPixel", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageLoader) RepeatCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "repeatCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoader) SetFieldRepeatCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "repeatCount", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageLoader) Compression() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "compression", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoader) SetFieldCompression(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "compression", val)
	if err != nil {
		panic(err)
	}

}


