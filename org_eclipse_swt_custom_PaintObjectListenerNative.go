package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomPaintObjectListenerNative_PaintObject(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomPaintObjectListenerNative_PaintObject
func go_callback_CustomPaintObjectListenerNative_PaintObject(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/PaintObjectListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomPaintObjectListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/PaintObjectEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomPaintObjectEvent{}
	arg_a.Callable = dst_a
i.PaintObject(arg_a)
}

var CustomPaintObjectListenerNativeMap = make(map[int]CustomPaintObjectListenerInterface)

type CustomPaintObjectListenerNative struct {
	*javabind.Callable
	CustomPaintObjectListenerInterface
}

func NewCustomPaintObjectListenerNative(implementation CustomPaintObjectListenerInterface) *CustomPaintObjectListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/PaintObjectListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomPaintObjectListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomPaintObjectListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/PaintObjectListenerNative", "paintObject", javabind.Void, []interface{}{"org/eclipse/swt/custom/PaintObjectEvent"}, C.go_callback_CustomPaintObjectListenerNative_PaintObject)

        })
    }
    