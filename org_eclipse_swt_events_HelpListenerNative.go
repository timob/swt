package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsHelpListenerNative_HelpRequested(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsHelpListenerNative_HelpRequested
func go_callback_EventsHelpListenerNative_HelpRequested(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/HelpListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsHelpListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/HelpEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsHelpEvent{}
	arg_a.Callable = dst_a
i.HelpRequested(arg_a)
}

var EventsHelpListenerNativeMap = make(map[int]EventsHelpListenerInterface)

type EventsHelpListenerNative struct {
	*javabind.Callable
	EventsHelpListenerInterface
}

func NewEventsHelpListenerNative(implementation EventsHelpListenerInterface) *EventsHelpListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/HelpListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsHelpListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsHelpListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/HelpListenerNative", "helpRequested", javabind.Void, []interface{}{"org/eclipse/swt/events/HelpEvent"}, C.go_callback_EventsHelpListenerNative_HelpRequested)

        })
    }
    