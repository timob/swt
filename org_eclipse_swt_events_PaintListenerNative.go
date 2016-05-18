package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsPaintListenerNative_PaintControl(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsPaintListenerNative_PaintControl
func go_callback_EventsPaintListenerNative_PaintControl(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/PaintListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsPaintListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/PaintEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsPaintEvent{}
	arg_a.Callable = dst_a
i.PaintControl(arg_a)
}

var EventsPaintListenerNativeMap = make(map[int]EventsPaintListenerInterface)

type EventsPaintListenerNative struct {
	*javabind.Callable
	EventsPaintListenerInterface
}

func NewEventsPaintListenerNative(implementation EventsPaintListenerInterface) *EventsPaintListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/PaintListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsPaintListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsPaintListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/PaintListenerNative", "paintControl", javabind.Void, []interface{}{"org/eclipse/swt/events/PaintEvent"}, C.go_callback_EventsPaintListenerNative_PaintControl)

        })
    }
    