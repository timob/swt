package swt

import "github.com/timob/javabind"

type GraphicsResourceInterface interface {
	JavaLangObjectInterface

	// public void dispose()
	Dispose() 

	// public org.eclipse.swt.graphics.Device getDevice()
	GetDevice() *GraphicsDevice

	// public abstract boolean isDisposed()
	IsDisposed() bool
}

type GraphicsResource struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.Resource()
func NewGraphicsResource() (*GraphicsResource) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Resource")
	if err != nil {
		panic(err)
	}
	x := &GraphicsResource{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *GraphicsResource) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Device getDevice()
func (jbobject *GraphicsResource) GetDevice() *GraphicsDevice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDevice", "org/eclipse/swt/graphics/Device")
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
	unique_x := &GraphicsDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract boolean isDisposed()
func (jbobject *GraphicsResource) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


