package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsMouseMoveListenerNative_MouseMove(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsMouseMoveListenerNative_MouseMove
func go_callback_EventsMouseMoveListenerNative_MouseMove(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseMoveListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseMoveListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseMove(arg_a)
}

var EventsMouseMoveListenerNativeMap = make(map[int]EventsMouseMoveListenerInterface)

type EventsMouseMoveListenerNative struct {
	*javabind.Callable
	EventsMouseMoveListenerInterface
}

func NewEventsMouseMoveListenerNative(implementation EventsMouseMoveListenerInterface) *EventsMouseMoveListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseMoveListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsMouseMoveListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsMouseMoveListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseMoveListenerNative", "mouseMove", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseMoveListenerNative_MouseMove)

        })
    }
    