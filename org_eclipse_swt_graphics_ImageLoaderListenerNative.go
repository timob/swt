package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_GraphicsImageLoaderListenerNative_ImageDataLoaded(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_GraphicsImageLoaderListenerNative_ImageDataLoaded
func go_callback_GraphicsImageLoaderListenerNative_ImageDataLoaded(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/graphics/ImageLoaderListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := GraphicsImageLoaderListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/graphics/ImageLoaderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &GraphicsImageLoaderEvent{}
	arg_a.Callable = dst_a
i.ImageDataLoaded(arg_a)
}

var GraphicsImageLoaderListenerNativeMap = make(map[int]GraphicsImageLoaderListenerInterface)

type GraphicsImageLoaderListenerNative struct {
	*javabind.Callable
	GraphicsImageLoaderListenerInterface
}

func NewGraphicsImageLoaderListenerNative(implementation GraphicsImageLoaderListenerInterface) *GraphicsImageLoaderListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageLoaderListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &GraphicsImageLoaderListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    GraphicsImageLoaderListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/graphics/ImageLoaderListenerNative", "imageDataLoaded", javabind.Void, []interface{}{"org/eclipse/swt/graphics/ImageLoaderEvent"}, C.go_callback_GraphicsImageLoaderListenerNative_ImageDataLoaded)

        })
    }
    