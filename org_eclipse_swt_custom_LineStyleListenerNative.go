package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomLineStyleListenerNative_LineGetStyle(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomLineStyleListenerNative_LineGetStyle
func go_callback_CustomLineStyleListenerNative_LineGetStyle(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/LineStyleListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomLineStyleListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/LineStyleEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomLineStyleEvent{}
	arg_a.Callable = dst_a
i.LineGetStyle(arg_a)
}

var CustomLineStyleListenerNativeMap = make(map[int]CustomLineStyleListenerInterface)

type CustomLineStyleListenerNative struct {
	*javabind.Callable
	CustomLineStyleListenerInterface
}

func NewCustomLineStyleListenerNative(implementation CustomLineStyleListenerInterface) *CustomLineStyleListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/LineStyleListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomLineStyleListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomLineStyleListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/LineStyleListenerNative", "lineGetStyle", javabind.Void, []interface{}{"org/eclipse/swt/custom/LineStyleEvent"}, C.go_callback_CustomLineStyleListenerNative_LineGetStyle)

        })
    }
    