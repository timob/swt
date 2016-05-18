package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsDragDetectListenerNative_DragDetected(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsDragDetectListenerNative_DragDetected
func go_callback_EventsDragDetectListenerNative_DragDetected(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/DragDetectListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsDragDetectListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/DragDetectEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsDragDetectEvent{}
	arg_a.Callable = dst_a
i.DragDetected(arg_a)
}

var EventsDragDetectListenerNativeMap = make(map[int]EventsDragDetectListenerInterface)

type EventsDragDetectListenerNative struct {
	*javabind.Callable
	EventsDragDetectListenerInterface
}

func NewEventsDragDetectListenerNative(implementation EventsDragDetectListenerInterface) *EventsDragDetectListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/DragDetectListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsDragDetectListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsDragDetectListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/DragDetectListenerNative", "dragDetected", javabind.Void, []interface{}{"org/eclipse/swt/events/DragDetectEvent"}, C.go_callback_EventsDragDetectListenerNative_DragDetected)

        })
    }
    