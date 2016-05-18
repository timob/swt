package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsTraverseListenerNative_KeyTraversed(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsTraverseListenerNative_KeyTraversed
func go_callback_EventsTraverseListenerNative_KeyTraversed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/TraverseListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsTraverseListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/TraverseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsTraverseEvent{}
	arg_a.Callable = dst_a
i.KeyTraversed(arg_a)
}

var EventsTraverseListenerNativeMap = make(map[int]EventsTraverseListenerInterface)

type EventsTraverseListenerNative struct {
	*javabind.Callable
	EventsTraverseListenerInterface
}

func NewEventsTraverseListenerNative(implementation EventsTraverseListenerInterface) *EventsTraverseListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TraverseListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsTraverseListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsTraverseListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/TraverseListenerNative", "keyTraversed", javabind.Void, []interface{}{"org/eclipse/swt/events/TraverseEvent"}, C.go_callback_EventsTraverseListenerNative_KeyTraversed)

        })
    }
    