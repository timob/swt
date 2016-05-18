package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsGestureListenerNative_Gesture(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsGestureListenerNative_Gesture
func go_callback_EventsGestureListenerNative_Gesture(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/GestureListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsGestureListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/GestureEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsGestureEvent{}
	arg_a.Callable = dst_a
i.Gesture(arg_a)
}

var EventsGestureListenerNativeMap = make(map[int]EventsGestureListenerInterface)

type EventsGestureListenerNative struct {
	*javabind.Callable
	EventsGestureListenerInterface
}

func NewEventsGestureListenerNative(implementation EventsGestureListenerInterface) *EventsGestureListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/GestureListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsGestureListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsGestureListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/GestureListenerNative", "gesture", javabind.Void, []interface{}{"org/eclipse/swt/events/GestureEvent"}, C.go_callback_EventsGestureListenerNative_Gesture)

        })
    }
    