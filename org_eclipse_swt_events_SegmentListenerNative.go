package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsSegmentListenerNative_GetSegments(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsSegmentListenerNative_GetSegments
func go_callback_EventsSegmentListenerNative_GetSegments(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/SegmentListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsSegmentListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/SegmentEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsSegmentEvent{}
	arg_a.Callable = dst_a
i.GetSegments(arg_a)
}

var EventsSegmentListenerNativeMap = make(map[int]EventsSegmentListenerInterface)

type EventsSegmentListenerNative struct {
	*javabind.Callable
	EventsSegmentListenerInterface
}

func NewEventsSegmentListenerNative(implementation EventsSegmentListenerInterface) *EventsSegmentListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/SegmentListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsSegmentListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsSegmentListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/SegmentListenerNative", "getSegments", javabind.Void, []interface{}{"org/eclipse/swt/events/SegmentEvent"}, C.go_callback_EventsSegmentListenerNative_GetSegments)

        })
    }
    