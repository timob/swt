package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsExpandListenerNative_ItemCollapsed(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsExpandListenerNative_ItemExpanded(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsExpandListenerNative_ItemCollapsed
func go_callback_EventsExpandListenerNative_ItemCollapsed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ExpandListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsExpandListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ExpandEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsExpandEvent{}
	arg_a.Callable = dst_a
i.ItemCollapsed(arg_a)
}

//export go_callback_EventsExpandListenerNative_ItemExpanded
func go_callback_EventsExpandListenerNative_ItemExpanded(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ExpandListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsExpandListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ExpandEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsExpandEvent{}
	arg_a.Callable = dst_a
i.ItemExpanded(arg_a)
}

var EventsExpandListenerNativeMap = make(map[int]EventsExpandListenerInterface)

type EventsExpandListenerNative struct {
	*javabind.Callable
	EventsExpandListenerInterface
}

func NewEventsExpandListenerNative(implementation EventsExpandListenerInterface) *EventsExpandListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ExpandListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsExpandListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsExpandListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ExpandListenerNative", "itemCollapsed", javabind.Void, []interface{}{"org/eclipse/swt/events/ExpandEvent"}, C.go_callback_EventsExpandListenerNative_ItemCollapsed)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ExpandListenerNative", "itemExpanded", javabind.Void, []interface{}{"org/eclipse/swt/events/ExpandEvent"}, C.go_callback_EventsExpandListenerNative_ItemExpanded)

        })
    }
    