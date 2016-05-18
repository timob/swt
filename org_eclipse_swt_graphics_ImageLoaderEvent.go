package swt

import "github.com/timob/javabind"

type GraphicsImageLoaderEventInterface interface {

	// public java.lang.String toString()
	ToString() string
}

type GraphicsImageLoaderEvent struct {
	*javabind.Callable
}

// public org.eclipse.swt.graphics.ImageLoaderEvent(org.eclipse.swt.graphics.ImageLoader, org.eclipse.swt.graphics.ImageData, int, boolean)
func NewGraphicsImageLoaderEvent(a GraphicsImageLoaderInterface, b GraphicsImageDataInterface, c int, d bool) (*GraphicsImageLoaderEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageLoaderEvent", conv_a.Value().Cast("org/eclipse/swt/graphics/ImageLoader"), conv_b.Value().Cast("org/eclipse/swt/graphics/ImageData"), c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsImageLoaderEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *GraphicsImageLoaderEvent) ToString() string {
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

func (jbobject *GraphicsImageLoaderEvent) ImageData() *GraphicsImageData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "imageData", "org/eclipse/swt/graphics/ImageData")
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
	unique_x := &GraphicsImageData{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsImageLoaderEvent) SetFieldImageData(val GraphicsImageDataInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "imageData", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsImageLoaderEvent) IncrementCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "incrementCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageLoaderEvent) SetFieldIncrementCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "incrementCount", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageLoaderEvent) EndOfImage() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "endOfImage", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsImageLoaderEvent) SetFieldEndOfImage(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "endOfImage", val)
	if err != nil {
		panic(err)
	}

}


