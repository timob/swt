package swt

import "github.com/timob/javabind"

type GraphicsDrawableInterface interface {

	// public abstract long internal_new_GC(org.eclipse.swt.graphics.GCData)
	Internal_new_GC(a GraphicsGCDataInterface) int64

	// public abstract void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
	Internal_dispose_GC(a int64, b GraphicsGCDataInterface) 
}

type GraphicsDrawable struct {
	JavaLangObject
}

// public abstract long internal_new_GC(org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsDrawable) Internal_new_GC(a GraphicsGCDataInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "internal_new_GC", javabind.Long, conv_a.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public abstract void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsDrawable) Internal_dispose_GC(a int64, b GraphicsGCDataInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "internal_dispose_GC", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}


