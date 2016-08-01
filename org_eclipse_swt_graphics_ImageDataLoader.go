package swt

import "github.com/timob/javabind"

type GraphicsImageDataLoaderInterface interface {
	JavaLangObjectInterface

	// public static org.eclipse.swt.graphics.ImageData[] load(java.lang.String)
	Load(a string) []*GraphicsImageData
}

type GraphicsImageDataLoader struct {
	JavaLangObject
}

// public static org.eclipse.swt.graphics.ImageData[] load(java.lang.String)
func (jbobject *GraphicsImageDataLoader) Load(a string) []*GraphicsImageData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/ImageDataLoader", "load", javabind.ObjectArrayType("org/eclipse/swt/graphics/ImageData"), conv_a.Value().Cast("java/lang/String"))
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


