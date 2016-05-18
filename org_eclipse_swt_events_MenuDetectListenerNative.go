package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsMenuDetectListenerNative_MenuDetected(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsMenuDetectListenerNative_MenuDetected
func go_callback_EventsMenuDetectListenerNative_MenuDetected(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MenuDetectListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMenuDetectListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MenuDetectEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMenuDetectEvent{}
	arg_a.Callable = dst_a
i.MenuDetected(arg_a)
}

var EventsMenuDetectListenerNativeMap = make(map[int]EventsMenuDetectListenerInterface)

type EventsMenuDetectListenerNative struct {
	*javabind.Callable
	EventsMenuDetectListenerInterface
}

func NewEventsMenuDetectListenerNative(implementation EventsMenuDetectListenerInterface) *EventsMenuDetectListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MenuDetectListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsMenuDetectListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsMenuDetectListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MenuDetectListenerNative", "menuDetected", javabind.Void, []interface{}{"org/eclipse/swt/events/MenuDetectEvent"}, C.go_callback_EventsMenuDetectListenerNative_MenuDetected)

        })
    }
    