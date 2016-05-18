package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsArmListenerNative_WidgetArmed(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsArmListenerNative_WidgetArmed
func go_callback_EventsArmListenerNative_WidgetArmed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ArmListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsArmListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ArmEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsArmEvent{}
	arg_a.Callable = dst_a
i.WidgetArmed(arg_a)
}

var EventsArmListenerNativeMap = make(map[int]EventsArmListenerInterface)

type EventsArmListenerNative struct {
	*javabind.Callable
	EventsArmListenerInterface
}

func NewEventsArmListenerNative(implementation EventsArmListenerInterface) *EventsArmListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ArmListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsArmListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsArmListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ArmListenerNative", "widgetArmed", javabind.Void, []interface{}{"org/eclipse/swt/events/ArmEvent"}, C.go_callback_EventsArmListenerNative_WidgetArmed)

        })
    }
    