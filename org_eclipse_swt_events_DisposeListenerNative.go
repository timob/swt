package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsDisposeListenerNative_WidgetDisposed(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsDisposeListenerNative_WidgetDisposed
func go_callback_EventsDisposeListenerNative_WidgetDisposed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/DisposeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsDisposeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/DisposeEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsDisposeEvent{}
	arg_a.Callable = dst_a
i.WidgetDisposed(arg_a)
}

var EventsDisposeListenerNativeMap = make(map[int]EventsDisposeListenerInterface)

type EventsDisposeListenerNative struct {
	*javabind.Callable
	EventsDisposeListenerInterface
}

func NewEventsDisposeListenerNative(implementation EventsDisposeListenerInterface) *EventsDisposeListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/DisposeListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsDisposeListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsDisposeListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/DisposeListenerNative", "widgetDisposed", javabind.Void, []interface{}{"org/eclipse/swt/events/DisposeEvent"}, C.go_callback_EventsDisposeListenerNative_WidgetDisposed)

        })
    }
    