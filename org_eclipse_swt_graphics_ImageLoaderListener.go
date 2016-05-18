package swt

import "github.com/timob/javabind"

type GraphicsImageLoaderListenerInterface interface {

	// public abstract void imageDataLoaded(org.eclipse.swt.graphics.ImageLoaderEvent)
	ImageDataLoaded(a GraphicsImageLoaderEventInterface) 
}

type GraphicsImageLoaderListener struct {
	*javabind.Callable
}

// public abstract void imageDataLoaded(org.eclipse.swt.graphics.ImageLoaderEvent)
func (jbobject *GraphicsImageLoaderListener) ImageDataLoaded(a GraphicsImageLoaderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "imageDataLoaded", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/ImageLoaderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


