package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomBidiSegmentListenerNative_LineGetSegments(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomBidiSegmentListenerNative_LineGetSegments
func go_callback_CustomBidiSegmentListenerNative_LineGetSegments(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/BidiSegmentListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomBidiSegmentListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/BidiSegmentEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomBidiSegmentEvent{}
	arg_a.Callable = dst_a
i.LineGetSegments(arg_a)
}

var CustomBidiSegmentListenerNativeMap = make(map[int]CustomBidiSegmentListenerInterface)

type CustomBidiSegmentListenerNative struct {
	*javabind.Callable
	CustomBidiSegmentListenerInterface
}

func NewCustomBidiSegmentListenerNative(implementation CustomBidiSegmentListenerInterface) *CustomBidiSegmentListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/BidiSegmentListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomBidiSegmentListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomBidiSegmentListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/BidiSegmentListenerNative", "lineGetSegments", javabind.Void, []interface{}{"org/eclipse/swt/custom/BidiSegmentEvent"}, C.go_callback_CustomBidiSegmentListenerNative_LineGetSegments)

        })
    }
    