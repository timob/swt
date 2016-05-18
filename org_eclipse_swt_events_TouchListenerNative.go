package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsTouchListenerNative_Touch(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsTouchListenerNative_Touch
func go_callback_EventsTouchListenerNative_Touch(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/TouchListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsTouchListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/TouchEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsTouchEvent{}
	arg_a.Callable = dst_a
i.Touch(arg_a)
}

var EventsTouchListenerNativeMap = make(map[int]EventsTouchListenerInterface)

type EventsTouchListenerNative struct {
	*javabind.Callable
	EventsTouchListenerInterface
}

func NewEventsTouchListenerNative(implementation EventsTouchListenerInterface) *EventsTouchListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TouchListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsTouchListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsTouchListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/TouchListenerNative", "touch", javabind.Void, []interface{}{"org/eclipse/swt/events/TouchEvent"}, C.go_callback_EventsTouchListenerNative_Touch)

        })
    }
    