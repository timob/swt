package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsControlListenerNative_ControlMoved(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsControlListenerNative_ControlResized(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsControlListenerNative_ControlMoved
func go_callback_EventsControlListenerNative_ControlMoved(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ControlListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsControlListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ControlEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsControlEvent{}
	arg_a.Callable = dst_a
i.ControlMoved(arg_a)
}

//export go_callback_EventsControlListenerNative_ControlResized
func go_callback_EventsControlListenerNative_ControlResized(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ControlListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsControlListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ControlEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsControlEvent{}
	arg_a.Callable = dst_a
i.ControlResized(arg_a)
}

var EventsControlListenerNativeMap = make(map[int]EventsControlListenerInterface)

type EventsControlListenerNative struct {
	*javabind.Callable
	EventsControlListenerInterface
}

func NewEventsControlListenerNative(implementation EventsControlListenerInterface) *EventsControlListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ControlListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsControlListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsControlListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ControlListenerNative", "controlMoved", javabind.Void, []interface{}{"org/eclipse/swt/events/ControlEvent"}, C.go_callback_EventsControlListenerNative_ControlMoved)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ControlListenerNative", "controlResized", javabind.Void, []interface{}{"org/eclipse/swt/events/ControlEvent"}, C.go_callback_EventsControlListenerNative_ControlResized)

        })
    }
    